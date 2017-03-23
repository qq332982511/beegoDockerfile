FROM ubuntu
MAINTAINER xiong

RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt update
RUN apt-get -y install golang
RUN apt-get -y install git
RUN mkdir /root/gopath
ENV GOPATH /root/gopath
#RUN echo "export GOPATH=/root/gopath" >> /root/.bashrc
#RUN source /root/.bashrc
RUN go get github.com/astaxie/beego
