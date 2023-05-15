package resource

import "testing"

func TestGetFontResourceInfoList(t *testing.T) {
	resourceInfoList, err := GetFontResourceInfoList()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(resourceInfoList))
}
