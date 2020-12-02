FROM scratch

WORKDIR app

COPY ./helloworld ./helloworld

ENTRYPOINT ["./helloworld"]
