package mage

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mage"
	"github.com/manifoldco/promptui"
	"github.com/manifoldco/promptui/list"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/pubgo/lava/pkg/clix"
)

var Cmd = clix.Command(func(cmd *cobra.Command, flags *pflag.FlagSet) {
	cmd.Use = "mage"
	cmd.Short = "mage install"
	cmd.Run = func(cmd *cobra.Command, args []string) {
		cmds := exec.Command("mage", "-l")
		cmds.Stderr = os.Stderr
		out, err := cmds.Output()
		if err != nil {
			return
		}

		scan := bufio.NewScanner(bytes.NewBuffer(out))

		var targets []string
		for scan.Scan() {
			line := scan.Text()
			if strings.HasPrefix(line, "Targets:") {
				continue
			}
			line = strings.TrimSpace(line)
			targets = append(targets, line)
		}

		templates := &promptui.SelectTemplates{
			Label:    "{{.}}",
			Active:   promptui.IconSelect + " {{.}}",
			Inactive: "  {{.|faint}}",
			Selected: promptui.IconGood + " {{.}}",
		}

		size := maxSize
		if len(targets) < size {
			size = len(targets)
		}

		prompt := promptui.Select{
			Label:             "Select a mage target:",
			Items:             targets,
			Templates:         templates,
			HideHelp:          true,
			Size:              size,
			Searcher:          searcher(targets),
			StartInSearchMode: true,
		}

		_, result, err := prompt.Run()

		if err != nil {
			return
		}

		result = strings.Split(result, " ")[0]

		fmt.Printf("mage %s\n", result)

		os.Args = append(os.Args, result)
		os.Exit(mage.Main())
	}
})

const (
	maxSize = 10
)

func searcher(targets []string) list.Searcher {
	return func(input string, index int) bool {
		if strings.Contains(strings.ToLower(targets[index]), input) {
			return true
		}
		return false
	}
}
