package main

import (
  "fmt"
  "os"
  //"strconv"
  "syscall"
  "os/exec"
  "log"
  "github.com/jessevdk/go-flags"
  cli "github.com/soft-kill/v1/cmd/app/utils"
)


type Options struct {
   Pid int `short:"p" long:"pid" default:"" description:"PID"`
   Name string `short:"n" long:"name" default:"" description:"Name of script"`
   Signal int `short:"s" long:"signal" default:"9" description:"Posix signal, Default SIGILL=4"`
}
func main() {
    var opts Options

    parser := flags.NewParser(&opts, flags.Default)
    _, err := parser.Parse()

    if err != nil {
        log.Fatal(err)
    }

    if opts.Pid > 0 {
         terminateProcessByPid(opts.Pid)
    }
}

func terminateProcessByPid(pid int) {
     process, err := os.FindProcess(pid)
     if err != nil {
         msg := fmt.Sprintf("Failed to find process: %s\n", err)
         cli.Error(msg)
     } else {
         err := process.Signal(syscall.Signal(opts.Signal))
         msg := fmt.Sprintf("process.Signal on pid %d returned: %v\n", pid, err)
         cli.Info(msg)
     }
}

func findProcessByName(name string) {
    cmd := exec.Command("ps", "aux", "|", "grep", name)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(stdout)
}