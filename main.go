package main

import (
	"fontguru/internal/fontdl"
	"fontguru/internal/fontinstall"
	"fontguru/internal/resource"
	"github.com/briandowns/spinner"
	"log"
	"sync"
	"time"
)

func main() {
	resourceInfoList, err := resource.GetFontResourceInfoList()
	if err != nil {
		log.Fatalln(err)
	}
	wg := sync.WaitGroup{}
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "正在安装字体，请稍后..."
	s.Start()
	wg.Add(len(resourceInfoList))
	for _, info := range resourceInfoList {
		go processFont(info, wg)
	}
	wg.Wait()
	s.Stop()
}

func processFont(info resource.FontResourceInfo, wg sync.WaitGroup) {
	defer wg.Done()
	if err := fontdl.Download(info.DownloadUrl, info.FontFileName); err != nil {
		log.Println("download font:", info.FontFileName, "failed:", err)
		return
	}
	if err := fontinstall.InstallFont(info.FontFileName); err != nil {
		log.Println("install font failed:", info.FontFileName)
	}
	log.Println("install font success:", info.FontFileName)
}
