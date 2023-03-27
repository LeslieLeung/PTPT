package prompt

type Bundle struct {
	Version string   `yaml:"version"`
	Prompts []Prompt `yaml:"prompts"`
}

type Prompt struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	System      string `yaml:"system"`
}

var Lib map[string]Prompt
