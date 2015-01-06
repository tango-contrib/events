package events

import (
	"github.com/lunny/tango"
)

type Before interface{
	Before()
}

type After interface {
	After()
}

func Events() tango.HandlerFunc {
	return func(ctx *tango.Context) {
		action := ctx.Action()
		if action != nil {
			if b, ok := action.(Before); ok {
				b.Before()
			}
		}

		ctx.Next()

		if action != nil {
			if b, ok := action.(After); ok {
				b.After()
			}
		}
	}
}
