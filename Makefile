all:
	@echo "no default"

.PHONY: run
run:
	@go run main.go

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
