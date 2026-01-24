//go:build !windows

package agent

import (
	"fmt"

	"c2framework/internals/control"
	"c2framework/internals/models"
)

// doPersist stub for non-Windows systems
func doPersist(args control.PersistArgsAgent) models.PersistResult {
	return models.PersistResult{
		Method:  args.Method,
		Success: false,
		Message: fmt.Sprintf("Persistence not implemented for this platform (requested: %s)", args.Method),
	}
}
