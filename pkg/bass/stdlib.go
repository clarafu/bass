package bass

import (
	"context"
	"fmt"

	"github.com/mattn/go-colorable"
	"github.com/morikuni/aec"
	"github.com/vito/bass/pkg/ioctx"
	"github.com/vito/bass/pkg/zapctx"
	"github.com/vito/bass/std"
)

func init() {
	for _, lib := range []string{
		"root.bass",
		"lists.bass",
		"streams.bass",
		"run.bass",
		"paths.bass",
		"bool.bass",
	} {
		file, err := std.FS.Open(lib)
		if err != nil {
			panic(err)
		}

		stderr := colorable.NewColorableStderr()
		ctx := context.Background()
		ctx = ioctx.StderrToContext(ctx, stderr)
		ctx = zapctx.ToContext(ctx, Logger())

		_, err = EvalReader(ctx, Ground, file, lib)
		if err != nil {
			fmt.Fprintf(stderr, aec.YellowF.Apply("eval ground %s: %s\n"), lib, err)
		}

		_ = file.Close()
	}
}
