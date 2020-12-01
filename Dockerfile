FROM scratch

WORKDIR app

COPY ./ ./

ENTRYPOINT ["./helloworld"]