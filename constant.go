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
	BackgroundColor             int
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
		SingleProcess:               0,
		NoSandbox:                   0,
		BrowserSubprocessPath:       "",
		MultiThreadedMessageLoop:    0,
		WindowlessRenderingEnabled:  1,
		CommandLineArgsDisabled:     0,
		CachePath:                   "",
		PersistSessionCookies:       1,
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
		UncaughtExceptionStackSize:  0,
		ContextSafetyImplementation: 0,
		IgnoreCertificateErrors:     0,
		BackgroundColor:             -1,
		constructed_by_NewSetting:   true,
	}
}

func SettingsFromC(cefSettings *C.struct__cef_settings_t) Settings {
	ret := NewSettings()

	ret.SingleProcess = int(cefSettings.single_process)
	ret.NoSandbox = int(cefSettings.no_sandbox)
	ret.BrowserSubprocessPath = fromCefStringCopy(&cefSettings.browser_subprocess_path)
	ret.MultiThreadedMessageLoop = int(cefSettings.multi_threaded_message_loop)
	ret.WindowlessRenderingEnabled = int(cefSettings.windowless_rendering_enabled)
	ret.CommandLineArgsDisabled = int(cefSettings.command_line_args_disabled)
	ret.CachePath = fromCefStringCopy(&cefSettings.cache_path)
	ret.PersistSessionCookies = int(cefSettings.persist_session_cookies)
	ret.CachePath = fromCefStringCopy(&cefSettings.cache_path)
	ret.UserAgent = fromCefStringCopy(&cefSettings.user_agent)
	ret.ProductVersion = fromCefStringCopy(&cefSettings.product_version)
	ret.Locale = fromCefStringCopy(&cefSettings.locale)
	ret.LogFile = fromCefStringCopy(&cefSettings.log_file)
	ret.LogSeverity = int(cefSettings.log_severity)
	ret.JavascriptFlags = fromCefStringCopy(&cefSettings.javascript_flags)
	ret.ResourcesDirPath = fromCefStringCopy(&cefSettings.resources_dir_path)
	ret.LocalesDirPath = fromCefStringCopy(&cefSettings.locales_dir_path)

	ret.PackLoadingDisabled = int(cefSettings.pack_loading_disabled)
	ret.RemoteDebuggingPort = int(cefSettings.remote_debugging_port)
	ret.UncaughtExceptionStackSize = int(cefSettings.uncaught_exception_stack_size)
	ret.ContextSafetyImplementation = int(cefSettings.context_safety_implementation)
	ret.IgnoreCertificateErrors = int(cefSettings.ignore_certificate_errors)
	ret.PackLoadingDisabled = int(cefSettings.pack_loading_disabled)
	ret.BackgroundColor = int(cefSettings.background_color)

	return ret
}

type CefState int

var (
	STATE_DEFAULT  CefState = 0
	STATE_ENABLED  CefState = 1
	STATE_DISABLED CefState = 2
)

type BrowserSettings struct {
	StandardFontFamily               string
	FixedFontFamily                  string
	SerifFontFamily                  string
	SansSerifFontFamily              string
	CursiveFontFamily                string
	FantasyFontFamily                string
	DefaultFontSize                  int
	DefaultFixedFontSize             int
	MinimumFontSize                  int
	MinimumLogicalFontSize           int
	DefaultEncoding                  string
	RemoteFonts                      CefState
	Javascript                       CefState
	JavascriptOpenWindows            CefState
	JavascriptCloseWindows           CefState
	JavascriptAccessClipboard        CefState
	JavascriptDomPaste               CefState
	CaretBrowsing                    CefState
	Java                             CefState
	Plugins                          CefState
	UniversalAccessFromFileUrls      CefState
	FileAccessFromFileUrls           CefState
	WebSecurity                      CefState
	ImageLoading                     CefState
	ImageShrinkStandaloneToFit       CefState
	TextAreaResize                   CefState
	TabToLinks                       CefState
	LocalStorage                     CefState
	Databases                        CefState
	ApplicationCache                 CefState
	Webgl                            CefState
	BackgroundColor                  int
	constructed_by_NewBrowserSetting bool
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
	if b.constructed_by_NewBrowserSetting == false {
		log.Warn("BrowserSettings should be constructed by NewBrowserSettings() function")
	}
	var cefBrowserSettings *C.struct__cef_browser_settings_t
	cefBrowserSettings = (*C.struct__cef_browser_settings_t)(
		C.calloc(1, C.sizeof_struct__cef_browser_settings_t))
	cefBrowserSettings.size = C.sizeof_struct__cef_browser_settings_t

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

func NewBrowserSettings() BrowserSettings {
	return BrowserSettings{
		StandardFontFamily:               "",
		FixedFontFamily:                  "",
		SerifFontFamily:                  "",
		SansSerifFontFamily:              "",
		CursiveFontFamily:                "",
		FantasyFontFamily:                "",
		DefaultFontSize:                  0,
		DefaultFixedFontSize:             0,
		MinimumFontSize:                  0,
		MinimumLogicalFontSize:           0,
		DefaultEncoding:                  "ISO-8859-1",
		RemoteFonts:                      STATE_DEFAULT,
		Javascript:                       STATE_DEFAULT,
		JavascriptOpenWindows:            STATE_DEFAULT,
		JavascriptCloseWindows:           STATE_DEFAULT,
		JavascriptAccessClipboard:        STATE_DEFAULT,
		JavascriptDomPaste:               STATE_DEFAULT,
		CaretBrowsing:                    STATE_DEFAULT,
		Java:                             STATE_DEFAULT,
		Plugins:                          STATE_DEFAULT,
		UniversalAccessFromFileUrls:      STATE_DEFAULT,
		FileAccessFromFileUrls:           STATE_DEFAULT,
		WebSecurity:                      STATE_DEFAULT,
		ImageLoading:                     STATE_DEFAULT,
		ImageShrinkStandaloneToFit:       STATE_DEFAULT,
		TextAreaResize:                   STATE_DEFAULT,
		TabToLinks:                       STATE_DEFAULT,
		LocalStorage:                     STATE_DEFAULT,
		Databases:                        STATE_DEFAULT,
		ApplicationCache:                 STATE_DEFAULT,
		Webgl:                            STATE_DEFAULT,
		BackgroundColor:                  -1,
		constructed_by_NewBrowserSetting: true,
	}
}

func BrowserSettingsFromC(cefBrowserSettings *C.struct__cef_browser_settings_t) BrowserSettings {
	b := NewBrowserSettings()

	b.StandardFontFamily = fromCefStringCopy(&cefBrowserSettings.standard_font_family)
	b.FixedFontFamily = fromCefStringCopy(&cefBrowserSettings.fixed_font_family)
	b.SerifFontFamily = fromCefStringCopy(&cefBrowserSettings.serif_font_family)
	b.SansSerifFontFamily = fromCefStringCopy(&cefBrowserSettings.sans_serif_font_family)
	b.CursiveFontFamily = fromCefStringCopy(&cefBrowserSettings.cursive_font_family)
	b.FantasyFontFamily = fromCefStringCopy(&cefBrowserSettings.fantasy_font_family)

	b.DefaultFontSize = int(cefBrowserSettings.default_font_size)
	b.DefaultFixedFontSize = int(cefBrowserSettings.default_fixed_font_size)
	b.MinimumFontSize = int(cefBrowserSettings.minimum_font_size)
	b.MinimumLogicalFontSize = int(cefBrowserSettings.minimum_logical_font_size)

	b.DefaultEncoding = fromCefStringCopy(&cefBrowserSettings.default_encoding)

	b.RemoteFonts = CefState(cefBrowserSettings.remote_fonts)
	b.Javascript = CefState(cefBrowserSettings.javascript)
	b.JavascriptOpenWindows = CefState(cefBrowserSettings.javascript_open_windows)
	b.JavascriptCloseWindows = CefState(cefBrowserSettings.javascript_close_windows)
	b.JavascriptAccessClipboard = CefState(cefBrowserSettings.javascript_close_windows)
	b.JavascriptDomPaste = CefState(cefBrowserSettings.javascript_dom_paste)
	b.CaretBrowsing = CefState(cefBrowserSettings.caret_browsing)
	b.Java = CefState(cefBrowserSettings.java)
	b.Plugins = CefState(cefBrowserSettings.plugins)
	b.UniversalAccessFromFileUrls = CefState(cefBrowserSettings.universal_access_from_file_urls)
	b.FileAccessFromFileUrls = CefState(cefBrowserSettings.universal_access_from_file_urls)
	b.WebSecurity = CefState(cefBrowserSettings.web_security)
	b.ImageLoading = CefState(cefBrowserSettings.image_loading)
	b.ImageShrinkStandaloneToFit = CefState(cefBrowserSettings.image_shrink_standalone_to_fit)
	b.TextAreaResize = CefState(cefBrowserSettings.text_area_resize)
	b.TabToLinks = CefState(cefBrowserSettings.tab_to_links)
	b.LocalStorage = CefState(cefBrowserSettings.local_storage)
	b.Databases = CefState(cefBrowserSettings.databases)
	b.ApplicationCache = CefState(cefBrowserSettings.application_cache)
	b.Webgl = CefState(cefBrowserSettings.webgl)
	b.BackgroundColor = int(cefBrowserSettings.background_color)
	return b
}

type CefErrorCode int

const (
	ERR_NONE                            CefErrorCode = 0
	ERR_FAILED                          CefErrorCode = -2
	ERR_ABORTED                         CefErrorCode = -3
	ERR_INVALID_ARGUMENT                CefErrorCode = -4
	ERR_INVALID_HANDLE                  CefErrorCode = -5
	ERR_FILE_NOT_FOUND                  CefErrorCode = -6
	ERR_TIMED_OUT                       CefErrorCode = -7
	ERR_FILE_TOO_BIG                    CefErrorCode = -8
	ERR_UNEXPECTED                      CefErrorCode = -9
	ERR_ACCESS_DENIED                   CefErrorCode = -10
	ERR_NOT_IMPLEMENTED                 CefErrorCode = -11
	ERR_CONNECTION_CLOSED               CefErrorCode = -100
	ERR_CONNECTION_RESET                CefErrorCode = -101
	ERR_CONNECTION_REFUSED              CefErrorCode = -102
	ERR_CONNECTION_ABORTED              CefErrorCode = -103
	ERR_CONNECTION_FAILED               CefErrorCode = -104
	ERR_NAME_NOT_RESOLVED               CefErrorCode = -105
	ERR_INTERNET_DISCONNECTED           CefErrorCode = -106
	ERR_SSL_PROTOCOL_ERROR              CefErrorCode = -107
	ERR_ADDRESS_INVALID                 CefErrorCode = -108
	ERR_ADDRESS_UNREACHABLE             CefErrorCode = -109
	ERR_SSL_CLIENT_AUTH_CERT_NEEDED     CefErrorCode = -110
	ERR_TUNNEL_CONNECTION_FAILED        CefErrorCode = -111
	ERR_NO_SSL_VERSIONS_ENABLED         CefErrorCode = -112
	ERR_SSL_VERSION_OR_CIPHER_MISMATCH  CefErrorCode = -113
	ERR_SSL_RENEGOTIATION_REQUESTED     CefErrorCode = -114
	ERR_CERT_COMMON_NAME_INVALID        CefErrorCode = -200
	ERR_CERT_DATE_INVALID               CefErrorCode = -201
	ERR_CERT_AUTHORITY_INVALID          CefErrorCode = -202
	ERR_CERT_CONTAINS_ERRORS            CefErrorCode = -203
	ERR_CERT_NO_REVOCATION_MECHANISM    CefErrorCode = -204
	ERR_CERT_UNABLE_TO_CHECK_REVOCATION CefErrorCode = -205
	ERR_CERT_REVOKED                    CefErrorCode = -206
	ERR_CERT_INVALID                    CefErrorCode = -207
	ERR_CERT_END                        CefErrorCode = -208
	ERR_INVALID_URL                     CefErrorCode = -300
	ERR_DISALLOWED_URL_SCHEME           CefErrorCode = -301
	ERR_UNKNOWN_URL_SCHEME              CefErrorCode = -302
	ERR_TOO_MANY_REDIRECTS              CefErrorCode = -310
	ERR_UNSAFE_REDIRECT                 CefErrorCode = -311
	ERR_UNSAFE_PORT                     CefErrorCode = -312
	ERR_INVALID_RESPONSE                CefErrorCode = -320
	ERR_INVALID_CHUNKED_ENCODING        CefErrorCode = -321
	ERR_METHOD_NOT_SUPPORTED            CefErrorCode = -322
	ERR_UNEXPECTED_PROXY_AUTH           CefErrorCode = -323
	ERR_EMPTY_RESPONSE                  CefErrorCode = -324
	ERR_RESPONSE_HEADERS_TOO_BIG        CefErrorCode = -325
	ERR_CACHE_MISS                      CefErrorCode = -400
	ERR_INSECURE_RESPONSE               CefErrorCode = -501
)

func toCefStringCopy(s string, out *C.cef_string_t) {
	var asC *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(asC))
	C.cef_string_from_utf8(asC, C.strlen(asC), C.cefStringCastToCefString16(out))
}

func fromCefStringCopy(src *C.cef_string_t) string {
	ret := C.GoString(C.cefStringtoChar(src))
	return ret
}
