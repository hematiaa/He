//go:build isatty
// +build isatty

package isatty

/*
#include <stdlib.h>
#include <unistd.h>
*/
import "C"

import (
	"syscall"
)

// Isatty tries to determine whether `fd` is a TTY.
func Isatty(fd uintptr) (bool, error) {
	result, err := C.isatty(C.int(fd))
	if err != nil && err != syscall.EINVAL {
		return false, err
	}
	return result != 0, nil
}
