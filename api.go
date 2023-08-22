package danger

import (
	"encoding/json"
	"fmt"

	dangerJs "github.com/luno/danger-go/danger-js"
)

// DSL wraps the DSL received from danger JS. This allows dangerfiles to only
// import the root danger package.
type DSL = dangerJs.DSL

type T struct {
	results Results
}

func New() *T {
	return &T{
		results: Results{
			Fails:     []Violation{},
			Messages:  []Violation{},
			Warnings:  []Violation{},
			Markdowns: []Violation{},
		},
	}
}

// Results returns the JSON marshalled from the messages, warnings, failures,
// and markdowns that was added so far.
func (s *T) Results() (string, error) {
	bb, err := json.Marshal(s.results)
	if err != nil {
		return "", fmt.Errorf("marshalling results: %w", err)
	}
	return string(bb), nil
}

// Message adds the message to the Danger table. The only difference between
// this and Warn is the emoji which shows in the table.
func (s *T) Message(message string, file string, line int) {
	s.results.Messages = append(s.results.Messages,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}

// Warn adds the message to the Danger table. The message highlights
// low-priority issues, but does not fail the build.
func (s *T) Warn(message string, file string, line int) {
	s.results.Warnings = append(s.results.Warnings,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}

// Fail a build, outputting a specific reason for failing into an HTML table.
func (s *T) Fail(message string, file string, line int) {
	s.results.Fails = append(s.results.Fails,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}

// Markdown adds the message as raw markdown into the Danger comment, under the
// table.
func (s *T) Markdown(message string, file string, line int) {
	s.results.Markdowns = append(s.results.Markdowns,
		Violation{
			Message: message,
			File:    file,
			Line:    line,
		})
}
