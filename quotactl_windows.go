package fsquota

import (
	"golang.org/x/sys/windows"
)

var (
	dskquota32 = windows.NewLazyDLL("dskquota.dll")

	procAddUserName   = dskquota32.NewProc("AddUserName")
	procDeleteUser    = dskquota32.NewProc("DeleteUser")
	procInitialize    = dskquota32.NewProc("Initialize")
	procSetQuotaLimit = dskquota32.NewProc("SetQuotaLimit")
	procFindUserName  = dskquota32.NewProc("FindUserName")
)
