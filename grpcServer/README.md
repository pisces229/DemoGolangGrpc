## Generated Code

```

protoc `
--go_out=:. `
--go_opt=paths=source_relative `
--go-grpc_out=require_unimplemented_servers=false:. `
--go-grpc_opt=paths=source_relative `
--grpc-gateway_out=logtostderr=true:. `
--grpc-gateway_opt=paths=source_relative `
runner.proto group/*.proto

```
