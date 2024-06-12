```
go install github.com/google/wire/cmd/wire@latest
export PATH="~/go/bin:$PATH"
touch wire.go
wire
```

```
//go:build !wireinject
// +build !wireinject
toda vez que o build ocorrer e nao tiver haver com o wire, utilizar o wire_gen e ignorar o wire.go

//go:build wireinject
// +build wireinject
toda vez que o build ocorrer utilizar o arquivo wire.go
```
