.PHONY: test fuzz lint

lint:
	golangci-lint run -c .golangci.yml ./...

test:
	go test -coverprofile=profile.out -coverpkg=github.com/abdullahiqbal1996/goldmark,github.com/abdullahiqbal1996/goldmark/ast,github.com/abdullahiqbal1996/goldmark/extension,github.com/abdullahiqbal1996/goldmark/extension/ast,github.com/abdullahiqbal1996/goldmark/parser,github.com/abdullahiqbal1996/goldmark/renderer,github.com/abdullahiqbal1996/goldmark/renderer/html,github.com/abdullahiqbal1996/goldmark/text,github.com/abdullahiqbal1996/goldmark/util ./...

cov: test
	go tool cover -html=profile.out

fuzz:
	cd ./fuzz && go test -fuzz=Fuzz
