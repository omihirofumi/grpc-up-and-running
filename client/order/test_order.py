import grpc
from proto.v1 import order_pb2_grpc
from proto.v1 import order_pb2


def run():
    channel = grpc.insecure_channel('localhost:50052')

    stub = order_pb2_grpc.OrderManagementStub(channel)

    order = order_pb2.Order(items=['A', 'B', 'C'],
                            price=1000,
                            destination='Japan',
                            id="-1")

    res = stub.addOrder(order)
    print(res)

run()