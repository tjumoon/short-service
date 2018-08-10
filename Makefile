build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -o short-service -a -installsuffix cgo .
	docker build -t short-service .
run:
	docker stop short-service1
	docker rm short-service1
	docker run \
		--name short-service1 \
        --link redis:redis \
        -p 8000:8000 \
        -d short-service
	docker stop short-service2
	docker rm short-service2
	docker run \
    	--name short-service2 \
        --link redis:redis \
        -p 8000:8001 \
        -d short-service
