package app

import (
	"encoding/json"
	"log"
)

const (
	Name        = "clustermom-agent"
	Version     = "1.0.0"
	Description = "ClusterMom Agent"
	Repository  = "http://gitlab.xtc.home/xtc/clustermom-agent"
)

var (
	GitTag    string
	GitCommit string
	GitBranch string
	BuildTime string
)

func BuildInfo() string {
	info := map[string]string{}
	info["Version"] = Version
	info["BuildTime"] = BuildTime
	info["GitCommit"] = GitCommit
	info["GitBranch"] = GitBranch
	info["GitTag"] = GitTag
	j, err := json.Marshal(info)
	if err != nil {
		log.Panic(err)
	}
	return string(j)
}
