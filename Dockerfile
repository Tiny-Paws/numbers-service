FROM golang:1.15.2-alpine3.12

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
WORKDIR /go/src/app/cmd
RUN GOBIN=/bin go install

EXPOSE 8080
CMD ["cmd"]