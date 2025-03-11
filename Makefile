run: build
	@./bin/wits

install:
	go install golang.org/x/tools/cmd/godoc@latest
	go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

build:
	go build -v -o ./bin/wits ./cmd/tui/main.go

clean:
	rm -f ./bin/wits
	rm -f coverage.html
	rm -f coverage.out
	rm -rf log
	rm -rf tmp
	rm -rf vendor

doc:
	godoc

changelog:
	git-chglog -o CHANGELOG.md

test:
	go test -race -v ./... -coverprofile coverage.out

test-ci:
	go test -race -v ./... -coverprofile coverage.out -covermode=atomic
	bash -c "bash <(curl -s https://codecov.io/bash)"

cover: test
	go tool cover -html coverage.out -o coverage.html

show-cover: cover
	open coverage.html

vet:
	go vet ./...
