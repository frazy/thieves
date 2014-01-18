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
	if len(this.val) == 0 {
		return this
	}

	var i, j = strings.Index(this.val, from), strings.LastIndex(this.val, to)
	max := len(this.val)
	if i <= -1 {
		i = 0
	}
	if j <= -1 {
		j = 0
	}
	if j > max {
		j = max
	}
	if i == 0 && j == 0 {
		this.val = ""
		return this
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
	return this.TrimScript().TrimStyle().TrimConsecutiveBlank().TrimNewLine()
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
	return NewWithHeader(url, http.Header{"Accept": []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"}})
}

func NewWithHeader(url string, headers http.Header) *Thief {
	thief := &Thief{url, ""}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("NewRequest %s error: %s\r\n", url, err.Error())
		return thief
	}
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("GET %s error: %s\r\n", url, err.Error())
		return thief
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return thief
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("READ res.Body error: %s\r\n", err.Error())
		return thief
	}

	thief.val = string(bytes)
	return thief
}
