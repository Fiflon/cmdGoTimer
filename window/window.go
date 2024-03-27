package window

import (
	"fmt"
	"syscall"

	"github.com/kbinani/win"
)

func GetWindowHandle() win.HWND {
	return win.GetConsoleWindow()
}

func GetWindowPid(hwnd win.HWND) uint32 {
	var windowPid uint32
	win.GetWindowThreadProcessId(hwnd, &windowPid)
	return windowPid
}

func GetWindowTitle(hwnd win.HWND) string {
	windowTitle := make([]uint16, 256)
	win.GetWindowText(hwnd, &windowTitle[0], int32(len(windowTitle)))
	return syscall.UTF16ToString(windowTitle)
}

func SetWindowFocus(hwnd win.HWND) {
	win.ShowWindow(hwnd, 1)
	win.SetForegroundWindow(hwnd)
}

func FindAndFocusTerminalWindow() {
	hwnd := GetWindowHandle()
	if hwnd == 0 {
		fmt.Println("Nie można znaleźć uchwytu do okna terminala.")
		return
	}
	SetWindowFocus(hwnd)
}
