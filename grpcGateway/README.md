## grpc gateway

[grpc-ecosystem](https://github.com/grpc-ecosystem)

```
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway 
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway 
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

[googleapis](https://github.com/googleapis/googleapis)

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


> DefaultHeaderMatcher is used to pass http request headers to/from gRPC context. This adds permanent HTTP header keys (as specified by the IANA) to gRPC context with grpcgateway- prefix. HTTP headers that start with 'Grpc-Metadata-' are mapped to gRPC metadata after removing prefix 'Grpc-Metadata-'.

[0](https://www.readfog.com/a/1665967072847433728)
[1](https://www.cnblogs.com/hacker-linner/p/14618862.html)
