package kpi_dashboard

import (
	"net/http"

	"path/filepath"

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
	http.HandleFunc("/public/", func(w http.ResponseWriter, r *http.Request) {
		glog.Infof("serving public file: %s", r.URL.Path[1:])
		http.ServeFile(w, r, filepath.Join("static", r.URL.Path[1:]))
	})
	glog.Infof("listening %+v", context.config.Bind)
	glog.Fatal(http.ListenAndServe(context.config.Bind, nil))
}
