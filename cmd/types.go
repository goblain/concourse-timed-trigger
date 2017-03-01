package cmd

import (
	"encoding/json"
)

type BackendRedis struct {
	Host string `json:"host"`
}

type Source struct {
	Backend string       `json:"backend"`
	Redis   BackendRedis `json:"redis"`
}

type Version struct {
	Ref string `json:"ref"`
}

type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InResponse struct {
	Version  Version `json:"version"`
	Metadata string  `json:"metadata"`
}

type OutParams struct {
	After   string `json:"after"`
	Comment string `json:"comment"`
	Reset   bool   `json:"reset"`
}

type OutRequest struct {
	Request string
	Source  Source    `json:"source"`
	Params  OutParams `json:"params"`
}

func (self *OutRequest) Process() {
	json.Unmarshal([]byte(self.Request), &self)
}

type OutResponse struct {
	Version  Version `json:"version"`
	Metadata string  `json:"metadata"`
}
