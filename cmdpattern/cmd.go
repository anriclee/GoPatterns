package cmdpattern

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

// Execute the execute entrypoint
func Execute(args []string) {
	// merge cmds together
	cmds := []*Cmd{
		NewSiteCmd(),
		NewThemeCmd(),
	}
	// find cmds according to Args
	cmd, argments := findCommands(cmds, args)
	if cmd == nil {
		panic(errors.New("not found cmd"))
	}
	// execute it
	err := cmd.Run(argments)
	if err != nil {
		panic(err)
	}
}

func findCommands(cmds []*Cmd, args []string) (*Cmd, []string) {
	s := stripFlags(args)
	argments := trimEle(args, s)
	join := strings.Join(s, " ")
	fmt.Println(argments)
	for _, cmd := range cmds {
		if strings.TrimSpace(join) == cmd.Use {
			return cmd, argments
		}
	}
	return nil, nil
}

// remove elements from source if the elements of target appeared
func trimEle(source, target []string) []string {
	if len(source) == 0 || len(target) == 0 {
		return source
	}
	tmap := make(map[string]bool, len(target))
	for _, s := range target {
		tmap[s] = true
	}
	var ret []string
	for _, s := range source {
		if tmap[s] {
			continue
		}
		ret = append(ret, s)
	}
	return ret
}

func stripFlags(args []string) []string {
	if len(args) == 0 {
		return args
	}
	var commands []string
Loop:
	for len(args) > 0 {
		s := args[0]
		args = args[1:]
		switch {
		case s == "--":
			// "--" terminates the flags
			break Loop
		case strings.HasPrefix(s, "--") && !strings.Contains(s, "="):
			// If '--flag arg' then
			// delete arg from args.
			fallthrough
		case strings.HasPrefix(s, "-") && !strings.Contains(s, "=") && len(s) == 2:
			// If '-f arg' then
			// delete 'arg' from args or break the loop if len(args) <= 1.
			if len(args) <= 1 {
				break Loop
			} else {
				args = args[1:]
				continue
			}
		case s != "" && !strings.HasPrefix(s, "-"):
			commands = append(commands, s)
		}
	}

	return commands
}

// Cmd base command unit
type Cmd struct {
	Use   string
	Args  []string
	flags *flag.FlagSet
	Run   func(args []string) error
}

func (cmd *Cmd) Flags() *flag.FlagSet {
	if cmd.flags == nil {
		cmd.flags = flag.NewFlagSet(cmd.Use, flag.ContinueOnError)
	}
	return cmd.flags
}
