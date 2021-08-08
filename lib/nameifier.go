package lib

import (
	_ "embed"
	"encoding/json"
	"hash/fnv"
)

//go:embed data/nouns.json
var nouns []byte

//go:embed data/adjectives.json
var adjectives []byte

type NameGenerator interface {
	Nameify(seed string) (string, error)
}

func hash(s string, max int) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % uint32(max)
}

func loadJson(file []byte, nameifier *Nameifier) error {
	err := json.Unmarshal(file, &nameifier)
	if err != nil {
		return err
	}
	return nil
}

type Nameifier struct {
	Nouns      []string
	Adjectives []string
}

func NewNameifier() (*Nameifier, error) {
	n := &Nameifier{}
	err := loadJson(nouns, n)
	if err != nil {
		return n, err
	}
	err = loadJson(adjectives, n)
	if err != nil {
		return n, err
	}
	return n, nil
}

func (n *Nameifier) Nameify(seed string) (string, error) {
	return n.Adjectives[hash(seed, len(n.Adjectives))] + "-" + n.Nouns[hash(seed, len(n.Nouns))], nil
}
