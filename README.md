# [ColoredCobra](https://github.com/ivanpirog/coloredcobra) just with [lipgloss](https://github.com/charmbracelet/lipgloss) styles

> This library is Ivan Pirog's [ColoredCobra](https://github.com/ivanpirog/coloredcobra) just modified to work with lipgloss. All the credit to him for this awesome tool

[![Go Reference](https://pkg.go.dev/badge/github.com/auribuo/stylishcobra.svg)](https://pkg.go.dev/github.com/auribuo/stylishcobra)
[![Go Report Card](https://goreportcard.com/badge/github.com/auribuo/stylishcobra)](https://goreportcard.com/report/github.com/auribuo/stylishcobra)

**[Cobra](https://github.com/spf13/cobra)** library for creating powerful modern CLI doesn't support color settings for console output. `StylishCobra` is a small library that allows you to colorize the text output of the Cobra library, making the console output look better.

`StylishCobra` allows to customize the cobra text output with [lipgloss](https://github.com/charmbracelet/lipgloss) styles.

It's very easy to add `StylishCobra` to your project!

---

## Installing

To add the library to your project run

```bash
go get -u github.com/auribuo/stylishcobra
```

## Quick start

Open your `cmd/root.go` and insert this code:

```go
import (
    "github.com/auribuo/stylishcobra"
    "github.com/charmbracelet/lipgloss"
)
```

Then, before the call to `Execute()` of your root command put the calls to customize cobra

```go
stylishcobra.Setup(rootCmd). // This would be enough, it would do nothing
	StyleHeadings(lipgloss.NewStyle().Underline(true).Bold(true).Foreground(lipgloss.ANSIColor(termenv.ANSICyan))).
	StyleCommands(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIYellow)).Bold(true)).
	StyleAliases(lipgloss.NewStyle().Bold(true).Italic(true)).
	StyleCmdShortDescr(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIBrightRed))).
	StyleExample(lipgloss.NewStyle().Italic(true)).
	StyleExecName(lipgloss.NewStyle().Bold(true)).
	StyleFlags(lipgloss.NewStyle().Bold(true)).
	StyleFlagsDescr(lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(termenv.ANSIRed))).
	StyleFlagsDataType(lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Italic(true)).
	Init() // Call Init() to apply the styles

stylishcobra.Init(cfg) // Initialize with a *Config. Is a bit more tedious, bacuse lipgloss.NewStyle() returns no pointer but styles in the config are stored as pointers
```

You can also reuse styles etc. For more information refer to the lipgloss docs.

That's all. Now build your project and see the output of the help command.


### Available config parameters:

![Config Parameters](https://user-images.githubusercontent.com/8699212/159517553-7ef67fac-371b-4995-bebe-d702b6167fe1.png)

* `Headings:` headers style.

* `Commands:` commands style.

* `CmdShortDescr:` short description of commands style.

* `ExecName:` executable name style.

* `Flags:` short and long flag names (-f, --flag) style.

* `FlagsDataType:` style of flags data type.

* `FlagsDescr:` flags description text style.

* `Aliases:` list of command aliases style.

* `Example:` example text style.

* `NoExtraNewlines:` no line breaks before and after headings, if `true`. By default: `false`.

* `NoBottomNewline:` no line break at the end of Cobra's output, if `true`. By default: `false`.

The functions in the builder pattern are named `Style<Parameter>` except for the `NoExtraNewlines` and `NoBottomNewline`, these can be enabled by using `[En/Dis]ableExtraNewLines` and `[En/Dis]ableNoBottomNewline`

<br>

### `NoExtraNewlines` parameter results:

![extranewlines](https://user-images.githubusercontent.com/8699212/159517630-00855ffe-80df-4670-a054-e695f6c4fea7.png)


## How it works

`StylishCobra` patches Cobra's usage template and extends it with functions for text styling. The styles are [lipgloss](https://github.com/charmbracelet/lipgloss) styles allowing for a variety of customization.

## License

StylishCobra is released under the MIT license. See [LICENSE](https://github.com/auribuo/stylishcobra/blob/main/LICENSE).
