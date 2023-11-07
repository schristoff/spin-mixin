package spin

import (
	"context"

	"get.porter.sh/porter/pkg/exec/builder"
)

// Upgrade runs a Spin destroy
func (m *Mixin) Upgrade(ctx context.Context) error {
	action, err := m.loadAction(ctx)
	if err != nil {
		return err
	}
	step := action.Steps[0]

	err = m.commandPreRun(ctx, &step)
	if err != nil {
		return err
	}

	// unsure what happens if you deploy twice
	step.Arguments = []string{"deploy"}
	action.Steps[0] = step
	_, err = builder.ExecuteSingleStepAction(ctx, m.RuntimeConfig, action)
	if err != nil {
		return err
	}

	//If it's watching and we remove, it should go away
	if err != nil {
		return err
	}

	return nil
}
