package shortid

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

// Ignore all milliseconds before a certain time to reduce
// the size of the date entropy without sacrificing uniqueness.
// This number should be updated every year or so to keep the generated id short.
// To regenerate `time.Now().UnixNano()/1000000` and bump the version. Always bump the version!
const reduceTime int64 = 1403265799803

type Gen struct {
	// don't change unless we change the algos or reduceTime
	// must be an integer and less than 16
	version int

	// if you are using cluster or multiple servers use this to make each instance
	// has a unique value for worker
	clusterWorkerId int

	// Counter is used when shortId is called multiple times in one second.
	counter int

	// Remember the last time shortId was called in case counter is needed.
	previousSeconds float64

	chars *Chars

	randCache *rand.Rand
}

func Generator() *Gen {
	return &Gen{
		version:         0,
		clusterWorkerId: 0,
		counter:         0,
		previousSeconds: 0.0,
		chars:           newChars(),
		randCache:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Gen) Generate() string {
	str := ""
	var t int64 = time.Now().UnixNano() / 1000000
	seconds := math.Ceil(float64(t-reduceTime) * 0.01)

	if seconds == g.previousSeconds {
		g.counter++
	} else {
		g.counter = 0
		g.previousSeconds = seconds
	}

	str = str + g.encode(g.version)
	str = str + g.encode(g.clusterWorkerId)
	if g.counter > 0 {
		str = str + g.encode(g.counter)
	}
	str = str + g.encode(int(seconds))

	return str
}

func (g *Gen) SetWorker(id int) *Gen {
	g.clusterWorkerId = id
	return g
}

func (g *Gen) Decode(id string) map[string]int {
	var alphabet = g.chars.shuffle()

	return map[string]int{
		"version": strings.Index(alphabet, string(id[0])) & 0x0f,
		"worker":  strings.Index(alphabet, string(id[1])) & 0x0f,
	}
}

func (g *Gen) SetSeed(seed float64) {
	g.chars.setSeed(seed)
}

func (g *Gen) SetCharacters(char string) error {
	return g.chars.setCharacters(char)
}

//
// Private
//

func (g *Gen) encode(number int) string {
	done := false
	str := make([]string, 20)
	loopCounter := uint32(0)
	alphabetshuffled := strings.Split(g.chars.shuffle(), "")

	for !done {
		idx := ((uint32(number) >> (4 * loopCounter)) & 0x0f) | (g.randomByte() & 0x30)
		str = append(str, g.lookup(alphabetshuffled, idx))
		done = float64(number) < math.Pow(16, float64(loopCounter+1))
		loopCounter++
	}

	return strings.Join(str, "")
}

func (g *Gen) randomByte() uint32 {
	r := g.randCache
	return uint32(r.Intn(256))
}

func (g *Gen) lookup(alphabetshuffled []string, index uint32) string {
	return alphabetshuffled[index]
}
