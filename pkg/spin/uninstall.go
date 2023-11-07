package spin

import (
	"context"
	"os"

	"get.porter.sh/porter/pkg/exec/builder"
)

// Uninstall runs a Spin destroy
func (m *Mixin) Uninstall(ctx context.Context) error {
	action, err := m.loadAction(ctx)
	if err != nil {
		return err
	}
	step := action.Steps[0]

	err = m.commandPreRun(ctx, &step)
	if err != nil {
		return err
	}

	// unsure what happens if you watch twice
	step.Arguments = []string{"watch"}
	action.Steps[0] = step
	_, err = builder.ExecuteSingleStepAction(ctx, m.RuntimeConfig, action)
	if err != nil {
		return err
	}

	//If it's watching and we remove, it should go away
	err = os.Remove("spin.toml")
	if err != nil {
		return err
	}

	return nil
}
