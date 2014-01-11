package gostas

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
)

const VERSION = "0.0.1"

const (
	CONFIG_FILE  = "gostas.conf"
	CONFIG_ADDR  = "address"
	CONFIG_LOC   = "location"
	CONFIG_ROOT  = "root"
	CONFIG_SPLIT = "="
	CONFIG_USAGE = `config file - usage:
	  address=:8080 (or address=127.0.0.1:8080)
	  location=assets (request loaction)
	  root=.public (physical path on server)
					`
)

//serve static file
var StaticMapping map[string]string

func init() {
	StaticMapping = make(map[string]string)
	StaticMapping["/static"] = "static"
}

type GoStaticServer struct {
	address      string
	showDirIndex bool
}

var defaultGostas = &GoStaticServer{address: "", showDirIndex: false}

func (s *GoStaticServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	for prefix, dir := range StaticMapping {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := dir + r.URL.Path[len(prefix)-1:]
			fileInfo, err := os.Stat(file)
			if err != nil {
				http.NotFound(rw, r)
				return
			}

			if fileInfo.IsDir() && !s.showDirIndex {
				//log.Println(prefix)
				if prefix != "/" {
					http.Error(rw, "403 Forbidden", http.StatusForbidden)
					return
				}
			}

			http.ServeFile(rw, r, file)
		}
	}
}

func Mapping(handlerPrefix string, staticDir string) *GoStaticServer {
	if !strings.HasPrefix(handlerPrefix, "/") {
		handlerPrefix = "/" + handlerPrefix
	}
	StaticMapping[handlerPrefix] = staticDir
	return defaultGostas
}

func ShowDirs() *GoStaticServer {
	defaultGostas.showDirIndex = true
	return defaultGostas
}

func UseConfig() *GoStaticServer {
	f, err := os.OpenFile(CONFIG_FILE, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		log.Fatal("Read config file ", err.Error())
	}
	reader := bufio.NewReader(f)
	for {
		s, eof := reader.ReadString('\n')

		if strings.Contains(s, CONFIG_ADDR) {
			addr := strings.Split(s, CONFIG_SPLIT)
			if len(addr) == 2 {
				var addrStr string = strings.TrimFunc(addr[1], func(r rune) bool {
					return r == ' ' || r == '\n' || r == '\r'

				})

				defaultGostas.address = addrStr
			} else {
				log.Print(CONFIG_USAGE)
				os.Exit(-1)

			}
		}

		if eof != nil {
			break
		}
	}

	return defaultGostas

}

func Addr(addr string) *GoStaticServer {
	defaultGostas.address = addr
	return defaultGostas
}

func (goss *GoStaticServer) Run() {
	log.Println("running on " + goss.address)
	err := http.ListenAndServe(defaultGostas.address, goss)
	if err != nil {
		log.Fatal("ListenAndServe", err.Error())
	}

}

func Run() {
	defaultGostas.Run()
}
