# Summary

This example shows how to we build golang docker image

Highlights:
1. Reproducible: don't rely on anything local
2. Use docker multi-stage builds to make the target image clean
3. Avoid always dowload latest go depenecies by using `go dep`

# How To Test
```
# Build image
docker build -f Dockerfile -t denny/mytest:v1 .

# Test

export docker_image="denny/mytest:v1"
docker stop my-test; docker rm my-test
docker run -t -d --privileged -h mytest --name my-test "$docker_image"

docker exec my-test ls -lth
# Here we will see out_syslog.so
## ,-----------
## | bash-3.2$docker exec my-test ls -lth 
## | total 14M
## | -rw-r--r-- 1 root root  14M Jul 11 16:11 out_syslog.so
## | drwxrwxrwx 2 root root 4.0K Jun 27 22:23 bin
## | drwxrwxrwx 2 root root 4.0K Jun 27 22:23 src
## `-----------

# In Dockerfile, if we change base image from golang:latest to scratch, the image size will change from 809MB to 15MB
## ,-----------
## | bash-3.2$ docker images
## | REPOSITORY                         TAG                 IMAGE ID            CREATED             SIZE
## | denny/mytest                       v1                  0446d5b9924c        7 minutes ago       809MB
## `-----------
## 
```
# More resources
```
https://gist.github.com/subfuzion/12342599e26f5094e4e2d08e9d4ad50d

https://stackoverflow.com/questions/24855081/how-do-i-import-a-specific-version-of-a-package-using-go-get

https://github.com/golang/go/issues/21933

https://github.com/golang/dep

https://github.com/golang/go/wiki/PackagePublishing

https://github.com/prometheus/client_golang/blob/ae27198cdd90bf12cd134ad79d1366a6cf49f632/examples/simple/Dockerfile

https://github.com/pivotal-cf/fluent-bit-out-syslog
```
