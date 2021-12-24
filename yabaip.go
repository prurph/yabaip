package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/prurph/yabaip/internal/windowtype"
	"gopkg.in/yaml.v2"
)

var space = `
label: my_space
windows: 
  -
    command:
      - open
      - /Users/prescott
    window_type: managed
  -
    command:
      - open
      - /Applications/TextEdit.app
    layout:
      window_type: floating
`

type Window struct {
	Command []string `yaml:"command"`
	Layout  struct {
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
		log.Fatal(err)
		os.Exit(1)
	}
	var errbuf strings.Builder
	cmd := exec.Command("yabai", "-m", "space", "--create")
	cmd.Stderr = &errbuf
	_, err = cmd.Output()
	if err != nil {
		log.Fatalf("command '%s' failed: %s", cmd, cmd.Stderr)
		log.Fatal(err)
		os.Exit(1)
	}
	err = exec.Command("yabai", "-m", "space", "--focus", "last").Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = exec.Command("yabai", "-m", "space", "--label", s.Label).Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, w := range s.Windows {
		err = exec.Command(w.Command[0], w.Command[1:]...).Run()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
