package flower

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/AldieNightStar/golisper"
)

func builtinHTTP(s *Scope) {
	h := newBuitinDict()
	h.SetValue("request", SFunc(func(s *Scope, args []*golisper.Value) (any, error) {
		if len(args) < 3 {
			return nil, errNotEnoughArgs(s.LastLine, "request", 3, len(args))
		}
		method, err := EvalCast("request", s, args[0], "")
		if err != nil {
			return nil, newErrLineName(s.LastLine, "request", err.Error())
		}
		url, err := EvalCast("request", s, args[1], "")
		if err != nil {
			return nil, newErrLineName(s.LastLine, "request", err.Error())
		}
		body, err := EvalCast("request", s, args[2], "")
		if err != nil {
			return nil, newErrLineName(s.LastLine, "request", err.Error())
		}
		request, err := http.NewRequest(method, url, strings.NewReader(body))
		if err != nil {
			return nil, newErrLineName(s.LastLine, "request", err.Error())
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return nil, newErrLineName(s.LastLine, "request", err.Error())
		}
		dat, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, newErrLineName(s.LastLine, "request", err.Error())
		}
		return string(dat), nil
	}))
	s.Memory["http"] = h
}
