
```
kubectl apply -f https://raw.githubusercontent.com/dennyzhang/dennytest/master/k8s_crd/crd.yml
kubectl apply -f https://raw.githubusercontent.com/dennyzhang/dennytest/master/k8s_crd/project.yml

kubectl get projects
# ,----------- Sample Output
# | bash-3.2$ kubectl get projects
# | NAME              CREATED AT
# | example-project   7s
# `-----------

kubectl describe project example-project
# ,-----------
# | bash-3.2$ kubectl describe project example-project
# | Name:         example-project
# | Namespace:    default
# | Labels:       <none>
# | Annotations:  kubectl.kubernetes.io/last-applied-configuration={"apiVersion":"example.denny/v1alpha","kind":"Project","metadata":{"annotations":{},"name":"example-project","namespace":"default"},"spec":{"replicas":...
# | API Version:  example.denny/v1alpha
# | Kind:         Project
# | Metadata:
# |   Cluster Name:        
# |   Creation Timestamp:  2018-07-05T22:27:12Z
# |   Generation:          1
# |   Resource Version:    63787
# |   Self Link:           /apis/example.denny/v1alpha/namespaces/default/projects/example-project
# |   UID:                 8fc850ae-80a2-11e8-85ea-080027ede6a1
# | Spec:
# |   Replicas:  2
# | Events:      <none>
# `-----------
```

https://github.com/martin-helmich/kubernetes-crd-example

https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/#create-a-customresourcedefinition

Extend the Kubernetes API with CustomResourceDefinitions - Kubernetes
