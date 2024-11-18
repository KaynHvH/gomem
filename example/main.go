package main

import (
	"fmt"
	"gomem/gomem"
	"gomem/winapi"
	"log"
	"time"
	"unsafe"
)

func main() {
	hwnd, err := winapi.FindWindow("", "Minecraft 1.8.8")
	if err != nil || hwnd == 0 {
		log.Fatalf("Cannot find window: %v", err)
	}

	fmt.Printf("Found window. HWND: 0x%x\n", hwnd)

	var pid gomem.DWORD
	err = winapi.GetWindowThreadProcessId(hwnd, (*winapi.DWORD)(&pid))
	if err != nil {
		log.Fatalf("Cannot find PID: %v", err)
	}

	fmt.Printf("PID: %d\n", pid)

	handle, err := gomem.OpenProcess(gomem.PROCESS_ALL_ACCESS, false, pid)
	if err != nil {
		log.Fatalf("Cannot open: %v", err)
	}
	defer gomem.CloseHandle(handle)

	fmt.Printf("Process handle: 0x%x\n", handle)

	fmt.Println("Waiting 5 seconds")
	time.Sleep(5 * time.Second)

	address := uintptr(0xA86CDF84)

	value := int32(1234)

	data := make([]byte, unsafe.Sizeof(value))
	*(*int32)(unsafe.Pointer(&data[0])) = value

	err = gomem.WriteProcessMemory(handle, address, data)
	if err != nil {
		log.Fatalf("Error writing to memory: %v", err)
	}

	fmt.Println("Done")

	var i int
	fmt.Println("Press enter to end")
	fmt.Scanln("%d", &i)
}
