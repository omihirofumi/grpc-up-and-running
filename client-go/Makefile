.PHONY: proto
proto:
	@protoc -I .. \
 			--go_out=./order \
 			--go-grpc_out=./order \
 			--go_opt=paths=source_relative \
 			--go-grpc_opt=paths=source_relative \
 			../proto/v1/order.proto

.PHONY: server
server: order

.PHONY: order
order:
	@go build -o ./bin/order ./order

.PHONY: start
start: order
	./bin/order
