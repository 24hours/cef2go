// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi


#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_browser_capi.h"



struct _cef_browser_host_t* cef_browser_t_get_host(struct _cef_browser_t* self) {
    return self->get_host(self);
}


int cef_browser_t_can_go_back(struct _cef_browser_t* self) {
    return self->can_go_back(self);
}

void cef_browser_t_go_back(struct _cef_browser_t* self) {
    self->go_back(self);
}

int cef_browser_t_can_go_forward(struct _cef_browser_t* self) {
    return self->can_go_forward(self);
}

void cef_browser_t_go_forward(struct _cef_browser_t* self) {
    self->go_forward(self);
}

int cef_browser_t_is_loading(struct _cef_browser_t* self) {
    return self->is_loading(self);
}

void cef_browser_t_reload(struct _cef_browser_t* self) {
    self->reload(self);
}

void cef_browser_t_reload_ignore_cache(struct _cef_browser_t* self) {
    self->reload_ignore_cache(self);
}

void cef_browser_t_stop_load(struct _cef_browser_t* self) {
    self->stop_load(self);
}

int cef_browser_t_get_identifier(struct _cef_browser_t* self) {
    return self->get_identifier(self);
}

int cef_browser_t_is_same(struct _cef_browser_t* self, struct _cef_browser_t* that) {
    return self->is_same(self, that);
}

int cef_browser_t_is_popup(struct _cef_browser_t* self) {
    return self->is_popup(self);
}

int cef_browser_t_has_document(struct _cef_browser_t* self) {
    return self->has_document(self);
}

struct _cef_frame_t* cef_browser_t_get_main_frame(struct _cef_browser_t* self) {
    return self->get_main_frame(self);
}

struct _cef_frame_t* cef_browser_t_get_focused_frame(struct _cef_browser_t* self) {
    return self->get_focused_frame(self);
}

struct _cef_frame_t* cef_browser_t_get_frame_byident(struct _cef_browser_t* self, int64 identifier) {
    return self->get_frame_byident(self, identifier);
}

struct _cef_frame_t* cef_browser_t_get_frame(struct _cef_browser_t* self, char * nameChar) {
    cef_string_t * name = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(nameChar, strlen(nameChar), name);
    struct _cef_frame_t* result = self->get_frame(self, name);
    cef_string_userfree_utf16_free(name);
    return result;
}

size_t cef_browser_t_get_frame_count(struct _cef_browser_t* self) {
    return self->get_frame_count(self);
}


size_t cef_browser_t_get_frame_identifiers(struct _cef_browser_t* self, size_t count, int64 * ids) {
    self->get_frame_identifiers(self, &count, ids);
    return count;
}

void cef_browser_t_get_frame_names(struct _cef_browser_t* self, cef_string_list_t names) {
    self->get_frame_names(self, names);
}


int64 int64_array_get(int64* aa, size_t i) {
    return aa[i];
}