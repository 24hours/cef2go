// Copyright (c) 2014 The cef2go authors. All rights reserved.
// License: BSD 3-clause.
// Website: https://github.com/CzarekTomczak/cef2go

package chrome

/*
CEF capi fixes
--------------
1. In cef_string.h:
    this => typedef cef_string_utf16_t cef_string_t;
    to => #define cef_string_t cef_string_utf16_t
2. In cef_export.h:
    #elif defined(COMPILER_GCC)
    #define CEF_EXPORT __attribute__ ((visibility("default")))
    #ifdef OS_WIN
    #define CEF_CALLBACK __stdcall
    #else
    #define CEF_CALLBACK
    #endif
*/

/*
#include <stdlib.h>
#include "string.h"
#include "include/capi/cef_app_capi.h"
#include "cef_base.h"
#include "cef_app.h"
#include "cef_client.h"
*/
import "C"
import "unsafe"
import (
	"log"
	"os"
	"runtime"
)

var Logger2 SimpleLogger = defaultLogger{}

// A simple interface to wrap a basic leveled logger.
// The format strings to do not have newlines on them.
type SimpleLogger interface {
	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	// Log the panic and exit.
	Panicf(fmt string, args ...interface{})
}

type defaultLogger struct{}

func (d defaultLogger) Infof(fmt string, args ...interface{}) {
	log.Printf("[cef] "+fmt, args...)
}
func (d defaultLogger) Warnf(fmt string, args ...interface{}) {
	log.Printf("[cef] "+fmt, args...)
}
func (d defaultLogger) Errorf(fmt string, args ...interface{}) {
	log.Printf("[cef] "+fmt, args...)
}
func (d defaultLogger) Panicf(fmt string, args ...interface{}) {
	log.Panicf("[cef] "+fmt, args...)
}

var Logger *log.Logger = log.New(os.Stdout, "[cef] ", log.Lshortfile)

var _MainArgs *C.struct__cef_main_args_t

// Sandbox is disabled. Including the "cef_sandbox.lib"
// library results in lots of GCC warnings/errors. It is
// compatible only with VS 2010. It would be required to
// build it using GCC. Add -lcef_sandbox to LDFLAGS.
// capi doesn't expose sandbox functions, you need do add
// these before import "C":
// void* cef_sandbox_info_create();
// void cef_sandbox_info_destroy(void* sandbox_info);
var _SandboxInfo unsafe.Pointer

type Settings struct {
	CachePath        string
	LogSeverity      int
	LogFile          string
	ResourcesDirPath string
	LocalesDirPath   string
}

type CefState int

var (
	STATE_DEFAULT  CefState = 0
	STATE_ENABLED  CefState = 1
	STATE_DISABLED CefState = 2
)

type BrowserSettings struct {
	StandardFontFamily          string
	FixedFontFamily             string
	SerifFontFamily             string
	SansSerifFontFamily         string
	CursiveFontFamily           string
	FantasyFontFamily           string
	DefaultFontSize             int
	DefaultFixedFontSize        int
	MinimumFontSize             int
	MinimumLogicalFontSize      int
	DefaultEncoding             string
	RemoteFonts                 CefState
	Javascript                  CefState
	JavascriptOpenWindows       CefState
	JavascriptCloseWindows      CefState
	JavascriptAccessClipboard   CefState
	JavascriptDomPaste          CefState
	CaretBrowsing               CefState
	Java                        CefState
	Plugins                     CefState
	UniversalAccessFromFileUrls CefState
	FileAccessFromFileUrls      CefState
	WebSecurity                 CefState
	ImageLoading                CefState
	ImageShrinkStandaloneToFit  CefState
	TextAreaResize              CefState
	TabToLinks                  CefState
	LocalStorage                CefState
	Databases                   CefState
	ApplicationCache            CefState
	Webgl                       CefState
	BackgroundColor             uint32
}

const (
	LOGSEVERITY_DEFAULT = C.LOGSEVERITY_DEFAULT
	LOGSEVERITY_VERBOSE = C.LOGSEVERITY_VERBOSE
	LOGSEVERITY_INFO    = C.LOGSEVERITY_INFO
	LOGSEVERITY_WARNING = C.LOGSEVERITY_WARNING
	LOGSEVERITY_ERROR   = C.LOGSEVERITY_ERROR
	// LOGSEVERITY_ERROR_REPORT = C.LOGSEVERITY_ERROR_REPORT
	LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE
)

func SetLogger(logger *log.Logger) {
	Logger = logger
}

func _InitializeGlobalCStructures() {
	_MainArgs = (*C.struct__cef_main_args_t)(
		C.calloc(1, C.sizeof_struct__cef_main_args_t))
	CreateRef(unsafe.Pointer(_MainArgs), "MainArgs")
}

func ExecuteProcess(programHandle unsafe.Pointer, appHandle AppHandler) int {
	Logger.Println("ExecuteProcess, args=", os.Args)

	_InitializeGlobalCStructures()
	FillMainArgs(_MainArgs, programHandle)

	// Sandbox info needs to be passed to both cef_execute_process()
	// and cef_initialize().
	// OFF: _SandboxInfo = C.cef_sandbox_info_create()
	go_AddRef(unsafe.Pointer(appHandle.GetAppHandlerT().CStruct))
	go_AddRef(unsafe.Pointer(_MainArgs))
	go_AddRef(unsafe.Pointer(_SandboxInfo))
	var exitCode C.int = C.cef_execute_process(_MainArgs, appHandle.GetAppHandlerT().CStruct, _SandboxInfo)
	if exitCode >= 0 {
		os.Exit(int(exitCode))
	}
	return int(exitCode)
}

func Initialize(settings Settings, appHandler AppHandler) int {
	Logger.Println("Initialize")

	if _MainArgs == nil {
		// _MainArgs structure is initialized and filled in ExecuteProcess.
		// If cef_execute_process is not called, and there is a call
		// to cef_initialize, then it would result in creation of infinite
		// number of processes. See Issue 1199 in CEF:
		// https://code.google.com/p/chromiumembedded/issues/detail?id=1199
		Logger.Println("ERROR: missing a call to ExecuteProcess")
		return 0
	}

	// Initialize cef_settings_t structure.
	var cefSettings *C.struct__cef_settings_t
	cefSettings = (*C.struct__cef_settings_t)(
		C.calloc(1, C.sizeof_struct__cef_settings_t))
	cefSettings.size = C.sizeof_struct__cef_settings_t

	var cachePath *C.char = C.CString(settings.CachePath)
	defer C.free(unsafe.Pointer(cachePath))
	C.cef_string_from_utf8(cachePath, C.strlen(cachePath),
		&cefSettings.cache_path)

	// log_severity
	// ------------
	cefSettings.log_severity =
		(C.cef_log_severity_t)(C.int(settings.LogSeverity))

	var logFile *C.char = C.CString(settings.LogFile)
	defer C.free(unsafe.Pointer(logFile))
	C.cef_string_from_utf8(logFile, C.strlen(logFile),
		&cefSettings.log_file)

	// resources_dir_path
	// ------------------
	if settings.ResourcesDirPath == "" && runtime.GOOS != "darwin" {
		// Setting this path is required for the tests to run fine.
		cwd, _ := os.Getwd()
		settings.ResourcesDirPath = cwd
	}
	var resourcesDirPath *C.char = C.CString(settings.ResourcesDirPath)
	defer C.free(unsafe.Pointer(resourcesDirPath))
	C.cef_string_from_utf8(resourcesDirPath, C.strlen(resourcesDirPath),
		&cefSettings.resources_dir_path)

	// locales_dir_path
	// ----------------
	if settings.LocalesDirPath == "" && runtime.GOOS != "darwin" {
		// Setting this path is required for the tests to run fine.
		cwd, _ := os.Getwd()
		settings.LocalesDirPath = cwd + "/locales"
	}
	var localesDirPath *C.char = C.CString(settings.LocalesDirPath)
	defer C.free(unsafe.Pointer(localesDirPath))
	C.cef_string_from_utf8(localesDirPath, C.strlen(localesDirPath),
		&cefSettings.locales_dir_path)

	// no_sandbox
	// ----------
	cefSettings.no_sandbox = C.int(1)
	go_AddRef(unsafe.Pointer(appHandler.GetAppHandlerT().CStruct))
	go_AddRef(unsafe.Pointer(_MainArgs))
	go_AddRef(unsafe.Pointer(_SandboxInfo))
	// TODO : Figure out why second argument must be nil
	ret := C.cef_initialize(_MainArgs, cefSettings, nil, _SandboxInfo)
	return int(ret)
}

func CreateBrowser(hwnd unsafe.Pointer, clientHandler ClientHandler, browserSettings BrowserSettings,
	url string) {

	// Initialize cef_window_info_t structure.
	var windowInfo *C.cef_window_info_t
	windowInfo = (*C.cef_window_info_t)(
		C.calloc(1, C.sizeof_cef_window_info_t))
	FillWindowInfo(windowInfo, hwnd)

	// url
	var cefUrl *C.cef_string_t
	cefUrl = (*C.cef_string_t)(
		C.calloc(1, C.sizeof_cef_string_t))
	var charUrl *C.char = C.CString(url)
	defer C.free(unsafe.Pointer(charUrl))
	C.cef_string_from_utf8(charUrl, C.strlen(charUrl), C.cefStringCastToCefString16(cefUrl))

	// Initialize cef_browser_settings_t structure.
	cefBrowserSettings := browserSettings.toC()

	// Do not create the browser synchronously using the
	// cef_browser_host_create_browser_sync() function, as
	// it is unreliable. Instead obtain browser object in
	// life_span_handler::on_after_created. In that callback
	// keep CEF browser objects in a global map (cef window
	// handle -> cef browser) and introduce
	// a GetBrowserByWindowHandle() function. This function
	// will first guess the CEF window handle using for example
	// WinAPI functions and then search the global map of cef
	// browser objects.
	go_AddRef(unsafe.Pointer(clientHandler.GetClientHandlerT().CStruct))
	result := C.cef_browser_host_create_browser(
		windowInfo,
		clientHandler.GetClientHandlerT().CStruct,
		cefUrl,
		cefBrowserSettings,
		nil,
	)
	// TODO : this is bad
	_ = result
}

func RunMessageLoop() {
	Logger.Println("RunMessageLoop")
	C.cef_run_message_loop()
}

func QuitMessageLoop() {
	Logger.Println("QuitMessageLoop")
	C.cef_quit_message_loop()
}

func Shutdown() {
	Logger.Println("Shutdown")
	C.cef_shutdown()
	// OFF: cef_sandbox_info_destroy(_SandboxInfo)
}

func (b BrowserSettings) toC() *C.struct__cef_browser_settings_t {
	var cefBrowserSettings *C.struct__cef_browser_settings_t
	cefBrowserSettings = (*C.struct__cef_browser_settings_t)(
		C.calloc(1, C.sizeof_struct__cef_browser_settings_t))
	cefBrowserSettings.size = C.sizeof_struct__cef_browser_settings_t

	go_AddRef(unsafe.Pointer(cefBrowserSettings))

	if b.StandardFontFamily != "" {
		toCefStringCopy(b.StandardFontFamily, &cefBrowserSettings.standard_font_family)
	}
	if b.FixedFontFamily != "" {
		toCefStringCopy(b.FixedFontFamily, &cefBrowserSettings.fixed_font_family)
	}
	if b.SerifFontFamily != "" {
		toCefStringCopy(b.SerifFontFamily, &cefBrowserSettings.serif_font_family)
	}
	if b.SansSerifFontFamily != "" {
		toCefStringCopy(b.SansSerifFontFamily, &cefBrowserSettings.sans_serif_font_family)
	}
	if b.CursiveFontFamily != "" {
		toCefStringCopy(b.CursiveFontFamily, &cefBrowserSettings.cursive_font_family)
	}
	if b.FantasyFontFamily != "" {
		toCefStringCopy(b.FantasyFontFamily, &cefBrowserSettings.fantasy_font_family)
	}
	cefBrowserSettings.default_font_size = C.int(b.DefaultFontSize)
	cefBrowserSettings.default_fixed_font_size = C.int(b.DefaultFixedFontSize)
	cefBrowserSettings.minimum_font_size = C.int(b.MinimumFontSize)
	cefBrowserSettings.minimum_logical_font_size = C.int(b.MinimumLogicalFontSize)
	if b.DefaultEncoding != "" {
		toCefStringCopy(b.DefaultEncoding, &cefBrowserSettings.default_encoding)
	}
	cefBrowserSettings.remote_fonts = C.cef_state_t(b.RemoteFonts)
	cefBrowserSettings.javascript = C.cef_state_t(b.Javascript)
	cefBrowserSettings.javascript_open_windows = C.cef_state_t(b.JavascriptOpenWindows)
	cefBrowserSettings.javascript_close_windows = C.cef_state_t(b.JavascriptCloseWindows)
	cefBrowserSettings.javascript_access_clipboard = C.cef_state_t(b.JavascriptAccessClipboard)
	cefBrowserSettings.javascript_dom_paste = C.cef_state_t(b.JavascriptDomPaste)
	cefBrowserSettings.caret_browsing = C.cef_state_t(b.CaretBrowsing)
	cefBrowserSettings.java = C.cef_state_t(b.Java)
	cefBrowserSettings.plugins = C.cef_state_t(b.Plugins)
	cefBrowserSettings.universal_access_from_file_urls = C.cef_state_t(b.UniversalAccessFromFileUrls)
	cefBrowserSettings.file_access_from_file_urls = C.cef_state_t(b.FileAccessFromFileUrls)
	cefBrowserSettings.web_security = C.cef_state_t(b.WebSecurity)
	cefBrowserSettings.image_loading = C.cef_state_t(b.ImageLoading)
	cefBrowserSettings.image_shrink_standalone_to_fit = C.cef_state_t(b.ImageShrinkStandaloneToFit)
	cefBrowserSettings.text_area_resize = C.cef_state_t(b.TextAreaResize)
	cefBrowserSettings.tab_to_links = C.cef_state_t(b.TabToLinks)
	cefBrowserSettings.local_storage = C.cef_state_t(b.LocalStorage)
	cefBrowserSettings.databases = C.cef_state_t(b.Databases)
	cefBrowserSettings.application_cache = C.cef_state_t(b.ApplicationCache)
	cefBrowserSettings.webgl = C.cef_state_t(b.Webgl)
	cefBrowserSettings.background_color = C.cef_color_t(b.BackgroundColor)
	return cefBrowserSettings
}

func toCefStringCopy(s string, out *C.cef_string_t) {
	var asC *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(asC))
	C.cef_string_from_utf8(
		asC,
		C.strlen(asC),
		C.cefStringCastToCefString16(out),
	)
}

func extractCefMultiMap(cefMapPointer C.cef_string_multimap_t) map[string][]string {
	numKeys := C.cef_string_multimap_size(cefMapPointer)
	goMap := make(map[string][]string)
	for i := 0; i < int(numKeys); i++ {
		var key *C.cef_string_utf16_t = C.cef_string_userfree_utf16_alloc()
		C.cef_string_multimap_key(cefMapPointer, C.int(i), C.cefString16CastToCefString(key))
		charKeyUtf8 := C.cefStringToUtf8(C.cefString16CastToCefString(key))
		goKey := C.GoString(charKeyUtf8.str)
		if _, ok := goMap[goKey]; ok {
			continue
		}
		numValsForKey := C.cef_string_multimap_find_count(cefMapPointer, C.cefString16CastToCefString(key))

		if numValsForKey >= 0 {
			goVals := make([]string, numValsForKey)
			for k := 0; k < int(numValsForKey); k++ {
				var val *C.cef_string_utf16_t = C.cef_string_userfree_utf16_alloc()
				C.cef_string_multimap_enumerate(cefMapPointer,
					C.cefString16CastToCefString(key), C.int(k), C.cefString16CastToCefString(val))
				charValUtf8 := C.cefStringToUtf8(C.cefString16CastToCefString(val))
				goVals[k] = C.GoString(charValUtf8.str)
				C.cef_string_userfree_utf8_free(charValUtf8)
				C.cef_string_userfree_utf16_free(val)
			}
			goMap[goKey] = goVals
		}
		C.cef_string_userfree_utf8_free(charKeyUtf8)
		C.cef_string_userfree_utf16_free(key)
	}
	return goMap
}
