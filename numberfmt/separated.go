package numberfmt

import (
	"fmt"
	"strings"
)

func ThousandsSeparated(n uint64, separator string) string {
	return Separated(n, 3, separator)
}

func Separated(n, segmentLength uint64, separator string) string {
	segments := make([]uint64, 0)

	segmentRange := pow10(segmentLength)

	for n >= segmentRange {
		segment := n % segmentRange
		segments = append(segments, segment)
		n /= segmentRange
	}

	segments = append(segments, n)
	output := strings.Builder{}
	padFmt := fmt.Sprintf("%%0%dd", segmentLength)

	for idx, first := len(segments)-1, false; idx >= 0; idx-- {
		segment := segments[idx]

		if !first {
			first = true
			output.WriteString(fmt.Sprintf("%d", segment))
		} else {
			output.WriteString(fmt.Sprintf(padFmt, segment))
		}

		if idx > 0 {
			output.WriteString(separator)
		}
	}

	return output.String()
}

func pow10(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	result := uint64(10)
	for i := uint64(1); i < n; i++ {
		result *= 10
	}

	return result
}
