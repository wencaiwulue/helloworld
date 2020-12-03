FROM scratch

WORKDIR /

COPY Dockerfile Dockerfile
COPY main.go main.go
COPY deploy deploy
COPY helloworld helloworld
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /helloworld

ENTRYPOINT ["/helloworld"]
