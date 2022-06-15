package flower

import (
	"io/ioutil"
	"os"

	"github.com/AldieNightStar/golisper"
)

func builtinFiles(s *Scope) {
	fs := newBuitinDict()

	fs.m["read"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "fs read", 1, 0)
		}
		name, err := EvalCast("dict set", s, args[0], "")
		if err != nil {
			return nil, err
		}
		src, err := utilReadFileAsString(name)
		if err != nil {
			return nil, err
		}
		return src, nil
	})
	fs.m["write"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 2 {
			return nil, errNotEnoughArgs(s.LastLine, "fs write", 2, len(args))
		}
		name, err := EvalCast("fs write", s, args[0], "")
		if err != nil {
			return nil, err
		}
		src, err := EvalCast("fs write", s, args[1], "")
		if err != nil {
			return nil, err
		}
		f, err := os.Create(name)
		if err != nil {
			return nil, newErrLineName(s.LastLine, "fs write", err.Error())
		}
		defer f.Close()
		cnt, err := f.WriteString(src)
		if err != nil {
			return nil, newErrLineName(s.LastLine, "fs write", err.Error())
		}
		return float64(cnt), nil
	})
	fs.m["delete"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "fs delete", 1, 0)
		}
		name, err := EvalCast("fs delete", s, args[0], "")
		if err != nil {
			return nil, err
		}
		err = os.Remove(name)
		if err != nil {
			return false, nil
		}
		return true, nil
	})
	fs.m["list"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "fs delete", 1, 0)
		}
		path, err := EvalCast("fs list", s, args[0], "")
		if err != nil {
			return nil, err
		}
		dirArr, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, err
		}
		arr := make([]any, 0, len(dirArr))
		for _, d := range dirArr {
			dict := newBuitinDict()
			dict.m["name"] = d.Name()
			dict.m["size"] = float64(d.Size())
			dict.m["isfile"] = !d.IsDir()
			arr = append(arr, dict)
		}
		return &builtinList{arr}, nil
	})
	fs.m["import"] = SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 1 {
			return nil, errNotEnoughArgs(s.LastLine, "fs import", 1, 0)
		}
		name, err := EvalCast("fs import", s, args[0], "")
		if err != nil {
			return nil, err
		}
		src, err := utilReadFileAsString(name)
		if err != nil {
			return nil, newErrLineName(s.LastLine, "fs import", "Can't import: "+name+"\n\tReason: "+err.Error())
		}
		importScope, err := Load(src)
		if err != nil {
			return nil, newErrLineName(s.LastLine, "fs import", "Can't import: "+name+"\n\tReason: "+err.Error())
		}
		builtinAddReturn(importScope)
		res, err := importScope.Run()
		if err != nil {
			return nil, newErrLineName(s.LastLine, "fs import", "Can't import: "+name+"\n\tReason: "+err.Error())
		}
		return res, nil
	})

	s.Memory["fs"] = fs
}
