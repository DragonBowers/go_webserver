package main
 
import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"os"
	"log"
)
 
func main() {
	// 获取端口号
	port:=os.Args[1]
	http.HandleFunc("/api", helloHandler)

	// 静态服务
	http.Handle("/", http.FileServer(http.Dir("./static")))

    log.Fatal(http.ListenAndServe(":"+port, nil))
}

//将request转发给 http://www.tianqiapi.com/api?version=v9&appid=23035354&appsecret=8YvlPNrz
func helloHandler(w http.ResponseWriter, r *http.Request) {
    log.Println(r)
	trueServer := "http://www.tianqiapi.com"

	url, err := url.Parse(trueServer)

    if err != nil {
        log.Println(err)
        return
    }
	log.Println(url)
	proxy := httputil.NewSingleHostReverseProxy(url)
	d := proxy.Director
	proxy.Director = func(r *http.Request) {
		d(r)
		r.Host = url.Host
	}
    proxy.ServeHTTP(w, r)
}