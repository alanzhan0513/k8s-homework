客户端请求到访问 kubernetes 流程
clientrequest -> loadbalance -> pod
因为是自架设的 kubeadm ，所以必须要在 load balance 上写上 externalIP。

``` bash
kubectl apply -f httpserver-configmap.yml
kubectl apply -f httpserver-deployment.yml
kubectl apply -f loadbalance.yml

curl -XGET -i http://192.168.83.130:8080/healthz
```

result
```
$ curl -XGET -i http://192.168.83.130:8080/healthz
HTTP/1.1 200 OK
Version: V2
Date: Sun, 06 Mar 2022 07:05:17 GMT
Content-Length: 2
Content-Type: text/plain; charset=utf-8
```