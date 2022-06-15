package flower

import (
	"github.com/AldieNightStar/golisper"
)

func Load(src string) (*Scope, error) {
	values, err := golisper.Parse(src)
	if err != nil {
		return nil, err
	}
	return NewScopeWithBuiltIns(values, 0), nil
}

func LoadFromFile(name string) (*Scope, error) {
	src, err := utilReadFileAsString(name)
	if err != nil {
		return nil, err
	}
	return Load(
		utilIgnoreHashBangAtStart(src),
	)
}
