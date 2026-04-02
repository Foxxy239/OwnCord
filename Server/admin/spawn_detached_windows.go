//go:build windows

package admin

import (
	"os/exec"
	"syscall"
)

// applyDetachedProcessAttrs sets Windows creation flags for detached processes.
func applyDetachedProcessAttrs(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: 0x00000008, // DETACHED_PROCESS
	}
}
