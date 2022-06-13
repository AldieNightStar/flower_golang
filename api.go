package flower

import (
	"github.com/AldieNightStar/golisper"
)

func Load(src string) (*Scope, error) {
	tags, err := golisper.Parse(src)
	if err != nil {
		return nil, err
	}
	return NewScopeWithBuiltIns(tags, 0), nil
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
