package kpi_dashboard

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
)

var apiFuncs map[string]*apiFunc

func Start(context *Context) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		renderPage(context, w, r, ROOT_TEMPLATE_NAME)
	}
	initCmds(context)
	http.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		glog.V(4).Infof("URL path: %s", r.URL.Path[1:])

		fields := strings.Split(r.URL.Path[1:], "/")
		if len(fields) < 3 {
			glog.V(4).Infof("invalid endpoint path %s", r.URL.Path[1:])
			http.Error(w, "invalid endpoint", http.StatusBadRequest)
			return
		}
		ep := fields[2]
		apiFunc, ok := apiFuncs[ep]
		if !ok {
			glog.V(4).Infof("invalid endpoint %s", ep)
			http.Error(w, "invalid command", http.StatusBadRequest)
			return
		}

		var req = &Request{UriTunnel: fields[2:]}
		if r.Method == "POST" {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				glog.V(4).Infof("%+v", err)
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			glog.V(4).Infof("Retrieve body %s", string(b))
			if err := json.Unmarshal(b, &req); err != nil {
				glog.V(4).Infof("can't decode request msg")
				http.Error(w, "can't decode body", http.StatusBadRequest)
				return
			}
		}
		cmd := apiCmd{
			Context:    apiFunc.Context,
			RemoteAddr: r.RemoteAddr, Cmd: req, Req: r}
		resp, err := apiFuncs[ep].Handler(cmd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sendResp(w, resp)
	})
	http.HandleFunc("/", handler)
	if err := prepareContent(context); err != nil {
		glog.Fatalf("could not prepare content: %+v", err)
	}
	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		glog.Infof("serving public file: %s", r.URL.Path[1:])
		http.ServeFile(w, r, filepath.Join("static", r.URL.Path[1:]))
	})
	glog.Infof("listening %+v", context.config.Bind)
	glog.Fatal(http.ListenAndServe(context.config.Bind, nil))
}

func sendResp(w http.ResponseWriter, resp *Response) {
	result := resp.Result
	if resp.Error != nil {
		result = resp.Error
	}
	payload, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "resource error", http.StatusInternalServerError)
		return
	}
	glog.V(6).Infof("Payload: %s", string(payload))
	w.Write(payload)
}

func initCmds(context *Context) {
	apiFuncs = map[string]*apiFunc{
		"time-monthly": {Handler: timeMonthlyHandler, Context: context},
	}
}
