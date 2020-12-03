FROM scratch

WORKDIR /

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -mod=readonly -v -o /helloworld

ENTRYPOINT ["./helloworld"]
