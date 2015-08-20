from optparse import OptionParser, OptionGroup
from pprint import pprint
import utils, exporter, settings, cef_parser
from clang.cindex import Index, Config, CursorKind
Config.set_library_path('/Library/Developer/CommandLineTools/usr/lib')

if __name__ == '__main__':
  '''
  script use to generate c to golang binding
  it will contain alot of error, but will be modified as seen fit.
  '''
 
  parser = OptionParser("usage: %prog [options] {filename} [clang-args*]")
  parser.add_option("-g", "--go_file", dest="go_file",
                  help="write generated go code to FILE", metavar="FILE")
  parser.add_option("-c", "--c_file", dest="c_file",
                  help="write generated c code to FILE", metavar="FILE")
  parser.add_option("-d", "--h_file", dest="h_file",
                  help="write generated header code to FILE", metavar="FILE")
  parser.add_option("-p", "--h_path", dest="h_path",  
                  help="path for generated header code", metavar="FILE")
  parser.add_option("-r", "--reload", action="store_true", dest="refresh",
                  help="force refresh on the database")
  parser.add_option("-v", "--verbose", action="store_true", dest="verbose",
                  help="print the dumped node to console")
  
  parser.disable_interspersed_args()
  (opts, args) = parser.parse_args()

  exporter.prepare_file(file_c=opts.c_file, file_go=opts.go_file, file_h=opts.h_file, h_path=opts.h_path)

  if len(args) == 0:
      parser.error('invalid number arguments')

  if opts.refresh:
    index = Index.create()
    tu = index.parse(None, args)
    if not tu:
        parser.error("unable to load input")

    helper.reset_db()
    helper.save_node(tu.cursor, None)
    if opts.verbose:
      pprint(helper.dump_node(tu.cursor, None))
    
  struct_type = {}

  # starting from entry point : cef_app_t 
  # any struct referenced by "cef_app_t" will get parser.
  # # collect all public struct that should be parsed 
  # struct_list = recursive_search(app)
  # struct_list.union(recursive_search(client))
  
  # for p in struct_list:
  #   if isinstance(p, CefEnum):   
  #     # DumpGo(p)
  #     pass
  #   elif p == '_cef_app_t':
  #     DumpGo(p)

  handle_type = ['_cef_app_t', '_cef_client_t', '_cef_life_span_handler_t']

  for st_name in ['cef_app_t', 'cef_client_t', 'cef_life_span_handler_t']:
    st_id = utils.getIdbyName(st_name, settings.db_name)
    st = cef_parser.CefStructHandler(st_id)

    with open(opts.c_file, "a") as f:
      exporter.DumpC(st, file_d = f, handled_type = handle_type)
     
    with open(opts.go_file, "a") as f:
      exporter.DumpGoDefinition(st, file_d = f,  handled_type = handle_type)
      exporter.DumpGo(st, file_d = f,  handled_type = handle_type)

    with open(opts.h_path + opts.h_file, 'a') as f:
      exporter.DumpHeader(st, file_d = f,  handled_type = handle_type)
