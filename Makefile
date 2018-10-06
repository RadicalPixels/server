all:
	@echo "no default"

.PHONY: build/docker
build/docker:
	@docker build -t radpix:latest .

.PHONY: start
start:
	@go run main.go

.PHONY: start/docker
start/docker:
	@docker run -p 8000:8000 -t radpix:latest

.PHONY: clean
clean:
	@rm data.txt

.PHONY: test/upload
test/upload:
	@echo 'hello world' > data.txt
	@curl -X POST "http://localhost:8000/upload" -F file=@data.txt

.PHONY: test/ipfsutil
test/ipfsutil:
	@go test -v ipfsutil/*.go
