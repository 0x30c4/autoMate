package configparser

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Step struct {
	Name      string `yaml:"name"`
	XPath     string `yaml:"xpath"`
	Value     string `yaml:"value,omitempty"`
	WaitFor   int    `yaml:"wait_for,omitempty"`
	WaitUntil bool   `yaml:"wait_until,omitempty"`
	Select    bool   `yaml:"select,omitempty"`
	Fill      bool   `yaml:"fill,omitempty"`
	Click     bool   `yaml:"click,omitempty"`
}

// StepsContainer holds the list of steps
type StepsContainer struct {
	Steps []Step `yaml:"steps"`
}

func ParserSteps(fileName string) (StepsContainer, error) {
	var container StepsContainer

	file, err := os.Open(fileName)
	if err != nil {
		return container, err
	}

	yamlData, err := io.ReadAll(file)
	if err != nil {
		return container, err
	}

	err = yaml.Unmarshal(yamlData, &container)
	if err != nil {
		return container, err
	}

	return container, nil
}
