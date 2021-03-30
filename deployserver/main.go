package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)
//	"os/exec"//跑一些命令
//	"log"//打印一些东西

//在外面添加脚本文件

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")  //exec是go用来调用命令
	err := cmd.Start()
	if err != nil {  //有问题直接退出
		log.Fatal(err)
	}
	err = cmd.Wait()
}
func fistPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>hello, this is my deploy server!</h1>")
	//在之后运行，放在io后面防止阻塞，异步过程体验会好点
	reLaunch()
}

func main(){
	http.HandleFunc("/", fistPage)
	http.ListenAndServe("8080",nil)
}