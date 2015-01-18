// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi


#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_frame_capi.h"




int cef_frame_t_is_valid(struct _cef_frame_t* self)
{
    return self->is_valid(self);
}

void cef_frame_t_undo(struct _cef_frame_t* self)
{
    self->undo(self);
}

void cef_frame_t_redo(struct _cef_frame_t* self)
{
    self->redo(self);
}

void cef_frame_t_cut(struct _cef_frame_t* self)
{
    self->cut(self);
}

void cef_frame_t_copy(struct _cef_frame_t* self)
{
    self->copy(self);
}

void cef_frame_t_paste(struct _cef_frame_t* self)
{
    self->paste(self);
}

void cef_frame_t_del(struct _cef_frame_t* self)
{
    self->del(self);
}

void cef_frame_t_select_all(struct _cef_frame_t* self)
{
    self->select_all(self);
}

void cef_frame_t_view_source(struct _cef_frame_t* self)
{
    self->view_source(self);
}

void cef_frame_t_get_source(struct _cef_frame_t* self,
      struct _cef_string_visitor_t* visitor)
{
    self->get_source(self, visitor);
}

void cef_frame_t_get_text(struct _cef_frame_t* self,
      struct _cef_string_visitor_t* visitor)
{
    self->get_text(self, visitor);
}

void cef_frame_t_load_request(struct _cef_frame_t* self,
      struct _cef_request_t* request)
{
    self->load_request(self, request);
}

void cef_frame_t_load_url(struct _cef_frame_t* self,
      const char* url)
{
    cef_string_t * output = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(url, strlen(url), output);
    self->load_url(self, output);
    cef_string_userfree_utf16_free(output);
}

void cef_frame_t_load_string(struct _cef_frame_t* self,
      const char* string_val, const char* url)
{
    cef_string_t * strVal = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(string_val, strlen(string_val), strVal);
    cef_string_t * urlVal = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(url, strlen(url), urlVal);

    self->load_string(self, strVal, urlVal);

    cef_string_userfree_utf16_free(urlVal);
    cef_string_userfree_utf16_free(strVal);
}

void cef_frame_t_execute_java_script(struct _cef_frame_t* self,
      const char* code, const char* script_url,
      int start_line)
{
    cef_string_t * codeCef = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(code, strlen(code), codeCef);
    cef_string_t * urlVal = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(script_url, strlen(script_url), urlVal);

    self->execute_java_script(self, codeCef, urlVal, start_line);

    cef_string_userfree_utf16_free(urlVal);
    cef_string_userfree_utf16_free(codeCef);
}

int cef_frame_t_is_main(struct _cef_frame_t* self)
{
    return self->is_main(self);
}


int cef_frame_t_is_focused(struct _cef_frame_t* self)
{
    return self->is_focused(self);
}

cef_string_utf8_t* cef_frame_t_get_name(struct _cef_frame_t* self)
{
    cef_string_t * str = self->get_name(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}


int64 cef_frame_t_get_identifier(struct _cef_frame_t* self)
{
    return self->get_identifier(self);
}

struct _cef_frame_t* cef_frame_t_get_parent(struct _cef_frame_t* self)
{
    return self->get_parent(self);
}


cef_string_utf8_t* cef_frame_t_get_url(struct _cef_frame_t* self)
{
    cef_string_t * str = self->get_url(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

struct _cef_browser_t* cef_frame_t_get_browser(struct _cef_frame_t* self)
{
    return self->get_browser(self);
}

struct _cef_v8context_t* cef_frame_t_get_v8context(
      struct _cef_frame_t* self)
{
    return self->get_v8context(self);
}

void cef_frame_t_visit_dom(struct _cef_frame_t* self,
      struct _cef_domvisitor_t* visitor)
{
    self->visit_dom(self, visitor);
}