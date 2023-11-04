package spin

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func (m *Mixin) Init(ctx context.Context) error {
	var cmd *exec.Cmd
	if m.config.FermyonCloud {
		// If we hit here, it'll output to users.
		// Need to check if there's a timeout set
		// on this
		cmd = m.NewCommand(ctx, "spin", "login")
	}
	// it's saying it can't find the platform command, and
	// is install the cloud plugin :/
	cmd = m.NewCommand(ctx, "spin", "platform", "login")

	cmd.Stdout = m.Out
	cmd.Stderr = m.Err

	prettyCmd := fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("could not execute command, %s: %s", prettyCmd, err)
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
