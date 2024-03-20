alias tw := test-watch
alias tb := bench-watch

run:
	go run main.go

cover:
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"
	rm c.out

test:
	go test ./...

bench:
	go test -bench=. ./...

bench-watch:
	watchexec -e go just bench

test-watch:
	watchexec -e go just test
