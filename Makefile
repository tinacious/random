build:
	go build -o random

lint:
	goimports -w .
	gci write .

test:
	./scripts/test_with_coverage.sh

version:
	./scripts/create_git_version_tag.sh

doc:
	GENERATE_RANDOM_DOCS=1 go run main.go

setup:
	go install golang.org/x/tools/gopls@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/daixiang0/gci@latest

coverage:
	go tool cover -html=coverage.out
