package chrome

//#include "include/capi/cef_app_capi.h"
import "C"

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
	LOGSEVERITY_DISABLE = C.LOGSEVERITY_DISABLE
)
