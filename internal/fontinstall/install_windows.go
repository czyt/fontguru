//go:build windows

package fontinstall

import (
	"errors"
	"syscall"
	"unsafe"
)

func InstallFont(localPath string) error {
	name, err := syscall.UTF16PtrFromString(localPath)
	if err != nil {
		return err
	}

	handle, _, _ := syscall.SyscallN(
		syscall.NewLazyDLL("gdi32.dll").NewProc("AddFontResourceW").Addr(),
		1,
		uintptr(unsafe.Pointer(name)),
		0,
		0,
	)

	if handle == 0 {
		return errors.New("install font failed")
	}
	return nil
}
