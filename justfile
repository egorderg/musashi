alias tw := test-watch

run:
	go run main.go

cover:
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"
	rm c.out

test:
	go test ./...

test-watch:
	watchexec -e go just test
