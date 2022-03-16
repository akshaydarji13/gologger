package gologger

import "fmt"

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorCyan = "\033[36m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"

func Logln(a ...interface{}) {
	fmt.Println(a...)
}

func LogErrorln(a ...interface{}) {
	fmt.Print(colorRed)
	fmt.Println(a...)
	fmt.Print(colorReset)
}

func LogInfoln(a ...interface{}) {
	fmt.Print(colorCyan)
	fmt.Println(a...)
	fmt.Print(colorReset)
}

func LogDebugln(a ...interface{}) {
	fmt.Print(colorGreen)
	fmt.Println(a...)
	fmt.Print(colorReset)
}

func LogWarningln(a ...interface{}) {
	fmt.Print(colorYellow)
	fmt.Println(a...)
	fmt.Print(colorReset)
}

func Log(a ...interface{}) {
	fmt.Print(a...)
}

func LogError(a ...interface{}) {
	a = append([]interface{}{colorRed}, a...)
	a = append(a, colorReset)
	fmt.Print(a...)
}

func LogInfo(a ...interface{}) {
	a = append([]interface{}{colorCyan}, a...)
	a = append(a, colorReset)
	fmt.Print(a...)
}

func LogDebug(a ...interface{}) {
	a = append([]interface{}{colorGreen}, a...)
	a = append(a, colorReset)
	fmt.Print(a...)
}

func LogWarning(a ...interface{}) {
	a = append([]interface{}{colorYellow}, a...)
	a = append(a, colorReset)
	fmt.Print(a...)
}

func Logf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func LogErrorf(format string, a ...interface{}) {
	format = colorRed + format + colorReset
	fmt.Printf(format, a...)
}

func LogInfof(format string, a ...interface{}) {
	format = colorCyan + format + colorReset
	fmt.Printf(format, a...)
}

func LogDebugf(format string, a ...interface{}) {
	format = colorGreen + format + colorReset
	fmt.Printf(format, a...)
}

func LogWarningf(format string, a ...interface{}) {
	format = colorYellow + format + colorReset
	fmt.Printf(format, a...)
}
