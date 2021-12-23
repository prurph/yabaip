package main

import (
	"fmt"
	"os"

	windowtype "github.com/prurph/yabaip/internal"
	"gopkg.in/yaml.v2"
)

var space = `
label: my_space
windows:
  - title: my_window
    layout:
      window_type: managed
  - title: my_other_window
    layout:
      window_type: floating
`

type Window struct {
	Title  string `yaml:"title"`
	Layout struct {
		WindowType windowtype.WindowType `yaml:"window_type"`
	} `yaml:"layout"`
}

type Space struct {
	Label   string   `yaml:"label"`
	Windows []Window `yaml:"windows"`
}

func main() {
	s := Space{}
	err := yaml.Unmarshal([]byte(space), &s)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("--- t:\n%v\n\n", s)
}
