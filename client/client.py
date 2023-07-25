import grpc

from v1 import product_info_pb2_grpc
from v1 import product_info_pb2


def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = product_info_pb2_grpc.ProductInfoStub(channel)
        response = stub.addProduct(product_info_pb2.Product(name='Soccer Ball', description='The soccer ball whose messi.', price=10000))
        print('add product: response', response)
        productInfo = stub.getProduct(product_info_pb2.ProductID(value=response.value))
        print('get product: response', productInfo)

run()