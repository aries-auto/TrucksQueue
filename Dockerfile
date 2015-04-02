FROM golang:1.4


ADD . /go/src/github.com/aries-auto/trucksplusapi

RUN mkdir -p /home/deployer/gosrc/src/github.com/aries-auto/trucksqueue
ADD . /home/deployer/gosrc/src/github.com/aries-auto/trucksqueue
WORKDIR /home/deployer/gosrc/src/github.com/aries-auto/trucksqueue


RUN mkdir -p /root/.ssh
ADD id_rsa.pub /root/.ssh/id_rsa.pub
ADD id_rsa /root/.ssh/id_rsa

RUN ssh-keyscan -t rsa github.com 2>&1 >> /root/.ssh/known_hosts #troubleshoot

RUN chmod 700 /root/.ssh/id_rsa
#RUN echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config 
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/ 
#RUN	git clone git@github.com:aries-auto/trucksplusapi.git 
RUN export GOPATH=/home/deployer/gosrc && go get 
RUN export GOPATH=/home/deployer/gosrc && go build -o trucksqueue ./main.go

ENTRYPOINT /home/deployer/gosrc/src/github.com/aries-auto/trucksqueue/trucksqueue

EXPOSE 8080