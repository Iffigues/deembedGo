package utils

import (
	"runtime/debug"
	"testing"
)

func TestGetBuildInfoSetting(t *testing.T) {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		t.Fatalf("debug.ReadBuildInfo() are false")
	}
	for _, val := range bi.Settings {
		if val.Value != GetBuildInfoSetting(bi, val.Key) {
			t.Fatalf("val.value != getBuildInfoSetting(bi, val.Key)")
		}
	}
}
