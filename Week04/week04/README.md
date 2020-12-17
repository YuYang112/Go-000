# cnoocData

#### 介绍
中海油数据统一采集接口

#### 软件架构
软件架构说明



### rabbitMq安装
```
    1、下载epel-release-7-11.noarch.rpm 上传到服务器上
        rpm -Uvh epel-release*rpm
    2、wget https://packages.erlang-solutions.com/erlang-solutions-1.0-1.noarch.rpm

```
### emqtt安装
```
    1、下载emqx-centos7-v3.1.0.x86_64.rpm 上传到服务器上
        rpm -ivh emqx-centos7-v3.1.0.x86_64.rpm
    2、直接启动
        emqx start

        emqx 3.1.0 is started successfully!

        emqx_ctl status

        Node 'emqx@127.0.0.1' is started
        emqx v3.1.0 is running

    3、systemctl 启动
        systemctl start emqx

    4、service 启动
        service emqx start


    http://127.0.0.1:18083/


```
## docker 部署
```
    1、部署go环境
        下载：
            wget https://storage.googleapis.com/golang/go1.12.linux-amd64.tar.gz
            tar -zxvf go1.12.linux-amd64.tar.gz -C /usr/local/

        配置环境变量：
            vim /etc/profile

            export GOROOT=/usr/local/go
            export GOPATH=/data/work/go
            export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

            :wq

        即时生效：
            source /etc/profile
        创建工作目录：
            mkdir -p /data/work/go/bin
            mkdir -p /data/work/go/pkg
            mkdir -p /data/work/go/src
        测试安装：
            cd /data/work/go/src
            mkdir hello && cd hello
            vim hello.go

                package main

                import "fmt"

                func main() {
                        fmt.Println("hello")
                }
             :wq

         go run hello.go
            显示hello则安装成功



        tar -zxvf go1.12.linux-amd64.tar.gz -C /search/odin/fze/golang/


    1、go-oci8环境安装
        yum -y install pkgconfig libaio

        将instantclient_12_1上传到/home/cnooc/下
        cd instantclient_12_1
        ln -s /usr/lib/instantclient_12_1/libclntsh.so.12.1 /usr/lib/libclntsh.so
        ln -s /usr/lib/instantclient_12_1/libocci.so.12.1 /usr/lib/libocci.so
        ln -s /usr/lib/instantclient_12_1/libociei.so /usr/lib/libociei.so
        ln -s /usr/lib/instantclient_12_1/libnnz12.so /usr/lib/libnnz12.so

       vim /etc/profile
           PKG_CONFIG_PATH=/home/instantclient_12_1
           LD_LIBRARY_PATH=/home/instantclient_12_1
           export PKG_CONFIG_PATH
           export LD_LIBRARY_PATH
       source /etc/profile

    1、项目部署
        注意：前提docker和docker-compose安装成功
            如未安装参考：https://www.jianshu.com/p/8b6ac4b18d75
                        https://www.jianshu.com/p/22143c6cc23d


        mkdir -p /home/cnooc
        cd /home/cnooc && vim update.sh

            #/bin/bash
            #从github上拉取项目
            cd /data/work/go/src/
            rm -rf cnooc
            git clone http://931737966%40qq.com:YUYANGyjn2019*@gitee.com/ettech/cnooc.git

            echo pwd

            #rm -rf /data/work/go/src/cnooc/vendor/*
            cd cnooc/

            #govendor init
            #govendor add +external

            rm -rf cnooc
            go build

            cd /home/cnooc

            rm -rf cnooc conf docker-compose.yml Dockerfile
            cd /data/work/go/src/cnooc/
            cp cnooc conf docker-compose.yml Dockerfile update.sh /home/cnooc/

            docker stop cnooc_cnooc_1 && docker rm cnooc_cnooc_1
            docker rmi cnoccproject
            docker build -t cnoccproject .
            docker-compose -f docker-compose.yml up -d


            :wq
        chmod +x update.sh
        如果没有安装git：
        yum install git -y
        安装包管理工具：
        go get  github.com/kardianos/govendor
        ./update.sh

        运行时如果报错：
        utils/redis_tool.go:6:2: cannot find package "github.com/garyburd/redigo/redis" in any of:
        	/data/work/go/src/cnooc/vendor/github.com/garyburd/redigo/redis (vendor tree)
        解决办法：
             cd /data/work/go/src/
             go get github.com/garyburd/redigo/redis
         类似问题同样解决
    2、检查是否部署成功
        docker ps -a

        CONTAINER ID        IMAGE                       COMMAND                  CREATED             STATUS              PORTS                                        NAMES
        e136080441c2        cnoccproject                "/go/src/cnooc"          About an hour ago   Up About an hour    0.0.0.0:9010->9010/tcp                       cnooc_cnooc_1

        STATUS为RUN则运行成功

        运行：docker logs e136080441c2
        显示以下内容则启动成功：
        Now listening on: http://0.0.0.0:9010
        Application started. Press CTRL+C to shut down.

```

##  项目更新
```
   1、   将更改后的代码提交到码云
   2、	登录121.42.187.220服务器，密码：1QA@WS3ed
   3、	cd /home/cnooc && ./update.sh
   4、	检查是否部署成功


##更新
   ## root权限
   cd /home/cnooc
   ps aux |grep cnooc
   kill
   nohup ./cnooc &

```


#### mongodb安装教程
```
    1.  将mongo.sh、mongodb.conf、mongodb-linux-x86_64-4.0.0.tgz上传到服务器上
    2.  修改环境变量，在export PATH USER LOGNAME MAIL HOSTNAME HISTSIZE HISTCONTRO上增加
            vim /etc/profile
                export PATH=/usr/mongodb/bin:$PATH
            :wq
            立即生效：
            source /etc/profile
    3.  给mongo.sh增加可执行权限
            chomd +x mongo.sh
    4.  运行脚本
            ./mongo.sh
    5.  测试安装是否成功
        运行mongo命令
```
#### 使用说明
```
     1、POST访问http://IP:PORT/kpi/inPut
            参数类型：json
            参数名：data
            参数实例：
            {"kipName":"kpiabc","kipData":[{"time":"20191024","value":50},{"time":"20191024","value":51}]}

            根据kipName在数据库中，以字段为id（主键）、time（json中下标）、value（json中下标）、insert_time创建名为kipName的数据表



    2、GET访问http://IP:port/kpi/outPut
            参数名：
                    length：     返回条数，例：4   int类型
                    start_time： 开始时间，例：2019-10-28  string类型
                    end_time：   结束时间，例：2019-10-28  string类型
            访问实例：
            IP:PORT/kpi/outPut/kpiabc?length=1&start_time=2019-10-28&end_time=2019-10-28

    3、POST访问http://IP:PORT/kpi/addColumn/kpiabc

            注意：路由中kpiabc为数据表名称

            参数名：
                column_name：    列名称，例：data
                type:            列类型，例：varchar
                length：         列长度，例：10

      在对应数据表中增加一列


   4、POST访问http://IP:PORT/kpi/kpiName

        返回kpi名称列表

```

```
#### 服务器mongo启动命令
    cd /home/TYJK-DPKSH-APP/mongodb/mongodb-421/bin/
    ./mongod -f mongodb.conf
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 码云特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  码云官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解码云上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是码云最有价值开源项目，是码云综合评定出的优秀开源项目
5.  码云官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  码云封面人物是一档用来展示码云会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
