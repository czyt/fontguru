package resource

import (
	"io"
	"net/http"
	"os"
)

type FontResourceInfo struct {
	FontFileName string `json:"font_file_name"`
	DownloadUrl  string `json:"download_url"`
	PutTime      int64  `json:"put_time"`
	Hash         string `json:"hash"`
	FileSize     int    `json:"file_size"`
	MimeType     string `json:"mime_type"`
	Md5          string `json:"md5"`
}

func downloadFont(fontUrl string, fontName string) error {
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
