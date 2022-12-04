proto:
	protoc --go_out=./.. --go-grpc_out=./.. \
	--go_opt=M./api/proto/auth/auth.proto=./internal/pb/auth \
	--go_opt=M./api/proto/users/users.proto=./internal/pb/users \
	--go_opt=M./api/proto/samples/samples.proto=./internal/pb/samples \
	--go_opt=M./api/proto/default.proto=./internal/pb \
./api/proto/auth/auth.proto ./api/proto/users/users.proto ./api/proto/samples/samples.proto ./api/proto/default.proto
