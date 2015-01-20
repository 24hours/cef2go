// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cef2go
// Website: https://github.com/CzarekTomczak/cef2go

#include "_cgo_export.h"
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_request_handler_capi.h"

// Provides pass throughs for _cef_request_handler_t to golang

///
// Called on the UI thread before browser navigation. Return true (1) to
// cancel the navigation or false (0) to allow the navigation to proceed. The
// |request| object cannot be modified in this callback.
// cef_load_handler_t::OnLoadingStateChange will be called twice in all cases.
// If the navigation is allowed cef_load_handler_t::OnLoadStart and
// cef_load_handler_t::OnLoadEnd will be called. If the navigation is canceled
// cef_load_handler_t::OnLoadError will be called with an |errorCode| value of
// ERR_ABORTED.
///
int CEF_CALLBACK on_before_browse(struct _cef_request_handler_t* self,
                                  struct _cef_browser_t* browser, 
                                  struct _cef_frame_t* frame,
                                  struct _cef_request_t* request, 
                                  int is_redirect) {
    return go_OnBeforeBrowse(self, browser, frame, request, is_redirect);
}

///
// Called on the IO thread before a resource request is loaded. The |request|
// object may be modified. To cancel the request return true (1) otherwise
// return false (0).
///
int CEF_CALLBACK on_before_resource_load(struct _cef_request_handler_t* self, 
                                         struct _cef_browser_t* browser,
                                         struct _cef_frame_t* frame, 
                                         struct _cef_request_t* request) {
    return go_OnBeforeResourceLoad(self, browser, frame, request);
}

///
// Called on the IO thread before a resource is loaded. To allow the resource
// to load normally return NULL. To specify a handler for the resource return
// a cef_resource_handler_t object. The |request| object should not be
// modified in this callback.
///
struct _cef_resource_handler_t* CEF_CALLBACK get_resource_handler(
                                        struct _cef_request_handler_t* self, 
                                        struct _cef_browser_t* browser,
                                        struct _cef_frame_t* frame, 
                                        struct _cef_request_t* request) {
    return go_GetResourceHandler(self, browser, frame, request);
}

///
// Called on the IO thread when a resource load is redirected. The |old_url|
// parameter will contain the old URL. The |new_url| parameter will contain
// the new URL and can be changed if desired.
///
void CEF_CALLBACK on_resource_redirect(struct _cef_request_handler_t* self,
                                       struct _cef_browser_t* browser, 
                                       struct _cef_frame_t* frame,
                                       const cef_string_t* old_url, 
                                       cef_string_t* new_url) {
    go_OnResourceRedirect(self, browser, frame, (char *) old_url, (char *) new_url);
}

///
// Called on the IO thread when the browser needs credentials from the user.
// |isProxy| indicates whether the host is a proxy server. |host| contains the
// hostname and |port| contains the port number. Return true (1) to continue
// the request and call cef_auth_callback_t::cont() when the authentication
// information is available. Return false (0) to cancel the request.
///
int CEF_CALLBACK get_auth_credentials(struct _cef_request_handler_t* self,
                                      struct _cef_browser_t* browser, 
                                      struct _cef_frame_t* frame, 
                                      int isProxy,
                                      const cef_string_t* host, 
                                      int port, 
                                      const cef_string_t* realm,
                                      const cef_string_t* scheme, 
                                      struct _cef_auth_callback_t* callback) {
    return go_GetAuthCredentials(self, browser, frame, isProxy, (char *) host, port, (char *) realm, (char *) scheme, callback);
}

///
// Called on the IO thread when JavaScript requests a specific storage quota
// size via the webkitStorageInfo.requestQuota function. |origin_url| is the
// origin of the page making the request. |new_size| is the requested quota
// size in bytes. Return true (1) and call cef_quota_callback_t::cont() either
// in this function or at a later time to grant or deny the request. Return
// false (0) to cancel the request.
///
int CEF_CALLBACK on_quota_request(struct _cef_request_handler_t* self,
                                  struct _cef_browser_t* browser, 
                                  const cef_string_t* origin_url,
                                  int64 new_size, 
                                  struct _cef_quota_callback_t* callback) {
    return go_OnQuotaRequest(self, browser, (char *) origin_url, new_size, callback);
}

///
// Called on the UI thread to handle requests for URLs with an unknown
// protocol component. Set |allow_os_execution| to true (1) to attempt
// execution via the registered OS protocol handler, if any. SECURITY WARNING:
// YOU SHOULD USE THIS METHOD TO ENFORCE RESTRICTIONS BASED ON SCHEME, HOST OR
// OTHER URL ANALYSIS BEFORE ALLOWING OS EXECUTION.
///
void CEF_CALLBACK on_protocol_execution(struct _cef_request_handler_t* self, 
                                        struct _cef_browser_t* browser,
                                        const cef_string_t* url, 
                                        int* allow_os_execution) {
    go_OnProtocolExecution(self, browser, (char *) url, allow_os_execution);
}

///
// Called on the UI thread to handle requests for URLs with an invalid SSL
// certificate. Return true (1) and call
// cef_allow_certificate_error_callback_t:: cont() either in this function or
// at a later time to continue or cancel the request. Return false (0) to
// cancel the request immediately. If |callback| is NULL the error cannot be
// recovered from and the request will be canceled automatically. If
// CefSettings.ignore_certificate_errors is set all invalid certificates will
// be accepted without calling this function.
///
int CEF_CALLBACK on_certificate_error(struct _cef_request_handler_t* self,
                                      cef_errorcode_t cert_error, 
                                      const cef_string_t* request_url,
                                      struct _cef_allow_certificate_error_callback_t* callback) {
    cef_string_utf8_t * out = cefStringToUtf8(request_url);
    int ret = go_OnCertificateError(self, cert_error, out->str, callback);
    cef_string_userfree_utf8_free(out);
    return ret;
}

///
// Called on the browser process IO thread before a plugin is loaded. Return
// true (1) to block loading of the plugin.
///
int CEF_CALLBACK on_before_plugin_load( struct _cef_request_handler_t* self,
                                        struct _cef_browser_t* browser, 
                                        const cef_string_t* url,
                                        const cef_string_t* policy_url, 
                                        struct _cef_web_plugin_info_t* info) {
    return go_OnBeforePluginLoad(self, browser, (char *) url, (char *) policy_url, info);
}

///
// Called on the browser process UI thread when a plugin has crashed.
// |plugin_path| is the path of the plugin that crashed.
///
void CEF_CALLBACK on_plugin_crashed(struct _cef_request_handler_t* self,
                                    struct _cef_browser_t* browser, 
                                    const cef_string_t* plugin_path) {
    go_OnPluginCrashed(self, browser, (char *) plugin_path);
}

///
// Called on the browser process UI thread when the render process terminates
// unexpectedly. |status| indicates how the process terminated.
///
void CEF_CALLBACK on_render_process_terminated( struct _cef_request_handler_t* self, 
                                                struct _cef_browser_t* browser,
                                                cef_termination_status_t status) {
    go_OnRenderProcessTerminated(self, browser, status);
}


void initialize_request_handler(struct _cef_request_handler_t * requestHandler) {
    goDebugLog("initializeRequestHandler");
    requestHandler->base.size = sizeof(cef_request_handler_t);
    initialize_cef_base((cef_base_t*) requestHandler, "request_handler");
    // callbacks
    requestHandler->on_before_browse = on_before_browse;
    requestHandler->on_before_resource_load = on_before_resource_load;
    requestHandler->get_resource_handler = get_resource_handler;
    requestHandler->on_resource_redirect = on_resource_redirect;
    requestHandler->get_auth_credentials = get_auth_credentials;
    requestHandler->on_quota_request = on_quota_request;
    requestHandler->on_protocol_execution = on_protocol_execution;
    requestHandler->on_certificate_error = on_certificate_error;
    requestHandler->on_before_plugin_load = on_before_plugin_load;
    requestHandler->on_plugin_crashed = on_plugin_crashed;
    requestHandler->on_render_process_terminated = on_render_process_terminated;
}


void cef_allow_certificate_error_callback_t_cont(struct _cef_allow_certificate_error_callback_t* self, int allow) {
  self->cont(self, allow);
}