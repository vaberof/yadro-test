PROJECT_NAME=yadro-test
WORK_DIR_LINUX=./cmd/yadro-test

run:
	go run $(WORK_DIR_LINUX)/main.go

tests.run:
	go test -v ./internal/domain/ballsorter