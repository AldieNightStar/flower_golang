package flower

import "github.com/AldieNightStar/golisper"

// ====================================
// ====================================

type builtinIterator interface {
	Iteration() builtinIteration
}

type builtinIteration interface {
	Iterate() (any, error)
}

// ====================================
// ====================================

// for command (of 1 2 3 ...)
type builtinIteratorOfArgs struct {
	scope *Scope
	vals  []*golisper.Value
}

func (it *builtinIteratorOfArgs) Iteration() builtinIteration {
	return &builtinIteratorOfArgsIteration{it, 0}
}

// for command (of 1 2 3 ...)
type builtinIteratorOfArgsIteration struct {
	iter *builtinIteratorOfArgs
	ptr  int
}

func (it *builtinIteratorOfArgsIteration) Iterate() (any, error) {
	if it.ptr >= len(it.iter.vals) {
		return nil, nil
	}
	res := it.iter.vals[it.ptr]
	evl, err := it.iter.scope.Eval(res)
	if err != nil {
		return nil, err
	}
	it.ptr += 1
	return evl, nil
}

// ====================================
// ====================================

type builtinRangeIterator struct {
	min int
	max int
}

func (it *builtinRangeIterator) Iteration() builtinIteration {
	return &builtinRangeIteration{it, it.min}
}

type builtinRangeIteration struct {
	iter  *builtinRangeIterator
	count int
}

func (it *builtinRangeIteration) Iterate() (any, error) {
	if it.count <= it.iter.max {
		res := it.count
		it.count += 1
		return float64(res), nil
	}
	return nil, nil
}

// ====================================
// ====================================

type builtinDictKeysIterator struct {
	dict builtinKeysHolder
}

func (it *builtinDictKeysIterator) Iteration() builtinIteration {
	arr := make([]any, 0, 32) // Convert []string -> []any
	for k := range it.dict.Keys() {
		arr = append(arr, k)
	}
	return &builtinArrayIteration{arr, 0}
}

// ====================================
// ====================================

type builtinArrayIteration struct {
	arr []any
	pos int
}

func (it *builtinArrayIteration) Iterate() (any, error) {
	if it.pos >= len(it.arr) {
		return nil, nil
	}
	res := it.arr[it.pos]
	it.pos += 1
	return res, nil
}

// ====================================
// ====================================

type builtinStringIterator struct {
	str string
}

func (it *builtinStringIterator) Iteration() builtinIteration {
	return &buitlinStringIteration{it, 0}
}

type buitlinStringIteration struct {
	iter *builtinStringIterator
	pos  int
}

func (it *buitlinStringIteration) Iterate() (any, error) {
	if it.pos >= len(it.iter.str) {
		return nil, nil
	}
	res := it.iter.str[it.pos : it.pos+1]
	it.pos += 1
	return res, nil
}

// ====================================
// ====================================

type builtinForeverIterator int8

func (builtinForeverIterator) Iteration() builtinIteration {
	r := new(builtinForeverIteration)
	*r = 0
	return r
}

type builtinForeverIteration int

func (n *builtinForeverIteration) Iterate() (any, error) {
	r := *n
	*n += 1
	return float64(r), nil
}

// ====================================
// ====================================

// Iteration() builtinIteration
// Iterate() (any, error)

type builtinBlockIterator struct {
	block  *codeBlock
	alias  string
	parent *Scope
}

func (it *builtinBlockIterator) Iteration() builtinIteration {
	return &builtinBlockIteration{it, 0}
}

type builtinBlockIteration struct {
	iter *builtinBlockIterator
	ptr  int
}

func (it *builtinBlockIteration) Iterate() (any, error) {
	scope := it.iter.block.Load(it.iter.parent, map[string]any{
		it.iter.alias: float64(it.ptr),
	})
	builtinAddReturn(scope)
	it.ptr += 1
	res, err := scope.Run()
	if err != nil {
		return nil, err
	}
	return res, nil
}
