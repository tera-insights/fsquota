package fsquota

import (
	"os/user"
)

func setUserQuota(path string, user *user.User, limits *Limits) (info *Info, err error) {
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
