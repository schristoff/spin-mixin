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
		//todo: make this work
		//when we do this it spits out a code we need to
		// give to users, along with a url
		// it will print out "Device authorized!"
		// when done
		cmd = m.NewCommand(ctx, "spin", "login")
	}
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
