package script

//go:generate ../generator/run.sh

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/bitfield/script"
	"mvdan.cc/sh/v3/shell"
)

var NewReadAutoCloser = script.NewReadAutoCloser

type Pipe struct {
	*script.Pipe

	stderr io.Writer // captured from WithStderr

	// wd is the working directory for current pipe.
	wd string
	// env is the environment variables for current pipe.
	// Non-empty value will be set to the exec.Command instance.
	env []string
}

func (p *Pipe) At(dir string) *Pipe {
	p.wd = dir
	return p
}

// WithCurrentEnv sets the environment variables to the current process's environment.
func (p *Pipe) WithCurrentEnv() *Pipe {
	return p.WithEnv(os.Environ())
}

// WithEnv sets the environment variables for the current pipe.
func (p *Pipe) WithEnv(env []string) *Pipe {
	p.env = env
	return p
}

// AppendEnv appends the environment variables for the current pipe.
func (p *Pipe) AppendEnv(env ...string) *Pipe {
	p.env = append(p.env, env...)
	return p
}

// WithEnvKV sets the environment variable key-value pair for the current pipe.
func (p *Pipe) WithEnvKV(key, value string) *Pipe {
	p.env = append(p.env, key+"="+value)
	return p
}

func (p *Pipe) WithStderr(w io.Writer) *Pipe {
	p.stderr = w
	p.Pipe = p.Pipe.WithStderr(w)
	return p
}

func (p *Pipe) resolvePath(path string) string {
	if p.wd == "" {
		return path
	}
	return filepath.Join(p.wd, path)
}

func (p *Pipe) AppendFile(path string) (int64, error) {
	return p.Pipe.AppendFile(p.resolvePath(path))
}

func (p *Pipe) WriteFile(path string) (int64, error) {
	return p.Pipe.WriteFile(p.resolvePath(path))
}

func (p *Pipe) openFile(path string) (*os.File, error) {
	return os.Open(p.resolvePath(path))
}

func (p *Pipe) Concat() *Pipe {
	var readers []io.Reader
	p.FilterScan(func(line string, w io.Writer) {
		input, err := p.openFile(line)
		if err == nil {
			readers = append(readers, script.NewReadAutoCloser(input))
		}
	}).Wait()

	return p.WithReader(io.MultiReader(readers...))
}

func (p *Pipe) Read(b []byte) (int, error) {
	if p.Pipe == nil {
		return 0, io.EOF
	}
	return p.Pipe.Read(b)
}

func (p *Pipe) SHA256Sums() *Pipe {
	p.Pipe = p.Pipe.FilterScan(func(line string, w io.Writer) {
		f, err := p.openFile(line)
		if err != nil {
			return // skip unopenable files
		}
		defer f.Close()
		h := sha256.New()
		_, err = io.Copy(h, f)
		if err != nil {
			return // skip unreadable files
		}
		fmt.Fprintln(w, hex.EncodeToString(h.Sum(nil)))
	})
	return p
}

func (p *Pipe) execContext(
	ctx context.Context,
	command string,
	args []string,
) *exec.Cmd {
	cmd := exec.CommandContext(ctx, command, args...)
	if p.wd != "" {
		cmd.Dir = p.wd
	}
	if len(p.env) > 0 {
		cmd.Env = p.env
	}

	return cmd
}

func (p *Pipe) ExecContext(ctx context.Context, cmdLine string) *Pipe {
	p.Pipe = p.Pipe.Filter(func(r io.Reader, w io.Writer) error {
		args, err := shell.Fields(cmdLine, nil)
		if err != nil {
			return err
		}

		cmd := p.execContext(ctx, args[0], args[1:])
		cmd.Stdin = r
		cmd.Stdout = w
		cmd.Stderr = w
		if p.stderr != nil {
			cmd.Stderr = p.stderr
		}

		if err := cmd.Start(); err != nil {
			fmt.Fprintln(cmd.Stderr, err)
			return err
		}
		return cmd.Wait()
	})
	return p
}

func (p *Pipe) Exec(cmdLine string) *Pipe {
	return p.ExecContext(context.Background(), cmdLine)
}

func (p *Pipe) ExecForEachContext(ctx context.Context, cmdLine string) *Pipe {
	tpl, err := template.New("").Parse(cmdLine)
	if err != nil {
		p.Pipe = p.Pipe.WithError(err)
		return p
	}
	p.Pipe = p.Pipe.Filter(func(r io.Reader, w io.Writer) error {
		scanner := newScanner(r)
		for scanner.Scan() {
			cmdLine := new(strings.Builder)
			err := tpl.Execute(cmdLine, scanner.Text())
			if err != nil {
				return err
			}
			args, err := shell.Fields(cmdLine.String(), nil)
			if err != nil {
				return err
			}

			cmd := p.execContext(ctx, args[0], args[1:])
			cmd.Stdout = w
			cmd.Stderr = w
			if p.stderr != nil {
				cmd.Stderr = p.stderr
			}
			err = cmd.Start()
			if err != nil {
				fmt.Fprintln(cmd.Stderr, err)
				continue
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Fprintln(cmd.Stderr, err)
				continue
			}
		}
		return scanner.Err()
	})
	return p
}

func (p *Pipe) ExecForEach(cmdLine string) *Pipe {
	return p.ExecForEachContext(context.Background(), cmdLine)
}

func newScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 4096), math.MaxInt)
	return scanner
}
