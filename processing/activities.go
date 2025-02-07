package processing

import (
	"context"
	"fmt"
	"math/rand"
)

// AddSuffixActivity appends a fixed suffix to the input data.
func AddSuffixActivity(ctx context.Context, data string) (string, error) {
	suffixes := []string{
		"-one", "-two", "-three", "-four", "-five",
		"-six", "-seven", "-eight", "-nine", "-ten",
	}
	suffix := suffixes[rand.Intn(len(suffixes))]
	result := fmt.Sprintf("%s%s", data, suffix)
	fmt.Println("Go Activity: Modified data to:", result)
	return result, nil
}
