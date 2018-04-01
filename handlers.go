package kpi_dashboard

import (
	"time"

	"math/rand"
)

func timeMonthlyHandler(af apiCmd) (*Response, error) {
	var results ChartTimeResult
	roles := []string{"QA Engineers", "Backend Engineers", "Frontend Engineers", "Machine learning Engineers"}
	for i := 0; i <= 100; i++ {
		var values []map[string]interface{}
		for _, role := range roles {
			value := map[string]interface{}{}
			value["role"] = role
			value["result"] = rand.Intn(100)
			values = append(values, value)
		}
		results.Result = append(results.Result, ChartTimeResultItem{Value: values})
	}

	for i := range results.Result {
		m := (i % 12) + 1
		d := time.Date(2017, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
		results.Result[i].Timeframe.Start = d.Format(time.RFC3339)
		results.Result[i].Timeframe.Start = d.Format(time.RFC3339)
	}
	return &Response{Result: results}, nil
}

func expensesHandler(af apiCmd) (*Response, error) {
	var results ChartNoTimeResult
	for _, kind := range []string{"Electricity", "Salary", "Other"} {
		value := map[string]interface{}{}
		value["kind"] = kind
		value["result"] = rand.Intn(10000)
		results.Result = append(results.Result, value)
	}

	return &Response{Result: results}, nil
}

func commitsHandler(af apiCmd) (*Response, error) {
	var results ChartTimeResult
	developers := []string{"John", "Michael", "Barbara", "Konstantin"}
	for i := 0; i <= 100; i++ {
		var values []map[string]interface{}
		for _, name := range developers {
			value := map[string]interface{}{}
			value["name"] = name
			value["result"] = rand.Intn(100)
			values = append(values, value)
		}
		m := (i % 12) + 1
		day := rand.Intn(28)
		d := time.Date(2017, time.Month(m), day, 0, 0, 0, 0, time.UTC)
		results.Result = append(results.Result, ChartTimeResultItem{
			Value: values,
			Timeframe: Time{
				Start: d.Format(time.RFC3339),
				End:   d.Format(time.RFC3339),
			}})
	}
	return &Response{Result: results}, nil
}
