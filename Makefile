.PHONY: build
build:
	docker build -t tt/truthtai-api .

.PHONY: start
start:
	docker run -d -p 3000:3000 tt/truthtai-api
