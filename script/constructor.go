package script

// TODO: generate constructrs from bitfield/script's source code

import (
	"context"
	"net/http"

	"github.com/bitfield/script"
)

func newPipeFrom(pipe *script.Pipe) *Pipe {
	return &Pipe{Pipe: pipe}
}

func NewPipe() *Pipe {
	return newPipeFrom(script.NewPipe())
}

func At(dir string) *Pipe {
	return NewPipe().At(dir)
}

func Args() *Pipe {
	return newPipeFrom(script.Args())
}

func Do(req *http.Request) *Pipe {
	return newPipeFrom(script.Do(req))
}

func File(path string) *Pipe {
	return newPipeFrom(script.File(path))
}

func Get(url string) *Pipe {
	return newPipeFrom(script.Get(url))
}

func Post(url string) *Pipe {
	return newPipeFrom(script.Post(url))
}

func Echo(s string) *Pipe {
	return newPipeFrom(script.Echo(s))
}

func Slice(s []string) *Pipe {
	return newPipeFrom(script.Slice(s))
}

func Stdin() *Pipe {
	return newPipeFrom(script.Stdin())
}

func Exec(cmdLine string) *Pipe {
	return newPipeFrom(script.Exec(cmdLine))
}

func ExecContext(ctx context.Context, cmdLine string) *Pipe {
	return NewPipe().ExecContext(ctx, cmdLine)
}

func ListFiles(path string) *Pipe {
	return newPipeFrom(script.ListFiles(path))
}

func FindFiles(path string) *Pipe {
	return newPipeFrom(script.FindFiles(path))
}

func IfExists(path string) *Pipe {
	return newPipeFrom(script.IfExists(path))
}
