package kpi_dashboard

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
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
