package kpi_dashboard

import (
	"net/http"

	"github.com/golang/glog"
)

func Start(context *Context) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		renderPage(context, w, r, ROOT_TEMPLATE_NAME)
	}
	http.HandleFunc("/", handler)
	if err := prepareContent(context); err != nil {
		glog.Fatalf("could not prepare content: %+v", err)
	}
	glog.Infof("listening %+v", context.config.Bind)
	glog.Fatal(http.ListenAndServe(context.config.Bind, nil))
}
