SHELL=/bin/bash

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
				 -keyout ./config/ca-key.pem -out ./config/ca-cert.pem \
				 -subj "/C=JP/ST=Aichi/L=Anjo/O=DEV/OU=TUTORIAL/CN=self"

.PHONY: server_csr
server_csr:
	@openssl req -newkey rsa:2048 -nodes -keyout ./config/server-key.pem -out ./config/server-req.pem \
			 -subj "/C=JP/ST=AI/L=NIS/O=DEV, Inc./CN=localhost"


.PHONY: sign_csr
sign_csr:
	@openssl x509 -req -in ./config/server-req.pem -CA ./config/ca-cert.pem -CAkey ./config/ca-key.pem \
				  -extfile <(printf "subjectAltName=DNS:localhost") -CAcreateserial -out ./config/server-cert.pem

.PHONY: csr
csr: self_cert server_csr sign_csr