// Copyright (c) 2014 The cefcapi authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cefcapi
// Website: https://github.com/fromkeith/cefcapi

#include "_cgo_export.h"
#include "cef_base.h"
#include <string.h>
#include "bridge.h"

///
// Increment the reference count.
///
void CEF_CALLBACK add_ref(cef_base_t* self) {
    if(self == NULL){
        return;
    }
    acquire_lock();
    struct MemoryManagedBridge *s = find_handler((void*)self);
    s->count = s->count + 1;
    release_lock();
}

///
// Decrement the reference count.  Delete this object when no references
// remain.
///
int CEF_CALLBACK release(cef_base_t* self) {
    // there should a lock here 
    if(self == NULL){
        return 1;
    }
    acquire_lock();
    struct MemoryManagedBridge *s = find_handler((void*)self);
    s->count = s->count - 1;
    if(s->count == 0){
        if(s->Deconstructor != NULL){
            s->Deconstructor((void*)self);
        }
        delete_handler((void*) self);
        return 1;
    } else {
        replace_handler((void*) self, s);
    }
    release_lock();
    return 0; 
}

///
// Returns the current number of references.
///
int CEF_CALLBACK has_one_ref(cef_base_t* self) {
  struct MemoryManagedBridge *s = find_handler((void*)self);
  if(s->count == 1){
    return 1;
  } else {
    return 0;
  }
}

void initialize_cef_base(cef_base_t* base) {
    // Check if "size" member was set.
    size_t size = base->size;
    // Let's print the size in case sizeof was used
    // on a pointer instead of a structure. In such
    // case the number will be very high.
    if (size <= 0) {
        _exit(1);
    }
    base->add_ref = add_ref;
    base->release = release;
    base->has_one_ref = has_one_ref;

    create_handler((void*) base);
}

//
// other base/shared items
//

// returns a utf8 encoded string that you need to delete
cef_string_utf8_t * cefStringToUtf8(const cef_string_t * source) {
    cef_string_utf8_t * output = cef_string_userfree_utf8_alloc();
    if (source == 0) {
        return output;
    }
    cef_string_to_utf8(source->str, source->length, output);
    return output;
}

cef_string_t * cefString16CastToCefString(cef_string_utf16_t * source) {
    return (cef_string_t *) source;
}
cef_string_utf16_t * cefStringCastToCefString16(cef_string_t * source) {
    return (cef_string_utf16_t *) source;
}