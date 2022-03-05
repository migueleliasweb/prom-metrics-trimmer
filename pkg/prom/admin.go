package prom

// import (
// 	"context"
// 	"net/http"
// )

// type AdminAPI struct {
// 	location string
// }

// func New(location string) *AdminAPI {
// 	return &AdminAPI{
// 		location: location,
// 	}
// }

// // More information about Prometheus Admin APIs can be found here: https://prometheus.io/docs/prometheus/latest/querying/api/#tsdb-admin-apis

// func (AA *AdminAPI) DeleteSeries(ctx context.Context, c http.Client) {
// 	req, err := http.NewRequestWithContext(
// 		ctx,
// 		http.MethodPost,
// 		AA.location,
// 	)
// }
