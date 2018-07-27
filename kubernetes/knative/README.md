
# Table of Contents

1.  [Deploy knative on minikube](#org9fba89a)
    1.  [Useful tips](#org852f042)
    2.  [hello world setup](#org8d211b0)
    3.  [minikube start: is super slow: more than 10 minutes](#org521e425)
    4.  [get pods stucks in ContainerCreating state: takes more than 15 minutes](#orgab72693)
    5.  [knative serving deployment takes more than 5 minutes](#org0aa65cf)
    6.  [warm-up takes 11 seconds](#org559a42d)
    7.  [More Resources](#org9780bd7)



<a id="org9fba89a"></a>

# DONE Deploy knative on minikube

https://github.com/knative/docs/blob/master/install/Knative-with-Minikube.md  


<a id="org852f042"></a>

## Useful tips

kubectl describe services.serving.knative.dev helloworld-go2  

watch "kubectl get pods -n istio-system; echo "\n"; kubectl get pods -n knative-serving"  
kubectl get pods -n knative-serving  

     /Users/zdenny  kubectl describe services.serving.knative.dev helloworld-go2                                                                          ✔ 0
    Name:         helloworld-go2
    Namespace:    default
    Labels:       <none>
    Annotations:  kubectl.kubernetes.io/last-applied-configuration={"apiVersion":"serving.knative.dev/v1alpha1","kind":"Service","metadata":{"annotations":{},"name":"helloworld-go2","namespace":"default"},"spec":{"runL...
    API Version:  serving.knative.dev/v1alpha1
    Kind:         Service
    Metadata:
     Cluster Name:
     Creation Timestamp:  2018-07-26T06:50:32Z
     Generation:          1
     Resource Version:    4814
     Self Link:           /apis/serving.knative.dev/v1alpha1/namespaces/default/services/helloworld-go2
     UID:                 313a146e-90a0-11e8-b2c6-080027a8db9e
    Spec:
     Generation:  1
     Run Latest:
       Configuration:
         Revision Template:
           Metadata:
             Creation Timestamp:  <nil>
           Spec:
             Concurrency Model:  Multi
             Container:
               Env:
                 Name:   TARGET
                 Value:  Go Sample v2
               Image:    docker.io/denny/helloworld-go
               Name:
               Resources:
    Status:
     Conditions:
       Last Transition Time:        2018-07-26T06:50:34Z
       Message:                     Revision "helloworld-go2-00001" failed with message: "UNAUTHORIZED: \"authentication required\"".
       Reason:                      RevisionFailed
       Status:                      False
       Type:                        ConfigurationsReady
       Last Transition Time:        2018-07-26T06:50:53Z
       Message:                     Configuration "helloworld-go2" does not have any ready Revision.
       Reason:                      RevisionMissing
       Status:                      False
       Type:                        RoutesReady
       Last Transition Time:        2018-07-26T06:51:23Z
       Message:                     Configuration "helloworld-go2" does not have any ready Revision.
       Reason:                      RevisionMissing
       Status:                      False
       Type:                        Ready
     Domain:                        helloworld-go2.default.example.com
     Domain Internal:               helloworld-go2.default.svc.cluster.local
     Latest Created Revision Name:  helloworld-go2-00001
     Observed Generation:           1


<a id="org8d211b0"></a>

## hello world setup


### Install virtualbox, minikube


### Start infra

    minikube start --memory=8192 --cpus=4 \
      --kubernetes-version=v1.10.5 \
      --vm-driver=virtualbox \
      --bootstrapper=kubeadm \
      --extra-config=controller-manager.cluster-signing-cert-file="/var/lib/localkube/certs/ca.crt" \
      --extra-config=controller-manager.cluster-signing-key-file="/var/lib/localkube/certs/ca.key" \
      --extra-config=apiserver.admission-control="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount,DefaultStorageClass,MutatingAdmissionWebhook"

    Every 1.0s: kubectl get pods -n istio-system                                                                                                          zdenny-a02.vmware.com: Wed Jul 25 23:29:20 2018
    
    NAME                                       READY     STATUS      RESTARTS   AGE
    istio-citadel-7bdc7775c7-ssdkj             1/1       Running     0          15m
    istio-cleanup-old-ca-gw2sk                 0/1       Completed   0          15m
    istio-egressgateway-795fc9b47-hsqrd        1/1       Running     0          15m
    istio-ingress-84659cf44c-5vtzd             1/1       Running     0          15m
    istio-ingressgateway-7d89dbf85f-nkcbc      1/1       Running     0          15m
    istio-mixer-post-install-cjxsx             0/1       Completed   0          15m
    istio-pilot-66f4dd866c-5q7kv               2/2       Running     0          15m
    istio-policy-76c8896799-29trn              2/2       Running     0          15m
    istio-sidecar-injector-645c89bc64-mv99l    1/1       Running     0          15m
    istio-statsd-prom-bridge-949999c4c-rqngn   1/1       Running     0          15m
    istio-telemetry-6554768879-mjqjw           2/2       Running     0          15m


### Deploy a sample application

https://github.com/knative/docs/blob/master/serving/samples/helloworld-go/README.md  

https://github.com/knative/docs/blob/master/install/getting-started-knative-app.md  

    docker build -t denny/knative:helloworld_go .
    
    docker push denny/knative:helloworld_go
    
    kubectl apply -f service.yaml
    
    kubectl get svc knative-ingressgateway -n istio-system
    
    kubectl get services.serving.knative.dev helloworld-go  -o=custom-columns=NAME:.metadata.name,DOMAIN:.status.domain
    
    curl -H "Host: helloworld-go.default.example.com" http://10.100.91.133
    
    https://github.com/knative/docs/blob/master/install/getting-started-knative-app.md
    
    curl -I -H "Host: helloworld-go.default.example.com" http://10.0.2.15:32380
    
    docker build -t denny/knative:helloworld_go .


<a id="org521e425"></a>

## DONE minikube start: is super slow: more than 10 minutes


<a id="orgab72693"></a>

## DONE get pods stucks in ContainerCreating state: takes more than 15 minutes


<a id="org0aa65cf"></a>

## DONE knative serving deployment takes more than 5 minutes


<a id="org559a42d"></a>

## DONE warm-up takes 11 seconds

    $ time  curl  -H "Host: helloworld-go4.default.example.com" http://${IP_ADDRESS}
    Hello World: Go Sample v4!
    
    real	0m11.426s
    user	0m0.003s
    sys	0m0.001s


<a id="org9780bd7"></a>

## More Resources

<div class="HTML">
<a href="https://www.dennyzhang.com"><img align="right" width="201" height="268" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/denny_201706.png"></a>  

<a href="https://www.dennyzhang.com"><img align="right" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/dns_small.png"></a>  

</div>

