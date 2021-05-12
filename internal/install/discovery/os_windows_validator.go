package discovery

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/newrelic/newrelic-cli/internal/install/types"
)

type OsWindowsValidator struct{}

var (
	windowsVersionNoLongerSupported = "This version of Windows is no longer supported"
	windowsNoVersionMessage         = "Failed to identified a valid version of Windows"
	windowsMinMajor                 = 6
	windowsMinMinor                 = 2
)

func NewOsWindowsValidator() *OsWindowsValidator {
	validator := OsWindowsValidator{}

	return &validator
}

func (v *OsWindowsValidator) Execute(m *types.DiscoveryManifest) error {
	if m.OS != "windows" {
		return nil
	}

	versions := strings.Split(m.PlatformVersion, ".")

	switch len(versions) {
	case 0:
		return fmt.Errorf(windowsNoVersionMessage)
	case 1:
		major, err := strconv.Atoi(versions[0])
		if err == nil {
			return ensureMinimumVersion(major, windowsMinMinor-1)
		}
		return fmt.Errorf(windowsNoVersionMessage)
	default:
		major, aerr := strconv.Atoi(versions[0])
		if aerr == nil {
			minor, ierr := strconv.Atoi(versions[1])
			if ierr == nil {
				return ensureMinimumVersion(major, minor)
			}
		}
	}

	return fmt.Errorf(windowsNoVersionMessage)
}

func ensureMinimumVersion(major int, minor int) error {
	if major < windowsMinMajor {
		return fmt.Errorf(windowsVersionNoLongerSupported)
	}
	if major == windowsMinMajor {
		if minor < windowsMinMinor {
			return fmt.Errorf(windowsVersionNoLongerSupported)
		}
	}
	return nil
}
