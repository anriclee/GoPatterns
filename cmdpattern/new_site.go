package cmdpattern

import "fmt"

func NewSiteCmd() *Cmd {
	cmd := &Cmd{
		Use: "new site",
		Run: func(args []string) error {
			fmt.Println("do new site command,args:", args)
			return nil
		},
	}
	cmd.Flags().Bool("force", false, "init inside non-empty directory")
	return cmd
}
