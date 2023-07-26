from google.protobuf import wrappers_pb2
import grpc
from proto.v1 import order_pb2_grpc
from proto.v1 import order_pb2


def run():
    channel = grpc.insecure_channel('localhost:50052')

    stub = order_pb2_grpc.OrderManagementStub(channel)

    order1 = order_pb2.Order(items=['Item - A', 'Item - B', 'Item - C'],
                                        price=2450.50,
                                        destination='San Jose, CA')

    order = stub.getOrder(order_pb2.Order(id='102'))
    print("Order service response", order)

    # Unary RPC : Adding an Order
    response = stub.addOrder(order1)
    print('Add order response :', response)

    # Server Streaming
    for order_search_result in stub.searchOrders(wrappers_pb2.StringValue(value='Item - A')):
        print('Search Result : ', order_search_result)

    # Client Streaming
    upd_order_iterator = generate_orders_for_updates()
    upd_status = stub.updateOrders(upd_order_iterator)
    print('Order update status : ', upd_status)

    # Bi-di Streaming
    proc_order_iterator = generate_orders_for_processing()
    for shipment in stub.processOrders(proc_order_iterator):
        print(shipment)


def generate_orders_for_updates():
    ord1 = order_pb2.Order(id='101', price=1000,
                                      items=['Item - A', 'Item - B', 'Item - C', 'Item - D'],
                                      destination='Mountain View, CA')
    ord2 = order_pb2.Order(id='102', price=1000,
                                      items=['Item - E', 'Item - Q', 'Item - R', 'Item - D'],
                                      destination='San Jose, CA')
    ord3 = order_pb2.Order(id='103', price=1000,
                                      items=['Item - A', 'Item - K'],
                                      destination='San Francisco, CA')
    list = []
    list.append(ord1)
    list.append(ord2)
    list.append(ord3)

    for updated_orders in list:
        yield updated_orders


def generate_orders_for_processing():
    ord1 = order_pb2.Order(
        id='104', price=2332,
        items=['Item - A', 'Item - B'],
        destination='San Jose, CA')
    ord2 = order_pb2.Order(
        id='105', price=3000,
        destination='San Francisco, CA')
    ord3 = order_pb2.Order(
        id='106', price=2560,
        destination='San Francisco, CA')
    ord4 = order_pb2.Order(
        id='107', price=2560,
        destination='Mountain View, CA')
    list = []
    list.append(ord1)
    list.append(ord1)
    list.append(ord3)
    list.append(ord4)

    for processing_orders in list:
        yield processing_orders


run()
