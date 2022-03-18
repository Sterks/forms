FROM golang:1.16-alpine3.15 AS builder

RUN go version

COPY . /forms/
WORKDIR /forms

RUN go mod download && GOOS=linux go build -o ./.bin/app ./cmd/app/main.go
RUN ls -la

FROM alpine:latest

RUN ls -la

WORKDIR /root/

COPY --from=0 forms/.bin/app .
COPY --from=0 forms/configs configs/
COPY --from=0 forms/templates templates/

CMD ["./app"]