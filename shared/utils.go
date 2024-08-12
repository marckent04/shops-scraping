package shared

import (
	"slices"
)

func SlicesFilter[S interface{ ~[]E }, E any](s S, f func(E) bool) (result S) {
	// idx := -1 just used in order to begin at first slice index
	for idx := -1; ; {
		sliceIdx := slices.IndexFunc(s[idx+1:], f)
		if sliceIdx == -1 {
			break
		}
		idx = sliceIdx + idx + 1
		result = append(result, s[idx])
	}
	return
}
