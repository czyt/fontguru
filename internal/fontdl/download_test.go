package fontdl

import "testing"

func TestDownloadFont(t *testing.T) {
	err := Download("http://fontguru.czyt.tech/3270NerdFont-Condensed.ttf?e=1684073625\\u0026token=lUkkMTqUK-fY7t6Tbg7zq-p3iaopntRMOQOkEEDW:sJ_b3lvDcqQ7KEkWYjUqHFGFUXI=",
		"test.ttf")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
