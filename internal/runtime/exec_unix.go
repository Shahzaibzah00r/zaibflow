//go:build !windows

package runtime

import "syscall"

func execReplace(path string, args []string, env []string) (int, error) {
	argv := append([]string{"claude"}, args...)
	if err := syscall.Exec(path, argv, env); err != nil {
		return 1, err
	}
	return 0, nil
}
