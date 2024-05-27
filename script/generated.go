// Code generated by script-contextual/generator DO NOT EDIT.
package script

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

func (p *Pipe) Basename() *Pipe {
	p.Pipe = p.Pipe.Basename()
	return p
}

func (p *Pipe) Column(a1 int) *Pipe {
	p.Pipe = p.Pipe.Column(a1)
	return p
}

func (p *Pipe) Dirname() *Pipe {
	p.Pipe = p.Pipe.Dirname()
	return p
}

func (p *Pipe) Do(a1 *http.Request) *Pipe {
	p.Pipe = p.Pipe.Do(a1)
	return p
}

func (p *Pipe) EachLine(a1 func(string, *strings.Builder)) *Pipe {
	p.Pipe = p.Pipe.EachLine(a1)
	return p
}

func (p *Pipe) Echo(a1 string) *Pipe {
	p.Pipe = p.Pipe.Echo(a1)
	return p
}

func (p *Pipe) Filter(a1 func(io.Reader, io.Writer) error) *Pipe {
	p.Pipe = p.Pipe.Filter(a1)
	return p
}

func (p *Pipe) FilterLine(a1 func(string) string) *Pipe {
	p.Pipe = p.Pipe.FilterLine(a1)
	return p
}

func (p *Pipe) FilterScan(a1 func(string, io.Writer)) *Pipe {
	p.Pipe = p.Pipe.FilterScan(a1)
	return p
}

func (p *Pipe) First(a1 int) *Pipe {
	p.Pipe = p.Pipe.First(a1)
	return p
}

func (p *Pipe) Freq() *Pipe {
	p.Pipe = p.Pipe.Freq()
	return p
}

func (p *Pipe) Get(a1 string) *Pipe {
	p.Pipe = p.Pipe.Get(a1)
	return p
}

func (p *Pipe) JQ(a1 string) *Pipe {
	p.Pipe = p.Pipe.JQ(a1)
	return p
}

func (p *Pipe) Join() *Pipe {
	p.Pipe = p.Pipe.Join()
	return p
}

func (p *Pipe) Last(a1 int) *Pipe {
	p.Pipe = p.Pipe.Last(a1)
	return p
}

func (p *Pipe) Match(a1 string) *Pipe {
	p.Pipe = p.Pipe.Match(a1)
	return p
}

func (p *Pipe) MatchRegexp(a1 *regexp.Regexp) *Pipe {
	p.Pipe = p.Pipe.MatchRegexp(a1)
	return p
}

func (p *Pipe) Post(a1 string) *Pipe {
	p.Pipe = p.Pipe.Post(a1)
	return p
}

func (p *Pipe) Reject(a1 string) *Pipe {
	p.Pipe = p.Pipe.Reject(a1)
	return p
}

func (p *Pipe) RejectRegexp(a1 *regexp.Regexp) *Pipe {
	p.Pipe = p.Pipe.RejectRegexp(a1)
	return p
}

func (p *Pipe) Replace(a1 string, a2 string) *Pipe {
	p.Pipe = p.Pipe.Replace(a1, a2)
	return p
}

func (p *Pipe) ReplaceRegexp(a1 *regexp.Regexp, a2 string) *Pipe {
	p.Pipe = p.Pipe.ReplaceRegexp(a1, a2)
	return p
}

func (p *Pipe) Tee(a1 ...io.Writer) *Pipe {
	p.Pipe = p.Pipe.Tee(a1...)
	return p
}

func (p *Pipe) WithError(a1 error) *Pipe {
	p.Pipe = p.Pipe.WithError(a1)
	return p
}

func (p *Pipe) WithHTTPClient(a1 *http.Client) *Pipe {
	p.Pipe = p.Pipe.WithHTTPClient(a1)
	return p
}

func (p *Pipe) WithReader(a1 io.Reader) *Pipe {
	p.Pipe = p.Pipe.WithReader(a1)
	return p
}

func (p *Pipe) WithStdout(a1 io.Writer) *Pipe {
	p.Pipe = p.Pipe.WithStdout(a1)
	return p
}
