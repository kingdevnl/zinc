package v1

import "strings"

func NewESInfo() *ESInfo {
	return &ESInfo{
		Name:        "zinc",
		ClusterName: "N/A",
		ClusterUUID: "N/A",
		Version: ESInfoVersion{
			Number:                    strings.TrimLeft(Version, "v"),
			BuildFlavor:               "default",
			BuildHash:                 CommitHash,
			BuildDate:                 BuildDate,
			BuildSnapshot:             false,
			LuceneVersion:             "N/A",
			MinimumWireVersion:        "N/A",
			MinimumIndexCompatibility: "N/A",
		},
		Tagline: "You Know, for Search",
	}
}

func NewESLicense() *ESLicense {
	return &ESLicense{
		Status: "active",
	}
}

func NewESXPack() *ESXPack {
	return &ESXPack{
		Build:    make(map[string]bool),
		Features: make(map[string]bool),
		License: ESLicense{
			Status: "active",
		},
	}
}

type ESInfo struct {
	Name        string        `json:"name"`
	ClusterName string        `json:"cluster_name"`
	ClusterUUID string        `json:"cluster_uuid"`
	Version     ESInfoVersion `json:"version"`
	Tagline     string        `json:"tagline"`
}

type ESInfoVersion struct {
	Number                    string `json:"number"`
	BuildFlavor               string `json:"build_flavor"`
	BuildHash                 string `json:"build_hash"`
	BuildDate                 string `json:"build_date"`
	BuildSnapshot             bool   `json:"build_snapshot"`
	LuceneVersion             string `json:"lucene_version"`
	MinimumWireVersion        string `json:"minimum_wire_version"`
	MinimumIndexCompatibility string `json:"minimum_index_compatibility"`
}

type ESLicense struct {
	Status string `json:"status"`
}

type ESXPack struct {
	Build    map[string]bool `json:"build"`
	Features map[string]bool `json:"features"`
	License  ESLicense       `json:"license"`
}
