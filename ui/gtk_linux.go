// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package ui

/*
#cgo CFLAGS: -I../
#cgo pkg-config: --libs --cflags gtk+-2.0
#include <gtk/gtk.h>
#include <gdk/gdk.h>
#include <gdk/gdkx.h>
#include <stdlib.h>
#include <string.h>
#include "include/capi/cef_app_capi.h"

static inline GtkWindow* ToGtkWindow(GtkWidget* w) { return GTK_WINDOW(w); }
static inline GtkContainer* ToGtkContainer(GtkWidget* w) { return GTK_CONTAINER(w); }
void TerminationSignal(int signatl) { cef_quit_message_loop(); }
void ConnectTerminationSignal() {
    signal(SIGINT, TerminationSignal);
    signal(SIGTERM, TerminationSignal);
}
void DestroySignal(GtkWidget* widget, gpointer data) {
    _GoDestroySignal(widget, data);
}
void ConnectDestroySignal(GtkWidget* window) {
    g_signal_connect(G_OBJECT(window), "destroy",
            G_CALLBACK(DestroySignal), NULL);
}
*/
import "C"
import "unsafe"
import "github.com/24hours/chrome"

func init() {
	C.gtk_init(nil, nil)
	C.ConnectTerminationSignal()
}

func CreateWindow(title string, width int, height int) chrome.WindowInfo {
	// Create window.
	window := C.gtk_window_new(C.GTK_WINDOW_TOPLEVEL)

	// Default size.
	C.gtk_window_set_default_size(C.ToGtkWindow(window), C.gint(width), C.gint(height))

	// Center.
	C.gtk_window_set_position(C.ToGtkWindow(window), C.GTK_WIN_POS_CENTER)

	// Title.
	csTitle := C.CString(title)
	defer C.free(unsafe.Pointer(csTitle))
	C.gtk_window_set_title(C.ToGtkWindow(window), (*C.gchar)(csTitle))

	// TODO: focus
	// g_signal_connect(window, "focus", G_CALLBACK(&HandleFocus), NULL);

	// CEF requires a container. Embedding browser in a top
	// level window fails.
	vbox := C.gtk_vbox_new(0, 0)
	C.gtk_container_add(C.ToGtkContainer(window), vbox)

	// Show.
	C.gtk_widget_show_all(window)

	ret := chrome.WindowInfo{}
	ret.Ptr = unsafe.Pointer(window)
	ret.Hdl = uint64(C.cef_window_handle_t(C.gdk_x11_drawable_get_xid(C.gtk_widget_get_window(window))))
	return ret
}

type DestroyCallback func()

var destroySignalCallbacks map[uintptr]DestroyCallback = make(map[uintptr]DestroyCallback)

func ConnectDestroySignal(window chrome.WindowInfo, callback DestroyCallback) {
	ptr := uintptr(window.Ptr)
	destroySignalCallbacks[ptr] = callback
	C.ConnectDestroySignal((*C.GtkWidget)(window.Ptr))
}
