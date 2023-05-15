//go:build windows

package fontinstall

import (
	"syscall"
	"unsafe"
)

func InstallFont(localPath string) error {
	var gdi32 = syscall.NewLazyDLL("gdi32.dll")
	var addFont = gdi32.NewProc("AddFontResourceExW")
	ret, _, err := addFont.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(localPath))),
		0,
		0,
	)
	if ret == 0 {
		return err
	}
	return nil
}
