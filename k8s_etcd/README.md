Table of Contents
=================

   * [Check in minikube](#check-in-minikube)
   * [Check etcd data folder in minikube](#check-etcd-data-folder-in-minikube)
   * [Query etcd](#query-etcd)
   * [Create mysql test](#create-mysql-test)
   * [Check etcd again for the DB service](#check-etcd-again-for-the-db-service)
   * [Useful link](#useful-link)

Explore what data is stored in kubernetes etcd

# Check in minikube
```
minikube ssh

sudo su -

ps -ef | grep curl

## ,----------- Sample Output
## | $ ps -ef | grep etcd
## | root      3226  3203  3 Jul18 ?        00:44:15 kube-apiserver --admission-control=Initializers,NamespaceLifecycle,LimitRanger,ServiceAccount,DefaultStorageClass,DefaultTolerationSeconds,NodeRestriction,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota --requestheader-group-headers=X-Remote-Group --advertise-address=192.168.99.100 --tls-cert-file=/var/lib/localkube/certs/apiserver.crt --tls-private-key-file=/var/lib/localkube/certs/apiserver.key --kubelet-client-certificate=/var/lib/localkube/certs/apiserver-kubelet-client.crt --proxy-client-cert-file=/var/lib/localkube/certs/front-proxy-client.crt --proxy-client-key-file=/var/lib/localkube/certs/front-proxy-client.key --insecure-port=0 --allow-privileged=true --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname --kubelet-client-key=/var/lib/localkube/certs/apiserver-kubelet-client.key --secure-port=8443 --requestheader-client-ca-file=/var/lib/localkube/certs/front-proxy-ca.crt --requestheader-username-headers=X-Remote-User --requestheader-extra-headers-prefix=X-Remote-Extra- --service-account-key-file=/var/lib/localkube/certs/sa.pub --client-ca-file=/var/lib/localkube/certs/ca.crt --enable-bootstrap-token-auth=true --requestheader-allowed-names=front-proxy-client --service-cluster-ip-range=10.96.0.0/12 --authorization-mode=Node,RBAC --etcd-servers=https://127.0.0.1:2379 --etcd-cafile=/var/lib/localkube/certs/etcd/ca.crt --etcd-certfile=/var/lib/localkube/certs/apiserver-etcd-client.crt --etcd-keyfile=/var/lib/localkube/certs/apiserver-etcd-client.key
## | root      3338  3272  1 Jul18 ?        00:15:21 etcd --data-dir=/data/minikube --key-file=/var/lib/localkube/certs/etcd/server.key --peer-cert-file=/var/lib/localkube/certs/etcd/peer.crt --peer-key-file=/var/lib/localkube/certs/etcd/peer.key --peer-trusted-ca-file=/var/lib/localkube/certs/etcd/ca.crt --client-cert-auth=true --peer-client-cert-auth=true --cert-file=/var/lib/localkube/certs/etcd/server.crt --trusted-ca-file=/var/lib/localkube/certs/etcd/ca.crt --listen-client-urls=https://127.0.0.1:2379 --advertise-client-urls=https://127.0.0.1:2379
## `-----------
```

# Check etcd data folder in minikube

```
ls -lth /data/minikube/*/*

## ,-----------
## | # ls -lth /data/minikube/*/*
## | /data/minikube/member/snap:
## | total 3.3M
## | -rw------- 1 root root  19M Jul 19 04:20 db
## | -rw-r--r-- 1 root root 7.3K Jul 19 04:13 0000000000000003-0000000000041ecb.snap
## | -rw-r--r-- 1 root root 7.3K Jul 18 23:32 0000000000000003-000000000003f7ba.snap
## | -rw-r--r-- 1 root root 7.3K Jul 18 21:03 0000000000000003-000000000003d0a9.snap
## | -rw-r--r-- 1 root root 7.3K Jul 18 18:34 0000000000000003-000000000003a998.snap
## | -rw-r--r-- 1 root root 7.3K Jul 18 06:30 0000000000000003-0000000000038287.snap
## | 
## | /data/minikube/member/wal:
## | total 367M
## | -rw------- 1 root root 62M Jul 19 04:20 0000000000000004-0000000000041514.wal
## | -rw------- 1 root root 62M Jul 19 01:22 0.tmp
## | -rw------- 1 root root 62M Jul 19 01:22 0000000000000003-0000000000031e3d.wal
## | -rw------- 1 root root 62M Jul 17 20:58 0000000000000002-0000000000022257.wal
## | -rw------- 1 root root 62M Jul 16 05:16 0000000000000001-00000000000110d1.wal
## | -rw------- 1 root root 62M Jul 13 20:41 0000000000000000-0000000000000000.wal
## `-----------
```

# Query etcd

```
# Inside minikube, found etcd container
docker ps | grep etcd
## ,-----------
## | $ docker ps | grep etcd
## | ec52f93666d2        52920ad46f5b                                    "etcd --data-dir=/da…"   36 hours ago        Up 36 hours                             k8s_etcd_etcd-minikube_kube-system_dc7c2a29e86f22cde1fcb7d6eaadc95e_0
## | ca241ffe10bf        k8s.gcr.io/pause-amd64:3.1                      "/pause"                 36 hours ago        Up 36 hours                             k8s_POD_etcd-minikube_kube-system_dc7c2a29e86f22cde1fcb7d6eaadc95e_0
## `-----------

docker exec -it ec52f93666d2 sh

# Run sample query: member list
command_prefix="etcdctl --endpoints 127.0.0.1:2379 --cacert /var/lib/localkube/certs/etcd/ca.crt --cert /var/lib/localkube/certs/etcd/peer.crt --key /var/lib/localkube/certs/etcd/peer.key"
ETCDCTL_API=3 $command_prefix member list
## ,----------- Sample Output
## | / # $command_prefix member list
## | 2018-07-19 04:52:38.396931 I | warning: ignoring ServerName for user-provided CA for backwards compatibility is deprecated
## | 8e9e05c52164694d: name=default peerURLs=http://localhost:2380 clientURLs=https://127.0.0.1:2379 isLeader=true
## `-----------

# https://kubernetes-v1-4.github.io/docs/admin/etcd/
# By default, Kubernetes objects are stored under the /registry key in etcd.
ETCDCTL_API=3 $command_prefix get /registry/namespaces/default -w=json

## ,----------- Sample Ouptut
## | / # ETCDCTL_API=3 $command_prefix get /registry/namespaces/default -w=json
## | 2018-07-19 04:58:49.768085 I | warning: ignoring ServerName for user-provided CA for backwards compatibility is deprecated
## | {"header":{"cluster_id":14841639068965178418,"member_id":10276657743932975437,"revision":266852,"raft_term":3},"kvs":[{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMvZGVmYXVsdA==","create_revision":6,"mod_revision":6,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEl8KRQoHZGVmYXVsdBIAGgAiACokMmM5MGY5ZDUtODVhMy0xMWU4LTg5YWQtMDgwMDI3Y2JhZWE0MgA4AEIICML/m9oFEAB6ABIMCgprdWJlcm5ldGVzGggKBkFjdGl2ZRoAIgA="}],"count":1}
## `-----------
```

# Create mysql test
[kubernetes.yaml](kubernetes.yaml):
- Create namespace: ns-test
- Create PV: 5GB local disk
- Create PVC: In ns-test namespace, create one PVC
- Create deployment and service: mysql. This db service will use the PVC

```
kubectl apply -f ./kubernetes.yaml

kubectl get pvc -n ns-test
## ,-----------
## | bash-3.2$ kubectl get pvc -n ns-test
## | NAME         STATUS    VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
## | mysql001-1   Bound     pvc-5a949e51-8b11-11e8-a8c8-080027cbaea4   5Gi        RWO            standard       5s
## `-----------

kubectl get pod -n ns-test
## ,-----------
## | bash-3.2$ kubectl get pod -n ns-test
## | NAME                                   READY     STATUS    RESTARTS   AGE
## | dbserver-deployment-7c76884dbf-xjq9t   1/1       Running   0          40s
## `-----------
```

# Check etcd again for the DB service

```
command_prefix="etcdctl --endpoints 127.0.0.1:2379 --cacert /var/lib/localkube/certs/etcd/ca.crt --cert /var/lib/localkube/certs/etcd/peer.crt --key /var/lib/localkube/certs/etcd/peer.key"
ETCDCTL_API=3 $command_prefix get /registry/namespaces/ns-test -w=json

## ,----------- Sample Output
## | / # ETCDCTL_API=3 $command_prefix get /registry/namespaces/ns-test -w=json
## | 2018-07-19 05:07:36.115210 I | warning: ignoring ServerName for user-provided CA for backwards compatibility is deprecated
## | {"header":{"cluster_id":14841639068965178418,"member_id":10276657743932975437,"revision":267495,"raft_term":3},"kvs":[{"key":"L3JlZ2lzdHJ5L25hbWVzcGFjZXMvbnMtdGVzdA==","create_revision":267291,"mod_revision":267291,"version":1,"value":"azhzAAoPCgJ2MRIJTmFtZXNwYWNlEvwBCuEBCgducy10ZXN0EgAaACIAKiQ1YTgxNWQyYy04YjExLTExZTgtYThjOC0wODAwMjdjYmFlYTQyADgAQggImLjA2gUQAGKZAQowa3ViZWN0bC5rdWJlcm5ldGVzLmlvL2xhc3QtYXBwbGllZC1jb25maWd1cmF0aW9uEmV7ImFwaVZlcnNpb24iOiJ2MSIsImtpbmQiOiJOYW1lc3BhY2UiLCJtZXRhZGF0YSI6eyJhbm5vdGF0aW9ucyI6e30sIm5hbWUiOiJucy10ZXN0IiwibmFtZXNwYWNlIjoiIn19CnoAEgwKCmt1YmVybmV0ZXMaCAoGQWN0aXZlGgAiAA=="}],"count":1}
## `-----------

```

# Useful link

```
https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/
https://coreos.com/etcd/docs/latest/v2/api.html
https://coreos.com/etcd/docs/latest/dev-guide/interacting_v3.html
https://stackoverflow.com/questions/47807892/how-to-access-kubernetes-keys-in-etcd
```
