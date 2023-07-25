.PHONY: proto
proto:
	@protoc --go_out=./server \
 			--go-grpc_out=./server \
 			--go_opt=paths=source_relative \
 			--go-grpc_opt=paths=source_relative \
 			./proto/v1/product_info.proto