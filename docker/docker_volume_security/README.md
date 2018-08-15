# Summary
In one shared VM, we might want to avoid user1 mounting volume of user2 for security concerns

# Details
How we can achieve that?

## Docker-in-Docker with sock file mounted

https://forums.docker.com/t/how-to-run-docker-inside-a-container-running-on-docker-for-mac/16268

https://jpetazzo.github.io/2015/09/03/do-not-use-docker-in-docker-for-ci/

![../../images/docker-volume.png](../../images/docker-volume.png)

Doesn't work, since inside the container, "docker ps" can see other containers.

## Docker-in-Docker without sock file mounted

https://hub.docker.com/_/docker/

```
docker run --privileged --name some-docker -d docker:stable-dind
docker exec -it some-docker sh

# Inside the container, start a new container
docker run -t -d -h mytest --name my-test --entrypoint=/bin/sh "nginx"

docker ps
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
