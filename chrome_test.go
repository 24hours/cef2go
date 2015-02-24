package chrome

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"runtime"
	"testing"
)

var cwd, test_dir string

func TestMain(m *testing.M) {
	// add some flags otherwise test will fail when cef fork itself
	flag.String("type", "foo", "-")
	flag.String("lang", "foo", "-")
	flag.String("locales-dir-path", "foo", "-")
	flag.String("resources-dir-path", "foo", "-")
	flag.String("user-agent", "foo", "-")
	flag.String("supports-dual-gpus", "foo", "-")
	flag.String("gpu-driver-bug-workarounds", "foo", "-")
	flag.String("disable-accelerated-video-decode", "foo", "-")
	flag.String("gpu-vendor-id", "foo", "-")
	flag.String("gpu-device-id", "foo", "-")
	flag.String("gpu-driver-vendor", "foo", "-")
	flag.String("gpu-driver-version", "foo", "-")
	flag.String("channel", "foo", "-")
	flag.String("no-sandbox", "foo", "-")

	gopath := os.Getenv("GOPATH")
	os.Chdir(fmt.Sprintf("%s/%s", gopath, "src/github/24hours/chrome/Release"))
	cwd, _ = os.Getwd()
	test_dir = "file://" + cwd + "/test"
	os.Exit(m.Run())
}

func TestSetting(t *testing.T) {
	settings := NewSettings()
	assert.Equal(t, 0, settings.NoSandbox)
	assert.Equal(t, cwd, settings.ResourcesDirPath)
	assert.Equal(t, cwd+"/locales", settings.LocalesDirPath)
}

func TestBasic(t *testing.T) {
	assert.Equal(t, -1, ExecuteProcess(nil, nil), "ExecuteProcess must return -1")
	settings := NewSettings()
	settings.NoSandbox = 1
	if runtime.GOOS != "darwin" {
		settings.ResourcesDirPath = cwd + "/Release"
		settings.LocalesDirPath = cwd + "/Release/locales"
	} else {
		settings.LocalesDirPath = ""
		settings.ResourcesDirPath = ""
	}
	settings.WindowlessRenderingEnabled = 1
	assert.Equal(t, 1, Initialize(settings, nil), "Initialize must return 1")
	go RunMessageLoop()
	browserBasic(t)
	if runtime.GOOS != "darwin" {
		QuitMessageLoop()
		Shutdown() // the test consider a success if it didn't crash immediately
	}
}

func browserBasic(t *testing.T) {
	window := NewWindowInfo(800, 640)
	window.WindowlessRendering = 1
	err := CreateBrowserAsync(window, nil, BrowserSettings{}, test_dir+"/index.html")
	assert.NotNil(t, err, "Creation must fail if renderhandler is not implemented")

	err = CreateBrowserAsync(window, &testClientHandler{}, BrowserSettings{}, test_dir+"/index.html")
	assert.Nil(t, err, "Creation must success with Render Handler")
}

type testClientHandler struct {
	BaseClientHandler
}

func (l *testClientHandler) OnAfterCreated(browser Browser) { fmt.Println("Created") }
