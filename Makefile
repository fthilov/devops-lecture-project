GO_CMD = go


all: build-all test start-all

build-all:
	$(GO_CMD) build -o auth-main ./auth-service/cmd/main.go 
	$(GO_CMD) build -o product-main ./product-service/cmd/main.go 
	$(GO_CMD) build -o checkout-main ./checkout-service/cmd/main.go

build:
	$(GO_CMD) build -o $(service)-main ./$(service)-service/cmd/main.go 

test:
	$(GO_CMD) test -v ./...

clean-all:
	rm -f auth-main 
	rm -f product-main 
	rm -f checkout-main

clean:
	rm -f $(service)-main

start-all:
	./auth-main &
	./product-main &
	./checkout-main &

start:
	./$(service)-main &
