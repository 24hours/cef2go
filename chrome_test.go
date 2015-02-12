package chrome_test

import (
	"github.com/24hours/chrome"
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"flag"
)

func TestMain(m *testing.M) { 
	// add some flags otherwise test will fail when cef fork it self 
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


	gopath := os.Getenv("GOPATH")
	os.Chdir(fmt.Sprintf("%s/%s", gopath, "src/github/24hours/chrome/Release"))
	os.Exit(m.Run()) 
}

func TestBasic(t *testing.T) {
	assert.Equal(t, -1, chrome.ExecuteProcess(nil, nil), "ExecuteProcess should return -1")
}

func TestBase(t *testing.T) {
	settings := chrome.NewSettings()
	assert.Equal(t, 1, chrome.Initialize(settings, nil), "Initialize")
	chrome.Shutdown()
}