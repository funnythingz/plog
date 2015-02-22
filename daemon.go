package main

import (
	"os"
	"syscall"
)

func Daemon(nochdir, noclose int) int {
	var ret uintptr
	var err syscall.Errno

	ret, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		return -1
	}
	switch ret {
	case 0:
		break
	default:
		os.Exit(0)
	}

	pid, _ := syscall.Setsid()
	if pid == -1 {
		return -1
	}

	if nochdir == 0 {
		os.Chdir("/")
	}

	syscall.Umask(0)

	if noclose == 0 {
		f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if e == nil {
			fd := int(f.Fd())
			syscall.Dup2(fd, int(os.Stdin.Fd()))
			syscall.Dup2(fd, int(os.Stdout.Fd()))
			syscall.Dup2(fd, int(os.Stderr.Fd()))
		}
	}

	return 0
}
