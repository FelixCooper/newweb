package main

import(
	"io"
	"net/http"
	"os/exec"//跑一些命令
	"log"//打印一些东西
)

func fistPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>hello, world!</h1>")
}

func main(){
	http.HandleFunc("/", fistPage)
	http.ListenAndServe("5000",nil)
}