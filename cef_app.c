// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi

#include "cef_base.h"
#include "include/capi/cef_app_capi.h"
#include "_cgo_export.h"

void CEF_CALLBACK on_before_command_line_processing(struct _cef_app_t* self, 
                                                    const cef_string_t* process_type,
                                                    struct _cef_command_line_t* command_line) {
    releaseVoid((void *) command_line);
    
}

void CEF_CALLBACK on_register_custom_schemes(struct _cef_app_t* self,
                                             struct _cef_scheme_registrar_t* registrar) {
    releaseVoid((void *) registrar);
}

struct _cef_resource_bundle_handler_t*
        CEF_CALLBACK get_resource_bundle_handler(struct _cef_app_t* self) {
    return NULL;
}

struct _cef_browser_process_handler_t*
        CEF_CALLBACK get_browser_process_handler(struct _cef_app_t* self) {
    return go_GetBrowserProcessHandler(self);
}

struct _cef_render_process_handler_t*
        CEF_CALLBACK get_render_process_handler(struct _cef_app_t* self) {
    return NULL;
}

void initialize_app_handler(cef_app_t* app) {
    app->base.size = sizeof(cef_app_t);
    initialize_cef_base((cef_base_t*) app, "app_handler");
    go_AddRef((cef_base_t*) app);
    // callbacks
    app->on_before_command_line_processing = on_before_command_line_processing;
    app->on_register_custom_schemes = on_register_custom_schemes;
    app->get_resource_bundle_handler = get_resource_bundle_handler;
    app->get_browser_process_handler = get_browser_process_handler;
    app->get_render_process_handler = get_render_process_handler;
}