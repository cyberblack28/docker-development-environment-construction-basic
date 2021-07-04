# 第5章 Kubernetesでコンテナアプリケーションを動かすまで

## 5.1 kubectlコマンド

### 5.1.1 Podの作成

#### Podの起動

コマンド
```
kubectl run nginx --image=nginx
```
コマンド結果
```
pod/nginx created
```

#### Podステータスの確認

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          17s
```

#### PodのIPアドレス確認

コマンド
```
kubectl describe pod nginx
```
コマンド結果
```
Name:         nginx
Namespace:    default
Priority:     0
Node:         gke-k8s-cluster-default-pool-8ad2d901-1k7g/10.146.15.232
Start Time:   Sun, 16 May 2021 03:38:49 +0000
Labels:       run=nginx
Annotations:  <none>
Status:       Running
IP:           10.0.1.7
IPs:
  IP:  10.0.1.7
Containers:
  nginx:
    Container ID:   docker://6f16f781e6870404be19d4283abe835c708fb1b66db5123970e9c30e2728d81a
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:df13abe416e37eb3db4722840dd479b00ba193ac6606e7902331dcea50f4f1f2
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Sun, 16 May 2021 03:38:59 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-7z9kn (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-7z9kn:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-7z9kn
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age    From               Message
  ----    ------     ----   ----               -------
  Normal  Scheduled  5m48s  default-scheduler  Successfully assigned default/nginx to gke-k8s-cluster-default-pool-8ad2d901-1k7g
  Normal  Pulling    5m47s  kubelet            Pulling image "nginx"
  Normal  Pulled     5m40s  kubelet            Successfully pulled image "nginx"
  Normal  Created    5m39s  kubelet            Created container nginx
  Normal  Started    5m38s  kubelet            Started container nginx
```

#### busyboxイメージの利用

コマンド
```
kubectl run busybox --image=busybox -it --rm -- sh
```
コマンド結果
```
If you don't see a command prompt, try pressing enter.
/ # wget -O- 10.0.1.7
Connecting to 10.0.1.7 (10.0.1.7:80)
writing to stdout
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
-                    100% |*****************************************************************************************************************************************************|   612  0:00:00 ETA
written to stdout
/ # exit
pod "busybox" deleted
```

#### Podの削除

コマンド
```
kubectl delete pod nginx
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

#### マニフェストの作成

コマンド
```
kubectl run nginx --image=nginx --dry-run=client -o yaml > nginx.yaml
```

コマンド
```
cat nginx.yaml
```
コマンド結果
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

#### Podの作成

コマンド
```
kubectl apply -f nginx.yaml
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl get pod nginx
```
コマンド結果
```
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          69s
```

#### マニフェストを指定したPodの削除

コマンド
```
kubectl delete -f nginx.yaml
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

### 5.1.2 Deployment

#### マニフェストの作成

コマンド
```
kubectl create deployment nginx --image=nginx:1.19.10 --dry-run=client -o yaml > nginx-deployment.yaml
```

#### レプリカ数の指定

コマンド
```
vim nginx-deployment.yaml
```
コマンド結果
```
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx
  name: nginx
spec:
  replicas: 2 #「replicas: 1」を「replicas: 2」に変更
  selector:
    matchLabels:
      app: nginx
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx:1.19.10
        name: nginx
        resources: {}
status: {}
```

#### Deploymentの作成

コマンド
```
kubectl apply -f nginx-deployment.yaml
```
コマンド結果
```
deployment.apps/nginx created
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                    READY   STATUS    RESTARTS   AGE
nginx-f589cd57f-5hsbt   1/1     Running   0          11s
nginx-f589cd57f-jkhmq   1/1     Running   0          11s
```

コマンド
```
kubectl get deployments
```
コマンド結果
```
NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   2/2     2            2           39s
```

#### セルフヒーリング機能の確認

コマンド
```
kubectl delete pod nginx-f589cd57f-5hsbt
```
コマンド結果
```
pod "nginx-f589cd57f-5hsbt" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                    READY   STATUS    RESTARTS   AGE
nginx-f589cd57f-dbkns   1/1     Running   0          32s
nginx-f589cd57f-jkhmq   1/1     Running   0          2m4s
```

#### ローリングアップデート

コマンド
```
kubectl set image deployment nginx nginx=nginx:1.20.0
```
コマンド結果
```
deployment.apps/nginx image updated
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                     READY   STATUS    RESTARTS   AGE
nginx-7d994954c4-bs8jm   1/1     Running   0          48s
nginx-7d994954c4-nwvfl   1/1     Running   0          47s
```

#### Podの詳細確認

コマンド
```
kubectl describe pod nginx-7d994954c4-bs8jm
```
コマンド結果
```
Name:         nginx-7d994954c4-bs8jm
Namespace:    default
Priority:     0
Node:         gke-k8s-cluster-default-pool-8ad2d901-1k7g/10.146.15.232
Start Time:   Sun, 16 May 2021 04:30:09 +0000
Labels:       app=nginx
              pod-template-hash=7d994954c4
Annotations:  <none>
Status:       Running
IP:           10.0.1.14
IPs:
  IP:           10.0.1.14
Controlled By:  ReplicaSet/nginx-7d994954c4
Containers:
  nginx:
    Container ID:   docker://39b36266384ed090e745343f0e78ce2f36ecd7dccc6ca626a4702149c98e9abf
    Image:          nginx:1.20.0
    Image ID:       docker-pullable://nginx@sha256:ea4560b87ff03479670d15df426f7d02e30cb6340dcd3004cdfc048d6a1d54b4
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Sun, 16 May 2021 04:30:10 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-7z9kn (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-7z9kn:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-7z9kn
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  3m7s  default-scheduler  Successfully assigned default/nginx-7d994954c4-bs8jm to gke-k8s-cluster-default-pool-8ad2d901-1k7g
  Normal  Pulled     3m7s  kubelet            Container image "nginx:1.20.0" already present on machine
  Normal  Created    3m7s  kubelet            Created container nginx
  Normal  Started    3m6s  kubelet            Started container nginx
```

#### リビジョンの確認

コマンド
```
kubectl rollout history deployment nginx
```
コマンド結果
```
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1         <none>
2         <none>
```

#### 各リビジョンの確認

コマンド
```
kubectl rollout history deployment nginx --revision=1
```
コマンド結果
```
deployment.apps/nginx with revision #1
Pod Template:
  Labels:       app=nginx
        pod-template-hash=f589cd57f
  Containers:
   nginx:
    Image:      nginx:1.19.10
    Port:       <none>
    Host Port:  <none>
    Environment:        <none>
    Mounts:     <none>
  Volumes:      <none>
```

コマンド
```
kubectl rollout history deployment nginx --revision=2
```
コマンド結果
```
deployment.apps/nginx with revision #2
Pod Template:
  Labels:       app=nginx
        pod-template-hash=7d994954c4
  Containers:
   nginx:
    Image:      nginx:1.20.0
    Port:       <none>
    Host Port:  <none>
    Environment:        <none>
    Mounts:     <none>
  Volumes:      <none>
```

#### ロールバックの実行

コマンド
```
kubectl rollout undo deployment nginx
```
コマンド結果
```
deployment.apps/nginx rolled back
```

#### kubectl describeコマンドによるDeploymentの確認

コマンド
```
kubectl describe deployment nginx
```
コマンド結果
```
Name:                   nginx
Namespace:              default
CreationTimestamp:      Sun, 16 May 2021 04:25:42 +0000
Labels:                 app=nginx
Annotations:            deployment.kubernetes.io/revision: 3
Selector:               app=nginx
Replicas:               2 desired | 2 updated | 2 total | 2 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  app=nginx
  Containers:
   nginx:
    Image:        nginx:1.19.10
    Port:         <none>
    Host Port:    <none>
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   nginx-f589cd57f (2/2 replicas created)
Events:
  Type    Reason             Age                From                   Message
  ----    ------             ----               ----                   -------
  Normal  ScalingReplicaSet  12m                deployment-controller  Scaled up replica set nginx-7d994954c4 to 1
  Normal  ScalingReplicaSet  12m                deployment-controller  Scaled down replica set nginx-f589cd57f to 1
  Normal  ScalingReplicaSet  12m                deployment-controller  Scaled up replica set nginx-7d994954c4 to 2
  Normal  ScalingReplicaSet  12m                deployment-controller  Scaled down replica set nginx-f589cd57f to 0
  Normal  ScalingReplicaSet  34s                deployment-controller  Scaled up replica set nginx-f589cd57f to 1
  Normal  ScalingReplicaSet  32s (x2 over 16m)  deployment-controller  Scaled up replica set nginx-f589cd57f to 2
  Normal  ScalingReplicaSet  32s                deployment-controller  Scaled down replica set nginx-7d994954c4 to 1
  Normal  ScalingReplicaSet  31s                deployment-controller  Scaled down replica set nginx-7d994954c4 to 0
```

#### 再度ロールバックを実行

コマンド
```
kubectl rollout undo deployment nginx --to-revision=2
```
コマンド結果
```
deployment.apps/nginx rolled back
```

コマンド
```
kubectl describe deployment nginx
```
コマンド結果
```
Name:                   nginx
Namespace:              default
CreationTimestamp:      Sun, 16 May 2021 04:25:42 +0000
Labels:                 app=nginx
Annotations:            deployment.kubernetes.io/revision: 4
Selector:               app=nginx
Replicas:               2 desired | 2 updated | 2 total | 2 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  app=nginx
  Containers:
   nginx:
    Image:        nginx:1.20.0
    Port:         <none>
    Host Port:    <none>
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   nginx-7d994954c4 (2/2 replicas created)
Events:
  Type    Reason             Age                  From                   Message
  ----    ------             ----                 ----                   -------
  Normal  ScalingReplicaSet  4m41s                deployment-controller  Scaled up replica set nginx-f589cd57f to 1
  Normal  ScalingReplicaSet  4m39s (x2 over 21m)  deployment-controller  Scaled up replica set nginx-f589cd57f to 2
  Normal  ScalingReplicaSet  4m39s                deployment-controller  Scaled down replica set nginx-7d994954c4 to 1
  Normal  ScalingReplicaSet  4m38s                deployment-controller  Scaled down replica set nginx-7d994954c4 to 0
  Normal  ScalingReplicaSet  100s (x2 over 16m)   deployment-controller  Scaled up replica set nginx-7d994954c4 to 1
  Normal  ScalingReplicaSet  98s (x2 over 16m)    deployment-controller  Scaled down replica set nginx-f589cd57f to 1
  Normal  ScalingReplicaSet  98s (x2 over 16m)    deployment-controller  Scaled up replica set nginx-7d994954c4 to 2
  Normal  ScalingReplicaSet  97s (x2 over 16m)    deployment-controller  Scaled down replica set nginx-f589cd57f to 0
```

#### Deploymentの削除

コマンド
```
kubectl delete deployment nginx
```
コマンド結果
```
deployment.apps "nginx" deleted
```

コマンド
```
kubectl delete -f nginx-deployment.yaml
```
コマンド結果
```
deployment.apps "nginx" deleted
```

コマンド
```
kubectl get deployments
```
コマンド結果
```
No resources found in default namespace.
```

### 5.1.3 Service

#### Podの作成

コマンド
```
kubectl run nginx --image=nginx:1.20.0 --port=80
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl describe pod nginx
```
コマンド結果
```
Name:         nginx
Namespace:    default
Priority:     0
Node:         gke-k8s-cluster-default-pool-8ad2d901-1k7g/10.146.15.232
Start Time:   Sun, 16 May 2021 14:48:24 +0000
Labels:       run=nginx
Annotations:  <none>
Status:       Running
IP:           10.0.6.6
IPs:
  IP:  10.0.6.6
Containers:
  nginx:
    Container ID:   docker://5eac6fee54bd6a2bc14b7e10379cc0c2af8fea40efad02816c32cd675bace4b2
    Image:          nginx:1.20.0
    Image ID:       docker-pullable://nginx@sha256:ea4560b87ff03479670d15df426f7d02e30cb6340dcd3004cdfc048d6a1d54b4
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sun, 16 May 2021 14:48:25 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-7z9kn (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-7z9kn:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-7z9kn
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  2m5s  default-scheduler  Successfully assigned default/nginx to gke-k8s-cluster-default-pool-8ad2d901-1k7g
  Normal  Pulled     2m4s  kubelet            Container image "nginx:1.20.0" already present on machine
  Normal  Created    2m4s  kubelet            Created container nginx
  Normal  Started    2m4s  kubelet            Started container nginx
```

#### Serviceの作成

コマンド
```
kubectl expose pod nginx --name=nginx --port=80 --target-port=80
```
コマンド結果
```
service/nginx exposed
```

#### ClusterIPアドレスの確認

コマンド
```
kubectl get services
```
コマンド結果
```
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.3.240.1     <none>        443/TCP   11h
nginx        ClusterIP   10.3.248.150   <none>        80/TCP    73s
```

#### PodのIPアドレスの確認

コマンド
```
kubectl describe pod nginx
```
コマンド結果
```
：
：（ 省略）
：
Labels: run=nginx
Annotations: <none>
Status: Running
IP: 10.0.6.6
IPs:
IP: 10.0.6.6
：
：（ 省略）
：
```

#### Busybox Podの作成

コマンド
```
kubectl run busybox --image=busybox --rm -it /bin/sh
```
コマンド結果
```
If you don't see a command prompt, try pressing enter.
/ # wget -O- 10.3.248.150
Connecting to 10.3.248.150 (10.3.248.150:80)
writing to stdout
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
-                    100% |*****************************************************************************************************************************************************|   612  0:00:00 ETA
written to stdout
/ # wget -O- 10.0.6.6:80
Connecting to 10.0.6.6:80 (10.0.6.6:80)
writing to stdout
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
-                    100% |*****************************************************************************************************************************************************|   612  0:00:00 ETA
written to stdout
/ # exit
pod "busybox" deleted
```

#### ServiceとPodの削除

コマンド
```
kubectl delete service nginx
```
コマンド結果
```
service "nginx" deleted
```

コマンド
```
kubectl delete pod nginx
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get service
```
コマンド結果
```
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.3.240.1   <none>        443/TCP   11h
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

コマンド
```
kubectl run nginx --image=nginx:1.20.0 --port=80 --expose
```
コマンド結果
```
service/nginx created
pod/nginx created
```

コマンド
```
kubectl delete service nginx
```
コマンド結果
```
service "nginx" deleted
```

コマンド
```
kubectl delete pod nginx
```
コマンド結果
```
pod "nginx" deleted
```

### 5.1.4 ConfigMap

#### ConfigMapの作成

コマンド
```
kubectl create configmap sample-config1 --from-literal=var1=docker --from-literal=var2=kubernetes
```
コマンド結果
```
configmap/sample-config1 created
```

コマンド
```
kubectl get configmaps
```
コマンド結果
```
NAME               DATA   AGE
kube-root-ca.crt   1      4m31s
sample-config1     2      82s
```

コマンド
```
kubectl get configmap sample-config1 -o yaml
```
コマンド結果
```
apiVersion: v1
data:
  var1: docker
  var2: kubernetes
kind: ConfigMap
metadata:
  creationTimestamp: "2021-05-20T15:01:12Z"
  name: sample-config1
  namespace: default
  resourceVersion: "5779"
  selfLink: /api/v1/namespaces/default/configmaps/sample-config1
  uid: 20bfdae8-5bfd-467a-a4b5-67923d5f5725
```

#### マニフェストの確認

コマンド
```
$ cd
```

コマンド
```
$ cd docker-development-environment-construction-basic/Chapter05/5-1-5-01
```

コマンド
``` 
cat nginx-configmap1.yaml
```
コマンド結果
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    resources: {}
    envFrom:　　　　　　　　　　# envFrom:以降の3 行が参照するための設定
    - configMapRef:
        name: sample-config1
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

#### Podの作成

コマンド
```
kubectl create -f nginx-configmap1.yaml
```
コマンド結果
```
pod/nginx created
```

#### 環境変数の確認

コマンド
```
kubectl exec -it nginx -- env
```
コマンド結果
```
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=nginx
TERM=xterm
var1=docker
var2=kubernetes
KUBERNETES_SERVICE_PORT=443
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT=tcp://10.3.240.1:443
KUBERNETES_PORT_443_TCP=tcp://10.3.240.1:443
KUBERNETES_PORT_443_TCP_PROTO=tcp
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_ADDR=10.3.240.1
KUBERNETES_SERVICE_HOST=10.3.240.1
NGINX_VERSION=1.19.10
NJS_VERSION=0.5.3
PKG_RELEASE=1~buster
HOME=/root
```

#### PodとConfigMapの削除

コマンド
```
kubectl delete -f nginx-configmap1.yaml
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

コマンド
```
kubectl delete configmap sample-config1
```
コマンド結果
```
configmap "sample-config1" deleted
```

コマンド
```
kubectl get configmaps
```
コマンド結果
```
NAME               DATA   AGE
kube-root-ca.crt   1      26m
```

#### 作成したファイルからConfigMapを作成

コマンド
```
cat sample-config2.env
```
コマンド結果
```
var3=istio
var4=envoy
```

コマンド
```
kubectl create configmap sample-config2 --from-env-file=sample-config2.env
```
コマンド結果
```
configmap/sample-config2 created
```

コマンド
```
kubectl get configmaps
```
コマンド結果
```
NAME               DATA   AGE
kube-root-ca.crt   1      91m
sample-config2     2      112s
```

コマンド
```
kubectl get configmap sample-config2 -o yaml
```
コマンド結果
```
apiVersion: v1
data:
  var3: istio
  var4: envoy
kind: ConfigMap
metadata:
  creationTimestamp: "2021-05-20T16:17:51Z"
  name: sample-config2
  namespace: default
  resourceVersion: "33035"
  selfLink: /api/v1/namespaces/default/configmaps/sample-config2
  uid: 1d9e9ce7-399a-42eb-b0b8-186e1486b01b
```

#### PodからKeyを参照

コマンド
```
cat nginx-configmap2.yaml
```
コマンド結果
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    resources: {}
    env:
    - name: sample-config
      valueFrom:
        configMapKeyRef:
          name: sample-config2
          key: var3
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

コマンド
```
kubectl apply -f nginx-configmap2.yaml
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl exec -it nginx -- env
```
コマンド結果
```
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=nginx
TERM=xterm
sample-config=istio
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_ADDR=10.3.240.1
KUBERNETES_SERVICE_HOST=10.3.240.1
KUBERNETES_SERVICE_PORT=443
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT=tcp://10.3.240.1:443
KUBERNETES_PORT_443_TCP=tcp://10.3.240.1:443
KUBERNETES_PORT_443_TCP_PROTO=tcp
NGINX_VERSION=1.19.10
NJS_VERSION=0.5.3
PKG_RELEASE=1~buster
HOME=/root
```

#### PodとConfigMapの削除

コマンド
```
kubectl delete -f nginx-configmap2.yaml
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

コマンド
```
kubectl delete configmap sample-config2
```
コマンド結果
```
configmap "sample-config2" deleted
```

コマンド
```
kubectl get configmaps
```
コマンド結果
```
NAME               DATA   AGE
kube-root-ca.crt   1      114m
```

### 5.1.5 ConfigMapとKeyの参照

コマンド
```
cd ../5-1-6-01
```

コマンド
```
cat sample-config3.txt
```
コマンド結果
```
rook
vitess
containerd
helm
```

コマンド
```
kubectl create configmap sample-cmvolume --from-file=sample-config3=sample-config3.txt
```
コマンド結果
```
configmap/sample-cmvolume created
```

コマンド
```
kubectl get configmaps
```
コマンド結果
```
NAME               DATA   AGE
kube-root-ca.crt   1      39h
sample-cmvolume    1      119s
```

コマンド
```
kubectl get configmap sample-cmvolume -o yaml
```
コマンド結果
```
apiVersion: v1
data:
  sample-config3: |-
    rook
    vitess
    containerd
    helm
    prometheus
kind: ConfigMap
metadata:
  creationTimestamp: "2021-05-22T05:51:49Z"
  name: sample-cmvolume
  namespace: default
  resourceVersion: "835111"
  selfLink: /api/v1/namespaces/default/configmaps/sample-cmvolume
  uid: 53ff63a7-fa03-4c82-ab05-de7cee63a9b7
```

#### マニフェストの確認

コマンド
```
cat nginx-configmap3.yaml
```
コマンド結果
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  volumes:
  - name: cmvolume
    configMap:
      name: sample-cmvolume
  containers:
  - image: nginx
    name: nginx
    resources: {}
    volumeMounts:
    - name: cmvolume
      mountPath: /configmap
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

#### Podの作成

コマンド
```
kubectl apply -f nginx-configmap3.yaml
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl exec -it nginx -- cat /configmap/sample-config3
```
コマンド結果
```
rook
vitess
containerd
helm
```

コマンド
```
kubectl delete pod nginx
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

### 5.1.6 Secret

#### Secretの作成

コマンド
```
kubectl create secret generic sample-secret1 --from-literal=password=testp@ss
```
コマンド結果
```
secret/sample-secret1 created
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      39h
sample-secret1        Opaque                                1      118s
```

#### マニフェストの確認

コマンド
```
kubectl get secret sample-secret1 -o yaml
```
コマンド結果
```
apiVersion: v1
data:
  password: dGVzdHBAc3M=
kind: Secret
metadata:
  creationTimestamp: "2021-05-22T06:17:23Z"
  name: sample-secret1
  namespace: default
  resourceVersion: "844182"
  selfLink: /api/v1/namespaces/default/secrets/sample-secret1
  uid: 2841a4c4-0deb-4227-8853-d192a4d98219
type: Opaque
```

コマンド
```
echo dGVzdHBAc3M= | base64 -d
```
コマンド結果
```
testp@ss
```

#### Secretの削除

コマンド
```
kubectl delete secret sample-secret1
```
コマンド結果
```
secret "sample-secret1" deleted
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      39h
```

#### envファイルからSecretを作成

コマンド
```
cd ../5-1-7-01
```

コマンド
```
cat sample-secret2.env
```
コマンド結果
```
password=p@ssw0rd
```

コマンド
```
kubectl create secret generic sample-secret2 --from-env-file=sample-secret2.env
```
コマンド結果
```
secret/sample-secret2 created
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      39h
sample-secret2        Opaque                                1      85s
```

コマンド
```
kubectl get secret sample-secret2 -o yaml
```
コマンド結果
```
apiVersion: v1
data:
  password: cEBzc3cwcmQ=
kind: Secret
metadata:
  creationTimestamp: "2021-05-22T06:46:20Z"
  name: sample-secret2
  namespace: default
  resourceVersion: "854491"
  selfLink: /api/v1/namespaces/default/secrets/sample-secret2
  uid: 4accb5b3-e51c-42cc-9bb5-d9a9402c6796
type: Opaque
```

#### マニフェストの作成

コマンド
```
cat nginx-secret.yaml
```
コマンド結果
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - image: nginx
    name: nginx
    resources: {}
    env:
    - name: PASSWORD
      valueFrom:
        secretKeyRef:
          name: sample-secret2
          key: password
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

#### Podの作成とSecretの参照確認

コマンド
```
kubectl apply -f nginx-secret.yaml
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl describe pod nginx
```
コマンド結果
```
Name:         nginx
Namespace:    default
Priority:     0
Node:         gke-k8s-cluster-default-pool-c2daf5f7-zz3p/10.146.0.6
Start Time:   Sat, 22 May 2021 07:07:37 +0000
Labels:       run=nginx
Annotations:  <none>
Status:       Running
IP:           10.0.1.10
IPs:
  IP:  10.0.1.10
Containers:
  nginx:
    Container ID:   docker://84bf3909c0bb84d1378e0d46e2a335fc2462feadd1d0f0afe532d0e400accf60
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:df13abe416e37eb3db4722840dd479b00ba193ac6606e7902331dcea50f4f1f2
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Sat, 22 May 2021 07:07:39 +0000
    Ready:          True
    Restart Count:  0
    Environment:
      PASSWORD:  <set to the key 'password' in secret 'sample-secret2'>  Optional: false
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-dmlps (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-dmlps:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-dmlps
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  23s   default-scheduler  Successfully assigned default/nginx to gke-k8s-cluster-default-pool-c2daf5f7-zz3p
  Normal  Pulling    22s   kubelet            Pulling image "nginx"
  Normal  Pulled     21s   kubelet            Successfully pulled image "nginx"
  Normal  Created    21s   kubelet            Created container nginx
  Normal  Started    21s   kubelet            Started container nginx
```

コマンド
```
kubectl exec -it nginx -- env | grep PASSWORD
```
コマンド結果
```
PASSWORD=p@ssw0rd
```

#### SecretとPodの削除

コマンド
```
kubectl delete secret sample-secret2
```
コマンド結果
```
secret "sample-secret2" deleted
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      40h
```

コマンド
```
kubectl delete pod nginx
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

#### Secretの作成とVolumeデータの参照

コマンド
```
cat sample-secret3.txt
```
コマンド結果
```
admin=86asnNlW
operator=po89SYin
user=oshu894LD
```

コマンド
```
kubectl create secret generic sample-secret3 --from-file=sample-secret3.txt
```
コマンド結果
```
secret/sample-secret3 created
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      40h
sample-secret3        Opaque                                1      39s
```

コマンド
```
kubectl get secret sample-secret3 -o yaml
```
コマンド結果
```
apiVersion: v1
data:
  sample-secret3.txt: YWRtaW49ODZhc25ObFcKb3BlcmF0b3I9cG84OVNZaW4KdXNlcj1vc2h1ODk0TEQK
kind: Secret
metadata:
  creationTimestamp: "2021-05-22T07:22:59Z"
  name: sample-secret3
  namespace: default
  resourceVersion: "867562"
  selfLink: /api/v1/namespaces/default/secrets/sample-secret3
  uid: 835beb6f-3d14-4ae8-9e43-3e7450ca69a7
type: Opaque
```

#### マニフェストの作成（Volumeの定義）

コマンド
```
cat nginx-secret2.yaml
```
コマンド結果
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  volumes:
  - name: secretvolume
    secret:
      secretName: sample-secret3
  containers:
  - image: nginx
    name: nginx
    resources: {}
    volumeMounts:
    - name: secretvolume
      mountPath: /secret
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

コマンド
```
kubectl apply -f nginx-secret2.yaml
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl exec -it nginx -- cat /secret/sample-secret3.txt
```
コマンド結果
```
admin=86asnNlW
operator=po89SYin
user=oshu894LD
```

#### SecretとPodの削除

コマンド
```
kubectl delete secret sample-secret3
```
コマンド結果
```
secret "sample-secret3" deleted
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      41h
```

コマンド
```
kubectl delete pod nginx
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

### 5.1.7 Multi Container Pod

コマンド
```
cd ../5-1-8-01
```

コマンド
```
cat multicontainer.yaml
```
コマンド結果
```
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx-pod
  name: nginx-pod
spec:
  volumes:
  - name: share-data
    emptyDir: {}
  containers:
  - image: nginx:1.20.0
    name: nginx
    volumeMounts:
    - name: share-data
      mountPath: /usr/share/nginx/html
    resources: {}
  - name: work-container
    image: busybox
    volumeMounts:
    - name: share-data
      mountPath: /data
    command: ["/bin/sh"]
    args: ["-c", "echo Hello from the work-container > /data/index.html;sleep 2400"]
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
```

コマンド
```
kubectl apply -f multicontainer.yaml
```
コマンド結果
```
pod/nginx-pod created
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME        READY   STATUS    RESTARTS   AGE
nginx-pod   2/2     Running   0          3m26s
```

コマンド
```
kubectl exec -it nginx-pod -c nginx -- /bin/sh
```
コマンド & 結果
```
# curl localhost
Hello from the work-container
# ls /usr/share/nginx/html
index.html
# exit
```

コマンド
```
kubectl delete pod nginx-pod
```
コマンド結果
```
pod "nginx-pod" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

### 5.1.8 createとapply

コマンド
```
cd ../5-1-9-01
```

コマンド
```
kubectl create -f nginx.yaml
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl create -f nginx.yaml
```
コマンド結果
```
Error from server (AlreadyExists): error when creating "nginx.yaml": pods "nginx" already exists
```

コマンド
```
kubectl delete -f nginx.yaml
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl apply -f nginx.yaml
```
コマンド結果
```
pod/nginx created
```

コマンド
```
kubectl apply -f nginx.yaml
```
コマンド結果
```
pod/nginx configured
```

コマンド
```
kubectl delete -f nginx.yaml
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl delete pod nginx
```
コマンド結果
```
pod "nginx" deleted
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

## 5・2 Kubernetesでアプリケーションを動かす

### 5.2.2 NFSサーバの作成

#### gcePersistentDiskの作成

コマンド
```
gcloud compute disks create nfs-disk --size=10GB --zone=asia-northeast1-a
```
コマンド結果
```
WARNING: You have selected a disk size of under [200GB]. This may result in poor I/O performance. For more information, see: https://developers.google.com/compute/docs/disks#performance.
Created [https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-a/disks/nfs-disk].
NAME      ZONE               SIZE_GB  TYPE         STATUS
nfs-disk  asia-northeast1-a  10       pd-standard  READY
New disks are unformatted. You must format and mount a disk before it
can be used. You can find instructions on how to do this at:

https://cloud.google.com/compute/docs/disks/add-persistent-disk#formatting
```

コマンド
```
gcloud compute disks describe --zone=asia-northeast1-a nfs-disk
```
コマンド結果
```
creationTimestamp: '2021-05-22T05:09:11.734-07:00'
id: '6161119284946094728'
kind: compute#disk
labelFingerprint: 42WmSpB8rSM=
name: nfs-disk
physicalBlockSizeBytes: '4096'
selfLink: https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-a/disks/nfs-disk
sizeGb: '10'
status: READY
type: https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-a/diskTypes/pd-standard
zone: https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-a
```

#### NFS サーバのPersistentVolumeとPersistentVolumeClaimの作成

コマンド
```
cd ../5-2-2-01
```

コマンド
```
cat nfs-pv.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: nfs-pv
spec:
  storageClassName: nfs
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  gcePersistentDisk: #「gcePersistentDisk」を指定する定義
    pdName: nfs-disk
    fsType: ext4
```

コマンド
```
kubectl apply -f nfs-pv.yaml
```
コマンド結果
```
persistentvolume/nfs-pv created
```

コマンド
```
kubectl get persistentvolume nfs-pv
```
コマンド結果
```
NAME     CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
nfs-pv   10Gi       RWO            Retain           Available           nfs                     4m50s
```

コマンド
```
cat nfs-pvc.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: nfs-pvc
spec:
  storageClassName: nfs
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```

コマンド
```
kubectl apply -f nfs-pvc.yaml
```
コマンド結果
```
persistentvolumeclaim/nfs-pvc created
```

コマンド
```
kubectl get persistentvolumes,persistentvolumeclaims
```
コマンド結果
```
NAME                      CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM             STORAGECLASS   REASON   AGE
persistentvolume/nfs-pv   10Gi       RWO            Retain           Bound    default/nfs-pvc   nfs                     7m30s

NAME                            STATUS   VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/nfs-pvc   Bound    nfs-pv   10Gi       RWO            nfs            11s
```

#### NFS サーバのDeploymentとServiceの作成

コマンド
```
cat nfs-server.yaml
```
コマンド結果
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nfs-server
spec:
  replicas: 1
  selector:
    matchLabels:
      role: nfs-server
  template:
    metadata:
      labels:
        role: nfs-server
    spec:
      containers:
      - name: nfs-server
        image: gcr.io/google_containers/volume-nfs:0.8
        ports:
          - name: nfs
            containerPort: 2049
          - name: mountd
            containerPort: 20048
          - name: rpcbind
            containerPort: 111
        securityContext:
          privileged: true
        volumeMounts:
          - mountPath: /exports
            name: nfs-local-storage
      volumes:  #「nfs-pvc」を指定する定義
        - name: nfs-local-storage
          persistentVolumeClaim:
            claimName: nfs-pvc
```

コマンド
```
kubectl apply -f nfs-server.yaml
```
コマンド結果
```
deployment.apps/nfs-server created
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                          READY   STATUS    RESTARTS   AGE
nfs-server-6f7fc97dfd-6kzmm   1/1     Running   0          2m45s
```

コマンド
```
cat nfs-service.yaml
```
コマンド結果
```
apiVersion: v1
kind: Service
metadata:
  name: nfs-service
spec:
  ports:
    - name: nfs
      port: 2049
    - name: mountd
      port: 20048
    - name: rpcbind
      port: 111
  selector:
    role: nfs-server
```

コマンド
```
kubectl apply -f nfs-service.yaml
```
コマンド結果
```
service/nfs-service created
```

コマンド
```
kubectl get services
```
コマンド結果
```
NAME          TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                      AGE
kubernetes    ClusterIP   10.48.0.1    <none>        443/TCP                      11m
nfs-service   ClusterIP   10.48.4.25   <none>        2049/TCP,20048/TCP,111/TCP   10s
```

### 5.2.3 Secretの作成

コマンド
```
kubectl create secret generic mysql --from-literal=password=mysqlp@ssw0d
```
コマンド結果
```
secret/mysql created
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-lm5kz   kubernetes.io/service-account-token   3      14m
mysql                 Opaque                                1      12s
```

コマンド
```
kubectl get secret mysql -o yaml
```
コマンド結果
```
apiVersion: v1
data:
  password: bXlzcWxwQHNzdzBk
kind: Secret
metadata:
  creationTimestamp: "2021-05-25T13:05:58Z"
  name: mysql
  namespace: default
  resourceVersion: "6035"
  selfLink: /api/v1/namespaces/default/secrets/mysql
  uid: de5b94bc-f58f-4c36-a372-8c1c303cbd68
type: Opaque
```

### 5.2.4 PersistentVolumeとPersistentVolumeClaimの作成

#### PersistentVolumeの作成

コマンド
```
cd ../5-2-4-01
```

コマンド
```
vim mysql-pv.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: mysql-pv
  labels:
    type: local
spec:
  storageClassName: mysql
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteMany
  nfs:
    server: 10.48.4.25 #「nfs-service」のCLUSTER-IP を定義
    path: /
```

コマンド
```
kubectl apply -f mysql-pv.yaml
```
コマンド結果
```
persistentvolume/mysql-pv created
```

コマンド
```
vim wordpress-pv.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: wordpress-pv
  labels:
    type: local
spec:
  storageClassName: wordpress
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteMany
  nfs:
    server: 10.48.4.25 #「nfs-service」のCLUSTER-IP を定義
    path: /
```

コマンド
```
kubectl apply -f wordpress-pv.yaml
```
コマンド結果
```
persistentvolume/wordpress-pv created
```

コマンド
```
kubectl get persistentvolumes
```
コマンド結果
```
NAME           CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM             STORAGECLASS   REASON   AGE
mysql-pv       10Gi       RWX            Retain           Available                     mysql                   2m15s
nfs-pv         10Gi       RWO            Retain           Bound       default/nfs-pvc   nfs                     15m
wordpress-pv   10Gi       RWX            Retain           Available                     wordpress               4s
```

#### PersistentVolumeClaimの作成

コマンド
```
cat mysql-pvc.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mysql-pvc
  labels:
    app: wordpress
    tier: mysql
spec:
  storageClassName: mysql
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 5Gi
```

コマンド
```
kubectl apply -f mysql-pvc.yaml
```
コマンド結果
```
persistentvolumeclaim/mysql-pvc created
```

コマンド
```
cat wordpress-pvc.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: wordpress-pvc
  labels:
    app: wordpress
    tier: wordpress
spec:
  storageClassName: wordpress
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 5Gi
```

コマンド
```
kubectl apply -f wordpress-pvc.yaml
```
コマンド結果
```
persistentvolumeclaim/wordpress-pvc created
```

コマンド
```
kubectl get persistentvolumes,persistentvolumeclaims
```
コマンド結果
```
NAME                            CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
persistentvolume/mysql-pv       10Gi       RWX            Retain           Bound    default/mysql-pvc       mysql                   3m21s
persistentvolume/nfs-pv         10Gi       RWO            Retain           Bound    default/nfs-pvc         nfs                     16m
persistentvolume/wordpress-pv   10Gi       RWX            Retain           Bound    default/wordpress-pvc   wordpress               70s

NAME                                  STATUS   VOLUME         CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/mysql-pvc       Bound    mysql-pv       10Gi       RWX            mysql          21s
persistentvolumeclaim/nfs-pvc         Bound    nfs-pv         10Gi       RWO            nfs            9m2s
persistentvolumeclaim/wordpress-pvc   Bound    wordpress-pv   10Gi       RWX            wordpress      8s
```

### 5.2.5 DeploymentとServiceの作成

#### MySQLのDeploymentとServiceの作成

コマンド
```
cd ../5-2-5-01
```

コマンド
```
cat mysql.yaml
```
コマンド結果
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql
  labels:
    app: mysql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
        - image: mysql:5.6
          name: mysql
          env:
            - name: MYSQL_ROOT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: mysql
                  key: password
          ports:
            - containerPort: 3306
              name: mysql
          volumeMounts:
            - name: mysql-local-storage
              mountPath: /var/lib/mysql
      volumes:
        - name: mysql-local-storage
          persistentVolumeClaim:
            claimName: mysql-pvc
```

コマンド
```
kubectl apply -f mysql.yaml
```
コマンド結果
```
deployment.apps/mysql created
```

コマンド
```
kubectl get pod -l app=mysql
```
コマンド結果
```
NAME                     READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-jw9kh   1/1     Running   0          13s
```

コマンド
```
cat mysql-service.yaml
```
コマンド結果
```
apiVersion: v1
kind: Service
metadata:
  name: mysql-service
  labels:
    app: mysql
spec:
  type: ClusterIP
  ports:
    - protocol: TCP
      port: 3306
      targetPort: 3306
  selector:
    app: mysql
```

コマンド
```
kubectl apply -f mysql-service.yaml
```
コマンド結果
```
service/mysql-service created
```

コマンド
```
kubectl get services
```
コマンド結果
```
NAME            TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)                      AGE
kubernetes      ClusterIP   10.48.0.1     <none>        443/TCP                      30m
mysql-service   ClusterIP   10.48.14.45   <none>        3306/TCP                     7s
nfs-service     ClusterIP   10.48.4.25    <none>        2049/TCP,20048/TCP,111/TCP   19m
```

#### WordPressのDeploymentとServiceの作成

コマンド
```
cat wordpress.yaml
```
コマンド結果
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  replicas: 1
  selector:
    matchLabels:
      app: wordpress
  template:
    metadata:
      labels:
        app: wordpress
    spec:
      containers:
        - image: wordpress:5.6.2
          name: wordpress
          env:
          - name: WORDPRESS_DB_HOST
            value: mysql-service
          - name: WORDPRESS_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: mysql
                key: password
          ports:
            - containerPort: 80
              name: wordpress
          volumeMounts:
            - name: wordpress-local-storage
              mountPath: /var/www/html
      volumes:
        - name: wordpress-local-storage
          persistentVolumeClaim:
            claimName: wordpress-pvc
```

コマンド
```
kubectl apply -f wordpress.yaml
```
コマンド結果
```
deployment.apps/wordpress created
```

コマンド
```
kubectl get pod -l app=wordpress
```
コマンド結果
```
NAME                         READY   STATUS    RESTARTS   AGE
wordpress-59fcd9d75f-n6jst   1/1     Running   0          29s
```

コマンド
```
cat wordpress-service.yaml
```
コマンド結果
```
apiVersion: v1
kind: Service
metadata:
  name: wordpress-service
  labels:
    app: wordpress
spec:
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
  selector:
    app: wordpress
```

コマンド
```
kubectl apply -f wordpress-service.yaml
```
コマンド結果
```
service/wordpress-service created
```

コマンド
```
kubectl get services
```
コマンド結果
```
NAME                TYPE           CLUSTER-IP    EXTERNAL-IP      PORT(S)                      AGE
kubernetes          ClusterIP      10.48.0.1     <none>           443/TCP                      34m
mysql-service       ClusterIP      10.48.14.45   <none>           3306/TCP                     4m3s
nfs-service         ClusterIP      10.48.4.25    <none>           2049/TCP,20048/TCP,111/TCP   23m
wordpress-service   LoadBalancer   10.48.6.192   35.xx.xx.xx      80:31367/TCP                 53s
```

### 5.2.6 アプリケーションのスケールアウト/スケールイン

#### kubectlコマンドによるPodの追加

コマンド
```
kubectl scale deployment wordpress --replicas 10
```
コマンド結果
```
deployment.apps/wordpress scaled
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                          READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-jw9kh        1/1     Running   0          8m14s
nfs-server-6f7fc97dfd-6kzmm   1/1     Running   0          26m
wordpress-59fcd9d75f-5q95f    1/1     Running   0          33s
wordpress-59fcd9d75f-84kv7    1/1     Running   0          33s
wordpress-59fcd9d75f-998n7    1/1     Running   0          33s
wordpress-59fcd9d75f-n6jst    1/1     Running   0          5m14s
wordpress-59fcd9d75f-sk6p6    1/1     Running   0          33s
wordpress-59fcd9d75f-tqzr8    1/1     Running   0          33s
wordpress-59fcd9d75f-x7gsd    1/1     Running   0          33s
wordpress-59fcd9d75f-xrwkn    1/1     Running   0          33s
wordpress-59fcd9d75f-xz5qb    1/1     Running   0          33s
wordpress-59fcd9d75f-z9mnw    1/1     Running   0          33s
```

#### マニフェストの編集によるPod数の変更

コマンド
```
cd ../5-2-6-01
```

コマンド
```
cat wordpress.yaml
```
コマンド結果
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  replicas: 5  #「replicas」を5に変更
  selector:
    matchLabels:
      app: wordpress
  template:
    metadata:
      labels:
        app: wordpress
    spec:
      containers:
        - image: wordpress:5.6.2
          name: wordpress
          env:
          - name: WORDPRESS_DB_HOST
            value: mysql-service
          - name: WORDPRESS_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: mysql
                key: password
          ports:
            - containerPort: 80
              name: wordpress
          volumeMounts:
            - name: wordpress-local-storage
              mountPath: /var/www/html
      volumes:
        - name: wordpress-local-storage
          persistentVolumeClaim:
            claimName: wordpress-pvc
```

コマンド
```
kubectl apply -f wordpress.yaml
```
コマンド結果
```
deployment.apps/wordpress configured
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                          READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-jw9kh        1/1     Running   0          12m
nfs-server-6f7fc97dfd-6kzmm   1/1     Running   0          31m
wordpress-67b55cd4bd-2k7hv    1/1     Running   0          31s
wordpress-67b55cd4bd-blmcx    1/1     Running   0          31s
wordpress-67b55cd4bd-bxglw    1/1     Running   0          14s
wordpress-67b55cd4bd-gnbnn    1/1     Running   0          13s
wordpress-67b55cd4bd-mn9hl    1/1     Running   0          31s
```

コマンド
```
kubectl edit deployment wordpress
```
コマンド結果
```
# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "2"
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"apps/v1","kind":"Deployment","metadata":{"annotations":{},"labels":{"app":"wordpress"},"name":"wordpress","namespace":"default"},"spec":{"replicas":5,"selector":{"matchLabels":{"app":"wordpress"}},"template":{"metadata":{"labels":{"app":"wordpress"}},"spec":{"containers":[{"env":[{"name":"WORDPRESS_DB_HOST","value":"mysql-service"},{"name":"WORDPRESS_DB_PASSWORD","valueFrom":{"secretKeyRef":{"key":"password","name":"mysql"}}}],"image":"wordpress","name":"wordpress","ports":[{"containerPort":80,"name":"wordpress"}],"volumeMounts":[{"mountPath":"/var/www/html","name":"wordpress-local-storage"}]}],"volumes":[{"name":"wordpress-local-storage","persistentVolumeClaim":{"claimName":"wordpress-pvc"}}]}}}}
  creationTimestamp: "2021-05-25T13:23:48Z"
  generation: 3
  labels:
    app: wordpress
  name: wordpress
  namespace: default
  resourceVersion: "16002"
  selfLink: /apis/apps/v1/namespaces/default/deployments/wordpress
  uid: 895c8cfb-59da-4a37-ac0e-a41e8dc67136
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: wordpress
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: wordpress
    spec:
      containers:
      - env:
        - name: WORDPRESS_DB_HOST
          value: mysql-service
        - name: WORDPRESS_DB_PASSWORD
          valueFrom:
            secretKeyRef:
              key: password
              name: mysql
        image: wordpress
        imagePullPolicy: IfNotPresent
        name: wordpress
        ports:
        - containerPort: 80
          name: wordpress
          protocol: TCP
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /var/www/html
          name: wordpress-local-storage
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
      volumes:
      - name: wordpress-local-storage
        persistentVolumeClaim:
          claimName: wordpress-pvc
status:
  availableReplicas: 5
  conditions:
  - lastTransitionTime: "2021-05-25T13:28:53Z"
    lastUpdateTime: "2021-05-25T13:28:53Z"
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  - lastTransitionTime: "2021-05-25T13:23:48Z"
    lastUpdateTime: "2021-05-25T13:33:06Z"
    message: ReplicaSet "wordpress-67b55cd4bd" has successfully progressed.
    reason: NewReplicaSetAvailable
    status: "True"
    type: Progressing
  observedGeneration: 3
  readyReplicas: 5
  replicas: 5
  updatedReplicas: 5

deployment.apps/wordpress edited
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                          READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-jw9kh        1/1     Running   0          34m
nfs-server-6f7fc97dfd-6kzmm   1/1     Running   0          53m
wordpress-67b55cd4bd-blmcx    1/1     Running   0          22m
```

### 5.2.7 アプリケーションの削除

コマンド
```
cd ../
```

コマンド
```
kubectl delete -f 5-2-5-01
```
コマンド結果
```
service "mysql-service" deleted
deployment.apps "mysql" deleted
service "wordpress-service" deleted
deployment.apps "wordpress" deleted
```

コマンド
```
kubectl delete -f 5-2-4-01
```
コマンド結果
```
persistentvolume "mysql-pv" deleted
persistentvolumeclaim "mysql-pvc" deleted
persistentvolume "wordpress-pv" deleted
persistentvolumeclaim "wordpress-pvc" deleted
```

コマンド
```
kubectl delete -f 5-2-2-01
```
コマンド結果
```
deployment.apps "nfs-server" deleted
service "nfs-service" deleted
```

コマンド
```
kubectl delete secret mysql
```
コマンド結果
```
secret "mysql" deleted
```

コマンド
```
kubectl get deployments
```
コマンド結果
```
No resources found in default namespace.
```

コマンド
```
kubectl get services
```
コマンド結果
```
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.48.0.1    <none>        443/TCP   73m
```

コマンド
```
kubectl get persistentvolumes,persistentvolumeclaims
```
コマンド結果
```
No resources found
```

コマンド
```
kubectl get secrets
```
コマンド結果
```
NAME                  TYPE                                  DATA   AGE
default-token-lm5kz   kubernetes.io/service-account-token   3      77m
```

コマンド
```
gcloud compute disks delete nfs-disk --zone=asia-northeast1-a
```
コマンド結果
```
The following disks will be deleted:
 - [nfs-disk] in [asia-northeast1-a]

Do you want to continue (Y/n)?   Y

Deleted [https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-a/disks/nfs-disk].
```

## 5.3 マニフェストの管理

### 5.3.2 Chartの作成とアプリケーションのインストール

#### Helm クライアントのインストール

コマンド
```
cd 5-3-2-01
```

コマンド
```
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
```

コマンド
```
chmod 700 get_helm.sh
```

コマンド
```
./get_helm.sh
```
コマンド結果
```
Helm v3.5.4 is available. Changing from version v3.5.0.
Downloading https://get.helm.sh/helm-v3.5.4-linux-amd64.tar.gz
Verifying checksum... Done.
Preparing to install helm into /usr/local/bin
helm installed into /usr/local/bin/helm
```

コマンド
```
helm version
```
コマンド結果
```
version.BuildInfo{Version:"v3.5.4", GitCommit:"1b5edb69df3d3a08df77c9902dc17af864ff05d1", GitTreeState:"clean", GoVersion:"go1.15.11"}
```

#### Chartの雛形の作成

コマンド
```
helm create wordpress
```
コマンド結果
```
Creating wordpress
```

コマンド
```
ls wordpress
```
コマンド結果
```
charts  Chart.yaml  templates  values.yaml
```

コマンド
```
ls wordpress/templates/
```
コマンド結果
```
deployment.yaml  _helpers.tpl  hpa.yaml  ingress.yaml  NOTES.txt  serviceaccount.yaml  service.yaml  tests
```

コマンド
```
ls wordpress/templates/tests
```
コマンド結果
```
test-connection.yaml
```

#### WordPressの独自Chartの作成

コマンド
```
rm -rf wordpress/templates/*
```

コマンド
```
ls wordpress/templates/
```

#### Secret マニフェストテンプレートの作成

コマンド
```
cp -p helm-yaml/mysql-secret.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/mysql-secret.yaml
```
コマンド結果
```
apiVersion: v1
kind: Secret
metadata:
  name: mysql
type: Opaque
data:
  password: {{ .Values.mysql_secret.password }}
```

#### PersistentVolume、PersistentVolumeClaimの作成

コマンド
```
cp -p helm-yaml/mysql-pv.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/mysql-pv.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: {{ .Values.mysql_pv.name }}
  labels:
    app: wordpress
    tire: mysql
    type: local
spec:
  storageClassName: {{ .Values.mysql_pv.storageClassName }}
  capacity:
    storage: {{ .Values.mysql_pv.storage }}
  accessModes:
    - {{ .Values.mysql_pv.accessModes }}
  hostPath:
    path: {{ .Values.mysql_pv.hostPath }}
```

コマンド
```
cp -p helm-yaml/wordpress-pv.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/wordpress-pv.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolume
metadata:
  name: {{ .Values.wordpress_pv.name }}
  labels:
    app: wordpress
    tire: wordpress
    type: local
spec:
  storageClassName: {{ .Values.wordpress_pv.storageClassName }}
  capacity:
    storage: {{ .Values.wordpress_pv.storage }}
  accessModes:
    - {{ .Values.wordpress_pv.accessModes }}
  hostPath:
    path: {{ .Values.wordpress_pv.hostPath }}
```

コマンド
```
cp -p helm-yaml/mysql-pvc.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/mysql-pvc.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{ .Values.mysql_pvc.name }}
  labels:
    app: wordpress
    tier: mysql
spec:
  storageClassName: {{ .Values.mysql_pvc.storageClassName }}
  accessModes:
    - {{ .Values.mysql_pvc.accessModes }}
  resources:
    requests:
      storage: {{ .Values.mysql_pvc.storage }}
```

コマンド
```
cp -p helm-yaml/wordpress-pvc.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/wordpress-pvc.yaml
```
コマンド結果
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: {{ .Values.wordpress_pvc.name }}
  labels:
    app: wordpress
    tier: wordpress
spec:
  storageClassName: {{ .Values.wordpress_pvc.storageClassName }}
  accessModes:
    - {{ .Values.mysql_pvc.accessModes }}
  resources:
    requests:
      storage: {{ .Values.wordpress_pvc.storage }}
```

#### DeploymentとServiceの作成

コマンド
```
cp -p helm-yaml/mysql.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/mysql.yaml
```
コマンド結果
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Values.mysql.name }}
  labels:
    app: mysql
spec:
  replicas: {{ .Values.mysql.replicas }}
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
        - image: {{ .Values.mysql.image }}
          name: mysql
          env:
            - name: MYSQL_ROOT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: mysql
                  key: password
          ports:
            - containerPort: {{ .Values.mysql.containerPort }}
              name: mysql
          volumeMounts:
            - name: mysql-local-storage
              mountPath: {{ .Values.mysql.mountPath }}
      volumes:
        - name: mysql-local-storage
          persistentVolumeClaim:
            claimName: {{ .Values.mysql.claimName }}
```

コマンド
```
cp -p helm-yaml/mysql-service.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/mysql-service.yaml
```
コマンド結果
```
apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.mysql_service.name }}
  labels:
    app: mysql
spec:
  type: {{ .Values.mysql_service.type }}
  ports:
    - protocol: {{ .Values.mysql_service.protocol }}
      port: {{ .Values.mysql_service.port }}
      targetPort: {{ .Values.mysql_service.targetPort }}
  selector:
    app: mysql
```

コマンド
```
cp -p helm-yaml/wordpress.yaml wordpress/templates
```

コマンド
```
cat helm-yaml/wordpress.yaml
```
コマンド結果
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Values.wordpress.name }}
  labels:
    app: wordpress
spec:
  replicas: {{ .Values.wordpress.replicas }}
  selector:
    matchLabels:
      app: wordpress
  template:
    metadata:
      labels:
        app: wordpress
    spec:
      containers:
        - image: {{ .Values.wordpress.image }}
          name: wordpress
          env:
          - name: WORDPRESS_DB_HOST
            value: {{ .Values.wordpress.value }}
          - name: WORDPRESS_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: mysql
                key: password
          ports:
            - containerPort: {{ .Values.wordpress.containerPort }}
              name: wordpress
          volumeMounts:
            - name: wordpress-local-storage
              mountPath: {{ .Values.wordpress.mountPath }}
      volumes:
        - name: wordpress-local-storage
          persistentVolumeClaim:
            claimName: {{ .Values.wordpress.claimName }}
```

コマンド
```
cp -p helm-yaml/wordpress-service.yaml wordpress/templates
```

コマンド
```
cat wordpress/templates/wordpress-service.yaml
```
コマンド結果
```
apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.wordpress_service.name }}
  labels:
    app: wordpress
spec:
  type: {{ .Values.wordpress_service.type }}
  ports:
    - protocol: {{ .Values.wordpress_service.protocol }}
      port: {{ .Values.wordpress_service.port }}
      targetPort: {{ .Values.wordpress_service.targetPort }}
  selector:
    app: wordpress
```

#### values.yaml の作成

コマンド
```
cp -p helm-yaml/values.yaml wordpress
```

コマンド
```
cat wordpress/values.yaml
```
コマンド結果
```
#mysql-secret
mysql_secret:
  password: bXlzcWxwQHNzdzBk

#mysql-pv
mysql_pv:
  name: mysql-pv
  storageClassName: mysql
  storage: 10Gi
  accessModes: ReadWriteOnce
  hostPath: /tmp/data/mysql

#mysql-pvc
mysql_pvc:
  name: mysql-pvc
  storageClassName: mysql
  accessModes: ReadWriteOnce
  storage: 5Gi

#wordpress-pv
wordpress_pv:
  name: wordpress-pv
  storageClassName: wordpress
  storage: 10Gi
  accessModes: ReadWriteOnce
  hostPath: /tmp/data/wordpress

#wordpress-pvc
wordpress_pvc:
  name: wordpress-pvc
  storageClassName: wordpress
  accessModes: ReadWriteOnce
  storage: 5Gi

#mysql
mysql:
  name: mysql
  replicas: 1
  image: mysql:5.6
  containerPort: 3306
  mountPath: /var/lib/mysql
  claimName: mysql-pvc

#mysql-service
mysql_service:
  name: mysql-service
  type: ClusterIP
  protocol: TCP
  port: 3306
  targetPort: 3306

#wordpress
wordpress:
  name: wordpress
  replicas: 1
  image: wordpress:5.6.2
  value: mysql-service
  containerPort: 80
  mountPath: /var/www/html
  claimName: wordpress-pvc

#wordpress-service
wordpress_service:
  name: wordpress-service
  type: LoadBalancer
  protocol: TCP
  port: 80
  targetPort: 80
```

#### Chartのデバッグ

コマンド
```
helm install wordpress --debug --dry-run wordpress
```
コマンド結果
```
install.go:173: [debug] Original chart version: ""
install.go:190: [debug] CHART PATH: /home/iyutaka2020/docker-development-environment-construction-basic/Chapter05/5-3-2-01/wordpress

NAME: wordpress
LAST DEPLOYED: Sun May 23 07:22:58 2021
NAMESPACE: default
STATUS: pending-install
REVISION: 1
TEST SUITE: None
USER-SUPPLIED VALUES:
{}

COMPUTED VALUES:
mysql:
  claimName: mysql-pvc
  containerPort: 3306
  image: mysql:5.6
  mountPath: /var/lib/mysql
  name: mysql
  replicas: 1
mysql_pv:
  accessModes: ReadWriteOnce
  hostPath: /tmp/data/mysql
  name: mysql-pv
  storage: 10Gi
  storageClassName: mysql
mysql_pvc:
  accessModes: ReadWriteOnce
  name: mysql-pvc
  storage: 5Gi
  storageClassName: mysql
mysql_secret:
  password: bXlzcWxwQHNzdzBk
mysql_service:
  name: mysql-service
  port: 3306
  protocol: TCP
  targetPort: 3306
  type: ClusterIP
wordpress:
  claimName: wordpress-pvc
  containerPort: 80
  image: wordpress
  mountPath: /var/www/html
  name: wordpress
  replicas: 1
  value: mysql-service
wordpress_pv:
  accessModes: ReadWriteOnce
  hostPath: /tmp/data/wordpress
  name: wordpress-pv
  storage: 10Gi
  storageClassName: wordpress
wordpress_pvc:
  accessModes: ReadWriteOnce
  name: wordpress-pvc
  storage: 5Gi
  storageClassName: wordpress
wordpress_service:
  name: wordpress-service
  port: 80
  protocol: TCP
  targetPort: 80
  type: LoadBalancer

HOOKS:
MANIFEST:
---
# Source: wordpress/templates/mysql-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysql
type: Opaque
data:
  password: bXlzcWxwQHNzdzBk
---
# Source: wordpress/templates/mysql-pv.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: mysql-pv
  labels:
    app: wordpress
    tire: mysql
    type: local
spec:
  storageClassName: mysql
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /tmp/data/mysql
---
# Source: wordpress/templates/wordpress-pv.yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: wordpress-pv
  labels:
    app: wordpress
    tire: wordpress
    type: local
spec:
  storageClassName: wordpress
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: /tmp/data/wordpress
---
# Source: wordpress/templates/mysql-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mysql-pvc
  labels:
    app: wordpress
    tier: mysql
spec:
  storageClassName: mysql
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
---
# Source: wordpress/templates/wordpress-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: wordpress-pvc
  labels:
    app: wordpress
    tier: wordpress
spec:
  storageClassName: wordpress
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 5Gi
---
# Source: wordpress/templates/mysql-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql-service
  labels:
    app: mysql
spec:
  type: ClusterIP
  ports:
    - port: 3305
      targetPort: 3306
      protocol: TCP
  selector:
    app: mysql
---
# Source: wordpress/templates/wordpress-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: wordpress-service
  labels:
    app: wordpress
spec:
  type: LoadBalancer
  ports:
    - port: 80
      targetPort: 80
      protocol: TCP
  selector:
    app: wordpress
---
# Source: wordpress/templates/mysql.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mysql
  labels:
    app: mysql
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
        - image: mysql:5.6
          name: mysql
          env:
            - name: MYSQL_ROOT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: mysql
                  key: password
          ports:
            - containerPort: 3306
              name: mysql
          volumeMounts:
            - name: mysql-local-storage
              mountPath: /var/lib/mysql
      volumes:
        - name: mysql-local-storage
          persistentVolumeClaim:
            claimName: mysql-pvc
---
# Source: wordpress/templates/wordpress.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  replicas: 1
  selector:
    matchLabels:
      app: wordpress
  template:
    metadata:
      labels:
        app: wordpress
    spec:
      containers:
        - image: wordpress:5.6.2
          name: wordpress
          env:
          - name: WORDPRESS_DB_HOST
            value: mysql-service
          - name: WORDPRESS_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: mysql
                key: password
          ports:
            - containerPort: 80
              name: wordpress
          volumeMounts:
            - name: wordpress-local-storage
              mountPath: /var/www/html
      volumes:
        - name: wordpress-local-storage
          persistentVolumeClaim:
            claimName: wordpress-pvc
```

コマンド
```
helm lint wordpress
```
コマンド結果
```
==> Linting wordpress
[INFO] Chart.yaml: icon is recommended

1 chart(s) linted, 0 chart(s) failed
```

#### Chartのパッケージ化

コマンド
```
helm package wordpress
```
コマンド結果
```
Successfully packaged chart and saved it to: /home/iyutaka2020/docker-development-environment-construction-basic/Chapter05/5-3-2-01/wordpress-0.1.0.tgz
```

コマンド
```
helm repo index .
```

コマンド
```
cat index.yaml
```
コマンド結果
```
apiVersion: v1
entries:
  wordpress:
  - apiVersion: v2
    appVersion: 1.16.0
    created: "2021-05-25T14:22:24.438645627Z"
    description: A Helm chart for Kubernetes
    digest: a8be872cec364cd4c23a22db237ea39eeebccbaa924f659769e33129e3648dd9
    name: wordpress
    type: application
    urls:
    - wordpress-0.1.0.tgz
    version: 0.1.0
generated: "2021-05-25T14:22:24.43791071Z"
```

#### WordPressのインストール

コマンド
```
helm install wordpress ./wordpress
```
コマンド結果
```
NAME: wordpress
LAST DEPLOYED: Tue May 25 14:24:13 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                         READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-cnvwz       1/1     Running   0          16s
wordpress-59fcd9d75f-6qbc9   1/1     Running   0          16s
```

コマンド
```
kubectl get persistentvolumes,persistentvolumeclaims
```
コマンド結果
```
NAME                            CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
persistentvolume/mysql-pv       10Gi       RWO            Retain           Bound    default/mysql-pvc       mysql                   48s
persistentvolume/wordpress-pv   10Gi       RWO            Retain           Bound    default/wordpress-pvc   wordpress               48s

NAME                                  STATUS   VOLUME         CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/mysql-pvc       Bound    mysql-pv       10Gi       RWO            mysql          48s
persistentvolumeclaim/wordpress-pvc   Bound    wordpress-pv   10Gi       RWO            wordpress      48s
```

コマンド
```
kubectl get services
```
コマンド結果
```
NAME                TYPE           CLUSTER-IP     EXTERNAL-IP      PORT(S)        AGE
kubernetes          ClusterIP      10.48.0.1      <none>           443/TCP        95m
mysql-service       ClusterIP      10.48.5.133    <none>           3306/TCP       2m48s
wordpress-service   LoadBalancer   10.48.15.236   35.243.121.176   80:30690/TCP   2m48s
```

#### WordPressのアンインストール

コマンド
```
helm uninstall wordpress
```
コマンド結果
```
release "wordpress" uninstalled
```

コマンド
```
kubectl get pods
```
コマンド結果
```
No resources found in default namespace.
```

コマンド
```
kubectl get persistentvolumes,persistentvolumeclaims
```
コマンド結果
```
No resources found
```

コマンド
```
kubectl get services
```
コマンド結果
```
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.48.0.1    <none>        443/TCP   96m
```

#### 5.3.3 公開Chartの利用

### リポジトリの操作

コマンド
```
helm repo add bitnami https://charts.bitnami.com/bitnami
```
コマンド結果
```
"bitnami" has been added to your repositories
```

コマンド
```
helm repo list
```
コマンド結果
```
NAME            URL
bitnami         https://charts.bitnami.com/bitnami
```

コマンド
```
helm repo update
```
コマンド結果
```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "bitnami" chart repository
Update Complete. ⎈Happy Helming!⎈
```

コマンド
```
helm search repo wordpress
```
コマンド結果
```
NAME                    CHART VERSION   APP VERSION     DESCRIPTION
bitnami/wordpress       11.0.9          5.7.2           Web publishing platform for building blogs and ...
```

コマンド
```
helm search hub wordpress
```
コマンド結果
```
URL                                                     CHART VERSION   APP VERSION     DESCRIPTION
https://artifacthub.io/packages/helm/bitnami/wo...      11.0.9          5.7.2           Web publishing platform for building blogs and ...
https://artifacthub.io/packages/helm/groundhog2...      0.3.7           5.7.2-apache    A Helm chart for Wordpress on Kubernetes
https://artifacthub.io/packages/helm/seccurecod...      2.6.1           4.0             Insecure & Outdated Wordpress Instance: Never e...
https://artifacthub.io/packages/helm/wordpressm...      1.0.0                           This is the Helm Chart that creates the Wordpre...
https://artifacthub.io/packages/helm/presslabs/...      0.11.0-alpha.1  0.11.0-alpha.1  Presslabs WordPress Operator Helm Chart
https://artifacthub.io/packages/helm/presslabs/...      0.11.0-rc.2     v0.11.0-rc.2    A Helm chart for deploying a WordPress site on ...
https://artifacthub.io/packages/helm/phntom/bin...      0.0.3           0.0.3           www.binaryvision.co.il static wordpress
https://artifacthub.io/packages/helm/gh-shessel...      1.0.22          5.7.2           Web publishing platform for building blogs and ...
https://artifacthub.io/packages/helm/wordpress/...      0.2.0           1.1.0           Wordpress for Kubernetes
https://artifacthub.io/packages/helm/seccurecod...      2.6.1           3.8.17          A Helm chart for the WordPress security scanner...
https://artifacthub.io/packages/helm/presslabs/...      0.11.0-rc.2     v0.11.0-rc.2    Open-Source WordPress Infrastructure on Kubernetes
https://artifacthub.io/packages/helm/wordpressm...      0.1.0           1.1
```

コマンド
```
helm show values bitnami/wordpress --version 10.6.10
```
コマンド結果
```
## Global Docker image parameters
## Please, note that this will override the image parameters, including dependencies, configured to use the global value
## Current available global Docker image parameters: imageRegistry and imagePullSecrets
##
# global:
#   imageRegistry: myRegistryName
#   imagePullSecrets:
#     - myRegistryKeySecretName
#   storageClass: myStorageClass

## Bitnami WordPress image version
## ref: https://hub.docker.com/r/bitnami/wordpress/tags/
##
image:
  registry: docker.io
  repository: bitnami/wordpress
  tag: 5.6.2-debian-10-r8
  ## Specify a imagePullPolicy
  ## Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent'
  ## ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images
  ##
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ##
  # pullSecrets:
  #   - myRegistryKeySecretName
  ## Set to true if you would like to see extra information on logs
  ##
  debug: false

## Force target Kubernetes version (using Helm capabilites if not set)
##
kubeVersion:

## String to partially override aspnet-core.fullname template (will maintain the release name)
##
# nameOverride:

## String to fully override aspnet-core.fullname template
##
# fullnameOverride:

## Add labels to all the deployed resources
##
commonLabels: {}

## Add annotations to all the deployed resources
##
commonAnnotations: {}

## Kubernetes Cluster Domain
##
clusterDomain: cluster.local

## Deployment pod host aliases
## https://kubernetes.io/docs/concepts/services-networking/add-entries-to-pod-etc-hosts-with-host-aliases/
##
hostAliases:
  # Necessary for apache-exporter to work
  - ip: "127.0.0.1"
    hostnames:
      - "status.localhost"

## Extra objects to deploy (value evaluated as a template)
##
extraDeploy: []

## Use a service account for the WordPress pod
##
serviceAccountName: default

## User of the application
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
wordpressUsername: user

## Application password
## Defaults to a random 10-character alphanumeric string if not set
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
# wordpressPassword:

## Use existing secret (does not create the WordPress Secret object)
## Must contain key `wordpress-secret`
## NOTE: When it's set, the `wordpressPassword` parameter is ignored
##
# existingSecret:

## Admin email
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
wordpressEmail: user@example.com

## First name
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
wordpressFirstName: FirstName

## Last name
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
wordpressLastName: LastName

## Blog name
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
wordpressBlogName: User's Blog!

## Table prefix
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
wordpressTablePrefix: wp_

## Scheme to generate application URLs
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
wordpressScheme: http

## Skip wizard installation (only if you use an external database that already contains WordPress data)
## ref: https://github.com/bitnami/bitnami-docker-wordpress#connect-wordpress-docker-container-to-an-existing-database
##
wordpressSkipInstall: false

## Add extra content to the default configuration file
##
wordpressExtraConfigContent: ""

## Set to `false` to allow the container to be started with blank passwords
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
allowEmptyPassword: true

## Set Apache allowOverride to None
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables
##
allowOverrideNone: false

## Persist the custom changes of the htaccess. It depends on the value of
## `.Values.allowOverrideNone`, when `yes` it will persist `/opt/bitnami/wordpress/wordpress-htaccess.conf`
## if `no` it will persist `/opt/bitnami/wordpress/.htaccess`
##
htaccessPersistenceEnabled: false

## ConfigMap with custom wordpress-htaccess.conf file (requires allowOverrideNone to true)
##
customHTAccessCM:

## Command and args for running the container (set to default if not set). Use array form
##
command: []
args: []

## Set up update strategy for wordpress installation. Set to Recreate if you use persistent volume that cannot be mounted by more than one pods to makesure the pods is destroyed first.
## ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy
## Example:
## updateStrategy:
##  type: RollingUpdate
##  rollingUpdate:
##    maxSurge: 25%
##    maxUnavailable: 25%
##
updateStrategy:
  type: RollingUpdate

## Use an alternate scheduler, e.g. "stork".
## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
##
# schedulerName:

## SMTP mail delivery configuration
## ref: https://github.com/bitnami/bitnami-docker-wordpress/#smtp-configuration
##
# smtpHost:
# smtpPort:
# smtpUser:
# smtpPassword:
# smtpProtocol:

## Use an existing secret for the SMTP Password
## Can be the same secret as existingSecret
## Must contain key `smtp-password`
## NOTE: When it's set, the `smtpPassword` parameter is ignored
##
# smtpExistingSecret:

## Number of replicas (requires ReadWriteMany PVC support)
##
replicaCount: 1

## An array to add extra env vars
## Example:
## extraEnvVars:
##   - name: FOO
##     value: "bar"
##
extraEnvVars: []

## ConfigMap with extra environment variables
##
extraEnvVarsCM:

## Secret with extra environment variables
##
extraEnvVarsSecret:

## Extra volumes to add to the deployment
##
extraVolumes: []

## Extra volume mounts to add to the container
##
extraVolumeMounts: []

## Add sidecars to the pod.
## Example:
## sidecars:
##   - name: your-image-name
##     image: your-image
##     imagePullPolicy: Always
##     ports:
##       - name: portname
##         containerPort: 1234
##
sidecars: {}

## Add init containers to the pod.
## ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
## Example:
## initContainers:
##  - name: your-image-name
##    image: your-image
##    imagePullPolicy: Always
##    command: ['sh', '-c', 'copy themes and plugins from git and push to /bitnami/wordpress/wp-content. Should work with extraVolumeMounts and extraVolumes']
##
initContainers: {}

## Pod Labels
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
##
podLabels: {}

## Pod annotations
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
##
podAnnotations: {}

## Pod affinity preset
## ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
## Allowed values: soft, hard
##
podAffinityPreset: ""

## Pod anti-affinity preset
## Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
## Allowed values: soft, hard
##
podAntiAffinityPreset: soft

## Node affinity preset
## Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
## Allowed values: soft, hard
##
nodeAffinityPreset:
  ## Node affinity type
  ## Allowed values: soft, hard
  ##
  type: ""
  ## Node label key to match
  ## E.g.
  ## key: "kubernetes.io/e2e-az-name"
  ##
  key: ""
  ## Node label values to match
  ## E.g.
  ## values:
  ##   - e2e-az1
  ##   - e2e-az2
  ##
  values: []

## Affinity for pod assignment
## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
## Note: podAffinityPreset, podAntiAffinityPreset, and  nodeAffinityPreset will be ignored when it's set
##
affinity: {}

## Node labels for pod assignment. Evaluated as a template.
## ref: https://kubernetes.io/docs/user-guide/node-selection/
##
nodeSelector: {}

## Tolerations for pod assignment. Evaluated as a template.
## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
##
tolerations: {}

## WordPress containers' resource requests and limits
## ref: http://kubernetes.io/docs/user-guide/compute-resources/
##
resources:
  limits: {}
  requests:
    memory: 512Mi
    cpu: 300m

## Configure Pods Security Context
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod
##
podSecurityContext:
  enabled: true
  fsGroup: 1001

## Configure Container Security Context (only main container)
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container
##
containerSecurityContext:
  enabled: true
  runAsUser: 1001

## WordPress containers' liveness and readiness probes.
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/#configure-probes
##
livenessProbe:
  enabled: true
  httpGet:
    path: /wp-admin/install.php
    port: http
    scheme: HTTP
    ## If using an HTTPS-terminating load-balancer, the probes may need to behave
    ## like the balancer to prevent HTTP 302 responses. According to the Kubernetes
    ## docs, 302 should be considered "successful", but this issue on GitHub
    ## (https://github.com/kubernetes/kubernetes/issues/47893) shows that it isn't.
    ##
    ## httpHeaders:
    ## - name: X-Forwarded-Proto
    ##   value: https
    ##
    httpHeaders: []
  initialDelaySeconds: 120
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 6
  successThreshold: 1
readinessProbe:
  enabled: true
  httpGet:
    path: /wp-login.php
    port: http
    scheme: HTTP
    ## If using an HTTPS-terminating load-balancer, the probes may need to behave
    ## like the balancer to prevent HTTP 302 responses. According to the Kubernetes
    ## docs, 302 should be considered "successful", but this issue on GitHub
    ## (https://github.com/kubernetes/kubernetes/issues/47893) shows that it isn't.
    ##
    ## httpHeaders:
    ## - name: X-Forwarded-Proto
    ##   value: https
    ##
    httpHeaders: []
  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 6
  successThreshold: 1

## Custom liveness and readiness probes, it overrides the default one (evaluated as a template)
##
customLivenessProbe: {}
customReadinessProbe: {}

## Container ports
##
containerPorts:
  http: 8080
  https: 8443

## Kubernetes configuration
## For minikube, set this to NodePort, elsewhere use LoadBalancer or ClusterIP
##
service:
  type: LoadBalancer
  ## HTTP Port
  ##
  port: 80
  ## HTTPS Port
  ##
  httpsPort: 443
  ## HTTPS Target Port
  ## defaults to https unless overridden to the specified port.
  ## if you want the target port to be "http" or "80" you can specify that here.
  ##
  httpsTargetPort: https
  ## Node Ports to expose
  ## nodePorts:
  ##   http: <to set explicitly, choose port between 30000-32767>
  ##   https: <to set explicitly, choose port between 30000-32767>
  ##
  nodePorts:
    http: ""
    https: ""
  ## Service clusterIP.
  ##
  # clusterIP: None
  ## loadBalancerIP for the SuiteCRM Service (optional, cloud specific)
  ## ref: http://kubernetes.io/docs/user-guide/services/#type-loadbalancer
  ##
  # loadBalancerIP:
  ## Load Balancer sources
  ## https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/#restrict-access-for-loadbalancer-service
  ## Example:
  ## loadBalancerSourceRanges:
  ##   - 10.10.10.0/24
  ##
  loadBalancerSourceRanges: []
  ## Enable client source IP preservation
  ## ref http://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
  ##
  externalTrafficPolicy: Cluster
  ## Provide any additional annotations which may be required (evaluated as a template).
  ##
  annotations: {}
  ## Extra ports to expose (normally used with the `sidecar` value)
  ##
  # extraPorts:

## Configure the ingress resource that allows you to access the
## WordPress installation. Set up the URL
## ref: http://kubernetes.io/docs/user-guide/ingress/
##
ingress:
  ## Set to true to enable ingress record generation
  ##
  enabled: false

  ## Set this to true in order to add the corresponding annotations for cert-manager
  ##
  certManager: false

  ## Ingress Path type
  ##
  pathType: ImplementationSpecific

  ## Override API Version (automatically detected if not set)
  ##
  apiVersion:

  ## When the ingress is enabled, a host pointing to this will be created
  ##
  hostname: wordpress.local

  ## The Path to WordPress. You may need to set this to '/*' in order to use this
  ## with ALB ingress controllers.
  ##
  path: /

  ## Ingress annotations done as key:value pairs
  ## For a full list of possible ingress annotations, please see
  ## ref: https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md
  ##
  ## If certManager is set to true, annotation kubernetes.io/tls-acme: "true" will automatically be set
  ##
  annotations: {}

  ## Enable TLS configuration for the hostname defined at ingress.hostname parameter
  ## TLS certificates will be retrieved from a TLS secret with name: {{- printf "%s-tls" .Values.ingress.hostname }}
  ## You can use the ingress.secrets parameter to create this TLS secret or relay on cert-manager to create it
  ##
  tls: false

  ## The list of additional hostnames to be covered with this ingress record.
  ## Most likely the hostname above will be enough, but in the event more hosts are needed, this is an array
  ## extraHosts:
  ## - name: wordpress.local
  ##   path: /
  ##

  ## Any additional arbitrary paths that may need to be added to the ingress under the main host.
  ## For example: The ALB ingress controller requires a special rule for handling SSL redirection.
  ## extraPaths:
  ## - path: /*
  ##   backend:
  ##     serviceName: ssl-redirect
  ##     servicePort: use-annotation
  ##

  ## The tls configuration for additional hostnames to be covered with this ingress record.
  ## see: https://kubernetes.io/docs/concepts/services-networking/ingress/#tls
  ## extraTls:
  ## - hosts:
  ##     - wordpress.local
  ##   secretName: wordpress.local-tls
  ##

  ## If you're providing your own certificates, please use this to add the certificates as secrets
  ## key and certificate should start with -----BEGIN CERTIFICATE----- or
  ## -----BEGIN RSA PRIVATE KEY-----
  ##
  ## name should line up with a tlsSecret set further up
  ## If you're using cert-manager, this is unneeded, as it will create the secret for you if it is not set
  ##
  ## It is also possible to create and manage the certificates outside of this helm chart
  ## Please see README.md for more information
  ##
  secrets: []
  ## - name: wordpress.local-tls
  ##   key:
  ##   certificate:
  ##

## Enable persistence using Persistent Volume Claims
## ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
##
persistence:
  enabled: true
  ## wordpress data Persistent Volume Storage Class
  ## If defined, storageClassName: <storageClass>
  ## If set to "-", storageClassName: "", which disables dynamic provisioning
  ## If undefined (the default) or set to null, no storageClassName spec is
  ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
  ##   GKE, AWS & OpenStack)
  ##
  # storageClass: "-"
  ##
  ## If you want to reuse an existing claim, you can pass the name of the PVC using
  ## the existingClaim variable
  # existingClaim: your-claim
  accessMode: ReadWriteOnce
  size: 10Gi
  ## Custom dataSource
  ##
  dataSource: {}

## Wordpress Pod Disruption Budget configuration
## ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/
##
pdb:
  create: false
  ## Min number of pods that must still be available after the eviction
  ##
  minAvailable: 1
  ## Max number of pods that can be unavailable after the eviction
  ##
  # maxUnavailable: 1

## Wordpress Autoscaling configuration
##
autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 11
  # targetCPU: 50
  # targetMemory: 50

## Prometheus Exporter / Metrics
##
metrics:
  enabled: false
  image:
    registry: docker.io
    repository: bitnami/apache-exporter
    tag: 0.8.0-debian-10-r313
    pullPolicy: IfNotPresent
    ## Optionally specify an array of imagePullSecrets.
    ## Secrets must be manually created in the namespace.
    ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
    ##
    # pullSecrets:
    #   - myRegistryKeySecretName

  ## Prometheus expoter service parameters
  ##
  service:
    ## Metrics port
    ##
    port: 9117
    ## Annotations for the Prometheus exporter service
    ##
    annotations:
      prometheus.io/scrape: "true"
      prometheus.io/port: "{{ .Values.metrics.service.port }}"

  ## Metrics exporter containers' resource requests and limits
  ## ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##
  resources:
    limits: {}
    requests: {}

  ## Prometheus Service Monitor
  ## ref: https://github.com/coreos/prometheus-operator
  ##      https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#endpoint
  ##
  serviceMonitor:
    ## If the operator is installed in your cluster, set to true to create a Service Monitor Entry
    ##
    enabled: false
    ## Specify the namespace in which the serviceMonitor resource will be created
    # namespace: ""
    ## Specify the interval at which metrics should be scraped
    ##
    interval: 30s
    ## Specify the timeout after which the scrape is ended
    # scrapeTimeout: 30s
    ## Specify Metric Relabellings to add to the scrape endpoint
    # relabellings:
    ## Specify honorLabels parameter to add the scrape endpoint
    ##
    honorLabels: false
    ## Used to pass Labels that are used by the Prometheus installed in your cluster to select Service Monitors to work with
    ## ref: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
    ##
    additionalLabels: {}

##
## MariaDB chart configuration
##
## https://github.com/bitnami/charts/blob/master/bitnami/mariadb/values.yaml
##
mariadb:
  ## Whether to deploy a mariadb server to satisfy the applications database requirements. To use an external database set this to false and configure the externalDatabase parameters
  ##
  enabled: true
  ## MariaDB architecture. Allowed values: standalone or replication
  ##
  architecture: standalone
  ## MariaDB Authentication parameters
  ##
  auth:
    ## MariaDB root password
    ## ref: https://github.com/bitnami/bitnami-docker-mariadb#setting-the-root-password-on-first-run
    ##
    rootPassword: ""
    ## MariaDB custom user and database
    ## ref: https://github.com/bitnami/bitnami-docker-mariadb/blob/master/README.md#creating-a-database-on-first-run
    ## ref: https://github.com/bitnami/bitnami-docker-mariadb/blob/master/README.md#creating-a-database-user-on-first-run
    ##
    database: bitnami_wordpress
    username: bn_wordpress
    password: ""
  primary:
    ## Enable persistence using Persistent Volume Claims
    ## ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
    ##
    persistence:
      enabled: true
      ## mariadb data Persistent Volume Storage Class
      ## If defined, storageClassName: <storageClass>
      ## If set to "-", storageClassName: "", which disables dynamic provisioning
      ## If undefined (the default) or set to null, no storageClassName spec is
      ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
      ##   GKE, AWS & OpenStack)
      ##
      # storageClass: "-"
      accessModes:
        - ReadWriteOnce
      size: 8Gi

##
## External Database Configuration
##
## All of these values are only used when mariadb.enabled is set to false
##
externalDatabase:
  ## Database host
  ##
  host: localhost
  ## non-root Username for Wordpress Database
  ##
  user: bn_wordpress
  ## Database password
  ##
  password: ""
  ## Database name
  ##
  database: bitnami_wordpress
  ## Database port number
  ##
  port: 3306
  ## Use existing secret (ignores previous password)
  ## must contain key `mariadb-password`
  ## NOTE: When it's set, the `externalDatabase.password` parameter is ignored
  ##
  # existingSecret:

## Make use of custom post-init.d user scripts functionality inside the bitnami/wordpress image
## ref: https://github.com/bitnami/bitnami-docker-wordpress/tree/master/5/debian-10/rootfs/post-init.d
##
## The logic of the post-init.d user scripts is that all is all files with extensions .sh, .sql or .php are executed for one time only, at the very first initialization of the pod as the very last step of entrypoint.sh.
## Example:
## customPostInitScripts:
##   enable-multisite.sh: |
##     #!/bin/bash
##     chmod +w /bitnami/wordpress/wp-config.php
##     wp core multisite-install --url=example.com --title="Welcome to the WordPress Multisite" --admin_user="doesntmatternotreallyused" --admin_password="doesntmatternotreallyused" --admin_email="user@example.com"
##     cat /docker-entrypoint-init.d/.htaccess > /bitnami/wordpress/.htaccess
##     chmod -w bitnami/wordpress/wp-config.php
##   .htaccess: |
##     RewriteEngine On
##     RewriteBase /
##     ...
##
## NOTE: Combined with extraVolume and extraVolumeMounts to mount the configmap to /docker-entrypoint-init.d where custom user init scripts are looked for
##
customPostInitScripts: {}

##
## Init containers parameters:
## volumePermissions: Change the owner of the persist volume mountpoint to RunAsUser:fsGroup
##
volumePermissions:
  enabled: false
  image:
    registry: docker.io
    repository: bitnami/bitnami-shell
    tag: "10"
    pullPolicy: Always
    ## Optionally specify an array of imagePullSecrets.
    ## Secrets must be manually created in the namespace.
    ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
    ##
    # pullSecrets:
    #   - myRegistryKeySecretName
  resources:
    # We usually recommend not to specify default resources and to leave this as a conscious
    # choice for the user. This also increases chances charts run on environments with little
    # resources, such as Minikube. If you do want to specify resources, uncomment the following
    # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
    limits: {}
    #   cpu: 100m
    #   memory: 128Mi
    requests: {}
    #   cpu: 100m
    #   memory: 128Mi

  ## Init container Security Context
  ## Note: the chown of the data folder is done to containerSecurityContext.runAsUser
  ## and not the below volumePermissions.securityContext.runAsUser
  ## When runAsUser is set to special value "auto", init container will try to chwon the
  ## data folder to autodetermined user&group, using commands: `id -u`:`id -G | cut -d" " -f2`
  ## "auto" is especially useful for OpenShift which has scc with dynamic userids (and 0 is not allowed).
  ## You may want to use this volumePermissions.securityContext.runAsUser="auto" in combination with
  ## podSecurityContext.enabled=false,containerSecurityContext.enabled=false
  ##
  securityContext:
    runAsUser: 0
```

```
helm install wordpress bitnami/wordpress --version 10.6.10 --set wordpressUsername=admin --set wordpressPassword=wpp@ss
```
```
NAME: wordpress
LAST DEPLOYED: Tue May 25 14:47:00 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
** Please be patient while the chart is being deployed **

Your WordPress site can be accessed through the following DNS name from within your cluster:

    wordpress.default.svc.cluster.local (port 80)

To access your WordPress site from outside the cluster follow the steps below:

1. Get the WordPress URL by running these commands:

  NOTE: It may take a few minutes for the LoadBalancer IP to be available.
        Watch the status with: 'kubectl get svc --namespace default -w wordpress'

   export SERVICE_IP=$(kubectl get svc --namespace default wordpress --template "{{ range (index .status.loadBalancer.ingress 0) }}{{.}}{{ end }}")
   echo "WordPress URL: http://$SERVICE_IP/"
   echo "WordPress Admin URL: http://$SERVICE_IP/admin"

2. Open a browser and access WordPress using the obtained URL.

3. Login with the following credentials below to see your blog:

  echo Username: admin
  echo Password: $(kubectl get secret --namespace default wordpress -o jsonpath="{.data.wordpress-password}" | base64 --decode)
```

コマンド
```
helm install wordpress bitnami/wordpress --values values.yaml
```
コマンド結果
```
NAME: wordpress
LAST DEPLOYED: Sun May 23 11:28:10 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
** Please be patient while the chart is being deployed **

Your WordPress site can be accessed through the following DNS name from within your cluster:

    wordpress.default.svc.cluster.local (port 80)

To access your WordPress site from outside the cluster follow the steps below:

1. Get the WordPress URL by running these commands:

  NOTE: It may take a few minutes for the LoadBalancer IP to be available.
        Watch the status with: 'kubectl get svc --namespace default -w wordpress'

   export SERVICE_IP=$(kubectl get svc --namespace default wordpress --template "{{ range (index .status.loadBalancer.ingress 0) }}{{.}}{{ end }}")
   echo "WordPress URL: http://$SERVICE_IP/"
   echo "WordPress Admin URL: http://$SERVICE_IP/admin"

2. Open a browser and access WordPress using the obtained URL.

3. Login with the following credentials below to see your blog:

  echo Username: admin
  echo Password: $(kubectl get secret --namespace default wordpress -o jsonpath="{.data.wordpress-password}" | base64 --decode)
```

コマンド
```
helm list
```
コマンド結果
```
NAME            NAMESPACE       REVISION        UPDATED                                 STATUS          CHART                   APP VERSION
wordpress       default         1               2021-05-25 14:47:00.71429811 +0000 UTC  deployed        wordpress-10.6.10       5.6.2
```

コマンド
```
export SERVICE_IP=$(kubectl get svc --namespace default wordpress --template "{{ range (index .status.loadBalancer.ingress 0) }}{{.}}{{ end }}")
```

コマンド
```
echo "WordPress Admin URL: http://$SERVICE_IP/admin"
```
コマンド結果
```
WordPress Admin URL: http://35.xx.xx.xx/admin
```

コマンド
```
kubectl get service
```
コマンド結果
```
NAME                TYPE           CLUSTER-IP    EXTERNAL-IP      PORT(S)                      AGE
kubernetes          ClusterIP      10.48.0.1     <none>           443/TCP                      124m
wordpress           LoadBalancer   10.48.1.22    35.xx.xx.xx      80:30633/TCP,443:31049/TCP   9m32s
wordpress-mariadb   ClusterIP      10.48.6.162   <none>           3306/TCP                     9m32s
```

コマンド
```
echo Password: $(kubectl get secret --namespace default wordpress -o jsonpath="{.data.wordpress-password}" | base64 --decode)
```
コマンド結果
```
Password: wpp@ss
```

コマンド
```
kubectl get pods
```
コマンド結果
```
NAME                         READY   STATUS    RESTARTS   AGE
wordpress-7c9fcf495f-mkkvn   1/1     Running   0          11m
wordpress-mariadb-0          1/1     Running   0          11m
```

コマンド
```
kubectl get persistentvolumes,persistentvolumeclaims
```
コマンド結果
```
NAME                                                        CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                              STORAGECLASS   REASON   AGE
persistentvolume/pvc-2aed2ea9-1077-4164-8d8e-9894ea85e093   10Gi       RWO            Delete           Bound    default/wordpress                  standard                11m
persistentvolume/pvc-b2217342-7e51-4ad9-8c85-b660f1393cad   8Gi        RWO            Delete           Bound    default/data-wordpress-mariadb-0   standard                11m

NAME                                             STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/data-wordpress-mariadb-0   Bound    pvc-b2217342-7e51-4ad9-8c85-b660f1393cad   8Gi        RWO            standard       11m
persistentvolumeclaim/wordpress                  Bound    pvc-2aed2ea9-1077-4164-8d8e-9894ea85e093   10Gi       RWO            standard       11m
```

コマンド
```
helm uninstall wordpress
```
コマンド結果
```
release "wordpress" uninstalled
```

コマンド
```
helm list
```
コマンド結果
```
NAME    NAMESPACE       REVISION        UPDATED STATUS  CHART   APP VERSION
```

