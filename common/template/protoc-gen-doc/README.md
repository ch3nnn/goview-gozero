# Google Protocol Buffers 的文档生成器插件(模板)

> https://github.com/pseudomuto/protoc-gen-doc

## 安装
```shell
go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest  
```

## 生成 md 文档(自定义模板)
```shell
protoc --doc_out=. --doc_opt=../../sa-micro-common/template/protoc-gen-doc/markdown.tmpl,index.md pb/*.proto
```