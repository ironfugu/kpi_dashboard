package kpi_dashboard

import (
	"time"

	"math/rand"

	"fmt"
)

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
	for _, kind := range []string{"Electricity", "Salary", "Other", "Coffee", "Computers"} {
		value := map[string]interface{}{}
		value["kind"] = kind
		value["result"] = rand.Intn(10000)
		results.Result = append(results.Result, value)
	}

	return &Response{Result: results}, nil
}

func profitHandler(af apiCmd) (*Response, error) {
	var results ChartTimeResult
	developers := []string{"Client1", "Client2"}
	from, to, err := parseDateParams(af.Cmd.Params)
	_ = to
	if err != nil {
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	for i := 0; i < 30; i++ {
		var values []map[string]interface{}
		for _, name := range developers {
			value := map[string]interface{}{}
			value["name"] = name
			value["result"] = rand.Intn(100)
			values = append(values, value)
		}
		d := from.Add(time.Hour * 24 * time.Duration(i))
		results.Result = append(results.Result, ChartTimeResultItem{
			Value: values,
			Timeframe: Time{
				Start: d.Format(time.RFC3339),
				End:   d.Format(time.RFC3339),
			}})
	}
	return &Response{Result: results}, nil
}

func qualityAndReleasesHandler(af apiCmd) (*Response, error) {
	table := Table{
		Name: "Quality and Releases",
		Data: [][]string{
			{"Sev 1 open bug", "4"},
			{"Sev 2 open bug", "5"},
			{"Number of releases", "2"},
		},
	}
	return &Response{Result: table}, nil
}

func otherKeyHandler(af apiCmd) (*Response, error) {
	table := Table{
		Name:   "Other key",
		Header: []string{"Key", "Value", "Change from yesterday"},
		Data: [][]string{
			{"Number of bug", "25", "-2"},
			{"Number of customers", "4", "+1"},
		},
	}
	return &Response{Result: table}, nil
}
