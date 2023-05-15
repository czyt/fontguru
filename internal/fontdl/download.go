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
	defer fFont.Close()
	_, err = io.Copy(fFont, resp.Body)
	if err != nil {
		return err
	}
	return nil

}
