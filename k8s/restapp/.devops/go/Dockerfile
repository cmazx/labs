FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY ./swagger ./swagger
COPY ./gen ./gen
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./.devops/bin/swagger /usr/bin/swagger
RUN chmod 555 /usr/bin/swagger
RUN swagger generate server -f swagger/swagger.yaml  -t ./gen --default-scheme http

RUN go mod download
COPY ./ ./

RUN GOOS=linux go build  -o .bin/main ./gen/cmd/crud-service-server/main.go


FROM alpine

WORKDIR /app

COPY --from=builder /app/.bin/main ./
CMD ["./main"]