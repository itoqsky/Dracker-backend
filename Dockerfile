FROM golang:alpine AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/itoqsky/money-tracker-backend
WORKDIR /github.com/itoqsky/money-tracker-backend

RUN go mod download && go get -u ./...
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/itoqsky/money-tracker-backend/.bin/app .
COPY --from=0 /github.com/itoqsky/money-tracker-backend/configs ./configs/

EXPOSE 8000

CMD ["./app"]