package winapi

import (
	"syscall"
	"unsafe"
)

var (
	user32                       = syscall.NewLazyDLL("user32.dll")
	procFindWindow               = user32.NewProc("FindWindowW")
	procGetWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
)

type (
	HWND   uintptr
	DWORD  uint32
	HANDLE uintptr
)

func UTF16PtrFromString(s string) (*uint16, error) {
	return syscall.UTF16PtrFromString(s)
}

func FindWindow(className, windowName string) (HWND, error) {
	var classPtr, windowPtr *uint16
	var err error

	if className != "" {
		classPtr, err = UTF16PtrFromString(className)
		if err != nil {
			return 0, err
		}
	}

	if windowName != "" {
		windowPtr, err = UTF16PtrFromString(windowName)
		if err != nil {
			return 0, err
		}
	}

	hwnd, _, _ := procFindWindow.Call(
		uintptr(unsafe.Pointer(classPtr)),
		uintptr(unsafe.Pointer(windowPtr)),
	)

	if hwnd == 0 {
		return 0, syscall.GetLastError()
	}

	return HWND(hwnd), nil
}

func GetWindowThreadProcessId(hwnd HWND, pid *DWORD) error {
	_, _, err := procGetWindowThreadProcessId.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(pid)),
	)

	if pid == nil {
		return err
	}

	return nil
}
