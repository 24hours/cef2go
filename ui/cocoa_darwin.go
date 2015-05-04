// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package ui

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <stdlib.h>
#include <string.h>
#include <Cocoa/Cocoa.h>
#include <mach-o/dyld.h>
extern void _GoDestroySignal(void*);

void InitializeApp() {
    [NSAutoreleasePool new];
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
}
@interface WindowDelegate : NSObject <NSWindowDelegate> {
    @private
        NSView* view_;
}
@property (nonatomic, assign) NSView* view;
@end
@implementation WindowDelegate
@synthesize view = view_;
- (void)windowWillClose:(NSNotification *)notification {
    [NSAutoreleasePool new];
    _GoDestroySignal((__bridge void*)view_);
}
@end
void* CreateWindow(char* title, int width, int height) {
    [NSAutoreleasePool new];
    WindowDelegate* delegate = [[WindowDelegate alloc] init];
    id window = [[NSWindow alloc]
              initWithContentRect:NSMakeRect(0, 0, width, height)
              styleMask:(NSTitledWindowMask |
                         NSClosableWindowMask |
                         NSMiniaturizableWindowMask |
                         NSResizableWindowMask |
                         NSUnifiedTitleAndToolbarWindowMask )
              backing:NSBackingStoreBuffered
              defer:NO];
    delegate.view = [window contentView];
    [window setDelegate:(id)delegate];
    [window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
    [window setTitle:[NSString stringWithUTF8String:title]];
    [window makeKeyAndOrderFront:nil];
    return (__bridge void*) [window contentView];
}
void ActivateApp() {
    [NSAutoreleasePool new];
    [NSApp activateIgnoringOtherApps:YES];
}
*/
import "C"
import "unsafe"
import (
	"github.com/24hours/chrome"
)

func init() {
	C.InitializeApp()
}

func CreateWindow(title string, width int, height int) chrome.WindowInfo {
	csTitle := C.CString(title)
	defer C.free(unsafe.Pointer(csTitle))
	window := chrome.WindowInfo{}
	window.Ptr = C.CreateWindow(csTitle, C.int(width), C.int(height))
	window.Height = height
	window.Width = width
	return window
}

func ActivateApp() {
	C.ActivateApp()
}

type DestroyCallback func()

var destroySignalCallbacks map[uintptr]DestroyCallback = make(map[uintptr]DestroyCallback)

func ConnectDestroySignal(window chrome.WindowInfo, callback DestroyCallback) {
	ptr := uintptr(window.Ptr)
	destroySignalCallbacks[ptr] = callback
}
