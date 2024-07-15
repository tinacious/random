build:
	go build -o random

lint:
	goimports -w .
	gci write .

test:
	go test -v ./... -cover

version:
	echo 'Copy the following tag:';
	echo "\n0.0.0-$$(date +%s)\n";
