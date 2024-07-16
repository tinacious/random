#!/usr/bin/env sh

go test -v ./... -coverprofile coverage.out
COVERAGE=`go tool cover -func=coverage.out | grep total: | grep -Eo '[0-9]+\.[0-9]+'`
echo $COVERAGE

# Average
COLOR=F59300

if (( $(echo "$COVERAGE <= 50" | bc -l) )) ; then
	# Bad
    COLOR=D00071
    elif (( $(echo "$COVERAGE > 70" | bc -l) )); then
	# Good
    COLOR=09D000
fi

curl "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > badge.svg
