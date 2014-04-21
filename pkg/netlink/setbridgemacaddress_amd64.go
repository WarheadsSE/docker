// +build amd64

package netlink

import (
	"math/rand"
	"syscall"
	"unsafe"
)

func setBridgeMacAddress(s int, name string) error {
	ifr := ifreqHwaddr{}
	ifr.IfruHwaddr.Family = syscall.ARPHRD_ETHER
	copy(ifr.IfrnName[:], name)

	for i := 0; i < 6; i++ {
		ifr.IfruHwaddr.Data[i] = int8(rand.Intn(128))
	}

	ifr.IfruHwaddr.Data[0] &^= 0x1 // clear multicast bit
	ifr.IfruHwaddr.Data[0] |= 0x2  // set local assignment bit (IEEE802)

	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(s), syscall.SIOCSIFHWADDR, uintptr(unsafe.Pointer(&ifr))); err != 0 {
		return err
	}
	return nil
}
