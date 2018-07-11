# Summary

This example shows how to build golang docker image in a reproducible way. Meanwhile keep the target image as small as possible

Highlights:
1. We use docker multi-stage builds to make the target image clean
2. We go dep to make sure we download specific go depenecies, instead of always the head.


# How To Test
```
# Build image
docker build -f Dockerfile -t denny/mytest:v1 .

# Test

export docker_image="denny/mytest:v1"
docker stop my-test; docker rm my-test
docker run -t -d --privileged -h mytest --name my-test "$docker_image"

docker exec my-test ls -lth

## ,-----------
## | bash-3.2$docker exec my-test ls -lth 
## | total 14M
## | -rw-r--r-- 1 root root  14M Jul 11 16:11 out_syslog.so
## | drwxrwxrwx 2 root root 4.0K Jun 27 22:23 bin
## | drwxrwxrwx 2 root root 4.0K Jun 27 22:23 src
## `-----------
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
