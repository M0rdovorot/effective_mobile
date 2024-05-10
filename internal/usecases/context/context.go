package context

import (
	"context"
	"net/http"
)

type writerKey struct{}
type routesVarsKey struct{}
type queryVars struct{}

func GetVars(ctx context.Context) map[string]string {
	if ctx.Value(routesVarsKey{}) != nil {
		return ctx.Value(routesVarsKey{}).(map[string]string)
	}

	return nil
}

func GetQueryVars(ctx context.Context) map[string]string {
	if ctx.Value(queryVars{}) != nil {
		return ctx.Value(queryVars{}).(map[string]string)
	}

	return nil
}

func AddVars(ctx context.Context, vars map[string]string) context.Context {
	return context.WithValue(ctx, routesVarsKey{}, vars)
}

func AddQueryVars(ctx context.Context, vars map[string]string) context.Context {
	return context.WithValue(ctx, queryVars{}, vars)
}

func AddWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, writerKey{}, w)
}

func GetWriter(ctx context.Context) http.ResponseWriter {
	if ctx.Value(writerKey{}) != nil {
		return ctx.Value(writerKey{}).(http.ResponseWriter)
	}

	return nil
}