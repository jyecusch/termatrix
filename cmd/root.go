package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jyecusch/termatrix/pkg/tui"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "termatrix",
		Short: "Termatrix Matrix Terminal Effect",
		Long:  `Termatrix`,
		Run: func(cmd *cobra.Command, args []string) {
			p := tea.NewProgram(tui.NewMatrixRain(), tea.WithAltScreen())
			_, err := p.Run()
			if err != nil {
				fmt.Println("Error running program:", err)
				os.Exit(1)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}
