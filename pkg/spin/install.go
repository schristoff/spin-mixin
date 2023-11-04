package spin

import (
	"context"
	"fmt"
	"os"

	"get.porter.sh/porter/pkg/exec/builder"
)

// used by uninstall & upgrade for rn
func (m *Mixin) Execute(ctx context.Context) error {
	return nil
}

// Install runs a spin deploy
func (m *Mixin) Install(ctx context.Context) error {

	action, err := m.loadAction(ctx)
	if err != nil {
		return err
	}
	step := action.Steps[0]

	err = m.commandPreRun(ctx, &step)
	if err != nil {
		return fmt.Errorf("unable to login, %s", err)
	}
	//this should log us in?
	m.Init(ctx)

	step.Arguments = []string{"deploy"}

	os.Chdir(step.WorkingDir)

	action.Steps[0] = step
	_, err = builder.ExecuteSingleStepAction(ctx, m.RuntimeConfig, action)
	if err != nil {
		return err
	}

	//todo: Outputs?

	return nil

}
