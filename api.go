package flower

import (
	"io/ioutil"
	"os"

	"github.com/AldieNightStar/golisper"
)

func Load(src string) (*Scope, error) {
	tags, err := golisper.Parse(src)
	if err != nil {
		return nil, err
	}
	return NewScope(tags, 0, nil), nil
}

func LoadFromFile(name string) (*Scope, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	dat, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return Load(string(dat))
}
