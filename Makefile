run:
    go run ./app/app1/cmd/main.go

build:
    GOOS=linux GOARCH=amd64 go build -o trolley-linux-amd64 ./app/admin/cmd/main.go