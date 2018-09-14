Start pod and mount dockerd socket file

```
docker run --privileged -v /var/run/docker.sock:/var/run/docker.sock --name dockerd-sockfile -d getintodevops/jenkins-withdocker:lts
docker exec -it dockerd-sockfile sh

docker ps

exit 

docker stop  dockerd-sockfile; docker rm  dockerd-sockfile
```
