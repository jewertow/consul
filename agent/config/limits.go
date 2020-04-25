// +build !windows

package config

import "golang.org/x/sys/unix"

// getrlimit return the max file descriptors allocated by system
// return the number of file descriptors max
func getrlimit() (uint64, error) {
	var limit unix.Rlimit
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &limit)
	return uint64(limit.Cur), err
}
