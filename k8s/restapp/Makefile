.PHONY: swagger
swagger:
	swagger generate server -f swagger/swagger.yaml -t ./gen --default-scheme http
build:
	docker build . -f .devops/go/Dockerfile -t $(tag)