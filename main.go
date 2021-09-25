package main

import (
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {

	wireGuardDLL := windows.MustLoadDLL("wireguard.dll")
	defer wireGuardDLL.Release()

	createAdapter := wireGuardDLL.MustFindProc("WireGuardCreateAdapter")
	pool := uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("WireGuard")))
	name := uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("default")))
	guid := uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("e8f218da-8209-4880-b36d-33eea78dbb9c")))
	r1, r2, err := createAdapter.Call(pool, name, guid)
	log.Print(r1)
	log.Print(r2)
	log.Print(err)

	deleteAdapter := wireGuardDLL.MustFindProc("WireGuardDeleteAdapter")
	r1, r2, err = deleteAdapter.Call(uintptr(unsafe.Pointer(&r2)))
	log.Print(r1)
	log.Print(r2)
	log.Print(err)
}
