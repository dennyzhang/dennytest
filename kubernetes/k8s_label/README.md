# explore etcd for label store

All labels are stored in k8s etcd

https://github.com/dennyzhang/dennytest/tree/master/k8s_etcd

# Search Resource by labels
In below, it will:
- Create 2 pv: pv1 and pv2. They are using two different labels: `app=label1` and `app=label2`.
- Then create two pvc, which use the corresponding volume by labels.

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

https://hackernoon.com/top-10-kubernetes-tips-and-tricks-27528c2d0222
