Explore what data is stored in kubernetes etcd

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

- Check etcd data folder in minikube

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
