package kpi_dashboard

import "html/template"

type Context struct {
	config       *Config
	contentPath  string
	htmlPages    []string
	pageTemplate *template.Template
}

type Config struct {
	Bind string
}
