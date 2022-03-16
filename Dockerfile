FROM golang:1.16-alpine3.15 AS builder

RUN go version

COPY . /forms/
WORKDIR /forms

RUN go mod download && go mod tidy && \
    GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 forms/.bin/app .
COPY --from=0 forms/configs configs/
COPY --from=0 forms/templates templates/

CMD ["./app"]