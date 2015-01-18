// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go


#include "_cgo_export.h"

#include "cef_base.h"
#include "include/capi/cef_download_handler_capi.h"



//_cef_before_download_callback_t
void callBeforeDownloadCallback_cont(struct _cef_before_download_callback_t* self,
      const char* download_path, int show_dialog) {

    cef_string_t * pathCefString = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(download_path, strlen(download_path), pathCefString);

    self->cont(self, pathCefString, show_dialog);

    cef_string_userfree_utf16_free(pathCefString);
}

//_cef_download_item_callback_t
void callDownloadItemCallback_cancel(struct _cef_download_item_callback_t* self) {
    self->cancel(self);
}




// _cef_download_item_t
int callCefDownloadItem_is_valid(struct _cef_download_item_t* self) {
    return self->is_valid(self);
}
int callCefDownloadItem_is_in_progress(struct _cef_download_item_t* self) {
    return self->is_in_progress(self);
}
int callCefDownloadItem_is_complete(struct _cef_download_item_t* self) {
    return self->is_complete(self);
}
int callCefDownloadItem_is_canceled(struct _cef_download_item_t* self) {
    return self->is_canceled(self);
}
int64 callCefDownloadItem_get_current_speed(struct _cef_download_item_t* self) {
    return self->get_current_speed(self);
}
int callCefDownloadItem_get_percent_complete(struct _cef_download_item_t* self) {
    return self->get_percent_complete(self);
}
int64 callCefDownloadItem_get_total_bytes(struct _cef_download_item_t* self) {
    return self->get_total_bytes(self);
}
int64 callCefDownloadItem_get_received_bytes(struct _cef_download_item_t* self) {
    return self->get_received_bytes(self);
}
cef_time_t callCefDownloadItem_get_start_time(struct _cef_download_item_t* self) {
    return self->get_start_time(self);
}
cef_time_t callCefDownloadItem_get_end_time(struct _cef_download_item_t* self) {
    return self->get_end_time(self);
}
cef_string_utf8_t* callCefDownloadItem_get_full_path(struct _cef_download_item_t* self) {
    cef_string_userfree_t str = self->get_full_path(self);
    if (str == NULL) {
        return NULL;
    }
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}
uint32 callCefDownloadItem_get_id(struct _cef_download_item_t* self) {
    return self->get_id(self);
}
cef_string_utf8_t* callCefDownloadItem_get_url(struct _cef_download_item_t* self) {
    cef_string_userfree_t str = self->get_url(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}
cef_string_utf8_t* callCefDownloadItem_get_suggested_file_name(struct _cef_download_item_t* self) {
    cef_string_userfree_t str = self->get_suggested_file_name(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}
cef_string_utf8_t* callCefDownloadItem_get_content_disposition(struct _cef_download_item_t* self) {
    cef_string_userfree_t str = self->get_content_disposition(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}
cef_string_utf8_t* callCefDownloadItem_get_mime_type(struct _cef_download_item_t* self) {
    cef_string_userfree_t str = self->get_mime_type(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}







// _cef_download_handler_t


///
// Called before a download begins. |suggested_name| is the suggested name for
// the download file. By default the download will be canceled. Execute
// |callback| either asynchronously or in this function to continue the
// download if desired. Do not keep a reference to |download_item| outside of
// this function.
///
void CEF_CALLBACK on_before_download(struct _cef_download_handler_t* self,
  struct _cef_browser_t* browser,
  struct _cef_download_item_t* download_item,
  const cef_string_t* suggested_name,
  struct _cef_before_download_callback_t* callback) {
    cef_string_utf8_t * out = cefStringToUtf8(suggested_name);
    go_OnBeforeDownload(self, browser, download_item, out, callback);
}

///
// Called when a download's status or progress information has been updated.
// This may be called multiple times before and after on_before_download().
// Execute |callback| either asynchronously or in this function to cancel the
// download if desired. Do not keep a reference to |download_item| outside of
// this function.
///
void CEF_CALLBACK on_download_updated(struct _cef_download_handler_t* self,
  struct _cef_browser_t* browser,
  struct _cef_download_item_t* download_item,
  struct _cef_download_item_callback_t* callback) {
    go_OnDownloadUpdated(self, browser, download_item, callback);
}


void initialize_download_handler(struct _cef_download_handler_t* self) {
    goDebugLog("initialize_download_handler\n");
    self->base.size = sizeof(cef_download_handler_t);
    initialize_cef_base((cef_base_t*) self, "download_handler");
    // callbacks
    self->on_before_download = on_before_download;
    self->on_download_updated = on_download_updated;
}