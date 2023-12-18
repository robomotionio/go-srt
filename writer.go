package gosrt

import (
	"fmt"
	"os"
	"time"
)

func WriteFile(fileName string, subtitles []Subtitle) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, subtitle := range subtitles {
		_, err := fmt.Fprintf(file, "%d\n%s --> %s\n%s\n\n", subtitle.Number, formatTime(subtitle.Start), formatTime(subtitle.End), subtitle.Text)
		if err != nil {
			return err
		}
	}

	return nil
}

func formatTime(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := int(duration.Milliseconds()) % 1000

	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, milliseconds)
}
