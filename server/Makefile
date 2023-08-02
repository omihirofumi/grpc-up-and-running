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
	./bin/order

.PHONY: self_cert
self_cert:
	@openssl req -x509 -newkey rsa:2048 -nodes -days 3650 \
				 -keyout ca-key.pem -out ca-cert.pem \
				 -subj "/C=JP/ST=Aichi/L=Anjo/O=DEV/OU=TUTORIAL/CN=localhost"

.PHONY: server_csr
server_csr:
	@openssl req -newkey rsa:2048 -nodes -keyout server-key.pem -out server-req.pem \
				 -config ../config/san.cnf

.PHONY: sign_csr
sign_csr:
	@openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem \
				  -CAcreateserial -out server-cert.pem -extfile ../config/san.cnf

.PHONY: csr
csr: self_cert server_csr sign_csr