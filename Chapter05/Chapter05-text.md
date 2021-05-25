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
$ cp -p helm-yaml/mysql-secret.yaml wordpress/templates
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
$ cp -p helm-yaml/wordpress-pv.yaml wordpress/templates
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
$ cp -p helm-yaml/mysql-pvc.yaml wordpress/templates
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
$ cp -p helm-yaml/wordpress-pvc.yaml wordpress/templates
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
$ cp -p helm-yaml/mysql.yaml wordpress/templates
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
$ cp -p helm-yaml/mysql-service.yaml wordpress/templates
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
$ cp -p helm-yaml/wordpress.yaml wordpress/templates
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
$ cp -p helm-yaml/wordpress-service.yaml wordpress/templates
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
install.go:173: [debug] Original chart version: ""
install.go:190: [debug] CHART PATH: /home/iyutaka2020/container-develop-environment-construction-guide/Chapter05/5-3-2-01/wordpress

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
    - port: 3306
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

```helmコマンド
$ helm lint wordpress
==> Linting wordpress
[INFO] Chart.yaml: icon is recommended

1 chart(s) linted, 0 chart(s) failed
```

#### Chartのパッケージ化

```helmコマンド
$ helm package wordpress
Successfully packaged chart and saved it to: /home/iyutaka2020/container-develop-environment-construction-guide/Chapter05/5-3-2-01/wordpress-0.1.0.tgz
```

```helmコマンド
$ helm repo index .
```

```linuxコマンド
$ cat index.yaml
apiVersion: v1
entries:
  wordpress:
  - apiVersion: v2
    appVersion: 1.16.0
    created: "2021-05-23T07:38:08.15113946Z"
    description: A Helm chart for Kubernetes
    digest: 7993b1c7198683a6c715f39ba2044a478016f2d11b18eccc978dbcae2d106874
    name: wordpress
    type: application
    urls:
    - wordpress-0.1.0.tgz
    version: 0.1.0
generated: "2021-05-23T07:38:08.15011907Z"
```

#### WordPressのインストール

```helmコマンド
$ helm install wordpress ./wordpress
NAME: wordpress
LAST DEPLOYED: Sun May 23 07:41:44 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
```

```kubectlコマンド
$ kubectl get pod
NAME                         READY   STATUS    RESTARTS   AGE
mysql-656fbb9446-5wkth       1/1     Running   0          58s
wordpress-598599d9d6-qkplh   1/1     Running   0          58s
```

```kubectlコマンド
$ kubectl get persistentvolume,persistentvolumeclaim
NAME                            CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
persistentvolume/mysql-pv       10Gi       RWO            Retain           Bound    default/mysql-pvc       mysql                   2m10s
persistentvolume/wordpress-pv   10Gi       RWO            Retain           Bound    default/wordpress-pvc   wordpress               2m10s

NAME                                  STATUS   VOLUME         CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/mysql-pvc       Bound    mysql-pv       10Gi       RWO            mysql          2m10s
persistentvolumeclaim/wordpress-pvc   Bound    wordpress-pv   10Gi       RWO            wordpress      2m10s
```

```kubectlコマンド
$ kubectl get service
NAME                TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)        AGE
kubernetes          ClusterIP      10.3.240.1     <none>          443/TCP        2d16h
mysql-service       ClusterIP      10.3.242.238   <none>          3306/TCP       3m12s
wordpress-service   LoadBalancer   10.3.250.23    34.84.235.193   80:30863/TCP   3m12s
```

#### WordPressのアンインストール

```helmコマンド
$ helm uninstall wordpress
release "wordpress" uninstalled
```

```kubectlコマンド
$ kubectl get pod
No resources found in default namespace.
```

```kubectlコマンド
$ kubectl get persistentvolume,persistentvolumeclaim
No resources found
```

```kubectlコマンド
$ kubectl get service
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.3.240.1   <none>        443/TCP   2d17h
```

#### 5.3.3 公開Chartの利用

### リポジトリの操作

```helmコマンド
$ helm repo add bitnami https://charts.bitnami.com/bitnami
"bitnami" has been added to your repositories
```

```helmコマンド
$ helm repo list
NAME            URL
bitnami         https://charts.bitnami.com/bitnami
```

```helmコマンド
$ helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "bitnami" chart repository
Update Complete. ⎈Happy Helming!⎈
```

```helmコマンド
$ helm search repo wordpress
NAME                    CHART VERSION   APP VERSION     DESCRIPTION
bitnami/wordpress       11.0.9          5.7.2           Web publishing platform for building blogs and ...
```

```helmコマンド
$ helm search hub wordpress
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

```helmコマンド
$ helm show values bitnami/wordpress
## @section Global parameters
## Global Docker image parameters
## Please, note that this will override the image parameters, including dependencies, configured to use the global value
## Current available global Docker image parameters: imageRegistry, imagePullSecrets and storageClass

## @param global.imageRegistry Global Docker image registry
## @param global.imagePullSecrets Global Docker registry secret names as an array
## @param global.storageClass Global StorageClass for Persistent Volume(s)
##
global:
  imageRegistry:
  ## E.g.
  ## imagePullSecrets:
  ##   - myRegistryKeySecretName
  ##
  imagePullSecrets: []
  storageClass:

## @section Common parameters

## @param kubeVersion Override Kubernetes version
##
kubeVersion:
## @param nameOverride String to partially override common.names.fullname
##
nameOverride:
## @param fullnameOverride String to fully override common.names.fullname
##
fullnameOverride:
## @param commonLabels Labels to add to all deployed objects
##
commonLabels: {}
## @param commonAnnotations Annotations to add to all deployed objects
##
commonAnnotations: {}
## @param clusterDomain Kubernetes cluster domain name
##
clusterDomain: cluster.local
## @param extraDeploy Array of extra objects to deploy with the release
##
extraDeploy: []

## @section WordPress Image parameters

## Bitnami WordPress image
## ref: https://hub.docker.com/r/bitnami/wordpress/tags/
## @param image.registry WordPress image registry
## @param image.repository WordPress image repository
## @param image.tag WordPress image tag (immutable tags are recommended)
## @param image.pullPolicy WordPress image pull policy
## @param image.pullSecrets WordPress image pull secrets
## @param image.debug Enable image debug mode
##
image:
  registry: docker.io
  repository: bitnami/wordpress
  tag: 5.7.2-debian-10-r8
  ## Specify a imagePullPolicy
  ## Defaults to 'Always' if image tag is 'latest', else set to 'IfNotPresent'
  ## ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images
  ##
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ## e.g:
  ## pullSecrets:
  ##   - myRegistryKeySecretName
  ##
  pullSecrets: []
  ## Enable debug mode
  ##
  debug: false

## @section WordPress Configuration parameters
## WordPress settings based on environment variables
## ref: https://github.com/bitnami/bitnami-docker-wordpress#environment-variables

## @param wordpressUsername WordPress username
##
wordpressUsername: user
## @param wordpressPassword WordPress user password
## Defaults to a random 10-character alphanumeric string if not set
##
wordpressPassword: ""
## @param existingSecret Name of existing secret containing WordPress credentials
## NOTE: Must contain key `wordpress-password`
## NOTE: When it's set, the `wordpressPassword` parameter is ignored
##
existingSecret:
## @param wordpressEmail WordPress user email
##
wordpressEmail: user@example.com
## @param wordpressFirstName WordPress user first name
##
wordpressFirstName: FirstName
## @param wordpressLastName WordPress user last name
##
wordpressLastName: LastName
## @param wordpressBlogName Blog name
##
wordpressBlogName: User's Blog!
## @param wordpressTablePrefix Prefix to use for WordPress database tables
##
wordpressTablePrefix: wp_
## @param wordpressScheme Scheme to use to generate WordPress URLs
##
wordpressScheme: http
## @param wordpressSkipInstall Skip wizard installation
## NOTE: useful if you use an external database that already contains WordPress data
## ref: https://github.com/bitnami/bitnami-docker-wordpress#connect-wordpress-docker-container-to-an-existing-database
##
wordpressSkipInstall: false
## @param wordpressExtraConfigContent Add extra content to the default wp-config.php file
## e.g:
## wordpressExtraConfigContent: |
##   @ini_set( 'post_max_size', '128M');
##   @ini_set( 'memory_limit', '256M' );
##
wordpressExtraConfigContent:
## @param wordpressConfiguration The content for your custom wp-config.php file (advanced feature)
## NOTE: This will override configuring WordPress based on environment variables (including those set by the chart)
## NOTE: Currently only supported when `wordpressSkipInstall=true`
##
wordpressConfiguration:
## @param existingWordPressConfigurationSecret The name of an existing secret with your custom wp-config.php file (advanced feature)
## NOTE: When it's set the `wordpressConfiguration` parameter is ignored
##
existingWordPressConfigurationSecret:
## @param wordpressConfigureCache Enable W3 Total Cache plugin and configure cache settings
## NOTE: useful if you deploy Memcached for caching database queries or you use an external cache server
##
wordpressConfigureCache: false
## @param wordpressAutoUpdateLevel Level of auto-updates to allow. Allowed values: `major`, `minor` or `none`.
##
wordpressAutoUpdateLevel: none
## @param wordpressPlugins Array of plugins to install and activate. Can be specified as `all` or `none`.
## NOTE: If set to all, only plugins that are already installed will be activated, and if set to none, no plugins will be activated
##
wordpressPlugins: none
## @param apacheConfiguration The content for your custom httpd.conf file (advanced feature)
##
apacheConfiguration:
## @param existingApacheConfigurationConfigMap The name of an existing secret with your custom httpd.conf file (advanced feature)
## NOTE: When it's set the `apacheConfiguration` parameter is ignored
##
existingApacheConfigurationConfigMap:
## @param customPostInitScripts Custom post-init.d user scripts
## ref: https://github.com/bitnami/bitnami-docker-wordpress/tree/master/5/debian-10/rootfs/post-init.d
## NOTE: supported formats are `.sh`, `.sql` or `.php`
## NOTE: scripts are exclusively executed during the 1st boot of the container
## e.g:
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
customPostInitScripts: {}
## SMTP mail delivery configuration
## ref: https://github.com/bitnami/bitnami-docker-wordpress/#smtp-configuration
## @param smtpHost SMTP server host
## @param smtpPort SMTP server port
## @param smtpUser SMTP username
## @param smtpPassword SMTP user password
## @param smtpProtocol SMTP protocol
##
smtpHost: ""
smtpPort: ""
smtpUser: ""
smtpPassword: ""
smtpProtocol: ""
## @param smtpExistingSecret The name of an existing secret with SMTP credentials
## NOTE: Must contain key `smtp-password`
## NOTE: When it's set, the `smtpPassword` parameter is ignored
##
smtpExistingSecret:
## @param allowEmptyPassword Allow the container to be started with blank passwords
##
allowEmptyPassword: true
## @param allowOverrideNone Configure Apache to prohibit overriding directives with htaccess files
##
allowOverrideNone: false
## @param htaccessPersistenceEnabled Persist custom changes on htaccess files
## If `allowOverrideNone` is `false`, it will persist `/opt/bitnami/wordpress/wordpress-htaccess.conf`
## If `allowOverrideNone` is `true`, it will persist `/opt/bitnami/wordpress/.htaccess`
##
htaccessPersistenceEnabled: false
## @param customHTAccessCM The name of an existing ConfigMap with custom htaccess rules
## NOTE: Must contain key `wordpress-htaccess.conf` with the file content
## NOTE: Requires setting `allowOverrideNone=false`
##
customHTAccessCM:
## @param command Override default container command (useful when using custom images)
##
command: []
## @param args Override default container args (useful when using custom images)
##
args: []
## @param extraEnvVars Array with extra environment variables to add to the WordPress container
## e.g:
## extraEnvVars:
##   - name: FOO
##     value: "bar"
##
extraEnvVars: []
## @param extraEnvVarsCM Name of existing ConfigMap containing extra env vars
##
extraEnvVarsCM:
## @param extraEnvVarsSecret Name of existing Secret containing extra env vars
##
extraEnvVarsSecret:
## @section WordPress Multisite Configuration parameters
## ref: https://github.com/bitnami/bitnami-docker-wordpress#multisite-configuration

## @param multisite.enable Whether to enable WordPress Multisite configuration.
## @param multisite.host WordPress Multisite hostname/address. This value is mandatory when enabling Multisite mode.
## @param multisite.networkType WordPress Multisite network type to enable. Allowed values: `subfolder`, `subdirectory` or `subdomain`.
## @param multisite.enableNipIoRedirect Whether to enable IP address redirection to nip.io wildcard DNS. Useful when running on an IP address with subdomain network type.
##
multisite:
  enable: false
  host: ""
  networkType: subdomain
  enableNipIoRedirect: false

## @section WordPress deployment parameters

## @param replicaCount Number of WordPress replicas to deploy
## NOTE: ReadWriteMany PVC(s) are required if replicaCount > 1
##
replicaCount: 1
## @param updateStrategy.type WordPress deployment strategy type
## @param updateStrategy.rollingUpdate WordPress deployment rolling update configuration parameters
## ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#strategy
## NOTE: Set it to `Recreate` if you use a PV that cannot be mounted on multiple pods
## e.g:
## updateStrategy:
##  type: RollingUpdate
##  rollingUpdate:
##    maxSurge: 25%
##    maxUnavailable: 25%
##
updateStrategy:
  type: RollingUpdate
  rollingUpdate: {}
## @param schedulerName Alternate scheduler
## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
##
schedulerName:
## @param serviceAccountName ServiceAccount name
##
serviceAccountName: default
## @param hostAliases [array] WordPress pod host aliases
## https://kubernetes.io/docs/concepts/services-networking/add-entries-to-pod-etc-hosts-with-host-aliases/
##
hostAliases:
  ## Required for apache-exporter to work
  - ip: "127.0.0.1"
    hostnames:
      - "status.localhost"
## @param extraVolumes Optionally specify extra list of additional volumes for WordPress pods
##
extraVolumes: []
## @param extraVolumeMounts Optionally specify extra list of additional volumeMounts for WordPress container(s)
##
extraVolumeMounts: []
## @param sidecars Add additional sidecar containers to the WordPress pod
## e.g:
## sidecars:
##   - name: your-image-name
##     image: your-image
##     imagePullPolicy: Always
##     ports:
##       - name: portname
##         containerPort: 1234
##
sidecars: {}
## @param initContainers Add additional init containers to the WordPress pods
## ref: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/
## e.g:
## initContainers:
##  - name: your-image-name
##    image: your-image
##    imagePullPolicy: Always
##    command: ['sh', '-c', 'copy themes and plugins from git and push to /bitnami/wordpress/wp-content. Should work with extraVolumeMounts and extraVolumes']
##
initContainers: {}
## @param podLabels Extra labels for WordPress pods
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
##
podLabels: {}
## @param podAnnotations Annotations for WordPress pods
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
##
podAnnotations: {}
## @param podAffinityPreset Pod affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard`
## ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
##
podAffinityPreset: ""
## @param podAntiAffinityPreset Pod anti-affinity preset. Ignored if `affinity` is set. Allowed values: `soft` or `hard`
## Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
##
podAntiAffinityPreset: soft
## Node affinity preset
## Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
##
nodeAffinityPreset:
  ## @param nodeAffinityPreset.type Node affinity preset type. Ignored if `affinity` is set. Allowed values: `soft` or `hard`
  ##
  type: ""
  ## @param nodeAffinityPreset.key Node label key to match. Ignored if `affinity` is set
  ##
  key: ""
  ## @param nodeAffinityPreset.values Node label values to match. Ignored if `affinity` is set
  ## E.g.
  ## values:
  ##   - e2e-az1
  ##   - e2e-az2
  ##
  values: []
## @param affinity Affinity for pod assignment
## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
## NOTE: podAffinityPreset, podAntiAffinityPreset, and  nodeAffinityPreset will be ignored when it's set
##
affinity: {}
## @param nodeSelector Node labels for pod assignment
## ref: https://kubernetes.io/docs/user-guide/node-selection/
##
nodeSelector: {}
## @param tolerations Tolerations for pod assignment
## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
##
tolerations: []
## WordPress containers' resource requests and limits
## ref: http://kubernetes.io/docs/user-guide/compute-resources/
## @param resources.limits The resources limits for the WordPress container
## @param resources.requests [object] The requested resources for the WordPress container
##
resources:
  limits: {}
  requests:
    memory: 512Mi
    cpu: 300m
## Container ports
## @param containerPorts.http WordPress HTTP container port
## @param containerPorts.https WordPress HTTPS container port
##
containerPorts:
  http: 8080
  https: 8443
## Configure Pods Security Context
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod
## @param podSecurityContext.enabled Enabled WordPress pods' Security Context
## @param podSecurityContext.fsGroup Set WordPress pod's Security Context fsGroup
##
podSecurityContext:
  enabled: true
  fsGroup: 1001
## Configure Container Security Context (only main container)
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container
## @param containerSecurityContext.enabled Enabled WordPress containers' Security Context
## @param containerSecurityContext.runAsUser Set WordPress container's Security Context runAsUser
## @param containerSecurityContext.runAsNonRoot Set WordPress container's Security Context runAsNonRoot
##
containerSecurityContext:
  enabled: true
  runAsUser: 1001
  runAsNonRoot: true
## Configure extra options for WordPress containers' liveness and readiness probes
## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/#configure-probes
## @param livenessProbe.enabled Enable livenessProbe
## @skip livenessProbe.httpGet
## @param livenessProbe.initialDelaySeconds Initial delay seconds for livenessProbe
## @param livenessProbe.periodSeconds Period seconds for livenessProbe
## @param livenessProbe.timeoutSeconds Timeout seconds for livenessProbe
## @param livenessProbe.failureThreshold Failure threshold for livenessProbe
## @param livenessProbe.successThreshold Success threshold for livenessProbe
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
    ## E.g.
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
## @param readinessProbe.enabled Enable readinessProbe
## @skip readinessProbe.httpGet
## @param readinessProbe.initialDelaySeconds Initial delay seconds for readinessProbe
## @param readinessProbe.periodSeconds Period seconds for readinessProbe
## @param readinessProbe.timeoutSeconds Timeout seconds for readinessProbe
## @param readinessProbe.failureThreshold Failure threshold for readinessProbe
## @param readinessProbe.successThreshold Success threshold for readinessProbe
##
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
    ## E.g.
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
## @param customLivenessProbe Custom livenessProbe that overrides the default one
##
customLivenessProbe: {}
## @param customReadinessProbe Custom readinessProbe that overrides the default one
#
customReadinessProbe: {}

## @section Traffic Exposure Parameters

## WordPress service parameters
##
service:
  ## @param service.type WordPress service type
  ##
  type: LoadBalancer
  ## @param service.port WordPress service HTTP port
  ##
  port: 80
  ## @param service.httpsPort WordPress service HTTPS port
  ##
  httpsPort: 443
  ## @param service.httpsTargetPort Target port for HTTPS
  ##
  httpsTargetPort: https
  ## Node ports to expose
  ## @param service.nodePorts.http Node port for HTTP
  ## @param service.nodePorts.https Node port for HTTPS
  ## NOTE: choose port between <30000-32767>
  ##
  nodePorts:
    http:
    https:
  ## @param service.clusterIP WordPress service Cluster IP
  ## e.g.:
  ## clusterIP: None
  ##
  clusterIP:
  ## @param service.loadBalancerIP WordPress service Load Balancer IP
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#type-loadbalancer
  ##
  loadBalancerIP:
  ## @param service.loadBalancerSourceRanges WordPress service Load Balancer sources
  ## ref: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/#restrict-access-for-loadbalancer-service
  ## e.g:
  ## loadBalancerSourceRanges:
  ##   - 10.10.10.0/24
  ##
  loadBalancerSourceRanges: []
  ## @param service.externalTrafficPolicy WordPress service external traffic policy
  ## ref http://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
  ##
  externalTrafficPolicy: Cluster
  ## @param service.annotations Additional custom annotations for WordPress service
  ##
  annotations: {}
  ## @param service.extraPorts Extra port to expose on WordPress service
  ##
  extraPorts: []
## Configure the ingress resource that allows you to access the WordPress installation
## ref: https://kubernetes.io/docs/concepts/services-networking/ingress/
##
ingress:
  ## @param ingress.enabled Enable ingress record generation for WordPress
  ##
  enabled: false
  ## @param ingress.certManager Add the corresponding annotations for cert-manager integration
  ##
  certManager: false
  ## @param ingress.pathType Ingress path type
  ##
  pathType: ImplementationSpecific
  ## @param ingress.apiVersion Force Ingress API version (automatically detected if not set)
  ##
  apiVersion:
  ## @param ingress.hostname Default host for the ingress record
  ##
  hostname: wordpress.local
  ## @param ingress.path Default path for the ingress record
  ## NOTE: You may need to set this to '/*' in order to use this with ALB ingress controllers
  ##
  path: /
  ## @param ingress.annotations Additional custom annotations for the ingress record
  ## NOTE: If `ingress.certManager=true`, annotation `kubernetes.io/tls-acme: "true"` will automatically be added
  ##
  annotations: {}
  ## @param ingress.tls Enable TLS configuration for the host defined at `ingress.hostname` parameter
  ## TLS certificates will be retrieved from a TLS secret with name: `{{- printf "%s-tls" .Values.ingress.hostname }}`
  ## You can:
  ##   - Use the `ingress.secrets` parameter to create this TLS secret
  ##   - Relay on cert-manager to create it by setting `ingress.certManager=true`
  ##   - Relay on Helm to create self-signed certificates by setting `ingress.tls=true` and `ingress.certManager=false`
  ##
  tls: false
  ## @param ingress.extraHosts An array with additional hostname(s) to be covered with the ingress record
  ## e.g:
  ## extraHosts:
  ##   - name: wordpress.local
  ##     path: /
  ##
  extraHosts: []
  ## @param ingress.extraPaths An array with additional arbitrary paths that may need to be added to the ingress under the main host
  ## e.g:
  ## extraPaths:
  ## - path: /*
  ##   backend:
  ##     serviceName: ssl-redirect
  ##     servicePort: use-annotation
  ##
  extraPaths: []
  ## @param ingress.extraTls TLS configuration for additional hostname(s) to be covered with this ingress record
  ## ref: https://kubernetes.io/docs/concepts/services-networking/ingress/#tls
  ## e.g:
  ## extraTls:
  ## - hosts:
  ##     - wordpress.local
  ##   secretName: wordpress.local-tls
  ##
  extraTls: []
  ## @param ingress.secrets Custom TLS certificates as secrets
  ## NOTE: 'key' and 'certificate' are expected in PEM format
  ## NOTE: 'name' should line up with a 'secretName' set further up
  ## If it is not set and you're using cert-manager, this is unneeded, as it will create a secret for you with valid certificates
  ## If it is not set and you're NOT using cert-manager either, self-signed certificates will be created valid for 365 days
  ## It is also possible to create and manage the certificates outside of this helm chart
  ## Please see README.md for more information
  ## e.g:
  ## secrets:
  ##   - name: wordpress.local-tls
  ##     key: |-
  ##       -----BEGIN RSA PRIVATE KEY-----
  ##       ...
  ##       -----END RSA PRIVATE KEY-----
  ##     certificate: |-
  ##       -----BEGIN CERTIFICATE-----
  ##       ...
  ##       -----END CERTIFICATE-----
  ##
  secrets: []

## @section Persistence Parameters

## Persistence Parameters
## ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
##
persistence:
  ## @param persistence.enabled Enable persistence using Persistent Volume Claims
  ##
  enabled: true
  ## @param persistence.storageClass Persistent Volume storage class
  ## If defined, storageClassName: <storageClass>
  ## If set to "-", storageClassName: "", which disables dynamic provisioning
  ## If undefined (the default) or set to null, no storageClassName spec is set, choosing the default provisioner
  ##
  storageClass:
  ## @param persistence.accessModes [array] Persistent Volume access modes
  ##
  accessModes:
    - ReadWriteOnce
  ## @param persistence.accessMode Persistent Volume access mode (DEPRECATED: use `persistence.accessModes` instead)
  ##
  accessMode: ReadWriteOnce
  ## @param persistence.size Persistent Volume size
  ##
  size: 10Gi
  ## @param persistence.dataSource Custom PVC data source
  ##
  dataSource: {}
  ## @param persistence.existingClaim The name of an existing PVC to use for persistence
  ##
  existingClaim:
## 'volumePermissions' init container parameters
## Changes the owner and group of the persistent volume mount point to runAsUser:fsGroup values
##   based on the podSecurityContext/containerSecurityContext parameters
##
volumePermissions:
  ## @param volumePermissions.enabled Enable init container that changes the owner/group of the PV mount point to `runAsUser:fsGroup`
  ##
  enabled: false
  ## Bitnami Shell image
  ## ref: https://hub.docker.com/r/bitnami/bitnami-shell/tags/
  ## @param volumePermissions.image.registry Bitnami Shell image registry
  ## @param volumePermissions.image.repository Bitnami Shell image repository
  ## @param volumePermissions.image.tag Bitnami Shell image tag (immutable tags are recommended)
  ## @param volumePermissions.image.pullPolicy Bitnami Shell image pull policy
  ## @param volumePermissions.image.pullSecrets Bitnami Shell image pull secrets
  ##
  image:
    registry: docker.io
    repository: bitnami/bitnami-shell
    tag: "10"
    pullPolicy: Always
    ## Optionally specify an array of imagePullSecrets.
    ## Secrets must be manually created in the namespace.
    ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
    ## e.g:
    ## pullSecrets:
    ##   - myRegistryKeySecretName
    ##
    pullSecrets: []
  ## Init container's resource requests and limits
  ## ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ## @param volumePermissions.resources.limits The resources limits for the init container
  ## @param volumePermissions.resources.requests The requested resources for the init container
  ##
  resources:
    limits: {}
    requests: {}
  ## Init container Container Security Context
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container
  ## @param volumePermissions.securityContext.runAsUser Set init container's Security Context runAsUser
  ## NOTE: when runAsUser is set to special value "auto", init container will try to chown the
  ##   data folder to auto-determined user&group, using commands: `id -u`:`id -G | cut -d" " -f2`
  ##   "auto" is especially useful for OpenShift which has scc with dynamic user ids (and 0 is not allowed)
  ##
  securityContext:
    runAsUser: 0

## @section Other Parameters

## Wordpress Pod Disruption Budget configuration
## ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/
## @param pdb.create Enable a Pod Disruption Budget creation
## @param pdb.minAvailable Minimum number/percentage of pods that should remain scheduled
## @param pdb.maxUnavailable Maximum number/percentage of pods that may be made unavailable
##
pdb:
  create: false
  minAvailable: 1
  maxUnavailable:
## Wordpress Autoscaling configuration
## ref: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
## @param autoscaling.enabled Enable Horizontal POD autoscaling for WordPress
## @param autoscaling.minReplicas Minimum number of WordPress replicas
## @param autoscaling.maxReplicas Maximum number of WordPress replicas
## @param autoscaling.targetCPU Target CPU utilization percentage
## @param autoscaling.targetMemory Target Memory utilization percentage
##
autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 11
  targetCPU: 50
  targetMemory: 50

## @section Metrics Parameters

## Prometheus Exporter / Metrics configuration
##
metrics:
  ## @param metrics.enabled Start a sidecar prometheus exporter to expose metrics
  ##
  enabled: false
  ## Bitnami Apache Exporter image
  ## ref: https://hub.docker.com/r/bitnami/apache-exporter/tags/
  ## @param metrics.image.registry Apache Exporter image registry
  ## @param metrics.image.repository Apache Exporter image repository
  ## @param metrics.image.tag Apache Exporter image tag (immutable tags are recommended)
  ## @param metrics.image.pullPolicy Apache Exporter image pull policy
  ## @param metrics.image.pullSecrets Apache Exporter image pull secrets
  ##
  image:
    registry: docker.io
    repository: bitnami/apache-exporter
    tag: 0.8.0-debian-10-r386
    pullPolicy: IfNotPresent
    ## Optionally specify an array of imagePullSecrets.
    ## Secrets must be manually created in the namespace.
    ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
    ## e.g:
    ## pullSecrets:
    ##   - myRegistryKeySecretName
    ##
    pullSecrets: []
  ## Prometheus exporter container's resource requests and limits
  ## ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ## @param metrics.resources.limits The resources limits for the Prometheus exporter container
  ## @param metrics.resources.requests The requested resources for the Prometheus exporter container
  ##
  resources:
    limits: {}
    requests: {}
  ## Prometheus exporter service parameters
  ##
  service:
    ## @param metrics.service.port Metrics service port
    ##
    port: 9117
    ## @param metrics.service.annotations [object] Additional custom annotations for Metrics service
    ##
    annotations:
      prometheus.io/scrape: "true"
      prometheus.io/port: "{{ .Values.metrics.service.port }}"
  ## Prometheus Service Monitor
  ## ref: https://github.com/coreos/prometheus-operator
  ##      https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#endpoint
  ##
  serviceMonitor:
    ## @param metrics.serviceMonitor.enabled Create ServiceMonitor Resource for scraping metrics using PrometheusOperator
    ##
    enabled: false
    ## @param metrics.serviceMonitor.namespace The namespace in which the ServiceMonitor will be created
    ##
    namespace:
    ## @param metrics.serviceMonitor.interval The interval at which metrics should be scraped
    ##
    interval: 30s
    ## @param metrics.serviceMonitor.scrapeTimeout The timeout after which the scrape is ended
    ##
    scrapeTimeout:
    ## @param metrics.serviceMonitor.relabellings Metrics relabellings to add to the scrape endpoint
    ##
    relabellings:
    ## @param metrics.serviceMonitor.honorLabels Labels to honor to add to the scrape endpoint
    ##
    honorLabels: false
    ## @param metrics.serviceMonitor.additionalLabels Additional custom labels for the ServiceMonitor
    ##
    additionalLabels: {}

## @section Database Parameters

## MariaDB chart configuration
## ref: https://github.com/bitnami/charts/blob/master/bitnami/mariadb/values.yaml
##
mariadb:
  ## @param mariadb.enabled Deploy a MariaDB server to satisfy the applications database requirements
  ## To use an external database set this to false and configure the `externalDatabase.*` parameters
  ##
  enabled: true
  ## @param mariadb.architecture MariaDB architecture. Allowed values: `standalone` or `replication`
  ##
  architecture: standalone
  ## MariaDB Authentication parameters
  ## @param mariadb.auth.rootPassword MariaDB root password
  ## @param mariadb.auth.database MariaDB custom database
  ## @param mariadb.auth.username MariaDB custom user name
  ## @param mariadb.auth.password MariaDB custom user password
  ## ref: https://github.com/bitnami/bitnami-docker-mariadb#setting-the-root-password-on-first-run
  ##      https://github.com/bitnami/bitnami-docker-mariadb/blob/master/README.md#creating-a-database-on-first-run
  ##      https://github.com/bitnami/bitnami-docker-mariadb/blob/master/README.md#creating-a-database-user-on-first-run
  auth:
    rootPassword: ""
    database: bitnami_wordpress
    username: bn_wordpress
    password: ""
  ## MariaDB Primary configuration
  ##
  primary:
    ## MariaDB Primary Persistence parameters
    ## ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
    ## @param mariadb.primary.persistence.enabled Enable persistence on MariaDB using PVC(s)
    ## @param mariadb.primary.persistence.storageClass Persistent Volume storage class
    ## @param mariadb.primary.persistence.accessModes [array] Persistent Volume access modes
    ## @param mariadb.primary.persistence.size Persistent Volume size
    ##
    persistence:
      enabled: true
      storageClass:
      accessModes:
        - ReadWriteOnce
      size: 8Gi
## External Database Configuration
## All of these values are only used if `mariadb.enabled=false`
##
externalDatabase:
  ## @param externalDatabase.host External Database server host
  ##
  host: localhost
  ## @param externalDatabase.port External Database server port
  ##
  port: 3306
  ## @param externalDatabase.user External Database username
  ##
  user: bn_wordpress
  ## @param externalDatabase.password External Database user password
  ##
  password: ""
  ## @param externalDatabase.database External Database database name
  ##
  database: bitnami_wordpress
  ## @param externalDatabase.existingSecret The name of an existing secret with database credentials
  ## NOTE: Must contain key `mariadb-password`
  ## NOTE: When it's set, the `externalDatabase.password` parameter is ignored
  ##
  existingSecret:
## Memcached chart configuration
## ref: https://github.com/bitnami/charts/blob/master/bitnami/memcached/values.yaml
##
memcached:
  ## @param memcached.enabled Deploy a Memcached server for caching database queries
  ##
  enabled: false
  ## Service parameters
  ##
  service:
    ## @param memcached.service.port Memcached service port
    ##
    port: 11211
## External Memcached Configuration
## All of these values are only used if `memcached.enabled=false`
##
externalCache:
  ## @param externalCache.host External cache server host
  ##
  host: localhost
  ## @param externalCache.port External cache server port
  ##
  port: 11211
```

```helmコマンド
$ helm install wordpress bitnami/wordpress --set wordpressUsername=admin --set wordpressPassword=wpp@ss
NAME: wordpress
LAST DEPLOYED: Sun May 23 11:19:16 2021
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

```helmコマンド
$ helm install wordpress bitnami/wordpress --values values.yaml
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

```kubectlコマンド
$ kubectl get secret
NAME                              TYPE                                  DATA   AGE
default-token-dmlps               kubernetes.io/service-account-token   3      2d20h
sh.helm.release.v1.wordpress.v1   helm.sh/release.v1                    1      55s
wordpress                         Opaque                                1      54s
wordpress-mariadb                 Opaque                                2      55s
wordpress-mariadb-token-h8cvk     kubernetes.io/service-account-token   3      55s
```

```helmコマンド
$ helm list
NAME            NAMESPACE       REVISION        UPDATED                                 STATUS          CHART                   APP VERSION
wordpress       default         1               2021-05-23 11:28:10.548192984 +0000 UTC deployed        wordpress-11.0.9        5.7.2
```

```linuxコマンド
$ export SERVICE_IP=$(kubectl get svc --namespace default wordpress --template "{{ range (index .status.loadBalancer.ingress 0) }}{{.}}{{ end }}")
```

```linuxコマンド
$ echo "WordPress Admin URL: http://$SERVICE_IP/admin"
WordPress Admin URL: http://34.84.235.193/admin
```

```kubectlコマンド
$ kubectl get service
NAME                TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)                      AGE
kubernetes          ClusterIP      10.3.240.1     <none>          443/TCP                      2d21h
wordpress           LoadBalancer   10.3.240.221   34.84.235.193   80:31433/TCP,443:31074/TCP   9m10s
wordpress-mariadb   ClusterIP      10.3.253.236   <none>          3306/TCP                     9m10s
```

```linuxコマンド
$ echo Password: $(kubectl get secret --namespace default wordpress -o jsonpath="{.data.wordpress-password}" | base64 --decode)
Password: wpp@ss
```


$ kubectl get pod


$ kubectl get persistentvolume,persistentvolumeclaim


```helmコマンド
$ helm uninstall wordpress
release "wordpress" uninstalled
```

```helmコマンド
$ helm list
NAME    NAMESPACE       REVISION        UPDATED STATUS  CHART   APP VERSION
```

