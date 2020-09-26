APP=hasura-firebase-auth

build:
	docker build -t $(APP) .

run: build
	docker run -it -e CACHE_DURATION=300 -v $(shell pwd)/service-account.json:/root/service-account.json -p 8080:8080 $(APP)