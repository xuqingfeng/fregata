package macos

import (
	"runtime"
	"testing"
)

func TestValidate(t *testing.T) {

	testConfig := NewConfig()
	err := testConfig.Validate()
	if runtime.GOOS != "darwin" && err == nil {
		t.Errorf("E! should return err got %s", err.Error())
	}
}
