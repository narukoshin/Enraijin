package proxy

import (
	"custom-bruteforce/pkg/structs"
	"custom-bruteforce/pkg/config"
	"golang.org/x/net/proxy"
	"h12.io/socks"
	"net/http"
	"net/url"
	"time"
	"net"
)

var Proxy structs.YAMLProxy =  config.YAMLConfig.P

func dial_socks() *http.Transport {
	dialSocks := socks.Dial(Proxy.Socks)
	return &http.Transport{Dial: dialSocks}
}

func IsProxy() bool {
	return Proxy.Socks != ""
}

func Drive() *http.Transport {
	if Proxy.Socks != "" {
		return dial_socks()
	}
	return &http.Transport{}
}

func Dialer() (proxy.Dialer, error){
	parsed, err := parse_proxy(Proxy.Socks)
	if err != nil {
		return nil, err
	}
	return proxy.SOCKS5("tcp", parsed.Host, nil, &net.Dialer{Timeout: 3 * time.Second})
}

func parse_proxy(proxy string) (*url.URL, error) {
	return url.ParseRequestURI(proxy)
}