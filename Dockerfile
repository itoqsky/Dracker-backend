FROM golang:alpine AS builder

RUN go version
RUN apk add git

COPY ./ /github.com/itoqsky/money-tracker-backend
WORKDIR /github.com/itoqsky/money-tracker-backend

RUN go mod download && go get -u ./...
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

RUN chmod +x wait-for-postgres.sh

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /github.com/itoqsky/money-tracker-backend/.bin/app .
COPY --from=0 /github.com/itoqsky/money-tracker-backend/configs/ ./configs/
# COPY --from=0 /github.com/itoqsky/money-tracker-backend/wait-for-postgres.sh ./wait-for-postgres.sh

EXPOSE 8000

CMD ["./app"]