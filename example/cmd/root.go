package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"

	"github.com/auribuo/stylishcobra"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "This is an example of using ColoreCobra library.",
	Long: "This is just an example of using the ColoredCobra library. \n" +
		"Project home: https://github.com/ivanpirog/coloredcobra",
	Example: "There is a simple example of the Examples section.\n" +
		"Just try commands:\n\n" +
		"example help\n" +
		"example help test",
	Aliases: []string{"alias1", "alias2", "alias3"},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("No commands given. Run 'example help' for usage help.\n" +
			"Also try commands:\n\n" +
			"example help\n" +
			"example help test")
	},
}

var (
	// Used for flags.
	flag1 string
	flag2 string
	abc   []bool
)

func Execute() {

	stylishcobra.Setup(rootCmd).
		StyleHeadings(lipgloss.NewStyle().Underline(true).Bold(true).Foreground(lipgloss.ANSIColor(termenv.ANSICyan))).
		StyleCommands(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIYellow)).Bold(true)).
		StyleAliases(lipgloss.NewStyle().Bold(true).Italic(true)).
		StyleCmdShortDescr(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIBrightRed))).
		StyleExample(lipgloss.NewStyle().Italic(true)).
		StyleExecName(lipgloss.NewStyle().Bold(true)).
		StyleFlags(lipgloss.NewStyle().Bold(true)).
		StyleFlagsDescr(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIRed))).
		StyleFlagsDataType(lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Italic(true)).
		Init()

	rootCmd.PersistentFlags().StringVarP(&flag1, "flag", "f", "", "some flag")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&flag2, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().BoolSliceVar(&abc, "zzz", []bool{true, false}, "usage of bools")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
