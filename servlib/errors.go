package servlib

import "errors"

const (
	HubAddrNoExistErr     = iota
	NoJobsExistForService = iota
)

func AServerError(code int) error {
	switch code {
	case HubAddrNoExistErr:
		return errors.New("Server error: hub address env var does not exist")
	case NoJobsExistForService:
		return errors.New("Info: no jobs exist in hub yet")
	default:
		return errors.New("Server error: unknown error")
	}
}
