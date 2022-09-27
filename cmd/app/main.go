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
   Pid int `short:"p" long:"pid" default:"0" description:"PID"`
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
    if len(opts.Name) > 0 {
        findProcessByName(opts.Name)
        return
    }
    if opts.Pid > 0 {
         terminateProcessByPid(opts)
    }

}

func terminateProcessByPid(opts Options) {
     pid := opts.Pid
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
    command := fmt.Sprintf("ps aux | grep %s", name)
    cmd := exec.Command("bash", "-c", command)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(string(stdout))
}