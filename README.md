# Memory Manipulation 

This simple program allows you to read and write values in a running process's memory, such as a game. 
It demonstrates how to use Go to interact with Windows API functions, like `FindWindow`, `GetWindowThreadProcessId`, `ReadProcessMemory` and `WriteProcessMemory`, to manipulate the memory of a specific process

### Features

- Find a specific window by its title
- Get the ProcessID of the associated application
- Open the process with full access
- Write a new value to a specific memory address within the process

### Prerequisites

- **Go** installed on your computer
- **gomem** and **winapi** packages

### Setup

1. **Clone my repository**: Clone my repository by using `git clone https://github.com/KaynHvH/gomem`.
2. **Install Dependencies**: Clone or copy the required Go packages (`gomem` and `winapi`) into your Go workspace. These packages are responsible for interacting with system memory and calling Windows API functions

### How to Use

1. **Clone the repository** or copy the program code into your Go workspace
2. **Modify the code**:
    - Set the correct window title (e.g., `"Minecraft 1.8.8"`)
    - Set the correct **PID** of the target process (you can find this using Task Manager or Cheat Engine)
    - Adjust the memory address and the value you wish to write into that address
3. **Run the program**:
    - Open a command prompt or terminal
    - Navigate to the folder where your program is located
    - Build the program using `go build main.go` then run it by using `go run main.go`

4. The program will:
    - Find the window based on its title
    - Get the process ID (PID) of the running process
    - Open the process with full access
    - Write the value `1234` (or any) to the specified memory address
    - Print success or error messages to the console