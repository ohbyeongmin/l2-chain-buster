package chainbuster

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Scenario struct {
	Users    int    `yaml:"users"`
	Duration int    `yaml:"duration"`
	RPS      int    `yaml:"rps"`
	Bridge   Bridge `yaml:"bridge"`
}

type Bridge struct {
	Users                 int `yaml:"users"`
	Interval              int `yaml:"interval"`
	Counts                int `yaml:"counts"`
	WithdrawProvedTime    int `yaml:"withraw_proved_time"`
	WithdrawFinalizedTime int `yaml:"withraw_finalized_time"`
}

type Scenarios struct {
	Scenarios []Scenario `yaml:"scenarios"`
}

func NewScenarios(filename string) (*Scenarios, error) {
	var scens Scenarios
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(yamlFile, &scens); err != nil {
		return nil, err
	}

	return &scens, nil
}
