from enum import Enum

class Struct(Enum):
  HANDLER_PROVIDER = 0
  HANDLER = 1
  CEF_TYPE = 2

class Function(Enum):
  CALLBACK_EVENT = 0
  HANDLER_GETTER = 1
  STRUCT_METHOD = 2
