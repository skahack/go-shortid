package shortid

import (
	"regexp"
	"testing"
)

func Testshuffle(t *testing.T) {
	var actual string

	c := newChars()

	c.setSeed(1)
	actual = c.shuffle()
	if actual != "ylZM7VHLvOFcohp01x-fXNr8P_tqin6RkgWGm4SIDdK5s2TAJebzQEBUwuY9j3aC" {
		t.Errorf("error %v", actual)
	}

	c.setSeed(1)
	actual = c.shuffle()
	if actual != "ylZM7VHLvOFcohp01x-fXNr8P_tqin6RkgWGm4SIDdK5s2TAJebzQEBUwuY9j3aC" {
		t.Errorf("error %v", actual)
	}

	c.setSeed(1234)
	actual = c.shuffle()
	if actual != "ef4w9iMboqLOQdWu3hKI72A0VZpCtzDlXk5_a6cFSNYGnH-gmsP1UBxvTRJjE8ry" {
		t.Errorf("error %v", actual)
	}

	c.setSeed(1)
	actual = c.shuffle()
	if actual != "ylZM7VHLvOFcohp01x-fXNr8P_tqin6RkgWGm4SIDdK5s2TAJebzQEBUwuY9j3aC" {
		t.Errorf("error %v", actual)
	}

	actual = c.shuffle()
	if actual != "ylZM7VHLvOFcohp01x-fXNr8P_tqin6RkgWGm4SIDdK5s2TAJebzQEBUwuY9j3aC" {
		t.Errorf("error %v", actual)
	}

	c.setSeed(1)
	c.setCharacters("")
	c.setCharacters(c.shuffle())

	actual = c.shuffle()
	if actual != "WN3JLu5ARbdoPx_ylgC09eqvzant-8HEX1YKr7BsIhTViZUm2pcGQD4wk6jOfMFS" {
		t.Errorf("error %v", actual)
	}
}

func TestCharactor(t *testing.T) {
	var err error
	matchStr := ""
	matched := false

	c := newChars()

	c.setCharacters("")

	err = c.setCharacters("-‾zʎxʍʌnʇsɹbdouɯlʞɾıɥƃɟǝpɔqɐzʎxʍʌnʇsɹbdouɯlʞɾıɥƃɟǝpɔqɐ9876543210")
	matchStr = "Custom alphabet for shortId must be 64 unique characters. These"
	if matched, _ = regexp.MatchString(matchStr, err.Error()); !matched {
		t.Errorf("%v", err)
	}

	err = c.setCharacters("abc")
	matchStr = "Custom alphabet for shortId must be 64 unique characters. You"
	if matched, _ = regexp.MatchString(matchStr, err.Error()); !matched {
		t.Errorf("%v", err)
	}
}

func TestshuffleWithUnicode(t *testing.T) {
	c := newChars()
	c.setSeed(1)
	c.setCharacters("①②③④⑤⑥⑦⑧⑨⑩⑪⑫ⒶⒷⒸⒹⒺⒻⒼⒽⒾⒿⓀⓁⓂⓃⓄⓅⓆⓇⓈⓉⓊⓋⓌⓍⓎⓏⓐⓑⓒⓓⓔⓕⓖⓗⓘⓙⓚⓛⓜⓝⓞⓟⓠⓡⓢⓣⓤⓥⓦⓧⓨⓩ")
	actual := c.shuffle()
	if actual != "ⓌⒿⓧⓚ⑧ⓣⓕⓙⓉⓜⓓⒶⓂⒻⓃ①②ⓋⓩⒹⓥⓛⓅ⑨ⓝⓨⓇⓄⒼⓁ⑦ⓟⒾⒺⓤⓔⓀ⑤ⓠⓖⓑⒷⓘ⑥Ⓠ③ⓡⓎⓗⒸ⑫ⓍⓞⓒⓏⓢⓊⓈⓦ⑩Ⓗ④⑪ⓐ" {
		t.Errorf("error %v", actual)
	}
}

func TestRandom(t *testing.T) {
	var r float64 = 0

	c := newChars()

	c.setSeed(0)
	r = c.rand.random()
	if r != 0.21132115912208504 {
		t.Errorf("expect 0.21132115912208504, but actual %v", r)
	}

	c.setSeed(0)
	r = c.rand.random()
	if r != 0.21132115912208504 {
		t.Errorf("expect 0.21132115912208504, but actual %v", r)
	}

	c.setSeed(1)
	r = c.rand.random()
	if r != 0.2511917009602195 {
		t.Errorf("expect 0.2511917009602195, but actual %v", r)
	}

	c.setSeed(2)
	r = c.rand.random()
	if r != 0.2910622427983539 {
		t.Errorf("expect 0.2910622427983539, but actual %v", r)
	}

	c.setSeed(0.21132115912208504)
	r = c.rand.random()
	if r != 0.2197466482381452 {
		t.Errorf("expect 0.2197466482381452, but actual %v", r)
	}

	c.setSeed(0)
	r = c.rand.random()
	if r != 0.21132115912208504 {
		t.Errorf("expect 0.21132115912208504, but actual %v", r)
	}
}
