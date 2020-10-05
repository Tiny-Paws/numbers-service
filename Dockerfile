FROM golang:1.15.2-alpine3.12

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

EXPOSE 8080
CMD ["go run cmd/main.go"]