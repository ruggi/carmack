package util

import (
	"fmt"
	"time"

	"github.com/ruggi/carmack/context"
)

func CommitMessage(ctx *context.Context) string {
	return fmt.Sprintf(`'%s: plan update %s'`, ctx.Username, time.Now().UTC().Format(time.RFC3339))
}
