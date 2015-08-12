import sqlite3, sys
from cef_util import *
from cef_parser import init_prefix, golang_prefix
import inspect

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

def DumpGoEnum(struct, file_d):
  # type CefValueTypeT int
  # const (
  #    VTYPE_DICTIONARY    CefValueTypeT = 7
  #    VTYPE_STRING        CefValueTypeT = 5
  #    VTYPE_BOOL          CefValueTypeT = 2
  #    VTYPE_BINARY        CefValueTypeT = 6
  #    VTYPE_INT           CefValueTypeT = 3
  #    VTYPE_DOUBLE        CefValueTypeT = 4
  #    VTYPE_NULL          CefValueTypeT = 1
  #    VTYPE_INVALID       CefValueTypeT = 0
  #    VTYPE_LIST          CefValueTypeT = 8
  # )
  
  file_d.write("\ntype %s int\n" % (struct.GoName))
  file_d.write("const (\n")
  for v, k in enumerate(struct.enum_value):
    file_d.write("\t %s \t %s = %s\n" % (k, struct.GoName, struct.enum_value[k]))
  file_d.write(")\n")
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])

def DumpGoInterface(struct, file_d, handled_type=None):
  # type App interface{
  #   GetAppT () AppT 
  #   SetAppT (AppT)
  # 
  #   OnBeforeCommandLineProcessing (process_type String, command_line CommandLine)  
  #   OnRegisterCustomSchemes (registrar SchemeRegistrar)  
  #   GetResourceBundleHandler () ResourceBundleHandler 
  #   GetBrowserProcessHandler () BrowserProcessHandler 
  #   GetRenderProcessHandler () RenderProcessHandler 
  # }

  file_d.write("\ntype %s interface{\n" % (struct.GoName))
  file_d.write('\t Get%sT() %sT \n' % (struct.GoName, struct.GoName ))
  file_d.write('\t Set%sT(%sT) \n' % (struct.GoName, struct.GoName ))
  
  for api in struct.Functions:
    handling = api.TypeStruct in handled_type or api.TypeStruct.instance == 'ctype'
    file_d.write('\t %s(%s) %s \n' % 
        ( api.GoName, 
          constructGoParamString([ param for param in api.Params if param.name not in ['self'] ], handled_type=handled_type), 
          api.TypeStruct.GoName if handling else api.TypeStruct.GoGeneric))
  file_d.write("}\n")
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])

def DumpCgoInitializer(struct, file_d, handled_type=None):
  # var AppMap = make(map[unsafe.Pointer]App)

  # type AppT struct{
  #   CStruct   *C.cef_app_t
  # }

  # func NewAppT(handler App) AppT {
  #   var ret AppT
  #   ret.CStruct = (*C.cef_app_t)(C.calloc(1, C.sizeof_cef_app_t))
  #   C.initializecef_app(ret.CStruct))
  #   handler.SetAppT(ret)
  #   AppMap[unsafe.Pointer(a.CStruct)] = handler
  #   return ret
  # }

  cstruct_name = "CStruct"

  file_d.write("var %sMap = make(map[unsafe.Pointer]%s)\n" % 
        ( struct.GoName,
          struct.GoName))

  file_d.write("\ntype %sT struct{\n" % struct.GoName)
  file_d.write("\t%s\t\t*C.%s\n" % (cstruct_name, getStructName(struct.name, type=None)))
  file_d.write("}\n")

  file_d.write("\nfunc New%sT(handler %s) %sT {\n" % (struct.GoName, struct.GoName, struct.GoName))
  file_d.write("\tvar ret %sT\n" % struct.GoName)
  file_d.write("\tret.%s = (*C.%s)(C.calloc(1, C.sizeof_%s))\n" % 
      ( cstruct_name, 
        getStructName(struct.name, type=None), 
        getStructName(struct.name, type=None)))
  file_d.write("\tC.%s(ret.%s)\n" % (init_prefix+getStructName(struct.name), cstruct_name))
  file_d.write("\thandler.Set%sT(ret)\n" % (struct.GoName))
  file_d.write("\t%sMap[unsafe.Pointer(ret.%s)] = handler\n" % (struct.GoName, cstruct_name))
  file_d.write("\treturn ret\n")
  file_d.write("}\n")
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])


def DumpCInitializer(struct, file_d, handled_type=None):
  # void initializecef_app ( cef_app_t * self){
  #   self->base.size = sizeof(cef_app_t);
  #   initialize_cef_base((cef_base_t*) self);
  #
  #   self->on_before_command_line_processing = on_before_command_line_processing;
  #   self->on_register_custom_schemes = on_register_custom_schemes;
  #   self->get_resource_bundle_handler = get_resource_bundle_handler;
  #   self->get_browser_process_handler = get_browser_process_handler;
  #   self->get_render_process_handler = get_render_process_handler;
  # }

  file_d.write('void %s ( %s * self){\n' % ( init_prefix+getStructName(struct.name), getStructName(struct.name, type=None) ))
  file_d.write('\tself->base.size = sizeof(%s);\n' % (getStructName(struct.name, type=None)) )
  file_d.write('\tinitialize_cef_base((cef_base_t*) self);\n') 
  for api in struct.Functions:
    file_d.write('\tself->%s = %s;\n' % (api.name , api.name))
  file_d.write('}\n')
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])

def DumpCFunctions(func_struct, file_d, handled_type=None):
  if func_struct.Purpose is Function.STRUCT_METHOD:
    DumpFunctionAsAPI(func_struct, file_d, handled_type=handled_type)
  else:
    DumpFunctionAsHandler(func_struct, file_d, handled_type=handled_type)


def DumpFunctionAsAPI(func_struct, file_d, handled_type=None):
  # void CEF_CALLBACK on_before_command_line_processing (
  #    struct _cef_app_t          *  self, 
  #    const cef_string_t         *  process_type, 
  #    struct _cef_command_line_t *  command_line){
  # }

  file_d.write('%s CEF_CALLBACK %s %s (%s){\n' % 
      ( func_struct.TypeStruct.name,
        '' if func_struct.pointer is False else '*', 
        func_struct.name, 
        constructParamString(func_struct.Params) ))
  # TODO : include this line back 
  #file_d.write('\treturn %s(%s)\n' % (golang_prefix + self.name , ','.join([ i.name for i in self.param_detail ])))
  # remove this line 
  if func_struct.TypeStruct.CReturn != '':
    file_d.write('\t%s\n' % func_struct.TypeStruct.CReturn) 
  file_d.write('}\n')
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])


def DumpFunctionAsHandler(func_struct, file_d, handled_type=None):
  # struct _cef_render_process_handler_t * CEF_CALLBACK get_render_process_handler(struct _cef_app_t *  self){
  #   return go_GetRenderProcessHandler(self);
  # }

  file_d.write("%s %s CEF_CALLBACK %s(%s){\n" % 
    ( func_struct.type, 
      func_struct.pointer or '', 
      func_struct.name, 
      constructParamString(func_struct.Params) ))
  file_d.write("")
  file_d.write("\treturn %s%s(%s);\n" % 
      ( golang_prefix, 
        func_struct.GoName, 
        # normally the param should really contain only self.
        # this method ensure I won't miss anything incase my assumption is wrong
        # however this methods is alittle hacky
        ','.join(p.name for p in func_struct.Params)))
  file_d.write("}\n")
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])

def DumpGoFunctions(struct, func_struct, file_d, handled_type=None):
  if func_struct.Purpose is Function.CALLBACK_EVENT: #or func_struct.Purpose is Function.STRUCT_METHOD:  
    DumpGoCallBackFunction(struct, func_struct, file_d, handled_type=handled_type)
  else:
    DumpGoGetterFunctions(struct, func_struct, file_d, handled_type=handled_type)

def DumpGoCallBackFunction(struct, func_struct, file_d, handled_type=None):
  # //export go_OnBeforeCommandLineProcessing
  # func go_OnBeforeCommandLineProcessing (self *C.cef_app_t){
  #   if handler, ok := knownApp[unsafe.Pointer(self)]; ok {
  #    return handler.OnBeforeCommandLineProcessing(parameters [TODO])
  #   }
  #   return nil
  # }

  file_d.write("//export %s%s\n" % (golang_prefix, func_struct.GoName))

  param_string = []
  for p in func_struct.Params:
    param_string.append(p.name + ' *C.' + p.TypeStruct.CgoName)

  file_d.write("func %s%s (%s) %s {\n" % 
      ( golang_prefix, func_struct.GoName, 
        ', '.join(param_string), 
        func_struct.TypeStruct.GoName))
  file_d.write("\tif handler, ok := %sMap[unsafe.Pointer(%s)]; ok {\n" % (struct.GoName, func_struct.Params[0].name))
  # TODO : preprocess before calling 
  for param in func_struct.Params[1:]:
    file_d.write("\t\t//processing %s %s \n" % (param.TypeStruct.name, param.name))
  # call method
  if func_struct.TypeStruct.name == 'void':
    file_d.write("\thandler.%s(%s)\n" % 
        ( func_struct.GoName, 
          ','.join([ param.name for param in func_struct.Params if param.name not in ['self'] ]) ))    
  else:
    file_d.write("\t return handler.%s(%s)\n" % 
        ( func_struct.GoName, 
          ','.join([ param.name for param in func_struct.Params if param.name not in ['self'] ]) ))

  file_d.write("\t}\n")
  file_d.write("\t%s\n" % func_struct.TypeStruct.CReturn)
  file_d.write("}\n")
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])

def DumpGoGetterFunctions(struct, func_struct, file_d, handled_type=None):

  # //export go_GetResourceBundleHandler
  # func go_GetResourceBundleHandler (self *C.cef_app_t) *C.struct___cef_app_t{
  #   if handler, ok := knownApp[unsafe.Pointer(self)]; ok {
  #     ret := handler.GetResourceBundleHandler()
  #     if ret != nil{
  #       return ret.GetResourceBundleHandler().CStruct
  #     }
  #   }
  #   return nil
  # }

  cstruct_name = "CStruct"

  file_d.write("//export %s%s\n" % (golang_prefix, func_struct.GoName))
  file_d.write("func %s%s (self *C.%s) *C.%s{\n" % 
      ( golang_prefix, 
        func_struct.GoName, 
        getStructName(struct.name, type=None), 
        func_struct.TypeStruct.CgoName))
  file_d.write("\tif handler, ok := %sMap[unsafe.Pointer(self)]; ok {\n" % (struct.GoName))
  file_d.write("\t\tret := handler.%s()\n" % (func_struct.GoName))
  file_d.write("\t\tif ret != nil{\n")
  if func_struct in handled_type:
    file_d.write("\t\t\treturn ret.%s().%s\n" % (func_struct.GoName, cstruct_name))
  else:
    file_d.write("\t\t\treturn %s\n" % (func_struct.TypeStruct.GoGenericVal))
  file_d.write("\t\t}\n")
  file_d.write("\t}\n")
  file_d.write("\treturn nil\n")
  file_d.write("}\n")
  file_d.write("//Generate by %s \n\n" % inspect.stack()[0][3])


def DumpExtern(struct, file_d, handled_type=None):
  file_d.write('void %s ( %s * self);\n' % 
      ( init_prefix+getStructName(struct.name), 
        getStructName(struct.name, type=None) ))







