#### 1. 安装grpc-gateway相关代码包
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

go get -u github.com/golang/protobuf/protoc-gen-go

会生成相应的3个二进制文件
> protoc-gen-grpc-gateway （生成go的反向代理脚本）

> protoc-gen-swagger

> protoc-gen-go （生成相关语言的pb实现代码）

#### 2. 环境准备，需要安装protoc
brew install protobuf@3.11
#### 3. 编写pb文件，并添加 google.api.http 批注到pb文件中
example/service.proto
#### 4. 使用`protoc-gen-go` 生成pb go语言的存根文件
protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --go_out=plugins=grpc:. example/service.proto
#### 4. 实现grpc中的rpc服务（rpc服务端要实现的）
server/main.go
```
type server struct {}

func (s *server) Post(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{Message: "hello post "+in.Name}, nil
}
func (s *server) Get(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	return &pb.TestResponse{Message: "hello get "+in.Name}, nil
}

```
#### 5. 使用`protoc-gen-grpc-gateway` 生成反向代理gateway go 文件
protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --grpc-gateway_out=logtostderr=true:. example/service.proto
#### 6. 使用`protoc-gen-swagger`生成swagger定义（json文件 可选）
protoc -I/usr/local/include -I.   -I$GOPATH/src   -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis   --swagger_out=logtostderr=true:. example/service.proto
#### 7. 编写http反向代理服务脚本
main.go

#### 8. 执行rpc server 文件
go run server/main.go

#### 9. 启动rpc gateway 反向代理
go run main.go

#### 10. 验证http接口
Post 接口：curl -X POST http://localhost:8081/example/post

Get 接口：curl http://localhost:8081/example/get

