build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o short-service -a -installsuffix cgo .
	docker build -t short-service .
run:
	docker run \
        --link redis:redis \
        -p 8000:8000 \
        short-service
