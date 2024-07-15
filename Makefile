build:
	go build -o random

lint:
	goimports -w .
	gci write .

test:
	go test -v ./... -cover
