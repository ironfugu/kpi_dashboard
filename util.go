package kpi_dashboard

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// ParseBind
//
// A common infrastructure to parse "-bind" flag.
//
// Parameters:
//	Bind: IP, IP:PORT, :PORT
//  Port: DEFAULT_PORT
//
func ParseBind(bind string, port int) (string, string, int, *net.Interface, error) {

	var localIntf *net.Interface = nil

	localAddr := "127.0.0.1"
	h := bind
	b := fmt.Sprintf("%s:%d", bind, port)
	p := ""
	if f := strings.Split(bind, ":"); len(f) == 2 {
		var err error
		h, p, err = net.SplitHostPort(bind)
		if err != nil {
			return "", "", UNDEFINED, nil, fmt.Errorf("error parsed bind info - %v", err)
		}
		if h == "" {
			h = "0.0.0.0"
		}
	}

	intfs, err := net.Interfaces()
	if err != nil {
		return "", "", UNDEFINED, nil, fmt.Errorf("error determined interface list - %v", err)
	}

OuterLoop:
	for _, intf := range intfs {
		cidrs, err := GetIPAddrs(intf.Name)
		if err != nil {
			return "", "", UNDEFINED, nil, fmt.Errorf("error looking interface - %v", err)
		}

		for _, cidr := range cidrs {
			ip, _, err := net.ParseCIDR(cidr.String())
			if err != nil {
				continue
			}
			if ip.String() == h {
				localAddr = h
				localIntf = &intf
				break OuterLoop
			}
			// Use 1st interface for bind = 0.0.0.0
			// XXX Bad idea
			if h == "0.0.0.0" {
				localAddr = ip.String()
				localIntf = &intf
				break OuterLoop
			}
		}
	}
	if localIntf == nil {
		return "", "", UNDEFINED, nil, fmt.Errorf("invalid bind address - %s", h)
	}
	if len(p) != 0 {
		port, err = strconv.Atoi(p)
		if err != nil {
			return "", "", UNDEFINED, nil, fmt.Errorf("error parsed bind port - %v", err)
		}
	}

	b = fmt.Sprintf("%s:%d", h, port)
	return b, localAddr, port, localIntf, nil
}

func GetIPAddrs(name string) ([]net.Addr, error) {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return nil, fmt.Errorf("can't get interface %s, %v", name, err)
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, fmt.Errorf("can't get addresses for %s, %v", name, err)
	}
	return addrs, nil
}

func InitFlag(context *Context) {
	c := context.config
	flag.StringVar(&c.Bind, "bind", c.Bind, "[IP][[:]PORT] web server listening")
	flag.Parse()
}

func defaultConfig() *Config {
	return &Config{
		Bind: fmt.Sprintf("0.0.0.0:%d", DEFAULT_PORT),
	}
}
func GetContext() *Context {
	config := defaultConfig()

	context := &Context{config: config}
	return context
}

func prepareContent(context *Context) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current working directory: %+v", err)
	}
	glog.Infof("working directory %+v", dir)
	context.contentPath = "./static/content"
	fi, err := os.Stat(context.contentPath)
	if err != nil {
		return fmt.Errorf("contentPath %s does not exist: %+v", context.contentPath, err)

	}
	if !fi.IsDir() {
		return fmt.Errorf("%+v is not a directory", fi.Name())

	}
	glog.V(4).Infof("walking root application dir for html files")
	filepath.Walk(context.contentPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error trying to get static html files by walking context.contentPath: %v. err: %v", path, err)

		}
		if !f.IsDir() && filepath.Ext(path) == ".html" {
			context.htmlPages = append(context.htmlPages, path)

		}
		return nil

	})
	glog.V(4).Infof("template parse html files %v", context.htmlPages)
	context.pageTemplate = template.Must(template.New("").Funcs(template.FuncMap{
		"noescape": func(value interface{}) template.HTML {
			return template.HTML(fmt.Sprint(value))

		},
		"CallTemplate": func(name string, data interface{}) (ret template.HTML, err error) {
			buf := bytes.NewBuffer([]byte{})
			err = context.pageTemplate.ExecuteTemplate(buf, name, data)
			ret = template.HTML(buf.String())
			return

		},
	}).ParseFiles(context.htmlPages...))

	return nil

}

func renderPage(context *Context, w http.ResponseWriter, r *http.Request, templateName string, p Page) {
	err := context.pageTemplate.ExecuteTemplate(w, templateName, p)
	if err != nil {
		glog.Errorf("template error %v", err)
		errmsg := fmt.Errorf("template %s error: %v", templateName, err)
		http.Error(w, errmsg.Error(), http.StatusInternalServerError)
	}
}
