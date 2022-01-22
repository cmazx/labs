FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY ./ ./
RUN go mod download
RUN GOOS=linux go build  -o .bin/main main.go


FROM alpine

WORKDIR /app

COPY --from=builder /app/.bin/main ./

CMD ["./main"]

