# DemoGolangGrpc

[Protocol Buffers](https://protobuf.dev/)

[Package]

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

[Go Tool Create SSL Self-Signed Certificate]

```
go env

go run "c:\Program Files\go\src\crypto\tls\generate_cert.go" -h

go run "c:\Program Files\go\src\crypto\tls\generate_cert.go" -host localhost
```

[Go Generated Code Guide](https://protobuf.dev/reference/go/go-generated/)

```
protoc --go_out=:. --go_opt=paths=source_relative <path>/<name>.proto
protoc --go-grpc_out=:. --go-grpc_opt=paths=source_relative <path>/<name>.proto
protoc --go_out=:. --go_opt=paths=source_relative --go-grpc_out=:. --go-grpc_opt=paths=source_relative <path>/<name>.proto

protoc --go_out=:. --go-grpc_out=:. <path>/<namp>.proto
protoc --go_out=:. --go-grpc_out=:. <path>/*.proto
protoc --go_out=:. --go-grpc_out=require_unimplemented_servers=false=:. <path>/*.proto

protoc --go_out=:. --go-grpc_out=require_unimplemented_servers=false:. pb/*.proto
```
