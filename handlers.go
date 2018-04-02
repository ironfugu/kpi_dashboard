package kpi_dashboard

import (
	"time"

	"math/rand"

	"fmt"

	"github.com/golang/glog"
)

func contributionHandler(af apiCmd) (*Response, error) {
	var results ChartTimeResult
	params := af.Cmd.Params
	glog.V(4).Infof("params %+v", params)
	from, to, err := parseDateParams(params)
	if err != nil {
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}

	roles := []string{"QA Engineers", "Backend Engineers", "Frontend Engineers", "Machine learning Engineers"}
	for i := 0; i < monthsAmount(from, to); i++ {
		var values []map[string]interface{}
		for _, role := range roles {
			value := map[string]interface{}{}
			value["role"] = role
			value["result"] = rand.Intn(100)
			values = append(values, value)
		}
		d := from.Add(time.Hour * 24 * 30 * time.Duration(i))
		results.Result = append(results.Result, ChartTimeResultItem{
			Value: values,
			Timeframe: Time{
				Start: d.Format(time.RFC3339),
				End:   d.Format(time.RFC3339),
			},
		})
	}

	return &Response{Result: results}, nil
}

func monthsAmount(from, to *time.Time) int {
	diff := to.Sub(*from)
	months := int(diff.Hours() / 24 / 30)
	if months == 0 {
		return 1
	}
	return months
}

func parseDateParams(params []string) (from, to *time.Time, err error) {
	if len(params) < 2 {
		return nil, nil, fmt.Errorf("missing parameters")
	}
	f, err := time.Parse(time.RFC3339, params[0])
	if err != nil {
		return nil, nil, fmt.Errorf("could not parse 'from' time: %+v", err)
	}
	t, err := time.Parse(time.RFC3339, params[1])
	if err != nil {
		return &f, nil, fmt.Errorf("could not parse 'to' time: %+v", err)
	}
	return &f, &t, nil
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
	from, to, err := parseDateParams(af.Cmd.Params)
	if err != nil {
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}

	for i := 0; i < monthsAmount(from, to); i++ {
		var values []map[string]interface{}
		for _, name := range developers {
			value := map[string]interface{}{}
			value["name"] = name
			value["result"] = rand.Intn(100)
			values = append(values, value)
		}
		d := from.Add(time.Hour * 24 * 30 * time.Duration(i))
		results.Result = append(results.Result, ChartTimeResultItem{
			Value: values,
			Timeframe: Time{
				Start: d.Format(time.RFC3339),
				End:   d.Format(time.RFC3339),
			}})
	}
	return &Response{Result: results}, nil
}

func profitHandler(af apiCmd) (*Response, error) {
	var results ChartTimeResult
	developers := []string{"Client1", "Client2", "Client3"}
	for i := 0; i < 12; i++ {
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
