// File: processing/activities.go
package processing

import (
	"context"
	"fmt"
	"math/rand"
)

// AddRandomPrefixActivity adds a random prefix to the input data.
func AddRandomPrefixActivity(ctx context.Context, data string) (string, error) {
	// Define a list of possible prefixes.
	prefixes := []string{"alpha-", "beta-", "gamma-", "delta-", "epsilon-", "zeta-", "eta-", "theta-", "iota-", "kappa-", "lambda-", "mu-", "nu-", "xi-", "omicron-", "pi-", "rho-", "sigma-", "tau-", "upsilon-", "phi-", "chi-", "psi-", "omega-"}
	// Choose a random prefix using Go's automatically seeded rand package.
	prefix := prefixes[rand.Intn(len(prefixes))]
	return fmt.Sprintf("%s%s", prefix, data), nil
}

// AddSuffixActivity appends a fixed suffix to the input data.
func AddSuffixActivity(ctx context.Context, data string) (string, error) {
	// Define a list of possible suffixes.
	suffix := []string{"-one", "-two", "-three", "-four", "-five", "-six", "-seven", "-eight", "-nine", "-ten", "-eleven", "-twelve", "-thirteen", "-fourteen", "-fifteen", "-sixteen", "-seventeen", "-eighteen", "-nineteen", "-twenty"}
	// Choose a fixed suffix using Go's automatically seeded rand package.
	suffixIndex := rand.Intn(len(suffix))
	return fmt.Sprintf("%s%s", data, suffix[suffixIndex]), nil
}
