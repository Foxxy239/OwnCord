//go:build !windows

package admin

import "os/exec"

// applyDetachedProcessAttrs is a no-op on non-Windows platforms.
func applyDetachedProcessAttrs(_ *exec.Cmd) {}
