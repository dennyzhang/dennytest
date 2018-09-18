# Start pod and mount dockerd socket file

See [pod.yaml](https://github.com/dennyzhang/dennytest/blob/master/kubernetes/k8s_security/pod.yaml#L14-L20)

```
# start minikube

kubectl apply -f ./pod.yaml

kubectl exec -it dummy sh

apt-get install -y jq

# Inside pod, run "docker ps" with a specific docker socket file

export DOCKER_HOST=unix:///myrun/docker.sock
docker ps

## ,-----------
## | # docker ps
## | CONTAINER ID        IMAGE                                      COMMAND                  CREATED             STATUS              PORTS               NAMES
## | d4c11b1e9ec2        getintodevops/jenkins-withdocker           "/bin/tini -- /usr..."   54 seconds ago      Up 53 seconds                           k8s_dummy_dummy_default_5c6846b0-ba19-11e8-9d98-0800274f164d_0
## | 9895d02774da        k8s.gcr.io/pause-amd64:3.1                 "/pause"                 4 minutes ago       Up 4 minutes                            k8s_POD_dummy_default_5c6846b0-ba19-11e8-9d98-0800274f164d_0
## | 2be67e884604        k8s.gcr.io/k8s-dns-sidecar-amd64           "/sidecar --v=2 --..."   2 days ago          Up 2 days                               k8s_sidecar_kube-dns-86f4d74b45-bqj27_kube-system_c575c318-b7ea-11e8-9d98-0800274f164d_0
## | 29437b39c863        k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64     "/dnsmasq-nanny -v..."   2 days ago          Up 2 days                               k8s_dnsmasq_kube-dns-86f4d74b45-bqj27_kube-system_c575c318-b7ea-11e8-9d98-0800274f164d_0
## | 1fa9b08b1b59        gcr.io/k8s-minikube/storage-provisioner    "/storage-provisioner"   2 days ago          Up 2 days                               k8s_storage-provisioner_storage-provisioner_kube-system_c728cae8-b7ea-11e8-9d98-0800274f164d_0
## | 7f24d8afd70f        k8s.gcr.io/kubernetes-dashboard-amd64      "/dashboard --inse..."   2 days ago          Up 2 days                               k8s_kubernetes-dashboard_kubernetes-dashboard-5498ccf677-6bv8q_kube-system_c66f1afb-b7ea-11e8-9d98-0800274f164d_0
## | c8003d3c7e07        k8s.gcr.io/metrics-server-amd64            "/metrics-server -..."   2 days ago          Up 2 days                               k8s_metrics-server_metrics-server-85c979995f-9xv8q_kube-system_c68044f8-b7ea-11e8-9d98-0800274f164d_0
## | a3dce891af3c        k8s.gcr.io/k8s-dns-kube-dns-amd64          "/kube-dns --domai..."   2 days ago          Up 2 days                               k8s_kubedns_kube-dns-86f4d74b45-bqj27_kube-system_c575c318-b7ea-11e8-9d98-0800274f164d_0
## | 3cddc05c8d2c        k8s.gcr.io/kube-proxy-amd64                "/usr/local/bin/ku..."   2 days ago          Up 2 days                               k8s_kube-proxy_kube-proxy-lv8hp_kube-system_c5397001-b7ea-11e8-9d98-0800274f164d_0
## | 1256e0ac23bf        k8s.gcr.io/pause-amd64:3.1                 "/pause"                 2 days ago          Up 2 days                               k8s_POD_storage-provisioner_kube-system_c728cae8-b7ea-11e8-9d98-0800274f164d_0
## | bb26b54aea83        k8s.gcr.io/pause-amd64:3.1                 "/pause"                 2 days ago          Up 2 days                               k8s_POD_metrics-server-85c979995f-9xv8q_kube-system_c68044f8-b7ea-11e8-9d98-0800274f164d_0
## | 8ed079426794        k8s.gcr.io/pause-amd64:3.1                 "/pause"                 2 days ago          Up 2 days                               k8s_POD_kubernetes-dashboard-5498ccf677-6bv8q_kube-system_c66f1afb-b7ea-11e8-9d98-0800274f164d_0
## | 1d5ead0ba2f8        k8s.gcr.io/pause-amd64:3.1                 "/pause"                 2 days ago          Up 2 days                               k8s_POD_kube-dns-86f4d74b45-bqj27_kube-system_c575c318-b7ea-11e8-9d98-0800274f164d_0
## | 86262d45ccfd        k8s.gcr.io/pause-amd64:3.1                 "/pause"                 2 days ago          Up 2 days                               k8s_POD_kube-proxy-lv8hp_kube-system_c5397001-b7ea-11e8-9d98-0800274f164d_0
## | b738b23b9a24        k8s.gcr.io/kube-controller-manager-amd64   "kube-controller-m..."   2 days ago          Up 2 days                               k8s_kube-controller-manager_kube-controller-manager-minikube_kube-system_a0c286a874b5cc23b80fa6e5452e2316_0
## | 58c0a10d1a43        k8s.gcr.io/kube-apiserver-amd64            "kube-apiserver --..."   2 days ago          Up 2 days                               k8s_kube-apiserver_kube-apiserver-minikube_kube-system_e644e5cd0a490152a8c0f9c316d474b8_0
## | 819fb736a44b        k8s.gcr.io/etcd-amd64                      "etcd --trusted-ca..."   2 days ago          Up 2 days                               k8s_etcd_etcd-minikube_kube-system_ec80f8b3827b5447e695a4044668d66b_0
## | d1c43c83fd22        k8s.gcr.io/kube-addon-manager              "/opt/kube-addons.sh"    2 days ago          Up 2 days                               k8s_kube-addon-manager_kube-addon-manager-minikube_kube-system_3afaf06535cc3b85be93c31632b765da_0
## `-----------
```
