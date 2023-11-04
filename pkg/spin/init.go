package spin

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

const platURL string = "https://github.com/fermyon/platform-plugin/releases/download/canary/platform.json"

func (m *Mixin) Init(ctx context.Context) error {
	var cmd *exec.Cmd
	if m.config.FermyonCloud {
		// If we hit here, it'll output to users.
		// Need to check if there's a timeout set
		// on this
		fmt.Println("inside ferm true?")
		cmd = m.NewCommand(ctx, "spin", "login")
	}

	// would you please install
	platCmd := exec.Command("spin", "plugin", "install", "-y", "--url", platURL)
	fmt.Println(platCmd)
	if err := platCmd.Run(); err != nil {
		return err
	}

	cmd = m.NewCommand(ctx, "spin", "platform", "login")

	cmd.Stdout = m.Out
	cmd.Stderr = m.Err

	prettyCmd := fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args, " "))

	fmt.Println(cmd)
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
