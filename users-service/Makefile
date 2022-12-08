proto:
	protoc --go_out=./.. --go-grpc_out=./.. \
	--go_opt=M./api/proto/users.proto=./internal/pb \
	--go_opt=M./api/proto/default.proto=./internal/pb \
./api/proto/users.proto ./api/proto/default.proto
