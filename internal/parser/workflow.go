package parser

type Workflow struct {
	ID            string                          `yaml:"id"`
	Name          string                          `yaml:"name"`
	Triggers      WorkflowTriggers                `yaml:"on"`
	Args          map[string]string               `yaml:"args"`
	Params        []WorkflowParameter             `yaml:"params"`
	Tasks         map[string]WorkflowTask         `yaml:"tasks"`
	Notifications map[string]WorkflowNotification `yaml:"notifications"`
}

type WorkflowTriggers struct {
	Schedule  string   `yaml:"schedule"`
	Manual    bool     `yaml:"manual"`
	Workflows []string `yaml:"workflow"`
}

type WorkflowParameter struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type WorkflowTask struct {
	Name   string            `yaml:"name"`
	Uses   string            `yaml:"uses"`
	Config map[string]string `yaml:"with"`
	Args   map[string]string `yaml:"args"`
}

type WorkflowNotification struct {
	User     bool            `yaml:"user"`
	Workflow map[string]bool `yaml:"workflow"`
	Schedule bool            `yaml:"schedule"`
}
