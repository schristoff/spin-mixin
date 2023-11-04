package spin

import (
	"context"
	"text/template"

	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const dockerfileLines = `RUN apt update && apt install -y wget tar git && \
wget https://github.com/fermyon/spin/releases/download/v2.0.0-rc.1/spin-v2.0.0-rc.1-linux-amd64.tar.gz --progress=dot:giga && \
tar -xvf spin-v2.0.0-rc.1-linux-amd64.tar.gz -C /usr/bin/ && \
spin plugin install -y --url https://github.com/fermyon/platform-plugin/releases/download/canary/platform.json`

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config *MixinConfig
}

type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
	FermyonCloud  bool   `yaml:"fermyonCloud"`
}

func (m *Mixin) Build(ctx context.Context) error {

	input := BuildInput{
		Config: &m.config, // Apply config directly to the mixin
	}

	err := builder.LoadAction(ctx, m.RuntimeConfig, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	tmpl, err := template.New("spin").Parse(dockerfileLines)
	if err != nil {
		return errors.Wrapf(err, "error parsing terraform mixin Dockerfile template")
	}

	return tmpl.Execute(m.Out, tmpl)
}
