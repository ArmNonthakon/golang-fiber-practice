run: openapi-combine openapi-gen jet-gen-database test-cover
	go run cmd/main.go

run-with-doc: openapi-combine openapi-gen openapi-doc jet-gen-database
	go run cmd/main.go

openapi-combine:
	redocly bundle openapi/root.yaml -o openapi/openapi.yaml 

openapi-gen:
	oapi-codegen -package=server -generate "spec,fiber" openapi/openapi.yaml > internal/generated/server/server.gen.go

openapi-doc:
	redocly build-docs openapi/openapi.yaml --output=openapi/openapi-static.html

jet-gen-database:
	jet -source=mysql -host=localhost -port=3306 -user=root -password=Ninjaarm-2003 -dbname=go_database -schema=public -path=./internal/data/database/jet_generated

deploy:
	serverless deploy

test:
	go test ./...

test-cover:
	go test -cover ./...