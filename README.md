# MEMO-GO

MEMO-Go is for Memo application backend part using Golang

## Generate proto to shared/gen

### Auth
```sh
  protoc `
  --proto_path=shared/proto `
  --go_out=shared/gen `
  --go-grpc_out=shared/gen `
  shared/proto/auth/v1/auth.proto
```

### Pos
```sh
  protoc `
  --proto_path=shared/proto `
  --go_out=shared/gen `
  --go-grpc_out=shared/gen `
  shared/proto/pos/v1/pos.proto
```

### Test register

```sh
cd D:\Mamduh\Personal\MEMO-Go
grpcurl -plaintext -d "{ \"email\": \"new@example.com\", \"password\": \"password123\" }" localhost:50051 auth.v1.AuthService/Register
```