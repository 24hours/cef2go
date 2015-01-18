// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi


#include <stdlib.h>
#include <string.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_request_handler_capi.h"



cef_string_utf8_t* cef_request_t_get_url(struct _cef_request_t * self) {
    cef_string_t * str = self->get_url(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}


int cef_request_t_is_read_only(struct _cef_request_t * self) {
    return self->is_read_only(self);
}

void cef_request_t_set_url(struct _cef_request_t * self, char * url) {
    cef_string_t * output = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(url, strlen(url), output);
    self->set_url(self, output);
    cef_string_userfree_utf16_free(output);
}

cef_string_utf8_t* cef_request_t_get_method(struct _cef_request_t * self) {
    cef_string_t * str = self->get_method(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

void cef_request_t_set_method(struct _cef_request_t * self, char * meth){
    cef_string_t * output = cef_string_userfree_utf16_alloc();
    cef_string_from_utf8(meth, strlen(meth), output);
    self->set_method(self, output);
    cef_string_userfree_utf16_free(output);
}

struct _cef_post_data_t * cef_request_t_get_post_data(struct _cef_request_t * self) {
    return self->get_post_data(self);
}

cef_string_multimap_t cef_request_t_get_header_map(struct _cef_request_t * self) {
    cef_string_multimap_t mapPointer = cef_string_multimap_alloc();
    self->get_header_map(self, mapPointer);
    return mapPointer;
}

int cef_request_t_get_flags(struct _cef_request_t * self) {
    return self->get_flags(self);
}

void cef_request_t_set_flags(struct _cef_request_t * self, int f) {
    self->set_flags(self, f);
}

cef_string_utf8_t* cef_request_t_get_first_party_for_cookies(struct _cef_request_t * self) {
    cef_string_t * str = self->get_first_party_for_cookies(self);
    cef_string_utf8_t* out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

int cef_request_t_get_resource_type(struct _cef_request_t * self) {
    return self->get_resource_type(self);
}

int cef_request_t_get_transition_type(struct _cef_request_t * self) {
    return self->get_transition_type(self);
}



/// post_data_t


///
// Returns true (1) if this object is read-only.
///
int cef_post_data_is_read_only(struct _cef_post_data_t* self) {
    return self->is_read_only(self);
}

///
// Returns the number of existing post data elements.
///
size_t cef_post_data_get_element_count(struct _cef_post_data_t* self) {
    return self->get_element_count(self);
}

///
// Retrieve the post data elements.
///
void cef_post_data_get_elements(struct _cef_post_data_t* self,
  size_t* elementsCount, struct _cef_post_data_element_t** elements) {
    return self->get_elements(self, elementsCount, elements);
}

///
// Remove the specified post data element.  Returns true (1) if the removal
// succeeds.
///
int cef_post_data_remove_element(struct _cef_post_data_t* self,
  struct _cef_post_data_element_t* element) {
    return self->remove_element(self, element);
}

///
// Add the specified post data element.  Returns true (1) if the add succeeds.
///
int cef_post_data_add_element(struct _cef_post_data_t* self,
  struct _cef_post_data_element_t* element) {
    return self->add_element(self, element);
}

///
// Remove all existing post data elements.
///
void cef_post_data_remove_elements(struct _cef_post_data_t* self) {
    return self->remove_elements(self);
}

struct _cef_post_data_element_t* get_element_pointer(int i, struct _cef_post_data_element_t** elements) {
    return elements[i];
}





/// post data element
int cef_post_data_element_is_read_only(struct _cef_post_data_element_t* self) {
    return self->is_read_only(self);
}

///
// Remove all contents from the post data element.
///
void cef_post_data_element_set_to_empty(struct _cef_post_data_element_t* self) {
    self->set_to_empty(self);
}

///
// The post data element will represent a file.
///
void cef_post_data_element_set_to_file(struct _cef_post_data_element_t* self, const cef_string_t* fileName) {
    self->set_to_file(self, fileName);
}

///
// The post data element will represent bytes.  The bytes passed in will be
// copied.
///
void cef_post_data_element_set_to_bytes(struct _cef_post_data_element_t* self, size_t size, const void* bytes) {
    self->set_to_bytes(self, size, bytes);
}

///
// Return the type of this post data element.
///
cef_postdataelement_type_t cef_post_data_element_get_type(struct _cef_post_data_element_t* self) {
    return self->get_type(self);
}

///
// Return the file name.
///
// The resulting string must be freed by calling cef_string_userfree_free().
cef_string_utf8_t* cef_post_data_element_get_file(struct _cef_post_data_element_t* self) {
    cef_string_userfree_t str = self->get_file(self);
    cef_string_utf8_t * out = cefStringToUtf8(str);
    cef_string_userfree_free(str);
    return out;
}

///
// Return the number of bytes.
///
size_t cef_post_data_element_get_bytes_count(struct _cef_post_data_element_t* self) {
    return self->get_bytes_count(self);
}

///
// Read up to |size| bytes into |bytes| and return the number of bytes
// actually read.
///
size_t cef_post_data_element_get_bytes(struct _cef_post_data_element_t* self, size_t size, void* bytes) {
    return self->get_bytes(self, size, bytes);
}