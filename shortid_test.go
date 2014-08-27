package shortid

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	g := Generator()

	g.SetCharacters("")
	g.SetSeed(1)

	ids := map[string]bool{}

	for i := 0; i < 50000; i++ {
		id := g.Generate()
		_, ok := ids[id]

		if ok == true {
			t.Errorf("got %v", id)
			break
		}

		if len(id) > 17 {
			t.Errorf("id should length below 17", id)
			break
		}

		ids[id] = true
	}
}

func TestGenerateUnicodeChars(t *testing.T) {
	g := Generator()
	chars := "①②③④⑤⑥⑦⑧⑨⑩⑪⑫ⒶⒷⒸⒹⒺⒻⒼⒽⒾⒿⓀⓁⓂⓃⓄⓅⓆⓇⓈⓉⓊⓋⓌⓍⓎⓏⓐⓑⓒⓓⓔⓕⓖⓗⓘⓙⓚⓛⓜⓝⓞⓟⓠⓡⓢⓣⓤⓥⓦⓧⓨⓩ"
	g.SetCharacters(chars)
	g.SetSeed(1)

	id := g.Generate()
	if !strings.ContainsAny(id, chars) {
		t.Errorf("should contain Unicode characters. %v", id)
	}
}

func TestDecodeWorker(t *testing.T) {
	var decode map[string]int
	g := Generator()
	g.SetCharacters("")
	g.SetSeed(1)

	g.SetWorker(0)
	decode = g.Decode(g.Generate())
	if decode["worker"] != 0 {
		t.Errorf("error, %v", decode)
	}

	g.SetWorker(1)
	decode = g.Decode(g.Generate())
	if decode["worker"] != 1 {
		t.Errorf("error, %v", decode)
	}

	g.SetWorker(2)
	decode = g.Decode(g.Generate())
	if decode["worker"] != 2 {
		t.Errorf("error, %v", decode)
	}

	g.SetWorker(3)
	decode = g.Decode(g.Generate())
	if decode["worker"] != 3 {
		t.Errorf("error, %v", decode)
	}

	g.SetWorker(15)
	decode = g.Decode(g.Generate())
	if decode["worker"] != 15 {
		t.Errorf("error, %v", decode)
	}
}

func BenchmarkGenerate(b *testing.B) {
	g := Generator()
	g.SetCharacters("")
	g.SetSeed(1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Generate()
	}
}

func BenchmarkLookup(b *testing.B) {
	g := Generator()
	g.SetCharacters("")
	g.SetSeed(1)
	chars := strings.Split(g.chars.shuffle(), "")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.lookup(chars, 15)
	}
}

func BenchmarkRandomByte(b *testing.B) {
	g := Generator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.randomByte()
	}
}
