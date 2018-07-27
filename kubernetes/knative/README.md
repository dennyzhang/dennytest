
# Table of Contents

1.  [Deploy knative on minikube](#org5ab9dc2)
    1.  [SNS link](#orgeb8ab7e)
    2.  [Useful tips](#org0343036)
    3.  [hello world setup](#orga52dd4e)
    4.  [Key Observations](#org391a47c)
    5.  [More Resources](#orgb65088a)



<a id="org5ab9dc2"></a>

# DONE Deploy knative on minikube

https://github.com/knative/docs/blob/master/install/Knative-with-Minikube.md  


<a id="orgeb8ab7e"></a>

## SNS link

<div class="HTML">
<a href="https://www.linkedin.com/in/dennyzhang001"><img src="https://www.dennyzhang.com/wp-content/uploads/sns/linkedin.png" alt="linkedin" /></a>  
<a href="https://github.com/DennyZhang"><img src="https://www.dennyzhang.com/wp-content/uploads/sns/github.png" alt="github" /></a>  
<a href="https://www.dennyzhang.com/slack" target="\_blank" rel="nofollow"><img src="http://slack.dennyzhang.com/badge.svg" alt="slack"/></a>  
<a href="https://github.com/DennyZhang"><img align="right" width="200" height="183" src="https://www.dennyzhang.com/wp-content/uploads/denny/watermark/github.png" /></a>  

</div>


<a id="org0343036"></a>

## Useful tips

-   kubectl describe services.serving.knative.dev helloworld-go2

-   watch "kubectl get pods -n istio-system; echo "\n"; kubectl get pods -n knative-serving"

-   kubectl get pods -n knative-serving

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


<a id="orga52dd4e"></a>

## hello world setup


### Install virtualbox, minikube


### Start infra

https://github.com/knative/docs/blob/master/install/Knative-with-Minikube.md#installing-knative-serving  

-   Start minikube vm

    minikube start --memory=8192 --cpus=4 \
      --kubernetes-version=v1.10.5 \
      --vm-driver=virtualbox \
      --bootstrapper=kubeadm \
      --extra-config=controller-manager.cluster-signing-cert-file="/var/lib/localkube/certs/ca.crt" \
      --extra-config=controller-manager.cluster-signing-key-file="/var/lib/localkube/certs/ca.key" \
      --extra-config=apiserver.admission-control="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount,DefaultStorageClass,MutatingAdmissionWebhook"

-   Check status

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

-   Build docker image

    docker build -t denny/knative:helloworld_go .
    
    docker push denny/knative:helloworld_go

-   Create service

    kubectl apply -f service.yaml
    
    kubectl get svc knative-ingressgateway -n istio-system
    
    kubectl get services.serving.knative.dev helloworld-go  -o=custom-columns=NAME:.metadata.name,DOMAIN:.status.domain

-   Get Access IP, since we're using NodePort, instead of loadbalance service

\`\`\`  
echo \((minikube ip):\)(kubectl get svc knative-ingressgateway -n istio-system -o 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')  
\`\`\`  

https://github.com/knative/docs/blob/master/install/getting-started-knative-app.md  

-   Validate the service

\`\`\`  
curl -I -H "Host: helloworld-go.default.example.com" http://10.0.2.15:32380  
\`\`\`  


<a id="org391a47c"></a>

## Key Observations


### DONE minikube start: is super slow: more than 10 minutes


### DONE get pods stucks in ContainerCreating state: takes more than 15 minutes


### DONE knative serving deployment takes more than 5 minutes


### DONE warm-up takes 11 seconds

    $ time  curl  -H "Host: helloworld-go4.default.example.com" http://${IP_ADDRESS}
    Hello World: Go Sample v4!
    
    real	0m11.426s
    user	0m0.003s
    sys	0m0.001s


### DONE Istio yaml and Knative Serving yaml files are 3K-16.7K lines

https://github.com/knative/docs/blob/master/install/Knative-with-Minikube.md#installing-istio  

https://github.com/knative/docs/blob/master/install/Knative-with-Minikube.md#installing-knative-serving  


<a id="orgb65088a"></a>

## More Resources

<div class="HTML">
<a href="https://www.dennyzhang.com"><img align="right" width="201" height="268" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/denny_201706.png"></a>  

<a href="https://www.dennyzhang.com"><img align="right" src="https://raw.githubusercontent.com/USDevOps/mywechat-slack-group/master/images/dns_small.png"></a>  

</div>

