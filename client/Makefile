.PHONY: proto
proto: proto_product proto_order

.PHONY: proto_product
proto_product:
	 @python -m grpc_tools.protoc \
			 -I .. \
			 --python_out=./product \
			 --pyi_out=./product \
			 --grpc_python_out=./product \
			 ../proto/v1/product_info.proto \

.PHONY: proto_order
proto_order:
	 @python -m grpc_tools.protoc \
			 -I .. \
			 --python_out=./order \
			 --pyi_out=./order \
			 --grpc_python_out=./order \
			 ../proto/v1/order.proto
