package chrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"flag"
	"fmt"
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
	flag.String("no-sandbox", "foo", "-")

	gopath := os.Getenv("GOPATH")
	os.Chdir(fmt.Sprintf("%s/%s", gopath, "src/github/24hours/chrome/Release"))
	os.Exit(m.Run()) 
}

func TestBasic(t *testing.T) {
	assert.Equal(t, -1, ExecuteProcess(nil, nil), "ExecuteProcess should return -1")
}

// TODO : make sure setting actually work 

func TestBase(t *testing.T) {
	settings := NewSettings()
	settings.NoSandbox = 1
	// TODO : insert path here 
	assert.Equal(t, 1, Initialize(settings, nil), "Initialize")
	//chrome.Shutdown()
}