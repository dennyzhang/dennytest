
```
kubectl apply -f https://raw.githubusercontent.com/dennyzhang/dennytest/master/k8s_crd/crd.yml
kubectl apply -f https://raw.githubusercontent.com/dennyzhang/dennytest/master/k8s_crd/project.yml

kubectl get projects
# ,----------- Sample Output
# | bash-3.2$ kubectl get projects
# | NAME              CREATED AT
# | example-project   7s
# `-----------
```

https://github.com/martin-helmich/kubernetes-crd-example
