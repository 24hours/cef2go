package chrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"flag"
	"fmt"
)

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
	os.Exit(m.Run()) 
}

func TestSetting(t *testing.T){
	cwd, _ := os.Getwd()
	settings := NewSettings()
	assert.Equal(t, 0, settings.NoSandbox)
	assert.Equal(t, cwd, settings.ResourcesDirPath)
	assert.Equal(t,  cwd + "/locales", settings.LocalesDirPath)
}

func TestBasic(t *testing.T) {
	cwd, _ := os.Getwd()
	assert.Equal(t, -1, ExecuteProcess(nil, nil), "ExecuteProcess must return -1")
	settings := NewSettings()
	settings.NoSandbox = 1
	settings.ResourcesDirPath = cwd + "/Release"
	settings.LocalesDirPath = cwd + "/Release/locales"
	//settings.WindowlessRenderingEnabled = 1
	assert.Equal(t, 1, Initialize(settings, nil), "Initialize must return 1")
	go RunMessageLoop()
	
	window := NewWindowInfo(800, 640)
	//window.WindowlessRendering = 1
	CreateBrowser(window, nil, BrowserSettings{}, "file:///home/u24/Desktop/index.html")
	
	QuitMessageLoop()
	Shutdown() // the test consider a success if it didn't crash immediately
}
