package thieves

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Thief struct {
	url string
	val string
}

func (this *Thief) Cut(from, to string) *Thief {
	var i, j = strings.Index(this.val, from), strings.LastIndex(this.val, to)
	max := len(this.val)
	if i <= -1 {
		i = 0
	}
	if j > max {
		j = max
	}
	this.val = this.val[i:j]
	return this
}

func (this *Thief) FindAll(expr string) [][]string {
	re, _ := regexp.Compile(expr)
	list := re.FindAllStringSubmatch(this.val, -1)
	return list
}

func (this *Thief) HTMLToLower() *Thief {
	re, _ := regexp.Compile(`\<[\S\s]+?\>`)
	this.val = re.ReplaceAllStringFunc(this.val, strings.ToLower)
	return this
}

func (this *Thief) Replace(expr string, repl string) *Thief {
	re, _ := regexp.Compile(expr)
	this.val = re.ReplaceAllString(this.val, repl)
	return this
}

func (this *Thief) Trim(expr string) *Thief {
	return this.Replace(expr, "")
}

func (this *Thief) TrimAll() *Thief {
	return this.TrimConsecutiveBlank().TrimNewLine().TrimScript().TrimStyle()
}

func (this *Thief) TrimConsecutiveBlank() *Thief {
	return this.Trim(`\s{2,}`)
}

func (this *Thief) TrimNewLine() *Thief {
	return this.Trim(`\r|\n`)
}

func (this *Thief) TrimScript() *Thief {
	return this.Trim(`\<(?i:script)[\S\s]+?\</(?i:script)\>`)
}

func (this *Thief) TrimStyle() *Thief {
	return this.Trim(`\<(?i:style)[\S\s]+?\</(?i:style)\>`)
}

func (this *Thief) Val() string {
	return this.val
}

func New(url string) *Thief {
	thief := &Thief{url, ""}

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("GET %s error: %s\r\n", url, err.Error())
		return thief
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("READ res.Body error: %s\r\n", err.Error())
		return thief
	}

	thief.val = string(bytes)
	return thief
}