Table of Contents
=================

   * [Explore etcd for label storage](#explore-etcd-for-label-storage)
   * [Search Resource by labels](#search-resource-by-labels)
   * [Use k8s go client to watch events based on labels](#use-k8s-go-client-to-watch-events-based-on-labels)

# Explore etcd for label storage

All labels are stored in k8s etcd

https://github.com/dennyzhang/dennytest/tree/master/kubernetes/k8s_etcd

# Search Resource by labels
In below, it will:
- Create 2 pv: pv1 and pv2. They are using two different labels: `app=label1` and `app=label2`.
- Then create two PVCs, which use the corresponding volumes by labels.

See more: [pv.yaml](https://github.com/dennyzhang/dennytest/blob/master/kubernetes/k8s_label/pv.yaml#L49-L51)

```
kubectl apply -f ./pv.yaml

## ,-----------
## | bash-3.2$  kubectl apply -f ./pv.yaml
## | namespace/ns-test created
## | persistentvolume/pv1 created
## | persistentvolume/pv2 created
## | persistentvolumeclaim/pvc1 created
## | persistentvolumeclaim/pvc2 created
## `-----------

# List pv and pvc
## ,-----------
## | bash-3.2$ kubectl get -n ns-test pv
## | NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM          STORAGECLASS   REASON    AGE
## | pv1                                        10Gi       RWO            Retain           Available                                           29s
## | pv2                                        10Gi       RWO            Retain           Available                                           29s
## | pvc-1ae0c63f-8ec3-11e8-80ab-080027b7ac6c   5Gi        RWO            Delete           Bound       ns-test/pvc1   standard                 29s
## | pvc-1ae22a81-8ec3-11e8-80ab-080027b7ac6c   6Gi        RWO            Delete           Bound       ns-test/pvc2   standard                 29s
## `-----------
```

- Search volume by labels

```
## ,-----------
## | bash-3.2$ kubectl get -n ns-test pv --selector app=label1
## | NAME      CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM     STORAGECLASS   REASON    AGE
## | pv1       10Gi       RWO            Retain           Available                                      1m
## | bash-3.2$ 
## `-----------
```

# Use k8s go client to watch events based on labels

- Get k8s go client

```
git clone git@github.com:vladimirvivien/k8s-client-examples.git
cd k8s-client-examples/pvcwatch

# build code
go build .
```

- Configure context

```
# Here we use minikube to simplify the tests
eval $(minikube docker-env)
# Start watcher in one terminal. Let's name it as pvc terminal
./pvcwatch
```

- Trigger some pvc deletion events, pvcwatch shall give us some notification

```
# In one terminal, delete PVCs
## ,-----------
## | bash-3.2$ kubectl delete -f ./pv.yaml
## | namespace "ns-test" deleted
## | persistentvolume "pv1" deleted
## | persistentvolume "pv2" deleted
## | persistentvolumeclaim "pvc1" deleted
## | persistentvolumeclaim "pvc2" deleted
## `-----------

# In the pvc terminal, we shall see below events
## ,----------- ./pvcwatch
## | 2018/07/23 15:45:30 PVC pvc2 removed, size 6Gi
## | 2018/07/23 15:45:30 Claim usage normal: max 200Gi at 5Gi
## | 2018/07/23 15:45:30 *** Taking action ***
## | 2018/07/23 15:45:30 
## | At 2.5% claim capcity (5Gi/200Gi)
## | 2018/07/23 15:45:30 PVC pvc1 removed, size 5Gi
## | 2018/07/23 15:45:30 Claim usage normal: max 200Gi at 0
## | 2018/07/23 15:45:30 *** Taking action ***
## | 2018/07/23 15:45:30 
## `-----------
```

- Trigger some pvc creation events

```
# In one terminal, create PVCs
## ,-----------
## | bash-3.2$ kubectl apply -f ./pv.yaml
## | namespace/ns-test created
## | persistentvolume/pv1 created
## | persistentvolume/pv2 created
## | persistentvolumeclaim/pvc1 created
## | persistentvolumeclaim/pvc2 created
## `-----------

# In the pvc terminal, we shall see below events
## ,-----------
## | 2018/07/23 15:47:35 PVC pvc1 added, claim size 5Gi
## | 2018/07/23 15:47:35 
## | At 2.5% claim capcity (5Gi/200Gi)
## | 2018/07/23 15:47:35 
## | At 2.5% claim capcity (5Gi/200Gi)
## | 2018/07/23 15:47:35 PVC pvc2 added, claim size 6Gi
## | 2018/07/23 15:47:35 
## | At 5.5% claim capcity (11Gi/200Gi)
## | 2018/07/23 15:47:35 
## | At 5.5% claim capcity (11Gi/200Gi)
## | 2018/07/23 15:47:35 
## | At 5.5% claim capcity (11Gi/200Gi)
## | 2018/07/23 15:47:35 
## | At 5.5% claim capcity (11Gi/200Gi)
## | 2018/07/23 15:47:35 
## | At 5.5% claim capcity (11Gi/200Gi)
## | 2018/07/23 15:47:35 
## | At 5.5% claim capcity (11Gi/200Gi)
## | 2018/07/23 15:47:35 
## | At 5.5% claim capcity (11Gi/200Gi)
## `-----------
```

- Useful links

https://github.com/kubernetes/client-go

https://github.com/vladimirvivien/k8s-client-examples/tree/master/go

https://hackernoon.com/top-10-kubernetes-tips-and-tricks-27528c2d0222
