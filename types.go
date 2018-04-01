package kpi_dashboard

import (
	"html/template"
	"net/http"
)

type Context struct {
	config       *Config
	contentPath  string
	htmlPages    []string
	pageTemplate *template.Template
}

type Config struct {
	Bind string
}

type apiFunc struct {
	Handler handler
	Context *Context
}

type handler func(af apiCmd) (*Response, error)

type apiCmd struct {
	Context    *Context
	RemoteAddr string
	Cmd        *Request
	Req        *http.Request
}

type Request struct {
	Directive string
	UriTunnel []string
	Params    []string
}

type Response struct {
	Error  *ErrorResp
	Result interface{}
}

type ErrorResp struct {
	Reason string `json:"reason",omitempty`
	Code   string `json:"code"`
}

type ChartTimeResult struct {
	Result []ChartTimeResultItem `json:"result"`
}

type ChartTimeResultItem struct {
	Value     []map[string]interface{} `json:"value"`
	Timeframe Time                     `json:"timeframe"`
}

type Time struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type ChartNoTimeResult struct {
	Result []map[string]interface{} `json:"result"`
}
