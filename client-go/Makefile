SHELL=/bin/bash

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

.PHONY: client_csr
client_csr:
	@openssl req -newkey rsa:2048 -nodes -keyout ./config/client-key.pem -out ./config/client-req.pem \
			 -subj "/C=JP/ST=AI/L=NIS/O=DEV, Inc./CN=localhost"


.PHONY: sign_csr
sign_csr:
	@openssl x509 -req -in ./config/client-req.pem -CA ../config/ca-cert.pem -CAkey ../config/ca-key.pem \
				  -extfile <(printf "subjectAltName=DNS:localhost") -CAcreateserial -out ./config/client-cert.pem

.PHONY: csr
csr:  client_csr sign_csr
