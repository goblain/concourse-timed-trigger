package resource

import (
	"os"
	"log"
	"encoding/json"
)

func Check() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	var request CheckRequest
	err := dec.Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}
	var refs []Version
	refs = append(refs, Version{Ref: "abc"})
	refs = append(refs, Version{Ref: "def"})
	enc.Encode(&refs)
}

func In() {
	//dec := json.NewDecoder(os.Stdin)
	//enc := json.NewEncoder(os.Stdout)
	//var request InRequest
	//err := dec.Decode(&request)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
}

func Out() {
	dec := json.NewDecoder(os.Stdin)
	//enc := json.NewEncoder(os.Stdout)
	var request OutRequest
	err := dec.Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}
}

type CheckRequest struct {
	Source Source `json:"source"`
	Version Version `json:"version"`
}

type OutRequest struct {
	Source  Source    `json:"source"`
	Params  OutParams `json:"params"`
}

// Params define how to configure added trigger
// After - set timer to trigger new resource version after given period (values like 1s, 4h, 2d)
// Comment - register provided comment with given trigger
// Reset - if previous timer was set reset the timer to the new value
type OutParams struct {
	After   string `json:"after"`
	Comment string `json:"comment"`
	Reset   bool   `json:"reset"`
}

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

type OutResponse struct {
	Version  Version `json:"version"`
	Metadata string  `json:"metadata"`
}
