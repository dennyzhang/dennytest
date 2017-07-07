git clone https://github.com/DennyZhang/dennytest.git
cd dennytest/nginx_upstream

docker-compose up -d

docker exec -it proxy_test ifconfig | grep inet
docker exec -it jenkins_test ip a | grep inet

docker exec -it proxy_test sh
vi /etc/nginx/nginx.conf

| Name                    | Summary                                 |
|-------------------------+-----------------------------------------|
| Query Jenkins directly  | curl -I http://172.21.0.1:8080/         |
| Query nginx proxy_test       | curl -I http://172.21.0.1:8082/         |
| Query Jenkins via proxy_test | curl -I http://172.21.0.1:8082/jenkins/ |

# Testcase1: nginx detect upstream failure

1. docker-compose up -d
2. docker stop jenkins
3. for((i=0; i< 100; i++)); do { sleep 1; curl -I http://172.21.0.1:8082/jenkins/ ;}; done

curl request will hang, when it hit the jenkins upstream. After quite a while, it will know jenkins upstream is unavailable.

# Testcase2: when any upstream is not running, nginx will run into "host not found in upstream" error
1. docker-compose up -d
2. docker stop jenkins_test
3. docker stop proxy_test
4. docker start proxy_test

proxy_test container fail to start
```
nginx: [emerg] host not found in upstream "jenkins:8080" in /etc/nginx/nginx.conf:33
```

# Testcase3: Be tolerant for some servers in upstream is unavailable
1. docker-compose up -d; docker-compose ps
2. docker exec -it jenkins_test ip a | grep inet
3. docker stop jenkins_test jenkins2_test
4. docker exec -it proxy_test ping jenkins_test
5. docker start jenkins2_test jenkins_test
6. docker exec -it jenkins_test ip a | grep inet
7. docker exec -it proxy_test ping jenkins_test
8. curl -I http://172.21.0.1:8082/jenkins/

proxy_test container fail to start
```
nginx: [emerg] host not found in upstream "jenkins:8080" in /etc/nginx/nginx.conf:33
```
