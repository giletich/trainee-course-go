package main

import (
	"context"
	"log/slog"
	"os"
)

type CompositeHandler struct{
	handlers []slog.Handler
}

func (h *CompositeHandler) Enabled(ctx context.Context, level slog.Level) bool {

	for _, handler := range h.handlers {
		b := handler.Enabled(ctx, level)

		if !b {
			return b
		}

	}
	return true
}

func (h *CompositeHandler) Handle(ctx context.Context, r slog.Record) error {

	for _, handler := range h.handlers {
		handler.Handle(ctx, r)
	}
	
	return nil
}

func (h *CompositeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {

	newHandlers := make([]slog.Handler, 0, len(h.handlers))

	for _, handler := range h.handlers {
		newHandlers = append(newHandlers, handler.WithAttrs(attrs))
	}

	return &CompositeHandler{handlers: newHandlers}
}

func (h *CompositeHandler) WithGroup(name string) slog.Handler {

	newHandlers := make([]slog.Handler, 0, len(h.handlers))

	for _, handler := range h.handlers {
		newHandlers = append(newHandlers, handler.WithGroup(name))
	}

	return &CompositeHandler{handlers: newHandlers}
}

func main() {

	textHandler := slog.NewTextHandler(os.Stdout, nil)
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)

	compositeHandler := CompositeHandler{handlers: []slog.Handler{textHandler, jsonHandler}}

	logger := slog.New(&compositeHandler)

	logger.Info("privet")
	
}