package gomem

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	kernel32               = windows.NewLazySystemDLL("kernel32.dll")
	procOpenProcess        = kernel32.NewProc("OpenProcess")
	procReadProcessMemory  = kernel32.NewProc("ReadProcessMemory")
	procWriteProcessMemory = kernel32.NewProc("WriteProcessMemory")
	procCloseHandle        = kernel32.NewProc("CloseHandle")
)

func OpenProcess(access DWORD, inheritHandle bool, processID DWORD) (HANDLE, error) {
	var inherit uintptr
	if inheritHandle {
		inherit = 1
	} else {
		inherit = 0
	}

	handle, _, err := procOpenProcess.Call(
		uintptr(access),
		inherit,
		uintptr(processID),
	)

	if handle == 0 {
		return 0, err
	}
	return HANDLE(handle), nil
}

func ReadProcessMemory(process HANDLE, address uintptr, size SIZE_T) ([]byte, error) {
	buffer := make([]byte, size)
	var bytesRead SIZE_T

	ret, _, err := procReadProcessMemory.Call(
		uintptr(process),
		address,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(size),
		uintptr(unsafe.Pointer(&bytesRead)),
	)

	if ret == 0 {
		return nil, err
	}

	return buffer, nil
}

func WriteProcessMemory(handle HANDLE, address uintptr, data []byte) error {
	var bytesWritten uintptr
	success, _, err := procWriteProcessMemory.Call(
		uintptr(handle),
		address,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)

	if success == 0 {
		return fmt.Errorf("Error writing to memory: %v", err)
	}

	fmt.Printf("%d bytes written to memory at address %x\n", bytesWritten, address)

	if bytesWritten != uintptr(len(data)) {
		return fmt.Errorf("only part of the WriteProcessMemory request completed")
	}

	return nil
}

// i had error with default writeprocessmemory handler so i made that as a test. UPDATE = fixed
func WriteProcessMemoryInChunks(handle HANDLE, address uintptr, data []byte) error {
	chunkSize := 1
	var bytesWritten uintptr
	totalBytesWritten := uintptr(0)

	for i := 0; i < len(data); i += chunkSize {
		chunk := data[i : i+chunkSize]
		success, _, err := procWriteProcessMemory.Call(
			uintptr(handle),
			address+uintptr(i),
			uintptr(unsafe.Pointer(&chunk[0])),
			uintptr(len(chunk)),
			uintptr(unsafe.Pointer(&bytesWritten)),
		)

		if success == 0 {
			return fmt.Errorf("Error: %v", err)
		}

		totalBytesWritten += bytesWritten
		fmt.Printf("Saved %d bytes on addr %x\n", bytesWritten, address+uintptr(i))

		if totalBytesWritten >= uintptr(len(data)) {
			break
		}
	}

	return nil
}

func CloseHandle(handle HANDLE) error {
	ret, _, err := procCloseHandle.Call(uintptr(handle))
	if ret == 0 {
		return err
	}
	return nil
}
