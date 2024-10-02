package football_data

import (
	"net/http"
	"os"
	"io"
	"fmt"
)


func DownloadIcsFile() (ics_file string, err error) {
	fmt.Println("Fetching data from sports.yahoo.com...")

	resp, err := http.Get("https://sports.yahoo.com/soccer/teams/man-utd/ical.ics")
	if err != nil {
        return "", err
    }

	defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
        return "", err
    }

	tmpFile, err := os.CreateTemp("", "fixtures.*.ics")
	if err != nil {
        return "", err
    }

    if err = os.WriteFile(tmpFile.Name(), bodyBytes, 0644); err != nil {
        return "", err
	}

	return tmpFile.Name(), nil
}