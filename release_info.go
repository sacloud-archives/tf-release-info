package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"regexp"
	"time"
)

type releaseInfo struct {
	TerraformVersion string
	ReleaseDate      time.Time
	PluginVersion    string
	PreRelease       bool
}

func newReleaseInfo(tfVersion string, pluginVersion string) *releaseInfo {

	info := &releaseInfo{
		TerraformVersion: tfVersion,
		PluginVersion:    pluginVersion,
		ReleaseDate:      time.Now(),
		PreRelease:       true,
	}

	if matched, err := regexp.Match(`^v\d?\.\d?\.\d?$`, []byte(info.PluginVersion)); err == nil && matched {
		info.PreRelease = false
	}

	return info
}

func readReleaseInfo(reader io.Reader) ([]*releaseInfo, error) {

	var releases []*releaseInfo
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return releases, err
	}
	err = json.Unmarshal(data, &releases)
	if err != nil {
		return releases, err
	}

	return releases, nil
}

func writeReleaseInfo(writer io.Writer, info []*releaseInfo) error {
	data, err := json.MarshalIndent(&info, "", "\t")
	if err != nil {
		return err
	}

	if _, err := writer.Write(data); err != nil {
		return err
	}
	return nil
}
