package gorequest

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// TraceNewSpan 创建OpenTelemetry链路追踪状态
func TraceNewSpan(ctx context.Context, name string, spanName string, spanValue string, version string, kind trace.SpanKind) (context.Context, trace.Span) {
	return otel.Tracer(name, trace.WithInstrumentationVersion(version)).Start(ctx, spanName+spanValue, trace.WithSpanKind(kind))
}

// TraceStartSpan 开始OpenTelemetry链路追踪状态
func TraceStartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return TraceNewSpan(ctx, "go.dtapp.net/gorequest", "gorequest.", spanName, Version, trace.SpanKindClient)
}

// TraceEndSpan 结束OpenTelemetry链路追踪状态
func TraceEndSpan(span trace.Span) {
	if span != nil {
		span.End()
	}
}

// TraceSetAttributes 设置OpenTelemetry链路追踪属性
func TraceSetAttributes(ctx context.Context, kv ...attribute.KeyValue) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.SetAttributes(kv...)
	}
}

// TraceSetStatus 设置OpenTelemetry链路追踪状态
func TraceSetStatus(ctx context.Context, code codes.Code, description string) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.SetStatus(code, description)
	}
}

// TraceRecordError 记录OpenTelemetry链路追踪错误
func TraceRecordError(ctx context.Context, err error, options ...trace.EventOption) {
	span := trace.SpanFromContext(ctx)
	if span != nil {
		span.RecordError(err, options...)
	}
}

// TraceGetTraceID 获取OpenTelemetry链路追踪TraceID
func TraceGetTraceID(ctx context.Context) (traceID string) {
	span := trace.SpanFromContext(ctx)
	if span != nil && span.IsRecording() {
		traceID = span.SpanContext().TraceID().String()
	}
	//if traceID == trace.TraceID([16]byte{}).String() {
	//	traceID = ""
	//}
	return traceID
}

// TraceGetSpanID 获取OpenTelemetry链路追踪SpanID
func TraceGetSpanID(ctx context.Context) (spanID string) {
	span := trace.SpanFromContext(ctx)
	if span != nil && span.IsRecording() {
		spanID = span.SpanContext().SpanID().String()
	}
	//if spanID == trace.SpanID([8]byte{}).String() {
	//	spanID = ""
	//}
	return spanID
}
