// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi


#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"


struct _cef_browser_t* cef_browser_host_t_get_browser(
  struct _cef_browser_host_t* self)
{
    return self->get_browser(self);
}

void cef_browser_host_t_close_browser(struct _cef_browser_host_t* self,
  int force_close)
{
    self->close_browser(self, force_close);
}

void cef_browser_host_t_set_focus(struct _cef_browser_host_t* self, int enable)
{
    self->set_focus(self, enable);
}

cef_window_handle_t cef_browser_host_t_get_window_handle(
  struct _cef_browser_host_t* self)
{
    return self->get_window_handle(self);
}

cef_window_handle_t cef_browser_host_t_get_opener_window_handle(
  struct _cef_browser_host_t* self)
{
    return self->get_opener_window_handle(self);
}

struct _cef_client_t* cef_browser_host_t_get_client(
  struct _cef_browser_host_t* self)
{
    return self->get_client(self);
}

struct _cef_request_context_t* cef_browser_host_t_get_request_context(
  struct _cef_browser_host_t* self)
{
    return self->get_request_context(self);
}

// not in the version of headers checked into this repo
/*cef_string_utf8_t* cef_browser_host_t_get_dev_tools_url(
  struct _cef_browser_host_t* self, int http_scheme)
{
    cef_string_t * str = self->get_dev_tools_url(self, http_scheme);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out
}
*/

double cef_browser_host_t_get_zoom_level(struct _cef_browser_host_t* self)
{
    return self->get_zoom_level(self);
}

void cef_browser_host_t_set_zoom_level(struct _cef_browser_host_t* self,
  double zoomLevel)
{
    self->set_zoom_level(self, zoomLevel);
}

void cef_browser_host_t_run_file_dialog(struct _cef_browser_host_t* self,
  cef_file_dialog_mode_t mode, const cef_string_t* title,
  const cef_string_t* default_file_name, cef_string_list_t accept_types,
  struct _cef_run_file_dialog_callback_t* callback)
{
    self->run_file_dialog(self, mode, title, default_file_name, accept_types, callback);
}

void cef_browser_host_t_start_download(struct _cef_browser_host_t* self,
  const cef_string_t* url)
{
    self->start_download(self, url);
}

void cef_browser_host_t_print(struct _cef_browser_host_t* self)
{
    self->print(self);
}

void cef_browser_host_t_find(struct _cef_browser_host_t* self, int identifier,
  const cef_string_t* searchText, int forward, int matchCase,
  int findNext)
{
    self->find(self, identifier, searchText, forward, matchCase, findNext);
}

void cef_browser_host_t_stop_finding(struct _cef_browser_host_t* self,
  int clearSelection)
{
    self->stop_finding(self, clearSelection);
}

void cef_browser_host_t_set_mouse_cursor_change_disabled(
  struct _cef_browser_host_t* self, int disabled)
{
    self->set_mouse_cursor_change_disabled(self, disabled);
}

int cef_browser_host_t_is_mouse_cursor_change_disabled(
  struct _cef_browser_host_t* self)
{
    return self->is_mouse_cursor_change_disabled(self);
}

int cef_browser_host_t_is_window_rendering_disabled(
  struct _cef_browser_host_t* self)
{
    return self->is_window_rendering_disabled(self);
}

void cef_browser_host_t_was_resized(struct _cef_browser_host_t* self)
{
    self->was_resized(self);
}

void cef_browser_host_t_was_hidden(struct _cef_browser_host_t* self, int hidden)
{
    self->was_hidden(self, hidden);
}

void cef_browser_host_t_notify_screen_info_changed(
  struct _cef_browser_host_t* self)
{
    self->notify_screen_info_changed(self);
}

void cef_browser_host_t_invalidate(struct _cef_browser_host_t* self, cef_paint_element_type_t type)
{
    self->invalidate(self, type);
}

void cef_browser_host_t_send_key_event(struct _cef_browser_host_t* self,
  const struct _cef_key_event_t* event)
{
    self->send_key_event(self, event);
}

void cef_browser_host_t_send_mouse_click_event(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event,
  cef_mouse_button_type_t type, int mouseUp, int clickCount)
{
    self->send_mouse_click_event(self, event, type, mouseUp, clickCount);
}

void cef_browser_host_t_send_mouse_move_event(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event, int mouseLeave)
{
    self->send_mouse_move_event(self, event, mouseLeave);
}

void cef_browser_host_t_send_mouse_wheel_event(struct _cef_browser_host_t* self,
  const struct _cef_mouse_event_t* event, int deltaX, int deltaY)
{
    self->send_mouse_wheel_event(self, event, deltaX, deltaY);
}

void cef_browser_host_t_send_focus_event(struct _cef_browser_host_t* self,
  int setFocus)
{
    self->send_focus_event(self, setFocus);
}

void cef_browser_host_t_send_capture_lost_event(
  struct _cef_browser_host_t* self)
{
    self->send_capture_lost_event(self);
}

cef_text_input_context_t cef_browser_host_t_get_nstext_input_context(
  struct _cef_browser_host_t* self)
{
    return self->get_nstext_input_context(self);
}

void cef_browser_host_t_handle_key_event_before_text_input_client(
  struct _cef_browser_host_t* self, cef_event_handle_t keyEvent)
{
    self->handle_key_event_before_text_input_client(self, keyEvent);
}

void cef_browser_host_t_handle_key_event_after_text_input_client(
  struct _cef_browser_host_t* self, cef_event_handle_t keyEvent)
{
    self->handle_key_event_after_text_input_client(self, keyEvent);
}