package skkdic

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	okuriAri  = ";; okuri-ari entries."
	okuriNasi = ";; okuri-nasi entries."
)

type Word struct {
	Text string
	Desc string
}

type Entry struct {
	Label string
	Words []Word
}

type Dict struct {
	okuriAri  []Entry
	okuriNasi []Entry
}

func (d *Dict) SearchOkuriAri(s string) []Entry {
	ret := make([]Entry, 0)
	for _, e := range d.okuriAri {
		if e.Label == s {
			ret = append(ret, e)
		}
	}
	return ret
}

func (d *Dict) SearchOkuriNasi(s string) []Entry {
	ret := make([]Entry, 0)
	for _, e := range d.okuriNasi {
		if e.Label == s {
			ret = append(ret, e)
		}
	}
	return ret
}

func (d *Dict) SearchOkuriNasiPrefix(s string) []Entry {
	ret := make([]Entry, 0)
	for _, e := range d.okuriNasi {
		if strings.HasPrefix(e.Label, s) {
			ret = append(ret, e)
		}
	}
	return ret
}

func New() *Dict {
	return &Dict{}
}

func (d *Dict) Load(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Text() != okuriAri {
			break
		}
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == okuriNasi {
			break
		}
		if strings.HasPrefix(line, ";;") {
			continue
		}
		e, err := parseEntry(line)
		if err != nil {
			continue
		}
		d.okuriAri = append(d.okuriAri, *e)
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ";;") {
			continue
		}
		e, err := parseEntry(line)
		if err != nil {
			continue
		}
		d.okuriNasi = append(d.okuriNasi, *e)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func parseWord(v string) Word {
	n := strings.Index(v, ";")
	if n < 0 {
		return Word{Text: v}
	}
	return Word{
		Text: v[0:n],
		Desc: v[n+1:],
	}
}

func parseEntry(s string) (*Entry, error) {
	if strings.HasPrefix(s, ";;") {
		return nil, nil
	}
	s = strings.TrimRight(s, " \t\r\n")
	items := strings.SplitN(s, " ", 2)
	if items == nil || len(items) != 2 {
		return nil, fmt.Errorf("invalid format: %s", s)
	}
	label := items[0]
	values := strings.Split(strings.Trim(items[1], "/"), "/")
	words := make([]Word, len(values))
	for i, v := range values {
		words[i] = parseWord(v)
	}
	return &Entry{
		Label: label,
		Words: words,
	}, nil
}
