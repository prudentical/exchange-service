package discovery

import "fmt"

type NoInstanceAvailable struct {
	service string
}

func (e NoInstanceAvailable) Error() string {
	return fmt.Sprintf("No instance is available for %s", e.service)
}
