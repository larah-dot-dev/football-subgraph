package football_data

import (
	"net/http"
	"os"
	"io"
	"github.com/rs/zerolog/log"
)


func DownloadIcsFile() (ics_file string, err error) {
	log.Print("Fetching data from manutd.com...")

	resp, err := http.Get("https://www.manutd.com/en/Manchester_United.ics")
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