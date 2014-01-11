package main

import (
	"fmt"
	"gostas"
	"log"
	"net/http"
)

func main() {
	gostas.Mapping("/assets/", "./public")
	gostas.Mapping("/pics", "./imgs")
	gostas.ShowDirs()
	//gostas.UseConfig()
	gostas.Addr(":8088").Run()
	//gostas.Run()
	log.Println("running ...")

}
