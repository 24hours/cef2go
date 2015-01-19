
// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi


#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_life_span_handler_capi.h"
#include "_cgo_export.h"

int CEF_CALLBACK cef_life_span_handler_t_on_before_popup(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser,
        struct _cef_frame_t* frame,
        const cef_string_t* target_url,
        const cef_string_t* target_frame_name,
        const struct _cef_popup_features_t* popupFeatures,
        struct _cef_window_info_t* windowInfo,
        struct _cef_client_t** client,
        struct _cef_browser_settings_t* settings,
        int* no_javascript_access) {
    return go_OnBeforePopup(
        self,
        browser,
        frame,
        (cef_string_t*)target_url,
        (cef_string_t*)target_frame_name,
        (struct _cef_popup_features_t*)popupFeatures,
        // for some reason warning for const to non-const parameter 
        // is not suppressed by cgo. 
        windowInfo,
        client,
        settings,
        no_javascript_access);
}


void CEF_CALLBACK cef_life_span_handler_t_on_after_created(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    go_OnAfterCreated(self, browser);
}

int CEF_CALLBACK cef_life_span_handler_t_run_modal(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    return go_RunModal(self, browser);
}

int CEF_CALLBACK cef_life_span_handler_t_do_close(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    return go_DoClose(self, browser);
}

void CEF_CALLBACK cef_life_span_handler_t_on_before_close(
        struct _cef_life_span_handler_t* self,
        struct _cef_browser_t* browser) {
    go_BeforeClose(self, browser);
}


void initialize_life_span_handler(struct _cef_life_span_handler_t* lifeHandler) {
    goDebugLog("initialize_life_span_handler\n");
    lifeHandler->base.size = sizeof(cef_life_span_handler_t);
    initialize_cef_base((cef_base_t*) lifeHandler, "life_span_handler");
    go_AddRef((cef_base_t*) lifeHandler);
    // callbacks
    lifeHandler->on_before_popup = cef_life_span_handler_t_on_before_popup;
    lifeHandler->on_after_created = cef_life_span_handler_t_on_after_created;
    lifeHandler->run_modal = cef_life_span_handler_t_run_modal;
    lifeHandler->do_close = cef_life_span_handler_t_do_close;
    lifeHandler->on_before_close = cef_life_span_handler_t_on_before_close;
}