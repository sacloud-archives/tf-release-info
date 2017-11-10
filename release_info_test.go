package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReleaseInfo(t *testing.T) {

	expects := []struct {
		releaseInfo *releaseInfo
		preRelease  bool
	}{
		{
			releaseInfo: newReleaseInfo("v0.0.1", "v0.0.1"),
			preRelease:  false,
		},
		{
			releaseInfo: newReleaseInfo("v0.0.1", "v0.0.1-rc"),
			preRelease:  true,
		},
	}

	for _, expect := range expects {
		assert.Equal(t, expect.preRelease, expect.releaseInfo.PreRelease, "pluginVersion:%s", expect.releaseInfo.PluginVersion)
	}
}
