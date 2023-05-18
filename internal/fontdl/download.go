package fontdl

import (
	"io"
	"net/http"
	"os"
)

func Download(fontUrl string, fontName string) error {
	resp, err := http.Get(fontUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fFont, err := os.Create(fontName)
	if err != nil {
		return err
	}

	_, err = io.Copy(fFont, resp.Body)
	if err != nil {
		fFont.Close()
		os.Remove(fontName)
		return err
	}
	fFont.Close()
	return nil

}
