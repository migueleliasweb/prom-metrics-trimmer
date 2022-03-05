package prom

import (
	"io"
	"time"

	prom_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

func sPtr(s string) *string {
	return &s
}

func sF64(f float64) *float64 {
	return &f
}

func sI64(i int64) *int64 {
	return &i
}

var metrics *prom_client.MetricFamily = &prom_client.MetricFamily{
	Name: sPtr("testmetric"),
	Help: sPtr("the value is always 1.0 and there will be one point per hour for the whole year 2021"),
	Type: prom_client.MetricType_GAUGE.Enum(),
	Metric: func() []*prom_client.Metric {
		data := []*prom_client.Metric{}

		// 01/01/2021 00:00:00
		iterateDate := time.Date(
			2021,
			time.January,
			1,
			0, 0, 0, 0,
			time.UTC,
		)

		// 31/12/2021 23:59:59
		finalDate := time.Date(
			2021,
			time.December,
			31,
			23, 59, 59, 0,
			time.UTC,
		)

		for !iterateDate.After(finalDate) {
			data = append(data, &prom_client.Metric{
				Gauge: &prom_client.Gauge{
					Value: sF64(1),
				},
				Label: []*prom_client.LabelPair{
					{
						Name:  sPtr("foo"),
						Value: sPtr("bar"),
					},
				},
				TimestampMs: sI64(iterateDate.Unix()),
			})

			iterateDate = iterateDate.Add(time.Hour)
		}

		return data
	}(),
}

func GenerateBackfillData(out io.Writer) (int, error) {
	return expfmt.MetricFamilyToOpenMetrics(
		out,
		metrics,
	)
}
