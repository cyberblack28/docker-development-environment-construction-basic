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
