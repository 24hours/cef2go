package main

import (
	"github.com/24hours/chrome"
	"github.com/24hours/chrome/ui"
	"os"
    "syscall"
    "unsafe"
)

func main() {
//   hInstance, e := wingui.GetModuleHandle(nil)
//    if e != nil { wingui.AbortErrNo("GetModuleHandle", e) }
    
//    cef.ExecuteProcess(unsafe.Pointer(hInstance), nil)
    chrome.ExecuteProcess(nil, nil)
	
    settings := chrome.Settings{}
    settings.LocalesDirPath = ""
	settings.ResourcesDirPath = ""
	chrome.Initialize(settings, nil)
    
    wndproc := syscall.NewCallback(WndProc)
    hwnd := ui.CreateWindow("cef2go example", wndproc)

    chrome.CreateBrowserAsync(hwnd, nil, chrome.BrowserSettings{}, "http://www.google.com")

    chrome.RunMessageLoop()
    //cef.Shutdown()
    os.Exit(0)
}

func WndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) (rc uintptr) {
    switch msg {
    case ui.WM_CREATE:
        rc = ui.DefWindowProc(hwnd, msg, wparam, lparam)
    case ui.WM_SIZE:
        chrome.WindowResized(unsafe.Pointer(hwnd))
    case ui.WM_CLOSE:
        ui.DestroyWindow(hwnd)
    case ui.WM_DESTROY:
        chrome.QuitMessageLoop()
    default:
        rc = ui.DefWindowProc(hwnd, msg, wparam, lparam)
    }
    return
}