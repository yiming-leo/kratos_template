# wire_gen.go
当业务发生变化时，不仅仅需要在各自的包内、server下的http.go或grpc.go中注册，还需要在这里注册那些业务组件
<br>具体可以直接运行wire_gen.go最上面的命令即可
```shell
go:generate go run github.com/google/wire/cmd/wire
```