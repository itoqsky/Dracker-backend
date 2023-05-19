# stage 1
FROM golang:alpine AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/itoqsky/money-tracker-backend
WORKDIR /github.com/itoqsky/money-tracker-backend

RUN go mod download && go get -u ./...
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

# stage 2

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/itoqsky/money-tracker-backend/.bin/app .
COPY --from=0 /github.com/itoqsky/money-tracker-backend/configs/ ./configs/

EXPOSE 8080

CMD ./app