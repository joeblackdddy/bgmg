package cli

import (
	"github.com/spf13/cobra"
	"github.com/joeblackdddy/bgmg/internal/art"
)

func Run() error {
	app := &cobra.Command{
		Use: "bgmg",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				art.Success()
				return
			}
			if args[0] == "0" {
				art.Success()
				return
			}
			art.Success()
		},
	}
	return app.Execute()
}