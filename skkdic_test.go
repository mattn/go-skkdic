package skkdic

import (
	"strings"
	"testing"
)

func TestSkkDict(t *testing.T) {
	r := strings.NewReader(`
;; test
;; okuri-ari entries.
われt /我/
わざわi /災/禍;(字義:落とし穴)/厄;<rare>/
;; okuri-nasi entries.
われ /我/吾/
われき /和暦/
    `)
	dict := NewDict()
	err := dict.Load(r)
	if err != nil {
		t.Fatal(err)
	}

	entries := dict.SearchOkuriAri("わざわi")
	if len(entries) != 1 {
		t.Fatal("only one entry should be returned")
	}
	if entries[0].Label != "わざわi" {
		t.Fatalf("label should be %q but %q", "わざわi", entries[0].Label)
	}
	if len(entries[0].Words) != 3 {
		t.Fatal("words should be 3")
	}
	if entries[0].Words[0].Text != "災" {
		t.Fatalf("first text should be %q but %q", "災", entries[0].Words[0].Text)
	}
	if entries[0].Words[1].Text != "禍" {
		t.Fatalf("second text should be %q but %q", "禍", entries[0].Words[1].Text)
	}
	if entries[0].Words[1].Desc != "(字義:落とし穴)" {
		t.Fatalf("second text should be %q but %q", "", entries[0].Words[1].Desc)
	}

	entries = dict.SearchOkuriAri("わさび")
	if len(entries) != 0 {
		t.Fatal("any entries should not be returned")
	}
	entries = dict.SearchOkuriNasi("われ")
	if len(entries) != 1 {
		t.Fatal("only one entry should be returned")
	}
	if entries[0].Label != "われ" {
		t.Fatalf("label should be %q but %q", "われ", entries[0].Label)
	}
	if len(entries[0].Words) != 2 {
		t.Fatal("words should be 2")
	}
	if entries[0].Words[0].Text != "我" {
		t.Fatalf("first text should be %q but %q", "我", entries[0].Words[0].Text)
	}
	if entries[0].Words[1].Text != "吾" {
		t.Fatalf("second text should be %q but %q", "吾", entries[0].Words[1].Text)
	}
}
