build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -o short-service -a -installsuffix cgo .
	docker build -t short-service .
run:
	docker run \
		--name short-service \
        --link redis:redis \
        -p 8000:8000 \
        -d short-service
