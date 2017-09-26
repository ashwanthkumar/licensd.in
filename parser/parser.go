package parser

import "bufio"

type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

type Dependency struct {
	Name        string     `json:"name"`
	Version     string     `json:"version"`
	BuildMatrix string     `json:"matrix"`
	CIURL       string     `json:"ci_url"`
	Licenses    []*License `json:"licenses,omitempty"`
}

// Parser parses a given input file generated for a build tool so licensed can understand it
type Parser interface {
	// Parse the license file generated by a plugin or by a tool like license_finder
	// and convert that to an []Dependency objects.
	Parse(bufio.Scanner) ([]*Dependency, error)
}
