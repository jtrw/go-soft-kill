package utils

import (
  "fmt"
)

func Info(msg string) {
    green := "0;32";
    color(msg, green)
}

func Error(msg string) {
    red := "0;31";
    color(msg, red)
}

func color(msg, color string) {
    fmt.Println("\033[" + color + "m" + msg + "\033[0m")
}