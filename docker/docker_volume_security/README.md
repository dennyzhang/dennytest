# Summary
In one shared VM, we might want to avoid user1 mounting volume of user2 for security concerns

# Details
How we can achieve that? Two different directions:

1. Docker-in-Docker

![../../images/docker-volume.png](../../images/docker-volume.png)

2. ABAC support for volumes

ABAC: https://kubernetes.io/docs/reference/access-authn-authz/abac/

Specify one volume can only be mounted by one specific user.
```
1. Add attributes/metadata to volumes.
2. Then the scheduling engine(dockerd, k8s schedule) load the volume, before starting pods/containers.
3. If it's not allowed, pods/containers refue to start.
```

Note: currently k8s volumes doesn't support ABAC
