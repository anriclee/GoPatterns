package cmdpattern

import "fmt"

func NewThemeCmd() *Cmd {
	cmd := &Cmd{
		Use: "new theme",
		Run: func(args []string) error {
			fmt.Println("do new theme command,args:", args)
			return nil
		},
	}
	cmd.Flags().Bool("force", false, "init inside non-empty directory")
	return cmd
}

