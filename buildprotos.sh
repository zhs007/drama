# export PATH="$PATH:/usr/local/go/bin"
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --proto_path=protos/ --go_out=./gamepb --go_opt=paths=source_relative protos/*.proto
protoc --proto_path=protos/ --go-grpc_out=./gamepb --go-grpc_opt=paths=source_relative protos/*.proto
