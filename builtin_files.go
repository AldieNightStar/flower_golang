package flower

import (
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
		f, err := os.Open(name)
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
		res, err := importScope.Run()
		if err != nil {
			return nil, newErrLineName(s.LastLine, "fs import", "Can't import: "+name+"\n\tReason: "+err.Error())
		}
		return res, nil
	})

	s.Memory["fs"] = fs
}
