.PHONY: proto
proto: proto_product proto_order

.PHONY: proto_product
proto_product:
	@protoc -I .. \
 			--go_out=./product \
 			--go-grpc_out=./product \
 			--go_opt=paths=source_relative \
 			--go-grpc_opt=paths=source_relative \
 			../proto/v1/product_info.proto \

.PHONY: proto_order
proto_order:
	@protoc -I .. \
 			--go_out=./order \
 			--go-grpc_out=./order \
 			--go_opt=paths=source_relative \
 			--go-grpc_opt=paths=source_relative \
 			../proto/v1/order.proto

.PHONY: server
server: product order

.PHONY: product
product:
	@go build -o ./bin/product ./product

.PHONY: order
order:
	@go build -o ./bin/order ./order

.PHONY: start_product
start_product: product
	./bin/product

.PHONY: start_order
start_order: order
	./bin/proruct
