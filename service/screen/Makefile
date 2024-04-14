.PHONY: rpc clang-format doc


rpc:
	@goctl rpc protoc pb/screen.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m --home "../../common/template/goctl"


# 格式化 protobuf 文件
clang-format:
	@clang-format ./pb/screen.proto -i -style="{AlignConsecutiveAssignments: true,AlignConsecutiveDeclarations: true,AllowShortFunctionsOnASingleLine: None,ColumnLimit: 0,IndentWidth: 4,Language: Proto,SpacesInLineCommentPrefix: {Minimum: 1,Maximum: 1}}"
	@echo "\033[32m Successfully formatted protobuf files \033[0m"


doc:
	# Google Protocol Buffers 的文档生成器插件
	# https://github.com/pseudomuto/protoc-gen-doc
	@protoc pb/*.proto --doc_out=./pb --doc_opt=../../sa-micro-common/template/protoc-gen-doc/markdown.tmpl,doc.md
