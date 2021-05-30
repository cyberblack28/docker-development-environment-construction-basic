# 第7章 コンテナアプリケーション開発におけるCI/CD

## 7.4 CI環境構築

### 7.4.1 GitHubとGitHub Actionsの環境構築

#### Codeリポジトリの作成

```linuxコマンド
$ cd ../../Chapter07/7-4-1-01
```

```linuxコマンド
$ cat code/app/main.go
package main

import (
  "fmt"
  "net/http"
  )

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello GitOps!!")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
```

```linuxコマンド
$ cat code/app/Dockerfile
#Stage-1
FROM golang:1.16 as builder
COPY ./app/main.go ./
RUN go build -o /gitops-go-app ./main.go

#Stage-2
FROM gcr.io/distroless/base
EXPOSE 8080
COPY --from=builder /gitops-go-app /.
ENTRYPOINT ["./gitops-go-app"]
```

```linuxコマンド
$ cd code
```

```gitコマンド
$ git init
Initialized empty Git repository in /home/iyutaka2021/container-develop-environment-construction-guide/Chapter07/7-4-1-01/code/.git/
```

```gitコマンド
$ git config --global user.email "you@example.com"
```

```gitコマンド
$ git config --global user.name "Your Name"
```

```gitコマンド
$ git add .
```

```gitコマンド
$ git commit -m "first commit"
[master (root-commit) 27593fc] first commit
 2 files changed, 25 insertions(+)
 create mode 100644 app/Dockerfile
 create mode 100644 app/main.go
```

```gitコマンド
$ git branch -M main
```

```gitコマンド
$ git remote add origin https://github.com/cyberblack28/code.git
```

```gitコマンド
$ git push -u origin main
Username for 'https://github.com': cyberblack28
Password for 'https://cyberblack28@github.com':
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 4 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (5/5), 616 bytes | 616.00 KiB/s, done.
Total 5 (delta 0), reused 0 (delta 0)
To https://github.com/cyberblack28/code.git
 * [new branch]      main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
```

```gitコマンド
$ git config --global credential.helper 'cache --timeout=28800'
```

#### Configリポジトリの作成

```linuxコマンド
$ cd ../
```

```helmコマンド
$ mkdir config
```

```helmコマンド
$ cd config
```

```helmコマンド
$ helm create gitops-helm
Creating gitops-helm
```

```linuxコマンド
$ rm -rf gitops-helm/templates/*
```

```linuxコマンド
$ cp -p ../helm-yaml/gitops-deployment.yaml gitops-helm/templates
```

```linuxコマンド
$ cat  gitops-helm/templates/gitops-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gitops-deployment
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
        image: {{ .Values.image.repository }}:{{ .Values.image.tag }}
        imagePullPolicy: {{ .Values.imageConfig.pullPolicy }}
```

```linuxコマンド
$ cp -p ../helm-yaml/gitops-service.yaml gitops-helm/templates
```

```linuxコマンド
$ cat  gitops-helm/templates/gitops-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: gitops-service
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
$ cp -p ../helm-yaml/values.yaml gitops-helm
```

```linuxコマンド
$ cat  gitops-helm/values.yaml
#Common
label: giops

#Deployment
replicas: 3
image:
  repository: docker.io/cyberblack28/gitops-go-app
  tag: 0

imageConfig:
  pullPolicy: IfNotPresent

#Service
service_type: LoadBalancer
```

```gitコマンド
$ git init
Initialized empty Git repository in /home/iyutaka2021/container-develop-environment-construction-guide/Chapter07/7-4-1-01/config/.git/
```

```gitコマンド
$ git add .
```

```gitコマンド
$ git commit -m "first commit"
[master (root-commit) a3d53a2] first commit
 5 files changed, 92 insertions(+)
 create mode 100644 gitops-helm/.helmignore
 create mode 100644 gitops-helm/Chart.yaml
 create mode 100644 gitops-helm/templates/gitops-deployment.yaml
 create mode 100644 gitops-helm/templates/gitops-service.yaml
 create mode 100644 gitops-helm/values.yaml
```

```gitコマンド
$ git branch -M main
```

```gitコマンド
$ git remote add origin https://github.com/cyberblack28/config.git
```

```gitコマンド
$ git push -u origin main
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 4 threads
Compressing objects: 100% (8/8), done.
Writing objects: 100% (9/9), 1.72 KiB | 1.72 MiB/s, done.
Total 9 (delta 0), reused 0 (delta 0)
To https://github.com/cyberblack28/config.git
 * [new branch]      main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
```

#### ローカルリポジトリとの同期

```linuxコマンド
$ cd ../code
```

```gitコマンド
$ git pull
remote: Enumerating objects: 6, done.
remote: Counting objects: 100% (6/6), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
Unpacking objects: 100% (5/5), done.
From https://github.com/cyberblack28/code
   27593fc..4386743  main       -> origin/main
Updating 27593fc..4386743
Fast-forward
 .github/workflows/main.yml | 36 ++++++++++++++++++++++++++++++++++++
 1 file changed, 36 insertions(+)
 create mode 100644 .github/workflows/main.yml
```

```linuxコマンド
$ ls -la
total 20
drwxr-xr-x 5 iyutaka2021 iyutaka2021 4096 May 30 11:11 .
drwxr-xr-x 6 iyutaka2021 iyutaka2021 4096 May 30 11:05 ..
drwxr-xr-x 2 iyutaka2021 iyutaka2021 4096 May 30 10:57 app
drwxr-xr-x 8 iyutaka2021 iyutaka2021 4096 May 30 11:11 .git
drwxr-xr-x 3 iyutaka2021 iyutaka2021 4096 May 30 11:11 .github
```

### 7.4.2 main.ymlの作成手順

#### main.ymlの作成

```linuxコマンド
$ cp -p ../github-actions/main.yml .github/workflows/
```

```linuxコマンド
$ cat .github/workflows/main.yml
name: GitHub Actions CI

on:
  push:
    branches: [ main ]

jobs:
  build:
    name: GitOps Workflow
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2

        # イメージビルド
      - name: Build an image from Dockerfile
        run: |
          # Dockerビルド
          DOCKER_BUILDKIT=1 docker build . -f app/Dockerfile --tag ${{ secrets.USERNAME }}/gitops-go-app:${{ github.run_number }}

        # Trivyによるイメージスキャン
      - name: Run Trivy vulnerability scanner
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: '${{ secrets.USERNAME }}/gitops-go-app:${{ github.run_number }}'
          format: 'table'
          exit-code: '1'
          ignore-unfixed: true
          severity: 'CRITICAL,HIGH'

        # イメージをDocker Hubにプッシュ
      - name: Push image to Docker Hub
        run: |
          # Docker Hub ログイン
          docker login docker.io --username ${{ secrets.USERNAME }} --password ${{ secrets.DOCKER_PASSWORD }}
          # イメージプッシュ
          docker push ${{ secrets.USERNAME }}/gitops-go-app:${{ github.run_number }}

        # values.yamlの更新、新規ブランチ作成、プッシュ、プルリクエスト
      - name: Update values.yaml & Pull Request to Config Repository
        run: |
          # GitHubログイン
          echo -e "machine github.com\nlogin ${{ secrets.USERNAME }}\npassword ${{ secrets.GH_PASSWORD }}" > ~/.netrc
          # 「config」リポジトリからクローン
          git clone https://github.com/${{ secrets.USERNAME }}/config.git
          # values.yamlファイルの更新処理
          cd config/gitops-helm
          git config --global user.email "${{ secrets.EMAIL }}"
          git config --global user.name "${{ secrets.USERNAME }}"
          # 新規ブランチ作成
          git branch feature/${{ github.run_number }}
          git checkout feature/${{ github.run_number }}
          # values.yamlのタグ番号を更新
          sed -i 's/tag: [0-9]*/tag: ${{ github.run_number }}/g' values.yaml
          # プッシュ処理
          git add values.yaml
          git commit -m "Update tag ${{ github.run_number }}"
          git push origin feature/${{ github.run_number }}
          # プルリクエスト処理
          echo ${{ secrets.PERSONAL_ACCESS_TOKEN }} > token.txt
          gh auth login --with-token < token.txt
          gh pr create  --title "Update Tag ${{ github.run_number }}" --body "Please Merge !!"
```

```gitコマンド
$ git add .
```

```gitコマンド
$ git commit -m "create main.yml"
[main 65c69be] create main.yml
 1 file changed, 62 insertions(+), 36 deletions(-)
 rewrite .github/workflows/main.yml (90%)
```

```gitコマンド
$ git branch -M main
```

```gitコマンド
$ git push -u origin main
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 4 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (5/5), 1.39 KiB | 1.39 MiB/s, done.
Total 5 (delta 0), reused 0 (delta 0)
To https://github.com/cyberblack28/code.git
   4386743..65c69be  main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
```

## 7.5 CD環境構築

### 7.5.2 Argo CDのインストール

#### Argo CD Clientのインストール

```linuxコマンド
$ VERSION=$(curl --silent "https://api.github.com/repos/argoproj/argo-cd/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')
```

```linuxコマンド
$ sudo curl -sSL -o /usr/local/bin/argocd https://github.com/argoproj/argo-cd/releases/download/$VERSION/argocd-linux-amd64
```

```linuxコマンド
$ sudo chmod +x /usr/local/bin/argocd
```

```argocdコマンド
$ argocd version
argocd: v2.0.3+8d2b13d
  BuildDate: 2021-05-27T17:38:37Z
  GitCommit: 8d2b13d733e1dff7d1ad2c110ed31be4804406e2
  GitTreeState: clean
  GoVersion: go1.16
  Compiler: gc
  Platform: linux/amd64
FATA[0000] Argo CD server address unspecified
```

#### Argo CDサーバのインストール

```helmコマンド
$ helm repo add argo https://argoproj.github.io/argo-helm
"argo" has been added to your repositories
```

```helmコマンド
$ helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "argo" chart repository
...Successfully got an update from the "bitnami" chart repository
Update Complete. ⎈Happy Helming!⎈
```

```helmコマンド
$ helm search repo argocd
NAME                            CHART VERSION   APP VERSION     DESCRIPTION
argo/argocd-applicationset      0.1.5           v0.1.0          A Helm chart for installing ArgoCD ApplicationSet
argo/argocd-notifications       1.3.2           1.1.1           A Helm chart for ArgoCD notifications, an add-o...
argo/argo-cd                    3.6.4           2.0.3           A Helm chart for ArgoCD, a declarative, GitOps ...
```

```kubectlコマンド
$ kubectl create namespace argocd
namespace/argocd created
```

```kubectlコマンド
$ kubectl get ns argocd
NAME     STATUS   AGE
argocd   Active   32s
```

```helmコマンド
$ helm install argo-cd -n argocd argo/argo-cd --version 3.6.4
manifest_sorter.go:192: info: skipping unknown hook: "crd-install"
manifest_sorter.go:192: info: skipping unknown hook: "crd-install"
NAME: argo-cd
LAST DEPLOYED: Sun May 30 11:44:09 2021
NAMESPACE: argocd
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
In order to access the server UI you have the following options:

1. kubectl port-forward service/argo-cd-argocd-server -n argocd 8080:443

    and then open the browser on http://localhost:8080 and accept the certificate

2. enable ingress in the values file `server.ingress.enabled` and either
      - Add the annotation for ssl passthrough: https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/ingress.md#option-1-ssl-passthrough
      - Add the `--insecure` flag to `server.extraArgs` in the values file and terminate SSL at your ingress: https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/ingress.md#option-2-multiple-ingress-objects-and-hosts


After reaching the UI the first time you can login with username: admin and the random password generated during the installation. You can find the password by running:

kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

(You should delete the initial secret afterwards as suggested by the Getting Started Guide: https://github.com/argoproj/argo-cd/blob/master/docs/getting_started.md#4-login-using-the-cli)
```

#### デプロイ状況の確認

```kubectlコマンド
$ kubectl get pods,services -n argocd
NAME                                                         READY   STATUS    RESTARTS   AGE
pod/argo-cd-argocd-application-controller-676599564b-25tp4   1/1     Running   0          55s
pod/argo-cd-argocd-dex-server-7b7d8d89b7-krqg8               1/1     Running   0          55s
pod/argo-cd-argocd-redis-685b9cf4b9-rtzsx                    1/1     Running   0          55s
pod/argo-cd-argocd-repo-server-5dd97dfcfd-f84ql              1/1     Running   0          55s
pod/argo-cd-argocd-server-796fcf56d4-cx9mp                   1/1     Running   0          55s

NAME                                            TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)             AGE
service/argo-cd-argocd-application-controller   ClusterIP   10.48.11.182   <none>        8082/TCP            56s
service/argo-cd-argocd-dex-server               ClusterIP   10.48.6.215    <none>        5556/TCP,5557/TCP   56s
service/argo-cd-argocd-redis                    ClusterIP   10.48.11.229   <none>        6379/TCP            57s
service/argo-cd-argocd-repo-server              ClusterIP   10.48.11.14    <none>        8081/TCP            57s
service/argo-cd-argocd-server                   ClusterIP   10.48.6.160    <none>        80/TCP,443/TCP      56s
```

#### Argo CD GUIの設定

```kubectlコマンド
$ kubectl patch service argo-cd-argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
service/argo-cd-argocd-server patched
```

```kubectlコマンド
$ kubectl get service -n argocd
NAME                                    TYPE           CLUSTER-IP     EXTERNAL-IP     PORT(S)                      AGE
argo-cd-argocd-application-controller   ClusterIP      10.48.11.182   <none>          8082/TCP                     20m
argo-cd-argocd-dex-server               ClusterIP      10.48.6.215    <none>          5556/TCP,5557/TCP            20m
argo-cd-argocd-redis                    ClusterIP      10.48.11.229   <none>          6379/TCP                     20m
argo-cd-argocd-repo-server              ClusterIP      10.48.11.14    <none>          8081/TCP                     20m
argo-cd-argocd-server                   LoadBalancer   10.48.6.160    34.84.196.132   80:30442/TCP,443:30360/TCP   20m
```

```kubectlコマンド
$ kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
inKunXwanIjNhqKs
```

```argocdコマンド
$ argocd --insecure login 34.84.196.132 --username admin
Password:
'admin:login' logged in successfully
Context '34.84.196.132' updated
```

```argocdコマンド
$ argocd  account update-password --account admin
*** Enter current password:
*** Enter new password:
*** Confirm new password:
Password updated
Context '34.84.196.132' updated
```

## 7.6 GitOpsの実行

```linuxコマンド
$ LB_EXIP=$(kubectl get service gitops-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

```linuxコマンド
$ curl http://${LB_EXIP}
Hello GitOps!!
```

```kubectlコマンド
$ kubectl get service gitops-service
NAME             TYPE           CLUSTER-IP    EXTERNAL-IP    PORT(S)        AGE
gitops-service   LoadBalancer   10.48.6.196   34.84.23.176   80:32756/TCP   2m24s
```

```linuxコマンド
$ vim app/main.go
package main

import (
        "fmt"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello Argo CD!!")
}

func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8080", nil)
}
```

```gitコマンド
$ git add .
```

```gitコマンド
$ git commit -m "Hello Argo CD"
git commit -m "Hello Argo CDgit add ."
[main b1b92e4] Hello Argo CDgit add .
 1 file changed, 2 insertions(+), 2 deletions(-)
```

```gitコマンド
$ git branch -M main
```

```gitコマンド
$ git push -u origin main
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 4 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 404 bytes | 404.00 KiB/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
To https://github.com/cyberblack28/code.git
   65c69be..b1b92e4  main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
```

```linuxコマンド
$ curl http://${LB_EXIP}
Hello Argo CD!!
```

```linuxコマンド
$ curl http://${LB_EXIP}
Hello GitOps!!
```

```linuxコマンド
$ curl http://${LB_EXIP}
Hello Argo CD!!
```