#! /bin/sh

# 用来瓷器服务,之前要kill之前的服务,为了防止重启失败
# kill -9 不太安全,为了演示而以
kill -9 $(pgrep webserver)
# cd到云服务器的项目文件夹
cd ~/workspace/src/felix_gitee/newweb/

git pull origin-f-gitee dev-felix

cd webserver/
//让服务问问跑在后台不会因为关闭窗口而退出
./webserver &
