.PHONY: api swagger format doc

format:
	@goctl api format --dir api --declare

api:
	@goctl api format --dir api --declare
	@goctl api go --api api/restful.api --dir . --home "../common/template/goctl"

swagger:
	@goctl api plugin -plugin goctl-swagger="swagger -filename api/rest.swagger.json  -host 127.0.0.1:8888"  -api api/restful.api -dir .

doc:
	@swag2md -s api/rest.swagger.json -o api/rest.swagger.md -t "接口文档"
