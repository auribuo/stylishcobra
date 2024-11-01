// StyledCobra allows you to colorize Cobra's text output using lipgloss,
// making it look better using simple settings to customize
// individual parts of console output.
//
// Usage example:
//
// 1. Insert in cmd/root.go file of your project :
//
//	import (
//	     "github.com/auribuo/stylishcobra"
//	     "github.com/charmbracelet/lipgloss"
//	 )
//
// 2. Put the following code to the beginning of the Execute() function:
//
//	 stylishcobra.Setup(rootCmd).
//		StyleHeadings(lipgloss.NewStyle().Underline(true).Bold(true).Foreground(lipgloss.ANSIColor(termenv.ANSICyan))).
//		StyleCommands(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIYellow)).Bold(true)).
//		StyleAliases(lipgloss.NewStyle().Bold(true).Italic(true)).
//		StyleCmdShortDescr(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIGreen))).
//		StyleExample(lipgloss.NewStyle().Italic(true)).
//		StyleExecName(lipgloss.NewStyle().Bold(true)).
//		StyleFlags(lipgloss.NewStyle().Bold(true)).
//		StyleFlagsDescr(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIRed))).
//		StyleFlagsDataType(lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Italic(true)).
//		Init()
//
// 3. Build & execute your code.
//
// Copyright © 2024 Aurelio Buonomo
// Released under the MIT license.
package stylishcobra

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// Config contains all the styles that can be applied using stylishcobra.
// Use `Setup()` to start a builder to configure. Optionally you can also instanciate your own instance
type Config struct {
	RootCmd         *cobra.Command
	Headings        *lipgloss.Style
	Commands        *lipgloss.Style
	CmdShortDescr   *lipgloss.Style
	ExecName        *lipgloss.Style
	Flags           *lipgloss.Style
	FlagsDataType   *lipgloss.Style
	FlagsDescr      *lipgloss.Style
	Aliases         *lipgloss.Style
	Example         *lipgloss.Style
	NoExtraNewlines bool
	NoBottomNewline bool
}

// Setup starts the builing process and supplies the required root command
func Setup(rootCmd *cobra.Command) *Config {
	return &Config{RootCmd: rootCmd}
}

func (cfg *Config) StyleHeadings(style lipgloss.Style) *Config {
	cfg.Headings = &style
	return cfg
}

func (cfg *Config) StyleCommands(style lipgloss.Style) *Config {
	cfg.Commands = &style
	return cfg
}

func (cfg *Config) StyleCmdShortDescr(style lipgloss.Style) *Config {
	cfg.CmdShortDescr = &style
	return cfg
}

func (cfg *Config) StyleExecName(style lipgloss.Style) *Config {
	cfg.ExecName = &style
	return cfg
}

func (cfg *Config) StyleFlags(style lipgloss.Style) *Config {
	cfg.Flags = &style
	return cfg
}

func (cfg *Config) StyleFlagsDataType(style lipgloss.Style) *Config {
	cfg.FlagsDataType = &style
	return cfg
}

func (cfg *Config) StyleFlagsDescr(style lipgloss.Style) *Config {
	cfg.FlagsDescr = &style
	return cfg
}

func (cfg *Config) StyleAliases(style lipgloss.Style) *Config {
	cfg.Aliases = &style
	return cfg
}

func (cfg *Config) StyleExample(style lipgloss.Style) *Config {
	cfg.Example = &style
	return cfg
}

func (cfg *Config) DisableExtraNewlines() *Config {
	cfg.NoExtraNewlines = true
	return cfg
}

func (cfg *Config) EnableExtraNewlines() *Config {
	cfg.NoExtraNewlines = false
	return cfg
}

func (cfg *Config) DisableBottomNewline() *Config {
	cfg.NoExtraNewlines = true
	return cfg
}

func (cfg *Config) EnableBottomNewline() *Config {
	cfg.NoExtraNewlines = false
	return cfg
}

// Init patches Cobra's usage template with configuration provided.
func (cfg *Config) Init() {

	if cfg.RootCmd == nil {
		panic("stylishcobra: Root command pointer is missing.")
	}

	// Get usage template
	tpl := cfg.RootCmd.UsageTemplate()
	// Debug line, uncomment when needed
	// fmt.Println(tpl)

	// Add extra line breaks for headings
	if cfg.NoExtraNewlines == false {
		tpl = strings.NewReplacer(
			"Usage:", "\nUsage:\n",
			"Aliases:", "\nAliases:\n",
			"Examples:", "\nExamples:\n",
			"Available Commands:", "\nAvailable Commands:\n",
			"Global Flags:", "\nGlobal Flags:\n",
			"Additional help topics:", "\nAdditional help topics:\n",
			"Use \"", "\nUse \"",
		).Replace(tpl)
		re := regexp.MustCompile(`(?m)^Flags:$`)
		tpl = re.ReplaceAllString(tpl, "\nFlags:\n")
	}

	// Styling headers
	if cfg.Headings != nil {
		cobra.AddTemplateFunc("HeadingStyle", sprintFunc(cfg.Headings))

		// Wrap template headers into a new function
		tpl = strings.NewReplacer(
			"Usage:", `{{HeadingStyle "Usage:"}}`,
			"Aliases:", `{{HeadingStyle "Aliases:"}}`,
			"Examples:", `{{HeadingStyle "Examples:"}}`,
			"Available Commands:", `{{HeadingStyle "Available Commands:"}}`,
			"Global Flags:", `{{HeadingStyle "Global Flags:"}}`,
			"Additional help topics:", `{{HeadingStyle "Additional help topics:"}}`,
		).Replace(tpl)

		re := regexp.MustCompile(`(?m)^(\s*)Flags:(\s*)$`)
		tpl = re.ReplaceAllString(tpl, `$1{{HeadingStyle "Flags:"}}$2`)
	}

	// Styling commands
	if cfg.Commands != nil {
		cobra.AddTemplateFunc("CommandStyle", sprintFunc(cfg.Commands))
		cobra.AddTemplateFunc("sum", func(a, b int) int {
			return a + b
		})

		re := regexp.MustCompile(`(?i){{\s*rpad\s+.Name\s+.NamePadding\s*}}`)
		tpl = re.ReplaceAllLiteralString(tpl, "{{rpad (CommandStyle .Name) (sum .NamePadding 12)}}")

		re = regexp.MustCompile(`(?i){{\s*rpad\s+.CommandPath\s+.CommandPathPadding\s*}}`)
		tpl = re.ReplaceAllLiteralString(tpl, "{{rpad (CommandStyle .CommandPath) (sum .CommandPathPadding 12)}}")
	}

	// Styling a short desription of commands
	if cfg.CmdShortDescr != nil {
		cobra.AddTemplateFunc("CmdShortStyle", sprintFunc(cfg.CmdShortDescr))
		tpl = strings.NewReplacer("{{.Short}}", "{{CmdShortStyle .Short}}").Replace(tpl)
	}

	// Styling executable file name
	if cfg.ExecName != nil {
		cobra.AddTemplateFunc("ExecStyle", sprintFunc(cfg.ExecName))
		cobra.AddTemplateFunc("UseLineStyle", func(s string) string {
			spl := strings.Split(s, " ")
			spl[0] = cfg.ExecName.Render(spl[0])
			return strings.Join(spl, " ")
		})

		re := regexp.MustCompile(`(?i){{\s*.CommandPath\s*}}`)
		tpl = re.ReplaceAllLiteralString(tpl, "{{ExecStyle .CommandPath}}")

		re = regexp.MustCompile(`(?i){{\s*.UseLine\s*}}`)
		tpl = re.ReplaceAllLiteralString(tpl, "{{UseLineStyle .UseLine}}")
	}

	// Styling flags
	if cfg.Flags != nil || cfg.FlagsDescr != nil || cfg.FlagsDataType != nil {
		cobra.AddTemplateFunc("FlagStyle", func(s string) string {

			// Flags info section is multi-line.
			// Let's split these lines and iterate them.
			lines := strings.Split(s, "\n")
			for k := range lines {

				// Styling short and full flags (-f, --flag)
				if cfg.Flags != nil {
					re := regexp.MustCompile(`(--?\S+)`)
					for _, flag := range re.FindAllString(lines[k], 2) {
						lines[k] = strings.Replace(lines[k], flag, cfg.Flags.Render(flag), 1)
					}
				}

				// If no styles for flag data types and description - continue
				if cfg.FlagsDescr == nil && cfg.FlagsDataType == nil {
					continue
				}

				// Split line into two parts: flag data type and description
				// Tip: Use debugger to understand the logic
				re := regexp.MustCompile(`\s{2,}`)
				spl := re.Split(lines[k], -1)
				if len(spl) != 3 {
					continue
				}

				// Styling the flag description
				if cfg.FlagsDescr != nil {
					lines[k] = strings.Replace(lines[k], spl[2], cfg.FlagsDescr.Render(spl[2]), 1)
				}

				// Styling flag data type
				// Tip: Use debugger to understand the logic
				if cfg.FlagsDataType != nil {
					re = regexp.MustCompile(`\s+(\w+)$`) // the last word after spaces is the flag data type
					m := re.FindAllStringSubmatch(spl[1], -1)
					if len(m) == 1 && len(m[0]) == 2 {
						lines[k] = strings.Replace(lines[k], m[0][1], cfg.FlagsDataType.Render(m[0][1]), 1)
					}
				}

			}
			s = strings.Join(lines, "\n")

			return s

		})

		re := regexp.MustCompile(`(?i)(\.(InheritedFlags|LocalFlags)\.FlagUsages)`)
		tpl = re.ReplaceAllString(tpl, "FlagStyle $1")
	}

	// Styling aliases
	if cfg.Aliases != nil {
		cobra.AddTemplateFunc("AliasStyle", sprintFunc(cfg.Aliases))

		re := regexp.MustCompile(`(?i){{\s*.NameAndAliases\s*}}`)
		tpl = re.ReplaceAllLiteralString(tpl, "{{AliasStyle .NameAndAliases}}")
	}

	// Styling the example text
	if cfg.Example != nil {
		cobra.AddTemplateFunc("ExampleStyle", sprintFunc(cfg.Example))

		re := regexp.MustCompile(`(?i){{\s*.Example\s*}}`)
		tpl = re.ReplaceAllLiteralString(tpl, "{{ExampleStyle .Example}}")
	}

	// Adding a new line to the end
	if !cfg.NoBottomNewline {
		tpl += "\n"
	}

	// Apply patched template
	cfg.RootCmd.SetUsageTemplate(tpl)
	// Debug line, uncomment when needed
	// fmt.Println("---")
	// fmt.Println(tpl)
}

func Init(cfg *Config) {
	cfg.Init()
}

func sprintFunc(st *lipgloss.Style) func(...interface{}) string {
	return func(a ...interface{}) string {
		return st.Render(fmt.Sprint(a...))
	}
}
