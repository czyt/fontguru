//go:build windows

package fontinstall

import (
	"syscall"
	"unsafe"
)

var (
	addFontResourceEx = syscall.NewLazyDLL("gdi32.dll").NewProc("AddFontResourceExW")
	sendMessage       = syscall.NewLazyDLL("user32.dll").NewProc("SendMessageW")
)

func InstallFont(localPath string) error {
	r1, _, err := addFontResourceEx.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(localPath))),
		0,
		0,
	)
	if r1 == 0 {
		return err
	}
	// Reload the font cache
	hwndBroadcast := uintptr(0xFFFF)
	wmFontchange := uintptr(0x001D)
	lparam := uintptr(0)
	r1, _, err = sendMessage.Call(
		hwndBroadcast,
		wmFontchange,
		0,
		lparam,
	)
	if r1 == 0 {
		return err
	}
	return nil
}
