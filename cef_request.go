// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/fromkeith/cefcapi

package chrome

/*
#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_request_handler_capi.h"
extern cef_string_utf8_t * cef_request_t_get_url(struct _cef_request_t * self);
extern int cef_request_t_is_read_only(struct _cef_request_t * self);
extern void cef_request_t_set_url(struct _cef_request_t * self, char * url);
extern cef_string_utf8_t * cef_request_t_get_method(struct _cef_request_t * self);
extern void cef_request_t_set_method(struct _cef_request_t * self, char * meth);
extern struct _cef_post_data_t * cef_request_t_get_post_data(struct _cef_request_t * self);
extern cef_string_multimap_t cef_request_t_get_header_map(struct _cef_request_t * self);
extern int cef_request_t_get_flags(struct _cef_request_t * self);
extern void cef_request_t_set_flags(struct _cef_request_t * self, int f);
extern cef_string_utf8_t * cef_request_t_get_first_party_for_cookies(struct _cef_request_t * self);
extern int cef_request_t_get_resource_type(struct _cef_request_t * self);
extern int cef_request_t_get_transition_type(struct _cef_request_t * self);
// post data
extern int cef_post_data_is_read_only(struct _cef_post_data_t* self);
extern size_t cef_post_data_get_element_count(struct _cef_post_data_t* self);
extern void cef_post_data_get_elements(struct _cef_post_data_t* self, size_t* elementsCount, struct _cef_post_data_element_t** elements);
extern int cef_post_data_remove_element(struct _cef_post_data_t* self, struct _cef_post_data_element_t* element);
extern int cef_post_data_add_element(struct _cef_post_data_t* self, struct _cef_post_data_element_t* element);
extern void cef_post_data_remove_elements(struct _cef_post_data_t* self);
extern struct _cef_post_data_element_t* get_element_pointer(int i, struct _cef_post_data_element_t** elements);
// post data element
extern int cef_post_data_element_is_read_only(struct _cef_post_data_element_t* self);
extern void cef_post_data_element_set_to_empty(struct _cef_post_data_element_t* self);
extern void cef_post_data_element_set_to_file(struct _cef_post_data_element_t* self, const cef_string_t* fileName);
extern void cef_post_data_element_set_to_bytes(struct _cef_post_data_element_t* self, size_t size, const void* bytes);
extern cef_postdataelement_type_t cef_post_data_element_get_type(struct _cef_post_data_element_t* self);
extern cef_string_utf8_t* cef_post_data_element_get_file(struct _cef_post_data_element_t* self);
extern size_t cef_post_data_element_get_bytes_count(struct _cef_post_data_element_t* self);
extern size_t cef_post_data_element_get_bytes(struct _cef_post_data_element_t* self, size_t size, void* bytes);
*/
import "C"

import (
	"unsafe"
	//"fmt"
)

type CefRequestT struct {
	CStruct *C.struct__cef_request_t
}

type CefPostDataT struct {
	CStruct *C.struct__cef_post_data_t
}

type CefPostDataElementT struct {
	CStruct *C.struct__cef_post_data_element_t
}

func (b CefPostDataElementT) Release() {
	C.releaseVoid(unsafe.Pointer(b.CStruct))
}

func (b CefPostDataElementT) AddRef() {
	C.add_refVoid(unsafe.Pointer(b.CStruct))
}

func (b CefPostDataT) Release() {
	C.releaseVoid(unsafe.Pointer(b.CStruct))
}

func (b CefPostDataT) AddRef() {
	C.add_refVoid(unsafe.Pointer(b.CStruct))
}

func (b CefRequestT) Release() {
	C.releaseVoid(unsafe.Pointer(b.CStruct))
}

func (b CefRequestT) AddRef() {
	C.add_refVoid(unsafe.Pointer(b.CStruct))
}

func (r CefRequestT) GetUrl() string {
	stringStruct := C.cef_request_t_get_url(r.CStruct)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}

func (r CefRequestT) IsReadOnly() bool {
	result := C.cef_request_t_is_read_only(r.CStruct)
	return result == 1
}

func (r CefRequestT) SetUrl(newUrl string) {
	var newUrlCString *C.char = C.CString(newUrl)
	defer C.free(unsafe.Pointer(newUrlCString))
	C.cef_request_t_set_url(r.CStruct, newUrlCString)
}

func (r CefRequestT) GetMethod() string {
	stringStruct := C.cef_request_t_get_method(r.CStruct)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}

func (r CefRequestT) SetMethod(method string) {
	var methodCString *C.char = C.CString(method)
	defer C.free(unsafe.Pointer(methodCString))
	C.cef_request_t_set_url(r.CStruct, methodCString)
}

func (r CefRequestT) GetPostData() CefPostDataT {
	postDataCStruct := C.cef_request_t_get_post_data(r.CStruct)
	return CefPostDataT{postDataCStruct}
}

func (r CefRequestT) SetPostData(data CefPostDataT) {
	// do nothing for now
}

func (r CefRequestT) GetHeaderMap() map[string][]string {
	cefMapPointer := C.cef_request_t_get_header_map(r.CStruct)
	if cefMapPointer == nil {
		return nil
	}
	defer C.cef_string_multimap_free(cefMapPointer)
	return extractCefMultiMap(cefMapPointer)
}

func (r CefRequestT) SetHeaderMap(headerMap map[string]string) {
	// do nothing for now
}

func (r CefRequestT) Set(url, method string, postData CefPostDataT, headerMap map[string]string) {
	// do nothing for now
}

func (r CefRequestT) GetFlags() int {
	return int(C.cef_request_t_get_flags(r.CStruct))
}

func (r CefRequestT) SetFlags(flags int) {
	C.cef_request_t_set_flags(r.CStruct, C.int(flags))
}

func (r CefRequestT) GetFirstPartyForCookies() string {
	stringStruct := C.cef_request_t_get_first_party_for_cookies(r.CStruct)
	defer C.cef_string_userfree_utf8_free(stringStruct)
	str := C.GoString(stringStruct.str)
	return str
}

func (r CefRequestT) SetFirstPartyForCookies(url string) {
	// do nothing cause i'm lazy right now
}

func (r CefRequestT) GetResourceType() int {
	return int(C.cef_request_t_get_resource_type(r.CStruct))
}

func (r CefRequestT) GetTransitionType() int {
	return int(C.cef_request_t_get_transition_type(r.CStruct))
}

func (p CefPostDataT) IsReadOnly() bool {
	return C.cef_post_data_is_read_only(p.CStruct) == 1
}
func (p CefPostDataT) GetElementCount() int {
	return int(C.cef_post_data_get_element_count(p.CStruct))
}
func (p CefPostDataT) GetElements(numToGet int) []CefPostDataElementT {
	var numElements C.size_t = C.size_t(numToGet)
	elements := (**C.struct__cef_post_data_element_t)(
		C.calloc(numElements, 32)) /// i should really have the size of a pointer....

	C.cef_post_data_get_elements(p.CStruct, &numElements, elements)

	numEle := int(numElements)
	out := make([]CefPostDataElementT, numEle)
	for i := range out {
		out[i].CStruct = C.get_element_pointer(C.int(i), elements)
	}
	C.free(unsafe.Pointer(elements))
	return out
}
func (p CefPostDataT) RemoveElement(element CefPostDataElementT) bool {
	element.AddRef()
	return C.cef_post_data_remove_element(p.CStruct, element.CStruct) == 1
}
func (p CefPostDataT) AddElement(element CefPostDataElementT) bool {
	element.AddRef()
	return C.cef_post_data_add_element(p.CStruct, element.CStruct) == 1
}
func (p CefPostDataT) RemoveElements() {
	C.cef_post_data_remove_elements(p.CStruct)
}

// post data element
func (e CefPostDataElementT) IsReadOnly() bool {
	return C.cef_post_data_element_is_read_only(e.CStruct) == 1
}
func (e CefPostDataElementT) SetToEmpty() {
	C.cef_post_data_element_set_to_empty(e.CStruct)
}
func (e CefPostDataElementT) SetToFile(fileName string) {
	panic("Not implemented")
	//C.cef_post_data_element_set_to_file(e.CStruct, const cef_string_t* fileName)
}
func (e CefPostDataElementT) SetToBytes(b []byte) {
	C.cef_post_data_element_set_to_bytes(e.CStruct, C.size_t(len(b)), unsafe.Pointer(&b[0]))
}
func (e CefPostDataElementT) GetType() C.cef_postdataelement_type_t {
	return C.cef_post_data_element_get_type(e.CStruct)
}
func (e CefPostDataElementT) GetFile() string {
	cefString := C.cef_post_data_element_get_file(e.CStruct)
	defer C.cef_string_userfree_utf8_free(cefString)
	return C.GoString(cefString.str)
}
func (e CefPostDataElementT) GetBytesCount() int {
	return int(C.cef_post_data_element_get_bytes_count(e.CStruct))
}
func (e CefPostDataElementT) GetBytes(b []byte) int {

	return int(C.cef_post_data_element_get_bytes(e.CStruct, C.size_t(len(b)), unsafe.Pointer(&b[0])))
}
