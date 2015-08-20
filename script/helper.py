from clang.cindex import Index, Config, CursorKind
import sqlite3
Config.set_library_path('/Library/Developer/CommandLineTools/usr/lib')

conn = sqlite3.connect(settings.db_name)
c = conn.cursor()

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


def reset_db():
  c.execute('DROP TABLE IF EXISTS node');
  c.execute('''CREATE TABLE node
             (hash Integer, parent_hash Integer, kind text, spelling text, type_spelling text, type_kind_spelling text, enum_value text, definition text, file text, line_no Integer, kind_id Integer)''')
  conn.commit()

def save_node(node, parent):
  c.execute('SELECT * FROM node WHERE hash=?', (node.hash, ))
  
  node_detail = node_info(node, parent)
  c.execute('INSERT INTO node VALUES (?,?,?,?,?,?,?,?,?,?, ?)', node_detail)
  conn.commit()
  for child in node.get_children():
    save_node(child, node)
  return