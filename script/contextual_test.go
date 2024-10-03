package script

import (
	"context"
	"testing"
)

func TestWithStdErr_IsConcurrencySafeAfterExec(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	err := ExecContext(ctx, "echo").WithStderr(nil).Wait()
	if err != nil {
		t.Fatal(err)
	}
}
