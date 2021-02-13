FROM golang:1.14-alpine

RUN apk update && apk add git

WORKDIR /go/src/family-board

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

EXPOSE 9000

ENTRYPOINT ["/go/bin/app"]
