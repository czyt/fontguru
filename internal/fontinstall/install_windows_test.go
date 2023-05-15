package fontinstall

import "testing"

func TestInstallFont(t *testing.T) {
	if err := InstallFont("testdata/mononoki-Regular.ttf"); err != nil {
		t.Fatal(err)
	}
	t.Log("done")
}
