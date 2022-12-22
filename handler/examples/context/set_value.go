package handler

import (
	"context"
	"fmt"
)

type TraceID string

const ZeroTraceID = ""

type traceIDKey struct{}

func SetTraceID(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tid)
}

func GetTraceID(ctx context.Context) TraceID {
	if v, ok := ctx.Value(traceIDKey{}).(TraceID); ok {
		fmt.Println(v)
		fmt.Println(ok)
		return v
	}
	return ZeroTraceID
}

func SetValue() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}
