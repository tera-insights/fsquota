package fsquota

import (
	"os"
	"os/user"
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func setUserQuota(path string, user *user.User, limits *Limits) (info *Info, err error) {

	usernamePtr, err := syscall.UTF16PtrFromString(user.Username)
	if err != nil {
		return
	}

	var diskQuotaUser windows.Handle

	r1, _, e1 := procFindUserName.Call(
		uintptr(unsafe.Pointer(usernamePtr)),
		uintptr(diskQuotaUser),
	)
	runtime.KeepAlive(usernamePtr)
	if int(r1) == 0 {
		return nil, os.NewSyscallError("FindUser", e1)
	}
	return nil, ErrNotSupported
}

func getUserInfo(path string, user *user.User) (info *Info, err error) {
	return nil, ErrNotSupported
}

func getUserReport(path string) (report *Report, err error) {
	return nil, ErrNotSupported
}

func setGroupQuota(path string, group *user.Group, limits *Limits) (info *Info, err error) {
	return nil, ErrNotSupported
}

func getGroupInfo(path string, group *user.Group) (info *Info, err error) {
	return nil, ErrNotSupported
}

func getGroupReport(path string) (report *Report, err error) {
	return nil, ErrNotSupported
}

func userQuotasSupported(path string) (supported bool, err error) {
	return false, ErrNotSupported
}

func groupQuotasSupported(path string) (supported bool, err error) {
	return false, ErrNotSupported
}
