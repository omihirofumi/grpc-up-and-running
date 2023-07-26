from google.protobuf import wrappers_pb2 as _wrappers_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Order(_message.Message):
    __slots__ = ["id", "items", "price", "destination"]
    ID_FIELD_NUMBER: _ClassVar[int]
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    DESTINATION_FIELD_NUMBER: _ClassVar[int]
    id: str
    items: _containers.RepeatedScalarFieldContainer[str]
    price: float
    destination: str
    def __init__(self, id: _Optional[str] = ..., items: _Optional[_Iterable[str]] = ..., price: _Optional[float] = ..., destination: _Optional[str] = ...) -> None: ...

class CombinedShipment(_message.Message):
    __slots__ = ["id", "status", "ordersList"]
    ID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    ORDERSLIST_FIELD_NUMBER: _ClassVar[int]
    id: str
    status: str
    ordersList: _containers.RepeatedCompositeFieldContainer[Order]
    def __init__(self, id: _Optional[str] = ..., status: _Optional[str] = ..., ordersList: _Optional[_Iterable[_Union[Order, _Mapping]]] = ...) -> None: ...
