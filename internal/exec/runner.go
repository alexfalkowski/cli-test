package exec

import "context"

// Runner will run a cmd and return the output.
type Runner interface {
	// Run cmd and return the output, unless an error occurs.
	Run(ctx context.Context, args string) (string, error)
}
