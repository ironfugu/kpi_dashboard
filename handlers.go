package kpi_dashboard

import (
	"time"

	"math/rand"

	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"

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
	for _, kind := range []string{"Electricity", "Salary", "Other", "Coffee", "Computers"} {
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

func projectsHandler(af apiCmd) (*Response, error) {
	type project struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}
	prjs, err := func() ([]project, error) {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/albums")
		if err != nil {
			return nil, fmt.Errorf("could not get projects: %+v", err)
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("could not parse response body: %+v", err)
		}
		var prjs []project
		if err := json.Unmarshal(b, &prjs); err != nil {
			return nil, fmt.Errorf("could not unmarshal result: %+v: %+v", err, string(b))
		}
		return prjs, nil
	}()
	if err != nil {
		glog.Error(err)
		return &Response{Error: &ErrorResp{Reason: err.Error()}}, nil
	}
	return &Response{Result: prjs}, nil
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
