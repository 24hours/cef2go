// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go

#include "_cgo_export.h"

#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_display_handler_capi.h"

// Provides pass throughs for _cef_display_handler_t to golang



void CEF_CALLBACK _cef_display_handler_t_on_address_change(struct _cef_display_handler_t* self,
      struct _cef_browser_t* browser, struct _cef_frame_t* frame,
      const cef_string_t* url) {

  cef_string_utf8_t * out = cefStringToUtf8(url);
  go_OnAddressChange(self, browser, frame, out->str);
  cef_string_userfree_utf8_free(out);
}

///
// Called when the page title changes.
///
void CEF_CALLBACK _cef_display_handler_t_on_title_change(struct _cef_display_handler_t* self,
    struct _cef_browser_t* browser, const cef_string_t* title) {

  cef_string_utf8_t * out = cefStringToUtf8(title);
  go_OnTitleChange(self, browser, out->str);
  cef_string_userfree_utf8_free(out);
}

///
// Called when the browser is about to display a tooltip. |text| contains the
// text that will be displayed in the tooltip. To handle the display of the
// tooltip yourself return true (1). Otherwise, you can optionally modify
// |text| and then return false (0) to allow the browser to display the
// tooltip. When window rendering is disabled the application is responsible
// for drawing tooltips and the return value is ignored.
///
int CEF_CALLBACK _cef_display_handler_t_on_tooltip(struct _cef_display_handler_t* self,
    struct _cef_browser_t* browser, cef_string_t* text) {

  cef_string_utf8_t * out = cefStringToUtf8(text);
  int ret = go_OnTooltip(self, browser, out->str);
  cef_string_userfree_utf8_free(out);
  return ret;
}

///
// Called when the browser receives a status message. |text| contains the text
// that will be displayed in the status message and |type| indicates the
// status message type.
///
void CEF_CALLBACK _cef_display_handler_t_on_status_message(struct _cef_display_handler_t* self,
    struct _cef_browser_t* browser, const cef_string_t* value) {

  cef_string_utf8_t * out = cefStringToUtf8(value);
  go_OnStatusMessage(self, browser, out->str);
  cef_string_userfree_utf8_free(out);
}

///
// Called to display a console message. Return true (1) to stop the message
// from being output to the console.
///
int CEF_CALLBACK _cef_display_handler_t_on_console_message(struct _cef_display_handler_t* self,
    struct _cef_browser_t* browser, const cef_string_t* message,
    const cef_string_t* source, int line) {

  cef_string_utf8_t * out1 = cefStringToUtf8(message);
  cef_string_utf8_t * out2 = cefStringToUtf8(source);
  int ret = go_OnConsoleMessage(self, browser, out1->str, out2->str, line);
  cef_string_userfree_utf8_free(out2);
  cef_string_userfree_utf8_free(out1);
  return ret;

}


void initialize_display_handler(struct _cef_display_handler_t * displayHandler) {
    goDebugLog("initialize_display_handler");
    displayHandler->base.size = sizeof(cef_display_handler_t);
    initialize_cef_base((cef_base_t*) displayHandler, "display_handler");
    // callbacks
    displayHandler->on_address_change = _cef_display_handler_t_on_address_change;
    displayHandler->on_title_change = _cef_display_handler_t_on_title_change;
    displayHandler->on_tooltip = _cef_display_handler_t_on_tooltip;
    displayHandler->on_status_message = _cef_display_handler_t_on_status_message;
    displayHandler->on_console_message = _cef_display_handler_t_on_console_message;
}