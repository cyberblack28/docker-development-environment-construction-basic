# 第6章 ローカル開発の準備

## 6.2 Skaffold を利用したローカル開発環境

### 6.2.2 Skaffold事前確認

#### Kubernetes クラスタの構築

```skaffoldコマンド
$ skaffold version
v1.24.1
```

```dockerコマンド
$ docker version
Client: Docker Engine - Community
 Version:           20.10.6
 API version:       1.41
 Go version:        go1.13.15
 Git commit:        370c289
 Built:             Fri Apr  9 22:46:45 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.6
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.13.15
  Git commit:       8728dd2
  Built:            Fri Apr  9 22:44:56 2021
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.4.4
  GitCommit:        05f951a3781f4f2c1911b05e61c160e9c30eaa8e
 runc:
  Version:          1.0.0-rc93
  GitCommit:        12644e614e25b05da6fd08a38ffa0cfe1903fdec
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
  ```

### 6.2.3 サンプルアプリケーションとDockerfileの作成

#### Dockerfile の作成

```linuxコマンド
$ cd ../Chapter06
$ cd 6-2-3-01
```

```linuxコマンド
$ cat main.go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Practice Skaffold!!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

```linuxコマンド
$ cat Dockerfile
#Stage-1
FROM golang:1.16 as builder
COPY ./main.go ./
RUN go build -o /sample-go-app ./main.go

#Stage-2
FROM gcr.io/distroless/base
EXPOSE 8080
COPY --from=builder /sample-go-app /.
ENTRYPOINT ["./sample-go-app"]
```

#### マニフェストの作成

```linuxコマンド
$ vim practice-skaffold-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: practice-skaffold-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: practice-skaffold
  template:
    metadata:
      labels:
        app: practice-skaffold
    spec:
      containers:
      - name: nginx
        image: DOCKERHUB_USER/practice-skaffold
```

```linuxコマンド
$ cat practice-skaffold-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: practice-skaffold-service
spec:
  type: LoadBalancer
  ports:
  - name: http
    protocol: TCP
    port: 80
    targetPort: 8080
  selector:
    app: practice-skaffold
```

#### Skaffoldの設定ファイル作成

```linuxコマンド
$ vim skaffold.yaml
apiVersion: skaffold/v2alpha3
kind: Config
build:
  artifacts:
  - image: cyberblack28/practice-skaffold
    docker:
      dockerfile: ./Dockerfile
  tagPolicy:
    dateTime: {}
  local:
    push: true
deploy:
  kubectl:
    manifests:
    - practice-skaffold-*
```

#### 6.2.4 Skaffoldの実行

```dockerコマンド
$ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID,
head over to https://hub.docker.com to create one.
Username: cyberblack28
Password:
WARNING! Your password will be stored unencrypted in /home/iyutaka2021/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
```

```skaffoldコマンド
$ skaffold dev -f skaffold.yaml
Listing files to watch...
 - cyberblack28/practice-skaffold
Generating tags...
 - cyberblack28/practice-skaffold -> cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC
Checking cache...
 - cyberblack28/practice-skaffold: Not found. Building
Starting build...
Building [cyberblack28/practice-skaffold]...
Sending build context to Docker daemon  3.072kB
Step 1/7 : FROM golang:1.16 as builder
1.16: Pulling from library/golang
d960726af2be: Pulling fs layer
e8d62473a22d: Pulling fs layer
8962bc0fad55: Pulling fs layer
65d943ee54c1: Pulling fs layer
f2253e6fbefa: Pulling fs layer
6d7fa7c7d5d3: Pulling fs layer
e2e442f7f89f: Pulling fs layer
65d943ee54c1: Waiting
f2253e6fbefa: Waiting
6d7fa7c7d5d3: Waiting
e2e442f7f89f: Waiting
e8d62473a22d: Verifying Checksum
e8d62473a22d: Download complete
8962bc0fad55: Verifying Checksum
8962bc0fad55: Download complete
d960726af2be: Verifying Checksum
d960726af2be: Download complete
65d943ee54c1: Verifying Checksum
65d943ee54c1: Download complete
f2253e6fbefa: Verifying Checksum
f2253e6fbefa: Download complete
e2e442f7f89f: Verifying Checksum
e2e442f7f89f: Download complete
6d7fa7c7d5d3: Verifying Checksum
6d7fa7c7d5d3: Download complete
d960726af2be: Pull complete
e8d62473a22d: Pull complete
8962bc0fad55: Pull complete
65d943ee54c1: Pull complete
f2253e6fbefa: Pull complete
6d7fa7c7d5d3: Pull complete
e2e442f7f89f: Pull complete
Digest: sha256:6f0b0a314b158ff6caf8f12d7f6f3a966500ec6afb533e986eca7375e2f7560f
Status: Downloaded newer image for golang:1.16
 ---> 96129f3766cf
Step 2/7 : COPY ./main.go ./
 ---> dfe2711a6618
Step 3/7 : RUN go build -o /sample-go-app ./main.go
 ---> Running in d27696b0ec69
 ---> cbf2081fff0f
Step 4/7 : FROM gcr.io/distroless/base
latest: Pulling from distroless/base
5dea5ec2316d: Pulling fs layer
bb771d6dc9a1: Pulling fs layer
5dea5ec2316d: Verifying Checksum
5dea5ec2316d: Download complete
bb771d6dc9a1: Verifying Checksum
bb771d6dc9a1: Download complete
5dea5ec2316d: Pull complete
bb771d6dc9a1: Pull complete
Digest: sha256:6ec6da1888b18dd971802c2a58a76a7702902b4c9c1be28f38e75e871cedc2df
Status: Downloaded newer image for gcr.io/distroless/base:latest
 ---> a4cf6da932ac
Step 5/7 : EXPOSE 8080
 ---> Running in dbabc63c319e
 ---> bec04a94cb6a
Step 6/7 : COPY --from=builder /sample-go-app /.
 ---> 58a915d510b4
Step 7/7 : ENTRYPOINT ["./sample-go-app"]
 ---> Running in 721a2400fc79
 ---> 513e4aabe403
Successfully built 513e4aabe403
Successfully tagged cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC
The push refers to repository [docker.io/cyberblack28/practice-skaffold]
6412f44e55af: Preparing
5d09c2db1d76: Preparing
417cb9b79ade: Preparing
417cb9b79ade: Pushed
5d09c2db1d76: Pushed
6412f44e55af: Pushed
2021-05-26_16-56-02.521_UTC: digest: sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b size: 949
Starting test...
WARN[0058] Ignoring image referenced by digest: [cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b]
Tags used in deployment:
 - cyberblack28/practice-skaffold -> cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b
Starting deploy...
 - deployment.apps/practice-skaffold-deployment created
 - service/practice-skaffold-service created
Waiting for deployments to stabilize...
 - deployment/practice-skaffold-deployment: creating container nginx
    - pod/practice-skaffold-deployment-5d988cb499-b8rsx: creating container nginx
    - pod/practice-skaffold-deployment-5d988cb499-ggrmr: creating container nginx
    - pod/practice-skaffold-deployment-5d988cb499-vmgr9: creating container nginx
 - deployment/practice-skaffold-deployment is ready.
Deployments stabilized in 8.672 seconds
Press Ctrl+C to exit
Watching for changes...
WARN[0069] Ignoring image referenced by digest: [cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b]
WARN[0069] Ignoring image referenced by digest: [cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b]
WARN[0069] Ignoring image referenced by digest: [cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b]
WARN[0069] Ignoring image referenced by digest: [cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b]
WARN[0069] Ignoring image referenced by digest: [cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b]
WARN[0069] Ignoring image referenced by digest: [cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b]
```

```kubectlコマンド
$ kubectl get deployment
NAME                           READY   UP-TO-DATE   AVAILABLE   AGE
practice-skaffold-deployment   3/3     3            3           3m59s
```

```kubectlコマンド
$ kubectl get pod
NAME                                            READY   STATUS    RESTARTS   AGE
practice-skaffold-deployment-5d988cb499-b8rsx   1/1     Running   0          2m54s
practice-skaffold-deployment-5d988cb499-ggrmr   1/1     Running   0          2m55s
practice-skaffold-deployment-5d988cb499-vmgr9   1/1     Running   0          2m55s
```

```kubectlコマンド
$ kubectl get deployments practice-skaffold-deployment -o jsonpath="{.spec.template.spec.containers[].image}"
cyberblack28/practice-skaffold:2021-05-26_16-56-02.521_UTC@sha256:21ac6053b7624e66cee5be05474c427d26723879ff0ba3690265646fee13739b
```

```kubectlコマンド
$ kubectl get service
NAME                        TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)        AGE
kubernetes                  ClusterIP      10.48.0.1      <none>          443/TCP        96m
practice-skaffold-service   LoadBalancer   10.48.10.160   35.200.102.72   80:30529/TCP   6m2s
```

```linuxコマンド
$ LB_EXIP=$(kubectl get service practice-skaffold-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

```linuxコマンド
$ curl http://${LB_EXIP}
Practice Skaffold!!
```

```linuxコマンド
$ cd docker-development-environment-construction-basic/Chapter06/6-2-3-01
```

```linuxコマンド
$ sed -i -e 's|Practice Skaffold!!|Convenience Skaffold!!|' main.go
```

```kubectlコマンド
$ curl http://${LB_EXIP}
Convenience Skaffold!!
```

```kubectlコマンド
$ kubectl get deployments practice-skaffold-deployment -o jsonpath="{.spec.template.spec.containers[].image}"
cyberblack28/practice-skaffold:d7bed2d-dirty@sha256:37888b93ca6f42fa4f6dd7d6cb97736fc2bf34a6c4cc5fb29c1f0b2fe9a89378
```

### 6.2.5 SkaffoldとHelmの連携

#### テンプレートの作成

```linuxコマンド
$ cd ../6-2-5-01
```

```helmコマンド
$ helm create skaffold-helm
Creating skaffold-helm
```

```linuxコマンド
$ rm -rf skaffold-helm/templates/*
```

```linuxコマンド
$ cp -p helm-yaml/practice-skaffold-deployment.yaml skaffold-helm/templates
```

```linuxコマンド
$ cat skaffold-helm/templates/practice-skaffold-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: practice-skaffold-deployment
spec:
  replicas: {{ .Values.replicas }}
  selector:
    matchLabels:
      app: {{ .Values.label }}
  template:
    metadata:
      labels:
        app: {{ .Values.label }}
    spec:
      containers:
      - name: {{ .Chart.Name }}
        image: {{ .Values.image }}
        imagePullPolicy: {{ .Values.imageConfig.pullPolicy }}
```

```linuxコマンド
$ cp -p helm-yaml/practice-skaffold-service.yaml skaffold-helm/templates
```

```linuxコマンド
$ cat skaffold-helm/templates/practice-skaffold-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: practice-skaffold-service
spec:
  type: {{ .Values.service_type }}
  ports:
  - name: {{ .Chart.Name }}
    protocol: TCP
    port: 80
    targetPort: 8080
  selector:
    app: {{ .Values.label }}
```

#### values.yamlの作成

```linuxコマンド
$ echo "" > skaffold-helm/values.yaml
```

```linuxコマンド
$ cp -p helm-yaml/values.yaml skaffold-helm
```

```linuxコマンド
$ cat skaffold-helm/values.yaml
#Common
label: practice-skaffold

#Deployment
replicas: 3
image:
  repository: docker.io/cyberblack28/
  name: practice-skaffold
imageConfig:
  pullPolicy: IfNotPresent

#Service
service_type: LoadBalancer
```

#### skaffold.yamlの編集

```linuxコマンド
$ cat skaffold.yaml
apiVersion: skaffold/v2beta10
kind: Config
build:
  artifacts:
  - image: docker.io/cyberblack28/practice-skaffold
    docker:
      dockerfile: ./Dockerfile
  tagPolicy:
    dateTime: {}
  local:
    push: true
deploy:
  helm:
    releases:
    - name: practice-skaffold
      chartPath: skaffold-helm
      valuesFiles: [ skaffold-helm/values.yaml ]
      artifactOverrides:
        image: docker.io/cyberblack28/practice-skaffold
```

```linuxコマンド
$ cat main.go
package main
import (
  "fmt"
  "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Skaffold & Helm!!")
}
func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
```

#### skaffoldの実行

```skaffoldコマンド
$ skaffold dev -f skaffold.yaml
Listing files to watch...
 - docker.io/cyberblack28/practice-skaffold
Generating tags...
 - docker.io/cyberblack28/practice-skaffold -> docker.io/cyberblack28/practice-skaffold:2021-05-27_09-37-37.748_UTC
Checking cache...
 - docker.io/cyberblack28/practice-skaffold: Found. Tagging
Starting test...
Tags used in deployment:
 - docker.io/cyberblack28/practice-skaffold -> docker.io/cyberblack28/practice-skaffold:2021-05-27_09-37-37.748_UTC@sha256:cedbecd4937b48b4a7bd7bf3274b748b76f4f6cf8ef1f23707e4b9496ab68eea
Starting deploy...
Helm release practice-skaffold not installed. Installing...
NAME: practice-skaffold
LAST DEPLOYED: Thu May 27 09:37:43 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
Waiting for deployments to stabilize...
 - deployment/practice-skaffold-deployment: waiting for rollout to finish: 0 of 3 updated replicas are available...
 - deployment/practice-skaffold-deployment is ready.
Deployments stabilized in 13.376 seconds
Press Ctrl+C to exit
Watching for changes...
```

```linuxコマンド
$ LB_EXIP=$(kubectl get service practice-skaffold-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

```linuxコマンド
$ curl http://${LB_EXIP}
Skaffold & Helm!!
```

```kubectlコマンド
$ kubectl get deployments practice-skaffold-deployment -o jsonpath="{.spec.template.spec.containers[].image}"
docker.io/cyberblack28/practice-skaffold:2021-05-27_09-37-37.748_UTC@sha256:cedbecd4937b48b4a7bd7bf3274b748b76f4f6cf8ef1f23707e4b9496ab68eea
```


```linuxコマンド
Listing files to watch...
 - docker.io/cyberblack28/practice-skaffold
Generating tags...
 - docker.io/cyberblack28/practice-skaffold -> docker.io/cyberblack28/practice-skaffold:2021-05-27_09-37-37.748_UTC
Checking cache...
 - docker.io/cyberblack28/practice-skaffold: Found. Tagging
Starting test...
Tags used in deployment:
 - docker.io/cyberblack28/practice-skaffold -> docker.io/cyberblack28/practice-skaffold:2021-05-27_09-37-37.748_UTC@sha256:cedbecd4937b48b4a7bd7bf3274b748b76f4f6cf8ef1f23707e4b9496ab68eea
Starting deploy...
Helm release practice-skaffold not installed. Installing...
NAME: practice-skaffold
LAST DEPLOYED: Thu May 27 09:37:43 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
Waiting for deployments to stabilize...
 - deployment/practice-skaffold-deployment: waiting for rollout to finish: 0 of 3 updated replicas are available...
 - deployment/practice-skaffold-deployment is ready.
Deployments stabilized in 13.376 seconds
Press Ctrl+C to exit
Watching for changes...

Cleaning up...
release "practice-skaffold" uninstalled
There is a new version (1.25.0) of Skaffold available. Download it from:
  https://github.com/GoogleContainerTools/skaffold/releases/tag/v1.25.0
```

