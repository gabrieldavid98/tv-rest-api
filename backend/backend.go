package backend

import jsoniter "github.com/json-iterator/go"

// Backend represents a basic backend behavior
type Backend interface {
	Start()
}

// JSON encode/decoder config
var JSON = jsoniter.ConfigFastest
