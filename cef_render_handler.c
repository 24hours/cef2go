#include <stdlib.h>
#include "cef_base.h"
#include "include/capi/cef_client_capi.h"
#include "include/capi/cef_render_handler_capi.h"
#include "_cgo_export.h"

///
// Called to retrieve the root window rectangle in screen coordinates. Return
// true (1) if the rectangle was provided.
///
int CEF_CALLBACK get_root_screen_rect(	struct _cef_render_handler_t* self,
  										struct _cef_browser_t* browser, 
  										cef_rect_t* rect){
	if(rect != NULL){
		return 1;
	} else {
		return 0;
	}
}

///
// Called to retrieve the view rectangle which is relative to screen
// coordinates. Return true (1) if the rectangle was provided.
///
int CEF_CALLBACK get_view_rect(struct _cef_render_handler_t* self,
							  struct _cef_browser_t* browser, 
							  cef_rect_t* rect){
	goDebugLog("get_view_rect\n");
	if(rect != NULL){
		return 1;
	} else {
		return 0;
	}
}
///
// Called to retrieve the translation from view coordinates to actual screen
// coordinates. Return true (1) if the screen coordinates were provided.
///
int CEF_CALLBACK get_screen_point(struct _cef_render_handler_t* self,
								  struct _cef_browser_t* browser, 
								  int viewX, 
								  int viewY, 
								  int* screenX,
								  int* screenY){
		goDebugLog("get_screen_point\n");
	return 1;	
}

///
// Called to allow the client to fill in the CefScreenInfo object with
// appropriate values. Return true (1) if the |screen_info| structure has been
// modified.
//
// If the screen info rectangle is left NULL the rectangle from GetViewRect
// will be used. If the rectangle is still NULL or invalid popups may not be
// drawn correctly.
///
int CEF_CALLBACK get_screen_info(struct _cef_render_handler_t* self,
  								struct _cef_browser_t* browser, 
  								struct _cef_screen_info_t* screen_info){
			goDebugLog("get_screen_info\n");
	return 1;
}

///
// Called when the browser wants to show or hide the popup widget. The popup
// should be shown if |show| is true (1) and hidden if |show| is false (0).
///
void CEF_CALLBACK on_popup_show(struct _cef_render_handler_t* self,
  								struct _cef_browser_t* browser, 
  								int show){
		goDebugLog("on_popup_show\n");
}

///
// Called when the browser wants to move or resize the popup widget. |rect|
// contains the new location and size in view coordinates.
///
void CEF_CALLBACK on_popup_size(struct _cef_render_handler_t* self,
  								struct _cef_browser_t* browser, 
  								const cef_rect_t* rect){
		goDebugLog("on_popup_size\n");
}

///
// Called when an element should be painted. Pixel values passed to this
// function are scaled relative to view coordinates based on the value of
// CefScreenInfo.device_scale_factor returned from GetScreenInfo. |type|
// indicates whether the element is the view or the popup widget. |buffer|
// contains the pixel data for the whole image. |dirtyRects| contains the set
// of rectangles in pixel coordinates that need to be repainted. |buffer| will
// be |width|*|height|*4 bytes in size and represents a BGRA image with an
// upper-left origin.
///
void CEF_CALLBACK on_paint(	struct _cef_render_handler_t* self,
  							struct _cef_browser_t* browser, 
  							cef_paint_element_type_t type,
  							size_t dirtyRectsCount, 
  							cef_rect_t const* dirtyRects, 
  							const void* buffer,
  							int width, 
  							int height){
			goDebugLog("on_paint\n");
}

///
// Called when the browser's cursor has changed. If |type| is CT_CUSTOM then
// |custom_cursor_info| will be populated with the custom cursor information.
///
void CEF_CALLBACK on_cursor_change(	struct _cef_render_handler_t* self,
  									struct _cef_browser_t* browser, 
  									cef_cursor_handle_t cursor,
  									cef_cursor_type_t type,
  									const struct _cef_cursor_info_t* custom_cursor_info){
			goDebugLog("on_cursor_change\n");
}

///
// Called when the user starts dragging content in the web view. Contextual
// information about the dragged content is supplied by |drag_data|. (|x|,
// |y|) is the drag start location in screen coordinates. OS APIs that run a
// system message loop may be used within the StartDragging call.
//
// Return false (0) to abort the drag operation. Don't call any of
// cef_browser_host_t::DragSource*Ended* functions after returning false (0).
//
// Return true (1) to handle the drag operation. Call
// cef_browser_host_t::DragSourceEndedAt and DragSourceSystemDragEnded either
// synchronously or asynchronously to inform the web view that the drag
// operation has ended.
///
int CEF_CALLBACK start_dragging(struct _cef_render_handler_t* self,
  								struct _cef_browser_t* browser, 
  								struct _cef_drag_data_t* drag_data,
  								cef_drag_operations_mask_t allowed_ops, 
  								int x, 
  								int y){
			goDebugLog("start_dragging\n");
	return 1;
}

///
// Called when the web view wants to update the mouse cursor during a drag &
// drop operation. |operation| describes the allowed operation (none, move,
// copy, link).
///
void CEF_CALLBACK update_drag_cursor(struct _cef_render_handler_t* self,
  									struct _cef_browser_t* browser, 
  									cef_drag_operations_mask_t operation){
			goDebugLog("update_drag_cursor\n");
}

///
// Called when the scroll offset has changed.
///
void CEF_CALLBACK on_scroll_offset_changed(	struct _cef_render_handler_t* self, 
											struct _cef_browser_t* browser){
			goDebugLog("on_scroll_offset_changed\n");
}


void initialize_render_handler(struct _cef_render_handler_t* renderHandler){
    renderHandler->base.size = sizeof(cef_render_handler_t);
    initialize_cef_base((cef_base_t*) renderHandler, "render_handler");
    go_AddRef((cef_base_t*) renderHandler);

	renderHandler->get_root_screen_rect = get_root_screen_rect;
	renderHandler->get_view_rect = get_view_rect;
	renderHandler->get_screen_point = get_screen_point;
	renderHandler->get_screen_info = get_screen_info;
	renderHandler->on_popup_show = on_popup_show;
	renderHandler->on_popup_size = on_popup_size;
	renderHandler->on_paint = on_paint;
	renderHandler->on_cursor_change = on_cursor_change;
	renderHandler->start_dragging = start_dragging;
	renderHandler->update_drag_cursor = update_drag_cursor;
	renderHandler->on_scroll_offset_changed = on_scroll_offset_changed;
}