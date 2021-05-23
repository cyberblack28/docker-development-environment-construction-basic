# 第5章 Kubernetesでコンテナアプリケーションを動かすまで

## 5.1 kubectlコマンド

### 5.1.2 Podの作成

#### リスタートポリシー

```kubectlコマンド
$ kubectl run nginx --image=nginx --restart=Never
PROJECT_ID              NAME              PROJECT_NUMBER
pod/nginx created
```

#### Podステータスの確認

```kubectlコマンド
$ kubectl get pod
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          17s
```

#### ClusterIPアドレスの確認

```kubectlコマンド
$ kubectl describe pod nginx
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

```kubectlコマンド
$ kubectl run busybox --image=busybox --restart=Never -it --rm -- sh
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

```kubectlコマンド
$ kubectl delete pod nginx
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

#### マニフェストファイルの作成

```kubectlコマンド
$ cat nginx.yaml
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

#### Pod の作成

```kubectlコマンド
$ kubectl apply -f nginx.yaml
pod/nginx created
```

```kubectlコマンド
$ kubectl get pod nginx
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          69s
```

#### マニフェストファイルを指定したPod の削除

```kubectlコマンド
$ kubectl delete -f nginx.yaml
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

### 5.1.3 Deployment

#### マニフェストファイルの作成

```kubectlコマンド
$ vim nginx-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx
  name: nginx
spec:
  replicas: 2
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

```kubectlコマンド
$ kubectl apply -f nginx-deployment.yaml
deployment.apps/nginx created
```

```kubectlコマンド
$ kubectl get pod
NAME                    READY   STATUS    RESTARTS   AGE
nginx-f589cd57f-5hsbt   1/1     Running   0          11s
nginx-f589cd57f-jkhmq   1/1     Running   0          11s
```

```kubectlコマンド
$ kubectl get deployment
NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   2/2     2            2           39s
```

#### セルフヒーリング機能の確認

```kubectlコマンド
$ kubectl delete pod nginx-f589cd57f-5hsbt
pod "nginx-f589cd57f-5hsbt" deleted
```

```kubectlコマンド
$ kubectl get pod
NAME                    READY   STATUS    RESTARTS   AGE
nginx-f589cd57f-dbkns   1/1     Running   0          32s
nginx-f589cd57f-jkhmq   1/1     Running   0          2m4s
```

#### ローリングアップデート

```kubectlコマンド
$ kubectl set image deployment nginx nginx=nginx:1.20.0
deployment.apps/nginx image updated
```

```kubectlコマンド
$ kubectl get pod
NAME                     READY   STATUS    RESTARTS   AGE
nginx-7d994954c4-bs8jm   1/1     Running   0          48s
nginx-7d994954c4-nwvfl   1/1     Running   0          47s
```

#### Podの詳細確認

```kubectlコマンド
$ kubectl describe pod nginx-7d994954c4-bs8jm
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

```kubectlコマンド
$ kubectl rollout history deployment nginx
deployment.apps/nginx
REVISION  CHANGE-CAUSE
1         <none>
2         <none>
```

#### 各リビジョンの確認

```kubectlコマンド
$ kubectl rollout history deployment nginx --revision=1
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

```kubectlコマンド
$ kubectl rollout history deployment nginx --revision=2
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

```kubectlコマンド
$ kubectl rollout undo deployment nginx
deployment.apps/nginx rolled back
```

#### kubectl describeコマンドによるDeploymentの確認

```kubectlコマンド
$ kubectl describe deployment nginx
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

```kubectlコマンド
$ kubectl rollout undo deployment nginx --to-revision=2
deployment.apps/nginx rolled back
```

```kubectlコマンド
$ kubectl describe deployment nginx
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

```kubectlコマンド
$ kubectl delete deployment nginx
deployment.apps "nginx" deleted
```

```kubectlコマンド
$ kubectl delete -f nginx-deployment.yaml
deployment.apps "nginx" deleted
```

```kubectlコマンド
$ kubectl get deployment
No resources found in default namespace.
```

### 5.1.4 Service

#### Podの作成

```kubectlコマンド
$ kubectl run nginx --image=nginx:1.20.0 --restart=Never --port=80
pod/nginx created
```

```kubectlコマンド
$ kubectl describe pod nginx
Name:         nginx
Namespace:    default
Priority:     0
Node:         gke-k8s-cluster-default-pool-8ad2d901-1k7g/10.146.15.232
Start Time:   Sun, 16 May 2021 14:48:24 +0000
Labels:       run=nginx
Annotations:  <none>
Status:       Running
IP:           10.0.1.19
IPs:
  IP:  10.0.1.19
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

```kubectlコマンド
$ kubectl expose pod nginx --name=nginx --port=80 --target-port=80
service/nginx exposed
```

#### ClusterIPアドレスの確認

```kubectlコマンド
$ kubectl get service
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.3.240.1     <none>        443/TCP   11h
nginx        ClusterIP   10.3.248.150   <none>        80/TCP    73s
```

#### PodのIPアドレスの確認

```kubectlコマンド
$ kubectl get ep
NAME         ENDPOINTS          AGE
kubernetes   35.200.23.27:443   11h
nginx        10.0.1.19:80       2m4s
```


#### busybox Podの作成

```kubectlコマンド
$ kubectl run busybox --image=busybox --restart=Never --rm -it /bin/sh
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
/ # wget -O- 10.0.1.19:80
Connecting to 10.0.1.19:80 (10.0.1.19:80)
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

```kubectlコマンド
$ kubectl delete service nginx
service "nginx" deleted
```

```kubectlコマンド
$ kubectl delete pod nginx
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get service
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.3.240.1   <none>        443/TCP   11h
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

```kubectlコマンド
$ kubectl run nginx --image=nginx:1.20.0 --port=80 --expose
service/nginx created
pod/nginx created
```

### 5.1.5 ConfigMap

#### ConfigMapの作成

```kubectlコマンド
$ kubectl create configmap sample-config1 --from-literal=var1=docker --from-literal=var2=kubernetes
configmap/sample-config1 created
```

```kubectlコマンド
$ kubectl get configmap
NAME               DATA   AGE
kube-root-ca.crt   1      4m31s
sample-config1     2      82s
```

```kubectlコマンド
$ kubectl get configmap sample-config1 -o yaml
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

#### マニフェストファイルの編集

```kubectlコマンド
$ kubectl run --restart=Never nginx --image=nginx -o yaml --dry-run=client > nginx-configmap1.yaml
```

```linuxコマンド
$ vim nginx-configmap1.yaml
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
    envFrom:
    - configMapRef:
        name: sample-config1
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}

```linuxコマンド
$ cd
$ cd container-develop-environment-construction-guide/Chapter05/5-1-5-01
```

#### Podの作成

```kubectlコマンド
$ kubectl create -f nginx-configmap1.yaml
pod/nginx created
```

#### 環境変数の確認

```kubectlコマンド
$ kubectl exec -it nginx -- env
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

```kubectlコマンド
$ kubectl delete -f nginx-configmap1.yaml
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

```kubectlコマンド
$ kubectl delete configmap sample-config1
configmap "sample-config1" deleted
```

```kubectlコマンド
$ kubectl get configmap
NAME               DATA   AGE
kube-root-ca.crt   1      26m
```

#### 作成したファイルからConfigMapを作成

```kubectlコマンド
$ kubectl create configmap sample-config2 --from-env-file=sample-config2.env
configmap/sample-config2 created
```

```kubectlコマンド
$ kubectl get configmap
NAME               DATA   AGE
kube-root-ca.crt   1      91m
sample-config2     2      112s
```

```kubectlコマンド
$ kubectl get configmap sample-config2 -o yaml
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

```kubectlコマンド
$ cat nginx-configmap2.yaml
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

```kubectlコマンド
$ kubectl apply -f nginx-configmap2.yaml
pod/nginx created
```

```kubectlコマンド
$ kubectl exec -it nginx -- env
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

```kubectlコマンド
$ kubectl delete -f nginx-configmap2.yaml
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

```kubectlコマンド
$ kubectl delete configmap sample-config2
configmap "sample-config2" deleted
```

```kubectlコマンド
$ kubectl get configmap
NAME               DATA   AGE
kube-root-ca.crt   1      114m
```

### 5.1.6 ConfigMapとKeyの参照

```linuxコマンド
$ cd ../5-1-6-01
```

```linuxコマンド
$ cat sample-config3.txt
rook
vitess
containerd
helm
```

```kubectlコマンド
$ kubectl create configmap sample-cmvolume --from-file=sample-config3=sample-config3.txt
configmap/sample-cmvolume created
```

```kubectlコマンド
$ kubectl get configmap
NAME               DATA   AGE
kube-root-ca.crt   1      39h
sample-cmvolume    1      119s
```

```kubectlコマンド
$ kubectl get configmap sample-cmvolume -o yaml
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

#### マニフェストファイルの編集

```linuxコマンド
$ cat nginx-configmap3.yaml
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

```kubectlコマンド
$ kubectl apply -f nginx-configmap3.yaml
pod/nginx created
```

```kubectlコマンド
$ kubectl exec -it nginx -- cat /configmap/sample-config3
rook
vitess
containerd
helm
```

### 5.1.7 Secret

#### Secretの作成

```kubectlコマンド
$ kubectl create secret generic sample-secret1 --from-literal=password=testp@ss
secret/sample-secret1 created
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      39h
sample-secret1        Opaque                                1      118s
```

#### マニフェストの確認

```kubectlコマンド
$ kubectl get secret sample-secret1 -o yaml
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

```linuxコマンド
$ echo dGVzdHBAc3M= | base64 -d
testp@ss
```

#### SecretとPodの削除

```kubectlコマンド
$ kubectl delete secret sample-secret1
secret "sample-secret1" deleted
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      39h
```

```kubectlコマンド
$ kubectl delete pod nginx
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

#### envファイルからSecretを作成

```linuxコマンド
$ cd ../5-1-7-01
```

```linuxコマンド
$ cat sample-secret2.env
password=p@ssw0rd
```

```kubectlコマンド
$ kubectl create secret generic sample-secret2 --from-env-file=sample-secret2.env
secret/sample-secret2 created
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      39h
sample-secret2        Opaque                                1      85s
```

```kubectlコマンド
$ kubectl get secret sample-secret2 -o yaml
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

```linuxコマンド
$ cat nginx-secret.yaml
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

```kubectlコマンド
$ kubectl apply -f nginx-secret.yaml
pod/nginx created
```

```kubectlコマンド
$ kubectl describe pod nginx
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

```kubectlコマンド
$ kubectl exec -it nginx -- env | grep PASSWORD
PASSWORD=p@ssw0rd
```

#### SecretとPodの削除

```kubectlコマンド
$ kubectl delete secret sample-secret2
secret "sample-secret2" deleted
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      40h
```

```kubectlコマンド
$ kubectl delete pod nginx
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

#### Secretの作成とVolumeデータの参照

```kubectlコマンド
$ cat sample-secret3.txt
admin=86asnNlW
operator=po89SYin
user=oshu894LD
```

```kubectlコマンド
$ kubectl create secret generic sample-secret3 --from-file=sample-secret3.txt
secret/sample-secret3 created
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      40h
sample-secret3        Opaque                                1      39s
```

```kubectlコマンド
$ kubectl get secret sample-secret3 -o yaml
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

```linuxコマンド
$ cat nginx-secret2.yaml
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

```kubectlコマンド
$ kubectl apply -f nginx-secret2.yaml
pod/nginx created
```

```kubectlコマンド
$ kubectl exec -it nginx -- cat /secret/sample-secret3.txt
admin=86asnNlW
operator=po89SYin
user=oshu894LD
```

#### SecretとPodの削除

```kubectlコマンド
$ kubectl delete secret sample-secret3
secret "sample-secret3" deleted
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      41h
```

```kubectlコマンド
$ kubectl delete pod nginx
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

### 5.1.8 Multi Container Pod

```linuxコマンド
$ cd ../5-1-8-01
```

```kubectlコマンド
$ cat multicontainer.yaml
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

```kubectlコマンド
$ kubectl apply -f multicontainer.yaml
pod/nginx-pod created
```

```kubectlコマンド
$ kubectl get pod
NAME        READY   STATUS    RESTARTS   AGE
nginx-pod   2/2     Running   0          3m26s
```

```kubectlコマンド
$ kubectl exec -it nginx-pod -c nginx -- /bin/sh
# curl localhost
Hello from the work-container
# ls /usr/share/nginx/html
index.html
# exit
```

```kubectlコマンド
$ kubectl delete pod nginx-pod
pod "nginx-pod" deleted
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

### 5.1.9 createとapply

```linuxコマンド
$ cd ../5-1-9-01
```

```kubectlコマンド
$ kubectl create -f nginx.yaml
pod/nginx created
```

```kubectlコマンド
$ kubectl create -f nginx.yaml
Error from server (AlreadyExists): error when creating "nginx.yaml": pods "nginx" already exists
```

```kubectlコマンド
$ kubectl delete -f nginx.yaml
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl apply -f nginx.yaml
pod/nginx created
```

```kubectlコマンド
$ kubectl apply -f nginx.yaml
pod/nginx configured
```

```kubectlコマンド
$ kubectl delete -f nginx.yaml
pod "nginx" deleted
```

```kubectlコマンド
$ kubectl delete pod nginx
pod "nginx" deleted
```

```kubectlコマンド
＄ kubectl get pod
No resources found in default namespace.
```

## 5・2 Kubernetesでアプリケーションを動かす

### 5.2.2 NFSサーバの作成

#### gcePersistentDiskの作成

```gcloudコマンド
$ gcloud compute disks create nfs-disk --size=10GB --zone=asia-northeast1-a
WARNING: You have selected a disk size of under [200GB]. This may result in poor I/O performance. For more information, see: https://developers.google.com/compute/docs/disks#performance.
Created [https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-a/disks/nfs-disk].
NAME      ZONE               SIZE_GB  TYPE         STATUS
nfs-disk  asia-northeast1-a  10       pd-standard  READY
New disks are unformatted. You must format and mount a disk before it
can be used. You can find instructions on how to do this at:

https://cloud.google.com/compute/docs/disks/add-persistent-disk#formatting
```

```gcloudコマンド
$ gcloud compute disks describe --zone=asia-northeast1-a nfs-disk
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

#### NFSサーバのDeploymentとServiceの作成

```linuxコマンド
$ cd ../5-2-2-01
```

```linuxコマンド
$ cat nfs-server.yaml
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
            name: nfs
      volumes:
        - name: nfs
          gcePersistentDisk:
            pdName: nfs-disk
            fsType: ext4
```

```kubectlコマンド
$ kubectl apply -f nfs-server.yaml
deployment.apps/nfs-server created
```

```kubectlコマンド
$ kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
nfs-server-6f7fc97dfd-n69bw   1/1     Running   0          12m
```

```linuxコマンド
$ cat nfs-service.yaml
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

```kubectlコマンド
$ kubectl apply -f nfs-service.yaml
service/nfs-service created
```

```kubectlコマンド
$ kubectl get service
NAME          TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)                      AGE
kubernetes    ClusterIP   10.3.240.1    <none>        443/TCP                      45h
nfs-service   ClusterIP   10.3.252.83   <none>        2049/TCP,20048/TCP,111/TCP   51s
```

### 5.2.3 Secretの作成

```kubectlコマンド
$ kubectl create secret generic mysql --from-literal=password=mysqlp@ssw0d
secret/mysql created
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      46h
mysql                 Opaque                                1      79s
```

```kubectlコマンド
$ kubectl get secret mysql -o yaml
apiVersion: v1
data:
  password: bXlzcWxwQHNzdzBk
kind: Secret
metadata:
  creationTimestamp: "2021-05-22T12:47:32Z"
  name: mysql
  namespace: default
  resourceVersion: "983091"
  selfLink: /api/v1/namespaces/default/secrets/mysql
  uid: 17258347-09ab-445b-b0dc-b12a1d031b62
type: Opaque
```

### 5.2.4 PersistentVolumeとPersistentVolumeClaimの作成

#### PersistentVolumeの作成

```linuxコマンド
$ cd ../5-2-4-01
```

```linuxコマンド
$ vim mysql-pv.yaml
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
    server: 10.3.252.83
    path: /
```

```kubectlコマンド
$ kubectl apply -f mysql-pv.yaml
persistentvolume/mysql-pv created
```

```linuxコマンド
$ vim wordpress-pv.yaml
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
    server: 10.3.252.83
    path: /
```

```kubectlコマンド
$ kubectl apply -f wordpress-pv.yaml
persistentvolume/wordpress-pv created
```

```kubectlコマンド
$ kubectl get persistentvolume
NAME           CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
mysql-pv       10Gi       RWX            Retain           Available           mysql                   9m23s
wordpress-pv   10Gi       RWX            Retain           Available           wordpress               44s
```

#### PersistentVolumeClaimの作成

```linuxコマンド
$ cat mysql-pvc.yaml
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

```kubectlコマンド
$ kubectl apply -f mysql-pvc.yaml
persistentvolumeclaim/mysql-pvc created
```

```linuxコマンド
$ cat wordpress-pvc.yaml
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

```kubectlコマンド
$ kubectl apply -f wordpress-pvc.yaml
persistentvolumeclaim/wordpress-pvc created
```

```kubectlコマンド
$ kubectl get persistentvolume,persistentvolumeclaim
NAME                            CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
persistentvolume/mysql-pv       10Gi       RWX            Retain           Bound    default/mysql-pvc       mysql                   27m
persistentvolume/wordpress-pv   10Gi       RWX            Retain           Bound    default/wordpress-pvc   wordpress               18m

NAME                                  STATUS   VOLUME         CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/mysql-pvc       Bound    mysql-pv       10Gi       RWX            mysql          4m9s
persistentvolumeclaim/wordpress-pvc   Bound    wordpress-pv   10Gi       RWX            wordpress      40s
```

### 5.2.5 DeploymentとServiceの作成

#### MySQLのDeploymentとServiceの作成

```linuxコマンド
$ cd ../5-2-5-01
```

```linuxコマンド
$ cat mysql.yaml
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

```kubectlコマンド
$ kubectl apply -f mysql.yaml
deployment.apps/mysql created
```

```kubectlコマンド
$ kubectl get pod -l app=mysql
NAME                     READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-hj4k9   1/1     Running   0          84s
```

```linuxコマンド
$ cat mysql-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: mysql-service
  labels:
    app: mysql
spec:
  type: ClusterIP
  ports:
    - port: 3306
  selector:
    app: mysql
```

```kubectlコマンド
$ kubectl apply -f mysql-service.yaml
service/mysql-service created
```

```kubectlコマンド
$ kubectl get service
NAME            TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
kubernetes      ClusterIP   10.3.240.1     <none>        443/TCP                      47h
mysql-service   ClusterIP   10.3.247.235   <none>        3306/TCP                     3m23s
nfs-service     ClusterIP   10.3.252.83    <none>        2049/TCP,20048/TCP,111/TCP   78m
```

#### WordPressのDeploymentとServiceの作成

```linuxコマンド
$ cat wordpress.yaml
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
        - image: wordpress
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

```kubectlコマンド
$ kubectl apply -f wordpress.yaml
deployment.apps/wordpress created
```

```linuxコマンド
$ cat wordpress-service.yaml
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
```

```kubectlコマンド
$ kubectl apply -f wordpress-service.yaml
service/wordpress-service created
```

```kubectlコマンド
$ kubectl get service
NAME                TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)                      AGE
kubernetes          ClusterIP      10.3.240.1     <none>          443/TCP                      47h
mysql-service       ClusterIP      10.3.247.235   <none>          3306/TCP                     18m
nfs-service         ClusterIP      10.3.252.83    <none>          2049/TCP,20048/TCP,111/TCP   93m
wordpress-service   LoadBalancer   10.3.247.206   34.84.235.193   80:31998/TCP                 2m52s
```

### 5.2.6 アプリケーションのスケールアウト/スケールイン

#### kubectlコマンドによるPodの追加

```kubectlコマンド
$ kubectl scale deployment wordpress --replicas 10
deployment.apps/wordpress scaled
```

```kubectlコマンド
$ kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-hj4k9        1/1     Running   0          117m
nfs-server-6f7fc97dfd-n69bw   1/1     Running   0          3h26m
wordpress-598599d9d6-2d8sp    1/1     Running   0          2m2s
wordpress-598599d9d6-5ngrl    1/1     Running   0          2m2s
wordpress-598599d9d6-7w6dh    1/1     Running   0          2m2s
wordpress-598599d9d6-jfb2z    1/1     Running   0          2m2s
wordpress-598599d9d6-mw9r2    1/1     Running   0          2m2s
wordpress-598599d9d6-p4pkt    1/1     Running   0          102m
wordpress-598599d9d6-q7rpc    1/1     Running   0          2m2s
wordpress-598599d9d6-rf77m    1/1     Running   0          2m2s
wordpress-598599d9d6-sg8nn    1/1     Running   0          2m2s
wordpress-598599d9d6-xthb5    1/1     Running   0          2m2s
```

#### マニフェストファイルの編集によるPod数の変更

```linuxコマンド
$ cd ../5-2-5-01
```

```linuxコマンド
$ cat wordpress.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: wordpress
  labels:
    app: wordpress
spec:
  replicas: 5
  selector:
    matchLabels:
      app: wordpress
  template:
    metadata:
      labels:
        app: wordpress
    spec:
      containers:
        - image: wordpress
          name: wordpress
          env:
          - name: WORDPRESS_DB_HOST         //cmt{# Service名「//tt{mysql-service//}」を定義//}
            value: mysql-service
          - name: WORDPRESS_DB_PASSWORD     //cmt{# MySQLのデータベースパスワードを参照する定義//}
            valueFrom:
              secretKeyRef:
                name: mysql
                key: password
          ports:
            - containerPort: 80
              name: wordpress
          volumeMounts:                     //cmt{# Podのマウントパス定義//}
            - name: wordpress-local-storage
              mountPath: /var/www/html
      volumes:                              //cmt{# 「//tt{wordpress-pvc//}」を指定する定義//}
        - name: wordpress-local-storage
          persistentVolumeClaim:
            claimName: wordpress-pvc
```

```kubectlコマンド
$ kubectl apply -f wordpress.yaml
deployment.apps/wordpress configured
```

```kubectlコマンド
$ kubectl edit deployment wordpress
deployment.apps/wordpress edited
```

```kubectlコマンド
$ kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-hj4k9        1/1     Running   0          128m
nfs-server-6f7fc97dfd-n69bw   1/1     Running   0          3h37m
wordpress-598599d9d6-jfb2z    1/1     Running   0          12m
```

```kubectlコマンド
$ kubectl delete -f wordpress-service.yaml
service "wordpress-service" deleted

```kubectlコマンド
$ kubectl delete -f wordpress.yaml
deployment.apps "wordpress" deleted
```

```kubectlコマンド
$ kubectl delete -f mysql-service.yaml
service "mysql-service" deleted
```

```kubectlコマンド
$ kubectl delete -f mysql.yaml
deployment.apps "mysql" deleted
```

```kubectlコマンド
$ kubectl delete -f wordpress-pvc.yaml
persistentvolumeclaim "wordpress-pvc" deleted
```

```kubectlコマンド
$ kubectl delete -f wordpress-pv.yaml
persistentvolume "wordpress-pv" deleted
```

```kubectlコマンド
$ kubectl delete -f mysql-pvc.yaml
persistentvolumeclaim "mysql-pvc" deleted
```

```kubectlコマンド
$ kubectl delete -f mysql-pv.yaml
persistentvolume "mysql-pv" deleted
```

```kubectlコマンド
$ kubectl delete -f nfs-service.yaml
service "nfs-service" deleted
```

```kubectlコマンド
$ kubectl delete -f nfs-server.yaml
deployment.apps "nfs-server" deleted
```

```kubectlコマンド
$ kubectl delete secret mysql
secret "mysql" deleted
```

```kubectlコマンド
$ kubectl get deployment
No resources found in default namespace.
```

```kubectlコマンド
$ kubectl get service
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.3.240.1   <none>        443/TCP   2d1h
```

```kubectlコマンド
$ kubectl get persistentvolume,persistentvolumeclaim
No resources found
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      2d1h
```

```kubectlコマンド
$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-dmlps   kubernetes.io/service-account-token   3      2d1h
```

```gcloudコマンド
$ gcloud compute disks delete nfs-disk --zone=asia-northeast1-a
The following disks will be deleted:
 - [nfs-disk] in [asia-northeast1-a]

Do you want to continue (Y/n)?   Y

Deleted [https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-a/disks/nfs-disk].
```

## 5.3 マニフェストの管理

### 5.3.2 Chartの作成とアプリケーションのインストール

#### Helm クライアントのインストール

```linuxコマンド
$ cd ../5-3-2-01
```

```linuxコマンド
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
```

```linuxコマンド
$ chmod 700 get_helm.sh
```

```linuxコマンド
$ ./get_helm.sh
```

```helmコマンド
$ helm version
version.BuildInfo{Version:"v3.5.4", GitCommit:"1b5edb69df3d3a08df77c9902dc17af864ff05d1", GitTreeState:"clean", GoVersion:"go1.15.11"}
```

#### Chartの雛形の作成

```helmコマンド
$ helm create wordpress
Creating wordpress
```

```linuxコマンド
$ ls wordpress
charts  Chart.yaml  templates  values.yaml
```

```linuxコマンド
$ ls wordpress/templates/
deployment.yaml  _helpers.tpl  hpa.yaml  ingress.yaml  NOTES.txt  serviceaccount.yaml  service.yaml  tests
```

```linuxコマンド
$ ls wordpress/templates/tests
test-connection.yaml
```

#### WordPressの独自Chartの作成

```linuxコマンド
$ rm -rf wordpress/templates/*
```

```linuxコマンド
$ ls wordpress/templates/
```

#### Secret マニフェストテンプレートの作成

```linuxコマンド
$ mv helm-yaml/mysql-secret.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/mysql-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysql
type: Opaque
data:
  password: {{ .Values.mysql_secret.password }}
```

#### PersistentVolume、PersistentVolumeClaimの作成

```linuxコマンド
$ cp -p helm-yaml/mysql-pv.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/mysql-pv.yaml
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

```linuxコマンド
$ mv helm-yaml/wordpress-pv.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/wordpress-pv.yaml
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

```linuxコマンド
$ mv helm-yaml/mysql-pvc.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/mysql-pvc.yaml
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

```linuxコマンド
$ mv helm-yaml/wordpress-pvc.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/wordpress-pvc.yaml
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

```linuxコマンド
$ mv helm-yaml/mysql.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/mysql.yaml
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

```linuxコマンド
$ mv helm-yaml/mysql-service.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/mysql-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.mysql_service.name }}
  labels:
    app: mysql
spec:
  type: {{ .Values.mysql_service.type }}
  ports:
    - port: {{ .Values.mysql_service.port }}
  selector:
    app: mysql
```

```linuxコマンド
$ mv helm-yaml/wordpress.yaml wordpress/templates
```

```linuxコマンド
$ cat helm-yaml/wordpress.yaml
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

```linuxコマンド
$ mv helm-yaml/wordpress-service.yaml wordpress/templates
```

```linuxコマンド
$ cat wordpress/templates/wordpress-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.wordpress_service.name }}
  labels:
    app: wordpress
spec:
  type: {{ .Values.wordpress_service.type }}
  ports:
    - port: {{ .Values.wordpress_service.port }}
      targetPort: {{ .Values.wordpress_service.targetPort }}
      protocol: {{ .Values.wordpress_service.protocol }}
  selector:
    app: wordpress
```

#### values.yaml の作成

```linuxコマンド
$ cp -p helm-yaml/values.yaml wordpress
```

```linuxコマンド
$ cat wordpress/values.yaml
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
  port: 3306

#wordpress
wordpress:
  name: wordpress
  replicas: 1
  image: wordpress
  value: mysql-service
  containerPort: 80
  mountPath: /var/www/html
  claimName: wordpress-pvc

#wordpress-service
wordpress_service:
  name: wordpress-service
  type: LoadBalancer
  port: 80
  targetPort: 80
  protocol: TCP
```

#### Chartのデバッグ

```helmコマンド
$ helm install wordpress --debug --dry-run wordpress
install wordpress --debug --dry-run wordpress
install.go:173: [debug] Original chart version: ""
install.go:190: [debug] CHART PATH: /home/iyutaka2020/container-develop-environment-construction-guide/Chapter05/5-3-2-01/wordpress

Error: template: wordpress/templates/tests/test-connection.yaml:14:61: executing "wordpress/templates/tests/test-connection.yaml" at <.Values.service.port>: nil pointer evaluating interface {}.port
helm.go:81: [debug] template: wordpress/templates/tests/test-connection.yaml:14:61: executing "wordpress/templates/tests/test-connection.yaml" at <.Values.service.port>: nil pointer evaluating interface {}.port