package main

import "os/exec"

type Event struct {
	id      string      `json:"id"`
	content string      `json:"content"`
	parent  string      `json:"parent"`
	extra   interface{} `json:"extra,omitempty"`
}
type TestResult struct {
	eventID string `json:"eventID"`
	// event   Event  `json:"eventID"`
	found  bool   `json:"eventID"`
	passed bool   `json:"passed"`
	line   string `json:"comment"`
	err    error  `json:"err"`
}

type CMD struct {
	stateID      string                   `json:"stateID"`
	homepath     string                   `json:"homepath"`
	cmd          string                   `json:"cmd"`
	filter       []string                 `json:"filter"`
	testing      bool                     `json:"testing"`
	test         Test                     `json:"test"`
	results      *map[string][]TestResult `json:"results"`
	dep          *State                   `json:"dep"`
	failed       *bool                    `json:"failed"`
	requireAlive bool                     `json:"requireAlive"`
	debug        bool                     `json:"debug"`
}

type State struct {
	id           string                   `json:"id"`
	finished     *bool                    `json:"finished"`
	awaiting     map[string]Await         `json:"awating"`
	testing      bool                     `json:"testing"`
	test         Test                     `json:"test"`
	results      *map[string][]TestResult `json:"results"`
	depending    *[]*State                `json:"depending"`
	cmd          *exec.Cmd                `json:"cmd"`
	failed       *bool                    `json:"failed"`
	requireAlive bool                     `json:"requireAlive"`
	lastLine     *string                  `json:"lastLine"`
	debug        bool                     `json:"debug"`
}

type Await struct {
	done *bool                     `json:"done"`
	pass *bool                     `json:"pass"`
	f    func(string) (bool, bool) `json:"f"`
	msg  string                    `json:"msg"`
}

type Test struct {
	expectedEvents   []string                                        `json:"expectedEvents"`
	unexpectedEvents []string                                        `json:"unexpectedEvents"`
	tests            map[string](func(string) (bool, string, error)) `json:"tests"`
	strict           bool                                            `json:"strict"`
}
