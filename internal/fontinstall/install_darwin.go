//go:build darwin

package fontinstall

import (
	"errors"
	"golang.org/x/text/fontconfig"
	"io/ioutil"
)

func InstallFont(localPath string) error {
	fontData, err := ioutil.ReadFile(localPath)
	if err != nil {
		return errors.Wrap(err, "failed to read font file")
	}

	fontURL := fontconfig.NewFileURL(localPath)
	fontURL.Normalize()
	defer fontURL.Close()

	fontDescriptors, err := fontconfig.ParseFont(fontData, fontURL)
	if err != nil {
		return errors.Wrap(err, "failed to parse font file")
	}

	if len(fontDescriptors) < 1 {
		return errors.New("no font descriptors found in file")
	}

	fontName := fontDescriptors[0].Family
	if fontName == "" {
		return errors.New("no font family name found in file")
	}

	font := fontconfig.NewFont(fontName)
	defer font.Close()

	err = fontconfig.InstallFont(fontData, font)
	if err != nil {
		return errors.Wrap(err, "failed to install font")
	}

	return nil
}
