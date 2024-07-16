build:
	go build -o random

lint:
	goimports -w .
	gci write .

test:
	./scripts/test_with_coverage.sh

version:
	./scripts/create_git_version_tag.sh

