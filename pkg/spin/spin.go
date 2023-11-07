package spin

import (
	"context"
	"os"

	"get.porter.sh/porter/pkg/runtime"
)

const defaultClientVersion string = "v2.0.0"
const defaultFermyonCloud bool = false

type Mixin struct {
	runtime.RuntimeConfig
	ClientVersion string
	config        MixinConfig
}

func New() *Mixin {
	return &Mixin{
		RuntimeConfig: runtime.NewConfig(),
		ClientVersion: defaultClientVersion,
		config: MixinConfig{
			FermyonCloud: defaultFermyonCloud,
		},
	}
}

// commandPreRun runs setup tasks applicable for every action
func (m *Mixin) commandPreRun(ctx context.Context, step *Step) error {
	if m.config.FermyonCloud {
		return nil
	}

	os.Setenv("BINDLE_PASSWORD", step.BindlePassword)
	os.Setenv("BINDLE_URL", step.BindleServer)
	os.Setenv("BINDLE_USERNAME", step.BindleUsername)
	os.Setenv("HIPPO_PASSWORD", step.HippoPassword)
	os.Setenv("HIPPO_USERNAME", step.HippoUsername)
	// os.Setenv("SPIN_AUTH_TOKEN", step.AuthToken)

	if m.config.WorkingDir != "" {
		os.Chdir(m.config.WorkingDir)
	}

	if step.GetWorkingDir() != "" {
		os.Chdir(step.GetWorkingDir())

	}

	return nil
}
