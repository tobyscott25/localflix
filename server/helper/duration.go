package helper

import (
	"fmt"

	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
)

// func GetVideoDuration(filePath string) (string, error) {
// 	cmd := exec.Command("ffmpeg", "-i", filePath)

// 	fmt.Println("Running ffmpeg command: ", cmd.String())

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Printf("Error running ffmpeg command: %v\n", err)
// 		return "", err
// 	}

// 	duration := ExtractDurationFromFFmpegOutput(string(output))
// 	return duration, nil
// }

// func ExtractDurationFromFFmpegOutput(output string) string {
// 	duration := ""
// 	lines := strings.Split(output, "\n")
// 	for _, line := range lines {
// 		if strings.Contains(line, "Duration: ") {
// 			duration = strings.Split(line, "Duration: ")[1]
// 			duration = strings.Split(duration, ",")[0]
// 			break
// 		}
// 	}
// 	return duration
// }

func GetVideoDuration(filePath string) (string, error) {
	if avformat.AvformatNetworkInit() != 0 {
		return "", fmt.Errorf("failed to initialize network")
	}
	defer avformat.AvformatNetworkDeinit()

	avformat.AvRegisterAll()
	formatContext := avformat.AvformatAllocContext()

	if avformat.AvformatOpenInput(&formatContext, filePath, nil, nil) != 0 {
		return "", fmt.Errorf("failed to open file")
	}
	defer avformat.AvformatCloseInput(formatContext)

	if avformat.AvformatFindStreamInfo(formatContext, nil) < 0 {
		return "", fmt.Errorf("failed to find stream info")
	}

	stream := formatContext.Streams()[0]
	duration := stream.Duration() * avutil.AV_TIME_BASE / avutil.Second

	// Calculate the hours, minutes, and seconds from the duration
	hours := duration / 3600
	minutes := (duration % 3600) / 60
	seconds := duration % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds), nil
}
