package logging

import "go.uber.org/zap"
import "time"
import ."fmt"

type ss struct{
	name string
}

func mainlog() {
	var a  int =3
	var b = "hhhhhhhhhhhhhhhh"
	var s = &ss{name:"loo"}
	var e = [1]string{"hvag"}
	var m = map[int]string{3:"hvag"}
	Println(s)
	sugar := zap.NewExample().Sugar()
defer sugar.Sync()
sugar.Infow("failed to fetch URL",
  "url", "http://example.com",
  "attempt",a,
  "backoff", time.Second,
  "hvag",b,
  "fff",s,
  "list",e,
  "dict",m,
)
sugar.Debugw("failed to fetch URL: %s", "http://example.com")
zap.S().Infow("An INFO message", "iteration", 1)
}