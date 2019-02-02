gcgt: cmd/gcgt/main.go
	go build ./cmd/gcgt
	go install ./cmd/gcgt

test: gcgt
	go generate ./tests