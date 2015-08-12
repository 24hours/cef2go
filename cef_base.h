// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi
// Website: https://github.com/fromkeith/cefcapi

#ifndef CEF_BASE_H
#define CEF_BASE_H

#include <stdio.h>
#include <unistd.h>
#include "include/capi/cef_base_capi.h"

cef_string_utf8_t * cefStringToUtf8(const cef_string_t * source);
cef_string_t * 	cefString16CastToCefString(cef_string_utf16_t * source);
cef_string_utf16_t * cefStringCastToCefString16(cef_string_t * source);
char * cefStringtoChar( cef_string_t * source);

void CEF_CALLBACK add_ref(cef_base_t* self);
int CEF_CALLBACK release(cef_base_t* self);

void initialize_cef_base(cef_base_t* base);
void add_refVoid(void* self);
int releaseVoid(void* self);

#endif
