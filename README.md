# grpcをちゃんと学ぶ
* client: python
* server: go

## メモ
### docker
```terminal
$ docker build -t productinfo-server ${dockerfile dir}
$ docker network create my-net
$ docker run -d -p 50051:50051 --network=my-net --hostname=productinfo --name productinfo-server productinfo 
```

### kubernetes
```terminal
$ kubectrl apply -f server/productinfo-server.yaml
$ kubectrl delete -f server/productinfo-server.yaml
```