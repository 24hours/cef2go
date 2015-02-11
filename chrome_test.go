package chrome_test

import (
	"github.com/24hours/chrome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	assert.Equal(t, 1, 1, "1 == 1 should be always true")
}

func TestBase(t *testing.T) {
	assert.Equal(t, -1, chrome.ExecuteProcess(nil, nil), "ExecuteProcess should return -1")
	chrome.Shutdown()
}
