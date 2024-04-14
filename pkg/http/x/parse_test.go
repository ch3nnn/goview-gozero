package x

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueries(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "http://test.com/api/test?id=1&hash=a&hash=b", nil)
	assert.NoError(t, err)

	assert.Equal(t, "1", Query(r, "id"))
	assert.Equal(t, []string{"1"}, QueryArray(r, "id"))
	assert.Equal(t, "a", Query(r, "hash"))
	assert.Equal(t, []string{"a", "b"}, QueryArray(r, "hash"))
}

func TestParse(t *testing.T) {
	r, err := http.NewRequest(http.MethodPost, "http://test.com/api/chunk/upload", strings.NewReader(`
    { 
		"current_data": "abcdefgh",
    	"current_seq": 1,
    	"current_size": 8,
    	"file_name": "test.txt",
    	"file_hash": "ec3f5c9819f41ec8965587553fbe9935ec26ec440c5adc94ff6c10efadeba80f",
    	"total_seq": 1,
    	"total_size": 8 
	}`))
	require.NoError(t, err)
	r.Header.Set("Content-Type", "application/json")

	req := &_UploadFileChunkReq{}
	err = Parse(r, req)
	require.NoError(t, err)
	assert.Equal(t, "abcdefgh", req.CurrentData)
	assert.Equal(t, int32(1), req.CurrentSeq)
	assert.Equal(t, int32(8), req.CurrentSize)
	assert.Equal(t, "test.txt", req.FileName)
	assert.Equal(t, "ec3f5c9819f41ec8965587553fbe9935ec26ec440c5adc94ff6c10efadeba80f", req.FileHash)
	assert.Equal(t, int64(1), req.TotalSeq)
	assert.Equal(t, int64(8), req.TotalSize)
}

func TestGetClientIP(t *testing.T) {
	cases := []struct {
		k      string
		v      string
		expect string
	}{
		{k: "X-Forwarded-For", v: "127.0.0.1", expect: "127.0.0.1"},
		{k: "X-Real-Ip", v: "127.0.0.1", expect: "127.0.0.1"},
		{k: "X-Appengine-Remote-Addr", v: "127.0.0.1", expect: "127.0.0.1"},
		{k: "Test", v: "127.0.0.1", expect: ""},
		{k: "", v: "", expect: ""},
	}

	for _, c := range cases {
		r, err := http.NewRequest(http.MethodGet, "http://test.com/api/test?id=1&hash=a&hash=b", nil)
		require.NoError(t, err)
		r.Header.Set(c.k, c.v)
		assert.Equal(t, c.expect, GetClientIP(r))
	}
}

func TestGetExternalIP(t *testing.T) {
	eip, err := GetExternalIP()
	assert.NoError(t, err)
	assert.NotEmpty(t, eip)
	t.Log(eip)
}

func TestGetInternalIP(t *testing.T) {
	iip := GetInternalIP()
	assert.NotEmpty(t, iip)
	t.Log(iip)
}

type _UploadFileChunkReq struct {
	CurrentData string `json:"current_data" validate:"required" label:"当前块数据"`                  // 当前块数据（须 base64 编码）
	CurrentSeq  int32  `json:"current_seq" validate:"required" label:"当前块序号"`                   // 当前块序号（从 1 开始）
	CurrentSize int32  `json:"current_size" validate:"required" label:"当前块大小"`                  // 当前块大小
	FileName    string `json:"file_name" validate:"required" label:"文件名"`                       // 文件名
	FileHash    string `json:"file_hash" validate:"required,len=64,hexadecimal" label:"文件hash"` // 文件hash（sha256(文件内容+文件名称)）
	TotalSeq    int64  `json:"total_seq" validate:"required" label:"总序号"`                       // 总序号
	TotalSize   int64  `json:"total_size" validate:"required" label:"总大小"`                      // 总大小
}
