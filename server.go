package kpi_dashboard

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

func Start(context *Context) {
	http.HandleFunc("/", handler)
	glog.Infof("listening %+v", context.config.Bind)
	glog.Fatal(http.ListenAndServe(context.config.Bind, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
