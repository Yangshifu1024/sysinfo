package formats

import (
    "github.com/fatih/color"
)

var (
    Title func(... interface{}) string
    Default func(...interface{}) string
    Success func(...interface{}) string
    Warning func(...interface{}) string
    Danger func(...interface{}) string
    Dangerf func(string, ...interface{}) string
    Warningf func(string, ...interface{}) string
    Successf func(string, ...interface{}) string
)

func init() {
    Title = color.New(color.FgBlue, color.Bold).SprintFunc()
    Default = color.New(color.FgBlue,).SprintFunc()
    Success = color.New(color.FgGreen).SprintFunc()
    Warning = color.New(color.FgYellow).SprintFunc()
    Danger = color.New(color.FgRed).SprintFunc()
    Dangerf = color.New(color.FgRed).SprintfFunc()
    Warningf = color.New(color.FgYellow).SprintfFunc()
    Successf = color.New(color.FgGreen).SprintfFunc()
}
