package log

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
)

// StartEndLog is a function that outputs the start and end of the function.
func StartEnd(
	ctx context.Context,
) func() {
	var pc uintptr
	for i := 0; i <= 1; i++ {
		pc, _, _, _ = runtime.Caller(i)
		// 0 ->　StartEndLog
		// 1 ->　呼び出し元
	}

	fn := runtime.FuncForPC(pc)

	slog.InfoContext(ctx, fmt.Sprintf("start %s", fn.Name()))

	end := func() {
		slog.InfoContext(ctx, fmt.Sprintf("end %s", fn.Name()))
	}

	return end
}
