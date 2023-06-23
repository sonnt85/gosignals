//go:build windows
// +build windows

package gosignals

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// convert a signal name to signal
func ToSignal(signalName string) (os.Signal, error) {
	if !strings.HasPrefix(signalName, "SIG") {
		signalName = fmt.Sprintf("SIG%s", signalName)
	}
	if signalName == "SIGHUP" {
		return syscall.SIGHUP, nil
	} else if signalName == "SIGINT" {
		return syscall.SIGINT, nil
	} else if signalName == "SIGQUIT" {
		return syscall.SIGQUIT, nil
	} else if signalName == "SIGKILL" {
		return syscall.SIGKILL, nil
	} else if signalName == "SIGUSR1" {
		return nil, errors.New("signal USR1 is not supported in windows")
	} else if signalName == "SIGUSR2" {
		return nil, errors.New("signal USR2 is not supported in windows")
	} else {
		return syscall.SIGTERM, nil
	}

}

// Args:
//
//	process - the process
//	sig - the signal
//	sigChildren - ignore in windows system
func Kill(process *os.Process, sig os.Signal, sigChilren bool) error {
	//Signal command can't kill children processes, call  taskkill command to kill them
	cmd := exec.Command("taskkill", "/F", "/T", "/PID", fmt.Sprintf("%d", process.Pid))
	err := cmd.Start()
	if err == nil {
		return cmd.Wait()
	}
	//if fail to find taskkill, fallback to normal signal
	return process.Signal(sig)
}
