import sqlite3, re
import settings, cef_enum, cef_parser

conn = sqlite3.connect(settings.db_name)
c = conn.cursor()

def recursive_search(struct, known_struct = set()):
  if struct in known_struct:
    return known_struct 

  if isinstance(struct, CType):
    known_struct.add(struct)
    return known_struct 

  for f in struct.Functions:
    known_struct.add(struct)
    known_struct = recursive_search(f.TypeStruct, known_struct)
    for p in f.Params:
      known_struct.add(struct)
      known_struct = recursive_search(p.TypeStruct, known_struct)

  return known_struct

def constructParamString(param_detail):
  param_list =[] 
  for param in param_detail:
    param_list.append(param.string + '  ' + param.name)

  return ', '.join(param_list)

def constructGoParamString(params, handled_type=None):
  param_list = [] 
  for param in params:
    if param.TypeStruct in handled_type:
      param_list.append(param.name + ' ' + param.TypeStruct.GoName )
    else:
      param_list.append(param.name + ' ' + param.TypeStruct.GoGeneric )
      
  return ', '.join(param_list)

def getStructHeader(hash_id, db_name):
  conn = sqlite3.connect(db_name)
  c = conn.cursor()
  c.execute("select * from node where hash=?", (hash_id,))
  row = c.fetchone()
  ret = {}
  ret['id'] = row[0]
  ret['name'] = row[3]
  ret['type'] = row[4]
  c.close()
  return ret    

def getStructName(struct_name, type=cef_enum.Struct.HANDLER_PROVIDER):
  if type is cef_enum.Struct.HANDLER_PROVIDER:
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

def ConstructStruct(hash_id):
  c.execute('SELECT DISTINCT * FROM node WHERE hash = ?', (hash_id,))
  record = c.fetchone()
  if record[2] == 'CursorKind.ENUM_DECL':
    return cef_parser.CefEnum(hash_id)
  else:
    return cef_parser.CefStructHandler(hash_id)