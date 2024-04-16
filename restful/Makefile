.PHONY: api swagger format doc fmt

format:
	@goctl api format --dir api --declare

api:
	@goctl api format --dir api --declare
	@goctl api go --api api/restful.api --dir . --home "../common/template/goctl"

swagger:
	@goctl api plugin -plugin goctl-swagger="swagger -filename api/rest.swagger.json  -host 127.0.0.1:8888"  -api api/restful.api -dir .

doc:
	@swag2md -s api/rest.swagger.json -o api/rest.swagger.md -t "接口文档"


fmt:
	@find . -name '*.go' -not -path "./vendor/*" -not -name "*.pb.go" -not -name "*.gen.go" | xargs gofumpt -w -extra
	@find . -name '*.go' -not -path "./vendor/*" -not -name "*.pb.go" -not -name "*.gen.go" | xargs -n 1 -t goimports-reviser -rm-unused -set-alias -company-prefixes "github.com/ch3nnn/goview-gozero" -project-name "github.com/ch3nnn/goview-gozero/restful"
