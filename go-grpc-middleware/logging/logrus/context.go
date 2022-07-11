package grpc_logrus

import (
	"github.com/hxx258456/ccgo/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/hxx258456/ccgo/net/context"
	"github.com/sirupsen/logrus"
)

// AddFields adds logrus fields to the logger.
// Deprecated: should use the ctxlogrus.Extract instead
func AddFields(ctx context.Context, fields logrus.Fields) {
	ctxlogrus.AddFields(ctx, fields)
}

// Extract takes the call-scoped logrus.Entry from grpc_logrus middleware.
// Deprecated: should use the ctxlogrus.Extract instead
func Extract(ctx context.Context) *logrus.Entry {
	return ctxlogrus.Extract(ctx)
}
