Table of Contents
=================

   * [Summary](#summary)
   * [Details](#details)
      * [Docker-in-Docker with sock file mounted](#docker-in-docker-with-sock-file-mounted)
      * [Docker-in-Docker without sock file mounted](#docker-in-docker-without-sock-file-mounted)
      * [ABAC support for volumes](#abac-support-for-volumes)

# Summary
In one shared VM, we might want to avoid user1 mounting volume of user2 for security concerns

# Details
How we can achieve that?

## Docker-in-Docker with sock file mounted

![../../images/docker-volume.png](../../images/docker-volume.png)

Doesn't work, since inside the container, "docker ps" can see other containers.

```
docker run --privileged -v /var/run/docker.sock:/var/run/docker.sock --name dockerd-sockfile -d getintodevops/jenkins-withdocker:lts
docker exec -it dockerd-sockfile sh

docker ps

# With sockfile mounted, inside the container, "docker ps" show all containers. So the segregation doesn't work
## ,-----------
## | # docker ps
## | CONTAINER ID        IMAGE                                  COMMAND                  CREATED             STATUS              PORTS                 NAMES
## | a7c4bf177510        getintodevops/jenkins-withdocker:lts   "/bin/tini -- /usr..."   2 minutes ago       Up 2 minutes        8080/tcp, 50000/tcp   dockerd-sockfile
## | 313c165f7041        docker:stable-dind                     "dockerd-entrypoin..."   7 minutes ago       Up 7 minutes        2375/tcp              docker-in-docker
## | e6a18dd5b461        nginx                                  "/bin/sh"                33 minutes ago      Up 33 minutes       80/tcp                container-outside
## | 98630c33c253        nginx                                  "/bin/sh"                39 minutes ago      Up 39 minutes       80/tcp                my-nginx
## `-----------
```

https://getintodevops.com/blog/the-simple-way-to-run-docker-in-docker-for-ci

https://forums.docker.com/t/how-to-run-docker-inside-a-container-running-on-docker-for-mac/16268

https://jpetazzo.github.io/2015/09/03/do-not-use-docker-in-docker-for-ci/

## Docker-in-Docker without sock file mounted

https://hub.docker.com/_/docker/

```
docker run -t -d -h mytest --name container-outside --entrypoint=/bin/sh "nginx"

docker run --privileged --name docker-in-docker -d docker:stable-dind
docker exec -it docker-in-docker sh

# Inside the container, start a new container
docker run -t -d -h mytest --name my-test --entrypoint=/bin/sh "nginx"

docker ps
## Previously we have started a container(container-outside). Inside current container, docker ps won't show it.
## ,-----------
## | CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
## | 6ac202f01812        nginx               "/bin/sh"           8 seconds ago       Up 7 seconds        80/tcp              my-test
## `-----------
```

## ABAC support for volumes

ABAC: https://kubernetes.io/docs/reference/access-authn-authz/abac/

Specify one volume can only be mounted by one specific user.
```
1. Add attributes/metadata to volumes.
2. Then the scheduling engine(dockerd, k8s schedule) load the volume, before starting pods/containers.
3. If it's not allowed, pods/containers refue to start.
```

Note: currently k8s volumes doesn't support ABAC
