from enum import Enum
from clang.cindex import Index, Config, CursorKind
import sqlite3, re
Config.set_library_path('/Library/Developer/CommandLineTools/usr/lib')

class Struct(Enum):
  HANDLER_PROVIDER = 0
  HANDLER = 1
  CEF_TYPE = 2

class Function(Enum):
  CALLBACK_EVENT = 0
  HANDLER_GETTER = 1
  STRUCT_METHOD = 2


def node_info(node, parent):
  if node is None:
    return (0, 0, "", "","","",0)
  else:
    if node.location.file is not None:
      filename = node.location.file.name 
      filename = filename.split('/')
      filename = filename[-1]
    else:
      filename = ""
    return (node.hash, parent.hash if parent is not None else 0, 
            str(node.kind), node.spelling, 
            node.type.spelling, 
            node.type.kind.spelling, 
            node.enum_value if node.kind is CursorKind.ENUM_CONSTANT_DECL else None,
            node.is_definition(),
            filename, node.location.line, node.type.get_declaration().hash)

def dump_node(node, parent):
  '''
  node : {clang.Cursor}
    a node in AST 
  depth : {int}
    depth of the node to be dumped
  call_back : {list}
    a list of function with the format function(node), return true and the node will be saved. 
  '''

  children = [dump_node(c, node) for c in node.get_children()]
  #if ( node.kind is CursorKind.TYPEDEF_DECL and node.is_definition() is True ) or node.kind is CursorKind.FUNCTION_DECL:
  #  print node.kind ,','  ,node.spelling,',', node.location.file.name, ',' , node.location.line
  

  return { 'parent' : parent.hash if parent is not None else 0,
           'hash' : node.hash,
           'kind_name' : type(node),
           'kind' : node.kind,
           'type' : node.type.get_declaration().hash,
           'type_kind' : node.type.kind.spelling, 
           'type_spelling' : node.type.spelling,
           'translation_spelling' : node.translation_unit.spelling,
           'spelling' : node.spelling,
           'location' : node.location,
           'is_definition' : node.is_definition(),
           'enum' : type(node.type) ,
           'enum_value' : node.enum_value if node.kind is CursorKind.ENUM_CONSTANT_DECL else None,
           'referenced' : node.referenced.hash if node.referenced is not None else 0,
           'children' : children }

def getStructName(struct_name, type=Struct.HANDLER_PROVIDER):
  if type is Struct.HANDLER_PROVIDER:
    name = struct_name.split("_")
    return '_'.join(name[1:-1])
  else:
    return re.search('cef[A-Za-z0-9\_]*', struct_name).group(0)

def camelCase(struct_name, capital=False):
  name = struct_name.split('_')
  ret_list = []
  for n in name:
    ret_list.append(n[0].upper() + n[1:])

  if capital is True:
    return ''.join(ret_list)
  else:
    ret = ''.join(ret_list)
    return ret[0].lower() + ret[1:]


def getIdbyName(name, db_name):
  conn = sqlite3.connect(db_name)
  c = conn.cursor()
  c.execute('SELECT DISTINCT * FROM node WHERE spelling = ? AND kind = ? ', (name,"CursorKind.TYPEDEF_DECL",))
  reference = c.fetchone()

  c.execute('SELECT DISTINCT * FROM node WHERE parent_hash = ? ', (reference[0],))
  hash_id = c.fetchone()
  c.close()
  if  hash_id is None:
    return reference[0]
  else:
    return hash_id[0]


