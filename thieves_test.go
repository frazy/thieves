package thieves

import (
	"testing"
)

func Test_Trim(t *testing.T) {
	url := "https://github.com/"
	thief := New(url)
	val := thief.HTMLToLower().TrimAll().Val()
	t.Logf("val=%s\r\n", val)
}

func Test_Trim_2(t *testing.T) {
	thief := Thief{"", "test<ScRipt>sdasdasdasd</Script>1234<STYLE>dddd</STYLE>lll"}
	val := thief.HTMLToLower().TrimAll().Val()
	t.Logf("val=%s\r\n", val)
}

func Test_Cut(t *testing.T) {
	url := "https://github.com/"
	thief := New(url)
	val := thief.HTMLToLower().TrimAll().Cut("<body", "</body>").Val()
	t.Logf("val=%s\r\n", val)
}

func Test_FindAll(t *testing.T) {
	url := "https://github.com/"
	thief := New(url)
	val := thief.HTMLToLower().TrimAll().Cut("<body", "</body>").FindAll(`<a href="(.+?)"`)
	for i, v := range val {
		t.Logf("i=%d, v=%s\r\n", i, v[1])
	}
}
