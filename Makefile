.PHONY: ann anntool all
all: ann anntool

ann:
	go build  -ldflags "-X github.com/annchain/annchain/chain/version.commitVer=`git rev-parse HEAD`" -o ./build/ann ./chain
anntool:
	go build -ldflags "-X github.com/annchain/annchain/client/main.version=`git rev-parse HEAD`" -o ./build/anntool ./client
test:
	go test ./tools/state
proto:
	protoc --proto_path=$(GOPATH)/src --proto_path=chain/app/remote --go_out=plugins=grpc:chain/app/remote chain/app/remote/*.proto
	protoc --proto_path=$(GOPATH)/src --proto_path=example/types --go_out=plugins=grpc:example/types example/types/*.proto
	#protoc --proto_path=$(GOPATH)/src --proto_path=chain/node/protos --gofast_out=plugins=grpc:chain/node/protos chain/node/protos/*.proto
	protoc --proto_path=types --go_out=types types/*.proto
