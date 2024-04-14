package errcode

import (
	stderrors "errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	cases := []struct {
		code           int
		msg            string
		httpCode       int
		expectHTTPCode int
	}{
		{code: 101, msg: "101 error", httpCode: 200, expectHTTPCode: 200},
		{code: 102, msg: "102 error", httpCode: 0, expectHTTPCode: 200},
		{code: 103, msg: "103 error", httpCode: 301, expectHTTPCode: 301},
	}

	for _, c := range cases {
		e := New(c.code, c.msg, c.httpCode)
		require.Error(t, e)

		var err *Err
		ok := errors.As(e, &err)

		assert.True(t, ok)
		assert.NotNil(t, err)
		assert.Equal(t, c.code, err.Code)
		assert.Equal(t, c.msg, err.Msg)
		assert.Equal(t, c.expectHTTPCode, err.HttpCode)
	}
}

func TestErr_Error(t *testing.T) {
	cases := []struct {
		code        int
		msg         string
		httpCode    int
		expectError string
	}{
		{code: 101, msg: "101 error", httpCode: 200, expectError: "101 error"},
		{code: 102, msg: "102 error", httpCode: 0, expectError: "102 error"},
		{code: 103, msg: "103 error", httpCode: 301, expectError: "103 error"},
	}

	for _, c := range cases {
		e := New(c.code, c.msg, c.httpCode)
		require.Error(t, e)

		var err *Err
		ok := errors.As(e, &err)

		assert.True(t, ok)
		assert.NotNil(t, err)
		assert.Equal(t, c.expectError, err.Error())
	}
}

func TestErr_String(t *testing.T) {
	cases := []struct {
		code         int
		msg          string
		httpCode     int
		expectString string
	}{
		{code: 101, msg: "101 error", httpCode: 200, expectString: "code: 101, msg: 101 error, http code: 200"},
		{code: 102, msg: "102 error", httpCode: 0, expectString: "code: 102, msg: 102 error, http code: 200"},
		{code: 103, msg: "103 error", httpCode: 301, expectString: "code: 103, msg: 103 error, http code: 301"},
	}

	for _, c := range cases {
		e := New(c.code, c.msg, c.httpCode)
		require.Error(t, e)

		var err *Err
		ok := errors.As(e, &err)

		assert.True(t, ok)
		assert.NotNil(t, err)
		assert.Equal(t, c.expectString, err.String())
	}
}

func TestErr_GRPCStatus(t *testing.T) {
	cases := []struct {
		code         int
		msg          string
		httpCode     int
		expectString string
	}{
		{code: 101, msg: "101 error", httpCode: 200},
		{code: 102, msg: "102 error", httpCode: 0},
		{code: 103, msg: "103 error", httpCode: 301},
	}

	for _, c := range cases {
		e := New(c.code, c.msg, c.httpCode)
		require.Error(t, e)

		var err *Err
		ok := errors.As(e, &err)

		assert.True(t, ok)
		assert.NotNil(t, err)

		grpcStatus := err.GRPCStatus()
		assert.NotNil(t, grpcStatus)
		assert.Equal(t, int(grpcStatus.Code()), err.Code)
		assert.Equal(t, grpcStatus.Message(), err.String())
		if uint32(grpcStatus.Code()) > GrpcMaxCode {
			assert.Equal(t, grpcStatus.Code().String(), fmt.Sprintf("Code(%d)", err.Code))
		}
	}
}

func TestFromMessage(t *testing.T) {
	cases := []struct {
		message   string
		expectErr error
		expectOK  bool
	}{
		{message: "code: 101, msg: 101 error, http code: 200", expectErr: New(101, "101 error"), expectOK: true},
		{message: "code: 102, msg: 102 error, http code: 200", expectErr: New(102, "102 error"), expectOK: true},
		{message: "", expectErr: ErrUnexpected, expectOK: false},
		{message: "test message", expectErr: ErrUnexpected, expectOK: false},
	}

	for _, c := range cases {
		err, ok := FromMessage(c.message)
		assert.Equal(t, c.expectOK, ok)
		assert.NotNil(t, err)
		assert.True(t, IsErr(err))
		assert.True(t, Is(err, c.expectErr))
	}
}

func TestFromError(t *testing.T) {
	cases := []struct {
		err    error
		expect bool
	}{
		{err: NoErr, expect: true},
		{err: ErrUnexpected, expect: true},
		{err: ErrInvalidParams, expect: true},
		{err: errors.WithMessage(ErrInvalidParams, "test wrap"), expect: true},
		{err: nil, expect: true},
		{err: New(CodeUnexpected, MsgUnexpected), expect: true},
		{err: errors.New("test"), expect: false},
		{err: stderrors.New("test"), expect: false},
		{err: New(101, "101 error"), expect: true},
	}

	for _, c := range cases {
		err, ok := FromError(c.err)
		assert.Equal(t, c.expect, ok)
		require.Error(t, err)
		assert.True(t, IsErr(err))
	}
}

func TestIs(t *testing.T) {
	cases := []struct {
		err    error
		target error
		expect bool
	}{
		{err: New(CodeOK, MsgOK), target: NoErr, expect: true},
		{err: New(CodeUnexpected, MsgUnexpected), target: ErrUnexpected, expect: true},
		{err: ErrInvalidParams, target: ErrUnexpected, expect: false},
		{err: New(CodeUnexpected, MsgUnexpected), target: ErrInvalidParams, expect: false},
		{err: New(CodeUnexpected, MsgUnexpected), target: New(CodeUnexpected, MsgUnexpected), expect: true},
		{err: New(CodeInvalidParams, MsgInvalidParams), target: ErrInvalidParams, expect: true},
		{err: ErrInvalidParams, target: errors.New("错误"), expect: false},
		{err: errors.New("错误"), target: ErrUnexpected, expect: false},
	}

	for _, c := range cases {
		assert.Equal(t, c.expect, Is(c.err, c.target))
	}
}
