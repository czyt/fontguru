package resource

import (
	"encoding/json"
	"net/http"
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

func GetFontResourceInfoList() ([]FontResourceInfo, error) {
	resp, err := http.Get("https://hub.czyt.tech/fonts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fontResourceInfoList := make([]FontResourceInfo, 0, 100)
	err = json.NewDecoder(resp.Body).Decode(&fontResourceInfoList)
	if err != nil {
		return nil, err
	}
	return fontResourceInfoList, nil
}
