// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go


#include "_cgo_export.h"

#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_browser_process_handler_capi.h"


void CEF_CALLBACK cef_browser_process_handler_t_on_context_initialized(
      struct _cef_browser_process_handler_t* self) {
    OnContextInitialized(self);
}

void CEF_CALLBACK cef_browser_process_handler_t_on_before_child_process_launch(
  struct _cef_browser_process_handler_t* self,
  struct _cef_command_line_t* command_line) {
    OnBeforeChildProcessLaunch(self, command_line);
}

void CEF_CALLBACK cef_browser_process_handler_t_on_render_process_thread_created(
  struct _cef_browser_process_handler_t* self,
  struct _cef_list_value_t* extra_info) {
  printf("cef_browser_process_handler_t_on_render_process_thread_created %X", (size_t) extra_info);
    OnRenderProcessThreadCreated(self, extra_info);
}

struct _cef_print_handler_t* CEF_CALLBACK cef_browser_process_handler_t_get_print_handler(
      struct _cef_browser_process_handler_t* self) {
    return 0;
}


void intialize_cef_browser_process_handler(struct _cef_browser_process_handler_t* handler) {
    goDebugLog("initializeProcessHandler\n");
    handler->base.size = sizeof(cef_browser_process_handler_t);
    initialize_cef_base((cef_base_t*) handler, "browser_process_handler");

    handler->on_context_initialized = cef_browser_process_handler_t_on_context_initialized;
    handler->on_before_child_process_launch = cef_browser_process_handler_t_on_before_child_process_launch;
    handler->on_render_process_thread_created = cef_browser_process_handler_t_on_render_process_thread_created;
    handler->get_print_handler = cef_browser_process_handler_t_get_print_handler;
}