```
One interesting question:

Let’s say I have a Pod with one container inside. The container has created a new file, let’s say /root/test.log
Then somehow the Pod has crashed. Since I have configured the restart policy, kubelet will restart it.

So my question is will /root/test.log still be there, or it’s gone?
```

When a Container crashes, kubelet will restart it, but the files will be lost - the Container starts with a clean state

This is different from the docker behavior.

If you run “docker stop/start”, the file will still be there

https://kubernetes.io/docs/concepts/storage/volumes/

```
# The pod will append one line to: /root/test.log. Sleep for 60 seconds, then crash
kubectl apply -f pod-dummy.yaml

kubectl get pods

## ,-----------
## | bash-3.2$  kubectl get pods
## | NAME      READY     STATUS    RESTARTS   AGE
## | dummy     1/1       Running   0          30s
## `-----------

# First check, the log file will has only one message
kubectl exec dummy cat /root/test.log
## ,-----------
## | bash-3.2$ kubectl exec dummy cat /root/test.log
## | hello
## `-----------

# After waiting for one minute, the pod has crashed and restarted
## ,-----------
## | bash-3.2$ kubectl get pods
## | NAME      READY     STATUS    RESTARTS   AGE
## | dummy     1/1       Running   1          1m
## `-----------

# Second check, the log file still has only one message.
# The pod start will create one message. So this means the original message has been lost
## ,-----------
## | bash-3.2$ kubectl exec dummy cat /root/test.log
## | hello
## `-----------
```
