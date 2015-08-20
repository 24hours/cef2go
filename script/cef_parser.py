import sqlite3
import sys, re, os.path
import exporter, settings, cef_enum, utils

conn = sqlite3.connect(settings.db_name)
c = conn.cursor()

# List of special handling

GoToC = {
  'cef_string_t' : {'name' : 'String' ,
                    'converter' : None
                    } 
}

# CefStruct 
# Base object for all struct, do not serve any purpose. 

class CefStruct(object):
  def __init__(self, hash_id):
    self.hash = hash_id
    c.execute('SELECT * FROM node WHERE hash = ?', (hash_id,))
    row = c.fetchone()
    self.name = row[3]
    self.type = row[4]
    self.instance = 'cefstruct'

  def __hash__(self):
    return hash(self.name)

  def __eq__(self, other):
    if isinstance(other, basestring):
      return self.name == other
    else:
      return self.name == other.name

  @property
  def CgoName(self):
    if self.name[0] == '_':
      return 'struct_' + ''.join(self.name.split(' '))
    else:
      return self.name

  def DumpC(all=False):
    raise Exception('Dump C is Not Implemented')

  def DumpGo(all=False):
    raise Exception('Dump Go is Not Implemented')

# CefStructHandler 
# Represent every struct in Cef
class CefStructHandler(CefStruct):
  def __init__(self, hash_id):
    c.execute('SELECT DISTINCT * FROM node WHERE hash = ?', (hash_id,))
    record = c.fetchone()
    if record[2] == 'CursorKind.ENUM_DECL':
      raise Exception('CefStructHandler failed to handle Enumeration')

    super(CefStructHandler, self).__init__(hash_id)
    
  def __repr__(self):
    return 'CefStructHandler <%s>' % ( self.name )

  @property
  def GoName(self):
    if self.name in GoToC:
      return GoToC[self.name]['name']
    else:
      return utils.camelCase(utils.getStructName(self.name[1:]), capital=True)

  @property
  def Functions(self):
    c.execute('SELECT DISTINCT * FROM node WHERE kind != ? AND spelling != ? AND parent_hash = ?', ('CursorKind.UNEXPOSED_ATTR', 'base', self.hash,))
    return [ CefFunction(api[0]) for api in c.fetchall() ]

  @property
  def GoGeneric(self):
    return "interface{}"

  @property
  def GoGenericVal(self):
    return "nil"

  @property
  def CGoGeneric(self):
    return "unsafe.pointer"

  @property
  def Purpose(self):
    handler_provider = ['_cef_app_t', '_cef_client_t']
    if self in handler_provider:
      return cef_enum.Struct.HANDLER_PROVIDER
    elif 'handler' in self.name:
      return cef_enum.Struct.HANDLER
    else:
      return cef_enum.Struct.CEF_TYPE

  def CToGoConversion(self, target, ret):
    if self.Purpose is cef_enum.Struct.HANDLER_PROVIDER or self.Purpose is cef_enum.Struct.HANDLER: 
      return """ 
      {ret}, ok := {self.GoName}Map[unsafe.Pointer({target})]; 
      if ok == false {{
        panic("{target} is not found in Map")
      }}
      """.format(**locals())

# CefParam
# Represent every parameter in the function
class CefParam(CefStruct):
  def __init__(self, hash_id):
    self.hash = hash_id
    c.execute('SELECT DISTINCT * FROM node WHERE hash = ?', (hash_id,))  
    param = c.fetchone()
    self.string = param[4]
    self.name = param[3]
    c.execute('SELECT DISTINCT * FROM node WHERE parent_hash = ?', (hash_id,))
    param_type = c.fetchone()
    
    if param_type is not None:
      self.type_id = param_type[10]
      self.type = param_type[3]
      self.const = True if re.search('const', param[4]) is not None else False
      self.pointer = param[4].replace(param_type[3],'').strip().split(' ')[-1]
      self.native = False
    else:
      # print "Native type found in %d" % self.hash
      # I don't exactly know what happen here. 
      self.type_id = hash_id
      self.type = param[4]
      self.const = True if re.search('const', param[4]) is not None else False
      self.pointer = '*' if '*' in param[4] else ''
      self.native = True # native ctype

  def __repr__(self):
    return "CefParam <<%r> %s %s>" % (self.TypeStruct, self.name, self.pointer)

  @property
  def isPointer(self):
    return False if self.pointer is '' else True

  @property
  def isConstant(self):
    return self.const

  @property
  def TypeStruct(self):
    if not self.native:
      return utils.ConstructStruct(self.type_id)
    else:
      return CType(self.type)
      # raise Exception("Native type (%s)parsing not implemented" % self.type )

# CefFunction
# Represent every methods or function pointer in a struct 

class CefFunction(object):
  def __init__(self, function_id):
    self.id = function_id
    c.execute('SELECT DISTINCT * FROM node WHERE kind != ? and spelling != ? and hash = ?', ('CursorKind.UNEXPOSED_ATTR', 'base', function_id,))
    api = c.fetchone()
    self.name = api[3]
    if api[5] == 'Pointer':
      type_spelling = api[4].split('(*)')
      c.execute('SELECT DISTINCT * FROM node WHERE kind = ? AND parent_hash = ?', ('CursorKind.TYPE_REF', api[0],))
      function_type = c.fetchone()
      if function_type is not None:
        self.type = function_type[3]
        self.pointer = type_spelling[0].replace(function_type[3], '').strip() 
      else:
        self.type = type_spelling[0].strip()
        self.pointer = None
    else:
      self.type = api[4].strip()
      self.pointer = None
      # raise Exception(api[0], api[3], 'is not function pointer')
  
  def __repr__(self):
    return 'CefFunction < %r %s>' % (self.TypeStruct, self.name)

  @property
  def TypeStruct(self):
    # naive way to detect if type is native
    if 'struct' in self.type or 'cef' in self.type:
      struct_name = utils.getStructName(self.type, type=None)
      struct_id = utils.getIdbyName(struct_name, settings.db_name)
      struct = utils.ConstructStruct(struct_id)
      return struct
    else:
      return CType(self.type)

  @property
  def Purpose(self):
    if re.match('get_.*_handler',self.name) is not None:
      return cef_enum.Function.HANDLER_GETTER
    elif re.match('on_.*', self.name) is not None:
      return cef_enum.Function.CALLBACK_EVENT
    else:
      return cef_enum.Function.STRUCT_METHOD

  @property 
  def Params(self):
    c.execute('SELECT DISTINCT * FROM node WHERE kind = ? AND parent_hash = ?', ('CursorKind.PARM_DECL', self.id,))
    return [ CefParam(row[0]) for row in c.fetchall() ] 

  def DumpGo(self, all=False):
    raise Exception('Dump Go is Not Implemented')

  # this dont belong here
  @property
  def GoName(self):
    return utils.camelCase(self.name, capital=True)

# CType
# Represent C native type 
class CType(object):
  def __init__(self, name):
    known_type = ['int *', 'int', 'void']
    if name not in known_type:
      raise Exception("Unknown type (%s)" % name)
    self.name = name.strip()
    self.instance = 'ctype'
    if '*' in name:
      self.pointer = '*'
    else:
      self.pointer = ''

  def __hash__(self):
    return hash(self.name)

  def __eq__(self, other):
    if isinstance(other, basestring):
      return self.name == other
    else:
      return self.name == other.name

  def __repr__(self):
    return 'CType <%s>' % (self.name)

  @property
  def GoName(self):
    if self.name == 'void':
      return ''
    elif self.name == 'int':
      return 'int'

  @property
  def CReturn(self):
    if self.name == 'int':
      return 'return 1;'
    return ''

  def CToGoConversion(self, target, ret):
    if self.name == 'int *':
      return "{ret} := int(*{target})\n".format(**locals())

    return "//%r" % self


  @property
  def GoGeneric(self):
    if self.name == 'int *':
      return "*int"
    else:
      return self.name

  @property
  def CgoName(self):
    if self.name == 'int *':
      return 'int'
    else:
      return self.name

  # @property 
  # def GoGenericVal(self):
  #   return 

# This thing seems to be useless
class CefType(CType, CefStruct):
  def __init__(self, hash_id):
    super(CefType, self).__init__(hash_id)
    c.execute("SELECT DISTINCT * FROM node WHERE hash=?", (hash_id,))
    original  = c.fetchone()
    self.hash = original[0]
    c.execute("SELECT DISTINCT * FROM node WHERE kind != 'CursorKind.UNEXPOSED_ATTR' AND spelling != 'base' AND parent_hash=?", (original[0],))
    self.function_pointer = []
    for api in c.fetchall():
      self.function_pointer.append(CefFunction(api[0]))
    self.hash = original[0]

# CefEnum
# Represent Cef Enumeration

class CefEnum(CType, CefStruct):
  def __init__(self, hash_id):
    c.execute("SELECT DISTINCT * FROM node WHERE hash=?", (hash_id,))
    original  = c.fetchone()
    self.hash = original[0]
    self.name = original[4]
    c.execute("SELECT DISTINCT * FROM node WHERE parent_hash=?", (original[0],))
    self.enum_value = {}
    for api in c.fetchall():
      self.enum_value[api[3]] = api[6]

  def __repr__(self):
    return 'CefEnum <%s>' % self.name

  @property
  def GoName(self):
    return utils.camelCase(self.name, capital=True)
  
  @property
  def CgoName(self):
    return self.name

