package helper

import "fmt"

func HumanReadableFileSize(bytesCount int64) string {
	const unit = 1000
	if bytesCount < unit {
		return fmt.Sprintf("%d B", bytesCount)
	}
	div, exp := int64(unit), 0
	for n := bytesCount / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(bytesCount)/float64(div), "kMGTPE"[exp])
}
