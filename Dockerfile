FROM golang:1.14

WORKDIR /

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /helloworld

ENTRYPOINT ["/helloworld"]
