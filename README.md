# About gostas
======

  gostas is a simple static web server written by golang.you can use it as a static server
  alone ,also can use it as a plugin for your dynamic web server to process static resources.
  
## Get It
```bash
$ go get github.com/scottkiss/gostas
```

## Useage
```golang
import (
	"github.com/scottkiss/gostas"
	"log"
)
func main() {
	gostas.Mapping("/assets/", "./public")
	gostas.Mapping("/pics", "./imgs")
	//gostas.ShowDirs()  //use it show dirs 
	//gostas.UseConfig() //read config from file,e.g. specify listening address and port
	//gostas.Addr(":8088").Run() //run it on port 8088
	gostas.Run() //default running on localhost:80
	log.Println("running ...")

}
```

## License
View the [LICENSE](https://github.com/scottkiss/gostas/blob/master/LICENSE) file