test:
    go test -json -v ./... | gotestfmt

coverage:
    courtney ./...
    go tool cover -func=coverage.out
    go tool cover -html=coverage.out -o coverage.html

build *FLAGS:
    go build -o bin/dugg -ldflags="-s -w" {{FLAGS}} ./cmd/dugg

run *FLAGS:
    npm run build
    go run ./cmd/dugg {{FLAGS}}

dev *FLAGS:
    npm run dev &
    air