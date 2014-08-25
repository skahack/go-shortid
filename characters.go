package shortid

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

const (
	original string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
)

type Chars struct {
	alphabet string
	shuffled string
	rand     *Random
}

func newChars() *Chars {
	return &Chars{
		alphabet: "",
		shuffled: "",
		rand:     newRandom(),
	}
}

func (c *Chars) setSeed(newSeed float64) {
	c.rand.setSeed(newSeed)

	if c.rand.isNewSeed(newSeed) {
		c.reset()
		c.rand.setPrevSeed(newSeed)
	}
}

func (c *Chars) setCharacters(newChar string) error {
	if len(newChar) == 0 {
		if c.alphabet != original {
			c.alphabet = original
			c.reset()
		}
		return nil
	}

	if newChar == c.alphabet {
		return nil
	}

	if utf8.RuneCountInString(newChar) != len(original) {
		err := fmt.Errorf("Custom alphabet for shortId must be %d unique characters. You submitted %d characters: %s",
			len(original), len(newChar), newChar)
		return err
	}

	unique := ""
	ss := strings.Split(newChar, "")
	for _, v := range ss {
		if strings.Count(newChar, v) > 1 && !strings.Contains(unique, v) {
			unique += v
		}
	}

	if len(unique) > 0 {
		err := fmt.Errorf("Custom alphabet for shortId must be %d unique characters. These characters were not unique: %s",
			len(original), unique)
		return err
	}

	c.alphabet = newChar
	c.reset()

	return nil
}

func (c *Chars) shuffle() string {
	if len(c.shuffled) > 0 {
		return c.shuffled
	}

	if len(c.alphabet) == 0 {
		c.setCharacters(original)
	}

	var target []string
	r := c.rand.random()
	source := strings.Split(c.alphabet, "")

	for len(source) > 0 {
		r = c.rand.random()
		idx := int(math.Floor(r * float64(len(source))))
		target = append(target, source[idx])
		if idx == 0 {
			source = source[1:]
		} else if idx == len(source)-1 {
			source = source[:idx]
		} else {
			source = append(source[:idx], source[idx+1:]...)
		}
	}

	c.shuffled = strings.Join(target, "")
	return c.shuffled
}

func (c *Chars) reset() {
	c.shuffled = ""
}
