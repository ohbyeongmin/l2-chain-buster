package chainbuster

type Scenario struct {
	Users    int    `yaml:"users"`
	Duration int    `yaml:"duration"`
	RPS      int    `yaml:"rps,omitempty"`
	Bridge   Bridge `yaml:"bridge,omitempty"`
}

type Bridge struct {
	Users                 int `yaml:"users"`
	Interval              int `yaml:"interval"`
	Counts                int `yaml:"counts"`
	WithdrawProvedTime    int `yaml:"withraw_proved_time"`
	WithdrawFinalizedTime int `yaml:"withraw_finalized_time"`
}

type Scenarios struct {
	List []Scenario `yaml:"scenarios"`
}

func CheckScenarios(scenarios *Scenarios) error {
	return nil
}

func (ss *Scenarios) maxUsers() int {
	max := 0
	for _, s := range ss.List {
		total := s.Users + s.Bridge.Users
		if max < total {
			max = total
		}
	}
	return max
}
