FROM scratch

WORKDIR /

COPY ./helloworld ./helloworld
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o /helloworld

ENTRYPOINT ["./helloworld"]
