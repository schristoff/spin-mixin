package spin

import (
	"context"

	"get.porter.sh/porter/pkg/exec/builder"
	"gopkg.in/yaml.v2"
)

var _ builder.BuildableAction = Action{}
var _ builder.ExecutableAction = Action{}

type Action struct {
	Name  string
	Steps []Step // using UnmarshalYAML so that we don't need a custom type per action
}

func (m *Mixin) loadAction(ctx context.Context) (*Action, error) {
	var action Action
	err := builder.LoadAction(ctx, m.RuntimeConfig, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &action)
		return &action, err
	})
	return &action, err
}

// MarshalYAML converts the action back to a YAML representation
// install:
//
//	spin:
//	  ...
func (a Action) MarshalYAML() (interface{}, error) {
	return map[string]interface{}{a.Name: a.Steps}, nil
}

// MakeSteps builds a slice of Step for data to be unmarshaled into.
func (a Action) MakeSteps() interface{} {
	return &[]Step{}
}

// UnmarshalYAML takes any yaml in this form
// ACTION:
// - spin: ...
// and puts the steps into the Action.Steps field
func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, a)
	if err != nil {
		return err
	}

	for actionName, action := range results {
		a.Name = actionName
		for _, result := range action {
			step := result.(*[]Step)
			a.Steps = append(a.Steps, *step...)
		}
		break // There is only 1 action
	}
	return nil
}

type Step struct {
	Instruction `yaml:"spin"`
}

// Actions is a set of actions, and the steps, passed from Porter.
type Actions []Action

// UnmarshalYAML takes chunks of a porter.yaml file associated with this mixin
// and populates it on the current action set.
// install:
//
//	spin:
//	  ...
//	spin:
//	  ...
//
// upgrade:
//
//	spin:
//	  ...
func (a *Actions) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, Action{})
	if err != nil {
		return err
	}

	for actionName, action := range results {
		for _, result := range action {
			s := result.(*[]Step)
			*a = append(*a, Action{
				Name:  actionName,
				Steps: *s,
			})
		}
	}
	return nil
}

var _ builder.StepWithOutputs = Instruction{}

type Instruction struct {
	Name           string        `yaml:"name"`
	Description    string        `yaml:"description"`
	Arguments      []string      `yaml:"arguments,omitempty"`
	Outputs        []Output      `yaml:"outputs,omitempty"`
	SuppressOutput bool          `yaml:"suppress-output,omitempty"`
	Flags          builder.Flags `yaml:"flags,omitempty"`
	WorkingDir     string        `yaml:"workingDir, omitempty"`

	// Allow the user to ignore some errors
	// Adds the ignoreError functionality from the exec mixin
	// https://release-v1.porter.sh/mixins/exec/#ignore-error
	builder.IgnoreErrorHandler `yaml:"ignoreError,omitempty"`
	SpinFields                 `yaml:",inline"`
}

type SpinFields struct {
	FermyonCloud   bool   `yaml:"fermyonCloud"`
	PlatformURL    string `yaml:"platformURL"`
	HippoUsername  string `yaml:"hippoUsername"`
	HippoPassword  string `yaml:"hippoPassword"`
	Insecure       bool   `yaml:"insecure"`
	BindleUsername string `yaml:"bindleUsername"`
	BindleServer   string `yaml:"bindleServer"`
	BindlePassword string `yaml:"bindlePassword"`
}

func (s Instruction) GetCommand() string {
	return "spin"
}

func (s Instruction) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(s.Outputs))
	for i := range s.Outputs {
		outputs[i] = s.Outputs[i]
	}
	return outputs
}

var _ builder.OutputJsonPath = Output{}
var _ builder.OutputFile = Output{}
var _ builder.OutputRegex = Output{}

type Output struct {
	Name string `yaml:"name"`

	// See https://porter.sh/mixins/exec/#outputs
	// TODO: If your mixin doesn't support these output types, you can remove these and the interface assertions above, and from #/definitions/outputs in schema.json
	JsonPath string `yaml:"jsonPath,omitempty"`
	FilePath string `yaml:"path,omitempty"`
	Regex    string `yaml:"regex,omitempty"`
}

func (o Output) GetName() string {
	return o.Name
}

func (o Output) GetJsonPath() string {
	return o.JsonPath
}

func (o Output) GetFilePath() string {
	return o.FilePath
}

func (o Output) GetRegex() string {
	return o.Regex
}

var _ builder.ExecutableStep = Step{}

func (s Step) GetArguments() []string {
	return s.Arguments
}

func (s Step) GetFlags() builder.Flags {
	return s.Flags
}
func (s Step) GetWorkingDir() string {
	return s.WorkingDir
}

func (a Action) GetSteps() []builder.ExecutableStep {
	// Go doesn't have generics, nothing to see here...
	steps := make([]builder.ExecutableStep, len(a.Steps))
	for i := range a.Steps {
		steps[i] = a.Steps[i]
	}

	return steps
}
