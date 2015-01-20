package chrome

//#include <string.h>
//#include <stdlib.h>
//#include "include/capi/cef_app_capi.h"
//#include "cef_base.h"
import "C"
import (
	log "github.com/cihub/seelog"
	"os"
	"unsafe"
)

type Settings struct {
	SingleProcess               int
	NoSandbox                   int
	BrowserSubprocessPath       string
	MultiThreadedMessageLoop    int
	WindowlessRenderingEnabled  int
	CommandLineArgsDisabled     int
	CachePath                   string
	PersistSessionCookies       int
	UserAgent                   string
	ProductVersion              string
	Locale                      string
	LogFile                     string
	LogSeverity                 int
	JavascriptFlags             string
	ResourcesDirPath            string
	LocalesDirPath              string
	PackLoadingDisabled         int
	RemoteDebuggingPort         int
	UncaughtExceptionStackSize  int
	ContextSafetyImplementation int
	IgnoreCertificateErrors     int
	BackgroundColor             uint32
	constructed_by_NewSetting   bool
}

func (s Settings) toC() *C.struct__cef_settings_t {
	if s.constructed_by_NewSetting == false {
		log.Warn("Settings should be constructed by NewSettings() function")
	}

	var cefSettings *C.struct__cef_settings_t
	cefSettings = (*C.struct__cef_settings_t)(C.calloc(1, C.sizeof_struct__cef_settings_t))
	cefSettings.size = C.sizeof_struct__cef_settings_t

	cefSettings.single_process = C.int(s.SingleProcess)
	cefSettings.no_sandbox = C.int(s.NoSandbox)
	if s.BrowserSubprocessPath != "" {
		toCefStringCopy(s.BrowserSubprocessPath, &cefSettings.browser_subprocess_path)
	}

	cefSettings.multi_threaded_message_loop = C.int(s.MultiThreadedMessageLoop)
	cefSettings.windowless_rendering_enabled = C.int(s.WindowlessRenderingEnabled)
	cefSettings.command_line_args_disabled = C.int(s.CommandLineArgsDisabled)

	if s.CachePath != "" {
		toCefStringCopy(s.CachePath, &cefSettings.cache_path)
	}

	cefSettings.persist_session_cookies = C.int(s.PersistSessionCookies)

	if s.UserAgent != "" {
		toCefStringCopy(s.UserAgent, &cefSettings.user_agent)
	}

	if s.ProductVersion != "" {
		toCefStringCopy(s.ProductVersion, &cefSettings.product_version)
	}

	if s.Locale != "" {
		toCefStringCopy(s.Locale, &cefSettings.locale)
	}

	if s.LogFile != "" {
		toCefStringCopy(s.LogFile, &cefSettings.log_file)
	}

	cefSettings.log_severity = C.cef_log_severity_t(s.LogSeverity)

	if s.JavascriptFlags != "" {
		toCefStringCopy(s.JavascriptFlags, &cefSettings.javascript_flags)
	}

	if s.ResourcesDirPath != "" {
		toCefStringCopy(s.ResourcesDirPath, &cefSettings.resources_dir_path)
	}

	if s.LocalesDirPath != "" {
		toCefStringCopy(s.LocalesDirPath, &cefSettings.locales_dir_path)
	}

	cefSettings.pack_loading_disabled = C.int(s.PackLoadingDisabled)
	cefSettings.remote_debugging_port = C.int(s.RemoteDebuggingPort)
	cefSettings.uncaught_exception_stack_size = C.int(s.UncaughtExceptionStackSize)
	cefSettings.context_safety_implementation = C.int(s.ContextSafetyImplementation)
	cefSettings.ignore_certificate_errors = C.int(s.IgnoreCertificateErrors)
	cefSettings.pack_loading_disabled = C.int(s.PackLoadingDisabled)
	cefSettings.background_color = C.cef_color_t(s.BackgroundColor)

	return cefSettings
}

func NewSettings() Settings {
	cwd, _ := os.Getwd()
	// TODO : enable sandbox this when is is implemented

	return Settings{
		SingleProcess:               1,
		NoSandbox:                   1,
		BrowserSubprocessPath:       "",
		MultiThreadedMessageLoop:    0,
		WindowlessRenderingEnabled:  0,
		CommandLineArgsDisabled:     0,
		CachePath:                   "",
		PersistSessionCookies:       0,
		UserAgent:                   "cef2go",
		ProductVersion:              "0.1",
		Locale:                      "en-US",
		LogFile:                     "",
		LogSeverity:                 LOGSEVERITY_DEFAULT,
		JavascriptFlags:             "",
		ResourcesDirPath:            cwd,
		LocalesDirPath:              cwd + "/locales",
		PackLoadingDisabled:         0,
		RemoteDebuggingPort:         0,
		UncaughtExceptionStackSize:  10,
		ContextSafetyImplementation: 0,
		IgnoreCertificateErrors:     0,
		BackgroundColor:             16777215,
		constructed_by_NewSetting:   true,
	}
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
	LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE
)

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
	C.cef_string_from_utf8(asC, C.strlen(asC), C.cefStringCastToCefString16(out))
}
