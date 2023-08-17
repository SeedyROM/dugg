test:
    go test -json -v ./... | gotestfmt

coverage:
    courtney ./...
    go tool cover -func=coverage.out
    go tool cover -html=coverage.out -o coverage.html

build *FLAGS:
    go build -o bin/dugg -ldflags="-s -w" {{FLAGS}} ./cmd/dugg

build-web *FLAGS:
    cd web && npm run build {{FLAGS}}

run *FLAGS: build-web build
    go run ./cmd/dugg {{FLAGS}}

dev *FLAGS:
    cd web && npm run dev &
    air -- {{FLAGS}}