## grpc web

[improbable-eng](https://github.com/improbable-eng)

```
go get github.com/improbable-eng/grpc-web

go install github.com/improbable-eng/grpc-web
```

[Generated Code]

```
protoc `
--go_out=:. `
--go_opt=paths=source_relative `
--go-grpc_out=require_unimplemented_servers=false:. `
--go-grpc_opt=paths=source_relative `
--grpc-gateway_out=logtostderr=false:. `
--grpc-gateway_opt=paths=source_relative `
runner.proto
```

[gRPC for Web Clients](https://github.com/grpc/grpc-web)

`npm i grpc-web`

```
protoc `
--js_out=import_style=commonjs:../frontend/src `
--plugin=protoc-gen-grpc=./protoc-gen-grpc-web.exe `
--grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:../frontend/src `
runner.proto
```
