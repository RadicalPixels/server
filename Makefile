all:
	@echo "no default"

.PHONY: build/docker
build/docker:
	@docker build -t radpix:latest .

.PHONY: build/contract
build/contract:
	@abigen --abi="contract/contract.abi" --pkg=contract --out=contract/contract.go

.PHONY: start
start:
	@go run main.go start --node-url $(URL) --address $(ADDRESS)

.PHONY: start/docker
start/docker:
	@docker run -p 8000:8000 -t radpix:latest $(ARGS)

.PHONY: test/upload
test/upload:
	@echo 'hello world' > data.txt
	@curl -X POST "http://localhost:8000/upload" -F file=@data.txt

.PHONY: test/grid
test/grid:
	@curl -X GET "http://localhost:8000/grid"

.PHONY: test/colors
test/colors:
	@curl -X POST "http://localhost:8000/content" -H 'Content-Type: application/json' -d '["#fff","#000"]'

.PHONY: test/ipfsutil
test/ipfsutil:
	@go test -v ipfsutil/*.go

.PHONY: test/client
test/client:
	@go test -v radpixclient/*.go $(ARGS)

.PHONY: clean
clean:
	@rm data.txt

.PHONY: kill
kill:
	@docker rm -f $$(docker ps -aq)
