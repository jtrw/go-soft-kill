package main

import (
  //"fmt"
  "log"
  "github.com/jessevdk/go-flags"
  cli "github.com/soft-kill/v1/cmd/app/utils"
)


type Options struct {
   Pid int `short:"p" long:"pid" default:"" description:"PID"`
   Name string `short:"n" long:"name" default:"" description:"Name of script"`
}
func main() {
    var opts Options

    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()

    if err != nil {
        log.Fatal(err)
    }
    cli.Info("[OK]")
    cli.Error("[ERROR]")
}