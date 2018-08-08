FROM scratch

WORKDIR $GOPATH/src/short-service
COPY . $GOPATH/src/short-service

EXPOSE 8000
ENTRYPOINT ["./short-service"]

