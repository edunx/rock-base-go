package base

import (
	"github.com/edunx/lua"
	"os"
)

type stdout struct {
	lua.Super
	count int
}

func (self *stdout) Write(v interface{} ) error {
	var str string
	var bytes []byte
	var msg lua.Message

	str, ok := v.(string)
	if ok {
		os.Stdout.WriteString(str)
		goto DONE
	}

	bytes, ok = v.([]byte)
	if ok {
		os.Stdout.Write(bytes)
		goto DONE
	}

	msg, ok = v.(lua.Message)
	if ok {
		os.Stdout.Write(msg.Byte())
		goto DONE
	}

DONE:
	os.Stdout.WriteString("\n")
	return nil
}

func injectStdout(L *lua.LState) lua.LValue {
	out := &stdout{count: 0}
	ud := &lua.LightUserData{Value: out}
	return ud
}
