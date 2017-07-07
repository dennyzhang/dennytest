docker-compose up -d

docker exec -it proxy ifconfig | grep inet
docker exec -it jenkins ip a | grep inet

docker exec -it proxy sh
vi /etc/nginx/nginx.conf

| Name                    | Summary                                 |
|-------------------------+-----------------------------------------|
| Query Jenkins directly  | curl -I http://172.21.0.1:8080/         |
| Query nginx proxy       | curl -I http://172.21.0.1:8082/         |
| Query Jenkins via proxy | curl -I http://172.21.0.1:8082/jenkins/ |

# Testcase1: nginx detect upstream failure

1. docker-compose up -d
2. docker stop jenkins
3. for((i=0; i< 100; i++)); do { sleep 1; curl -I http://172.21.0.1:8082/jenkins/ ;}; done

curl request will hang, when it hit the jenkins upstream. After quite a while, it will know jenkins upstream is unavailable.

# Testcase2: when any upstream is not running, nginx will run into "host not found in upstream" error
1. docker-compose up -d
2. docker stop jenkins
3. docker stop proxy
4. docker start proxy

proxy container fail to start
```
nginx: [emerg] host not found in upstream "jenkins:8080" in /etc/nginx/nginx.conf:33
```

# Testcase3: Be tolerant for some servers in upstream is unavailable
1. docker-compose up -d; docker-compose ps
2. docker exec -it jenkins ip a | grep inet
3. docker stop jenkins jenkins2
4. docker start jenkins2 jenkins
5. docker exec -it jenkins ip a | grep inet
6. docker exec -it proxy ping jenkins
7. curl -I http://172.21.0.1:8082/jenkins/

proxy container fail to start
```
nginx: [emerg] host not found in upstream "jenkins:8080" in /etc/nginx/nginx.conf:33
```
