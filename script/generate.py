from clang.cindex import Index, Config, CursorKind
from pprint import pprint
Config.set_library_path("/Library/Developer/CommandLineTools/usr/lib")
from optparse import OptionParser, OptionGroup
global opts

def get_diag_info(diag):
    return { 'severity' : diag.severity,
             'location' : diag.location,
             'spelling' : diag.spelling,
             'ranges' : diag.ranges,
             'fixits' : diag.fixits }

def get_info(node, depth=0):
  children = [get_info(c, depth+1) for c in node.get_children()]
  if ( node.kind is CursorKind.TYPEDEF_DECL and node.is_definition() is True ) or node.kind is CursorKind.FUNCTION_DECL:
    print node.kind ,','  ,node.spelling,',', node.location.file.name, ',' , node.location.line
  # return { 'kind_name' : type(node),
  #          'kind' : node.kind,
  #          'type' : node.type,
  #          'spelling' : node.spelling,
  #          'location' : node.location,
  #          'is_definition' : node.is_definition(),
  #          'enum' : type(node.type) ,
  #          'children' : children }

# 
#
# [debugging]
#
#
#
#
#

if __name__ == '__main__':
  '''
  script use to generate c to golang binding
  it will contain alot of error, but will be modified as fit.
  '''
  parser = OptionParser("usage: %prog {filename} [clang-args*]")
  parser.disable_interspersed_args()
  (opts, args) = parser.parse_args()

  if len(args) == 0:
      parser.error('invalid number arguments')

  index = Index.create()
  tu = index.parse(None, args)
  if not tu:
      parser.error("unable to load input")
  
  pprint(('diags', map(get_diag_info, tu.diagnostics)))
  pprint(('nodes', get_info(tu.cursor)))

