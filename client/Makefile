.PHONY: proto
proto:
	 @python -m grpc_tools.protoc \
			 -I .. \
			 --python_out=. \
			 --pyi_out=.\
			 --grpc_python_out=. \
			 ../proto/v1/product_info.proto \
			 ../proto/v1/order.proto