Start pod and mount dockerd socket file

```
# start minikube

kubectl apply -f ./pod.yaml

kubectl exec -it dummy sh

apt-get install -y jq

# docker ps use a different socker file

curl -XGET --unix-socket /myrun/docker.sock http://localhost/containers/json | jq

# curl -XGET --unix-socket /myrun/docker.sock http://localhost/containers/json | jq

## ,----------- Sample Output
## | curl -XGET --unix-socket /myrun/docker.sock http://localhost/containers/json | jq
## | 
## |   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
## |                                  Dload  Upload   Total   Spent    Left  Speed
## | 100 47019    0 47019    0     0  8426k      0 --:--:-- --:--:-- --:--:-- 9183k
## | [
## |   {
## |     "Id": "919c89751df4f65ad2b2e704db0e83c54c013781f899c6d7403b0d17aed95c4f",
## |     "Names": [
## |       "/k8s_dummy_dummy_default_95cf6d7d-b7e2-11e8-a43b-08002779ae1f_0"
## |     ],
## |     "Image": "sha256:7f15fe5ed50b65b0c4835db38c5b2a4d30b305695094f26e10b39c7f07081d8d",
## |     "ImageID": "sha256:7f15fe5ed50b65b0c4835db38c5b2a4d30b305695094f26e10b39c7f07081d8d",
## |     "Command": "/bin/tini -- /usr/local/bin/jenkins.sh /bin/sh -c 'i=0; while true; do echo \"$i: $(date)\"; i=$((i+1)); sleep 1; done'",
## |     "Created": 1536904444,
## |     "Ports": [],
## |     "Labels": {
## |       "annotation.io.kubernetes.container.hash": "38f494c3",
## |       "annotation.io.kubernetes.container.restartCount": "0",
## |       "annotation.io.kubernetes.container.terminationMessagePath": "/dev/termination-log",
## |       "annotation.io.kubernetes.container.terminationMessagePolicy": "File",
## |       "annotation.io.kubernetes.pod.terminationGracePeriod": "30",
## |       "io.kubernetes.container.logpath": "/var/log/pods/95cf6d7d-b7e2-11e8-a43b-08002779ae1f/dummy/0.log",
## |       "io.kubernetes.container.name": "dummy",
## |       "io.kubernetes.docker.type": "container",
## |       "io.kubernetes.pod.name": "dummy",
## |       "io.kubernetes.pod.namespace": "default",
## |       "io.kubernetes.pod.uid": "95cf6d7d-b7e2-11e8-a43b-08002779ae1f",
## |       "io.kubernetes.sandbox.id": "3ad96f76f804e85220c510898d24d12f25e9ee8ca896aee3107ca5dc0e6cf82c"
## |     },
## |     "State": "running",
## |     "Status": "Up 3 minutes",
## |     "HostConfig": {
## |       "NetworkMode": "container:3ad96f76f804e85220c510898d24d12f25e9ee8ca896aee3107ca5dc0e6cf82c"
## |     },
## |     "NetworkSettings": {
## |       "Networks": {}
## |     },
## |     "Mounts": [
## |       {
## |         "Type": "volume",
## |         "Name": "d958425923238869fba1b3a92da297bbf835d25c3f7b1b5673fef9c9ef025dfe",
## |         "Source": "",
## |         "Destination": "/var/jenkins_home",
## |         "Driver": "local",
## |         "Mode": "",
## |         "RW": true,
## |         "Propagation": ""
## |       },
## |       {
## |         "Type": "bind",
## |         "Source": "/var/lib/kubelet/pods/95cf6d7d-b7e2-11e8-a43b-08002779ae1f/volumes/kubernetes.io~secret/default-token-bzv6p",
## |         "Destination": "/var/run/secrets/kubernetes.io/serviceaccount",
## |         "Mode": "ro,rslave",
## |         "RW": false,
## |         "Propagation": "rslave"
## |       },
## |       {
## |         "Type": "bind",
## |         "Source": "/var/lib/kubelet/pods/95cf6d7d-b7e2-11e8-a43b-08002779ae1f/containers/dummy/82b08ee8",
## |         "Destination": "/dev/termination-log",
## |         "Mode": "",
## |         "RW": true,
## |         "Propagation": "rprivate"
## |       },
## `-----------
```
