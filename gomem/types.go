package gomem

type (
	HANDLE uintptr
	DWORD  uint32
	LPVOID uintptr
	SIZE_T uintptr
)

const (
	PROCESS_VM_READ      = 0x0010
	PROCESS_VM_WRITE     = 0x0020
	PROCESS_VM_OPERATION = 0x0008
	PROCESS_ALL_ACCESS   = 0x1F0FFF
)
