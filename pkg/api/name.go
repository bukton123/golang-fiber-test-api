package api

import "fmt"

func AppName (service string, version string) string {
	return fmt.Sprintf("Service: %s v%s", service, version)
}
