########## How To Use Docker Image ###############
##
##  Image Name: denny/devubuntu:v1
##  Install docker utility
##  Download docker image: docker pull denny/devubuntu:v1
##  Boot docker container: docker run -t -d denny/devubuntu:v1 /usr/sbin/sshd -D
##
##     ruby --version
##     gem --version
##     bundle --version
##     gem sources -l
##     python --version
##     java -version
##     chef-solo --version
##     nc -l 80
##     which pidstat
##################################################

FROM denny/sshd:v1
MAINTAINER DennyZhang.com <http://dennyzhang.com>

########################################################################################
apt-get update
apt-get install -y lsof vim strace ltrace tmux curl tar telnet
apt-get install -y software-properties-common python-software-properties tree
########################################################################################
