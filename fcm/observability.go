package fcm

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const (
	instrumentationName = "github.com/donnigundala/dg-firebase/fcm"
)

type observability struct {
	sentCounter       metric.Int64Counter
	durationHistogram metric.Float64Histogram
}

func newObservability() *observability {
	meter := otel.GetMeterProvider().Meter(instrumentationName)

	sentCounter, _ := meter.Int64Counter(
		"firebase.fcm.message.sent",
		metric.WithDescription("Total number of FCM messages sent"),
	)

	durationHistogram, _ := meter.Float64Histogram(
		"firebase.fcm.message.duration",
		metric.WithDescription("Duration of FCM send operations"),
		metric.WithUnit("ms"),
	)

	return &observability{
		sentCounter:       sentCounter,
		durationHistogram: durationHistogram,
	}
}

func (o *observability) record(ctx context.Context, operation string, msgType string, startTime time.Time, err error, count int64) {
	duration := float64(time.Since(startTime).Milliseconds())
	status := "success"
	if err != nil {
		status = "error"
	}

	attrs := metric.WithAttributes(
		attribute.String("operation", operation),
		attribute.String("type", msgType),
		attribute.String("status", status),
	)

	o.sentCounter.Add(ctx, count, attrs)
	o.durationHistogram.Record(ctx, duration, attrs)
}
