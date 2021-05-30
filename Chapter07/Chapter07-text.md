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
[master (root-commit) 7962980] first commit
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
[master (root-commit) 100b5d2] first commit
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
   7962980..2baf2ed  main       -> origin/main
Updating 7962980..2baf2ed
Fast-forward
 .github/workflows/main.yml | 36 ++++++++++++++++++++++++++++++++++++
 1 file changed, 36 insertions(+)
 create mode 100644 .github/workflows/main.yml
```

```linuxコマンド
$ ls -la
total 20
drwxr-xr-x 5 iyutaka2021 iyutaka2021 4096 May 30 08:20 .
drwxr-xr-x 6 iyutaka2021 iyutaka2021 4096 May 30 08:03 ..
drwxr-xr-x 2 iyutaka2021 iyutaka2021 4096 May 30 07:45 app
drwxr-xr-x 8 iyutaka2021 iyutaka2021 4096 May 30 08:20 .git
drwxr-xr-x 3 iyutaka2021 iyutaka2021 4096 May 30 08:20 .github
```

### 7.4.2 main.ymlの作成手順

#### main.ymlの作成

```linuxコマンド
$ cp -p ../github-actions/main.yml .github/workflows/
```

```linuxコマンド
$ cat config/.github/workflows/main.yml
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
[main b6a9802] create main.yml
 1 file changed, 65 insertions(+), 36 deletions(-)
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
Writing objects: 100% (5/5), 1.43 KiB | 1.43 MiB/s, done.
Total 5 (delta 0), reused 0 (delta 0)
To https://github.com/cyberblack28/code.git
   2baf2ed..b6a9802  main -> main
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
```

#### Argo CDサーバのインストール

```helmコマンド
$ helm repo add argo https://argoproj.github.io/argo-helm
```

```helmコマンド
$ helm repo update
```

```helmコマンド
$ helm search repo argocd
```

```kubectlコマンド
$ kubectl create namespace argocd
```

```kubectlコマンド
$ kubectl get ns argocd
```

```helmコマンド
$ helm install argo-cd -n argocd argo/argo-cd --version 2.11.0
```

#### デプロイ状況の確認

```kubectlコマンド
$ kubectl get pods,services -n argocd
```

#### Argo CD GUIの設定

```kubectlコマンド
$ kubectl patch service argo-cd-argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```

```kubectlコマンド
$ kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2
```

```argocdコマンド
$ argocd --insecure login 35.200.124.246 --username admin
```

```argocdコマンド
$ argocd  account update-password --account admin
```

## 7.6 GitOpsの実行

```linuxコマンド
$ LB_EXIP=$(kubectl get service gitops-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

```linuxコマンド
$ curl http://${LB_EXIP}
```

```kubectlコマンド
$ kubectl get service gitops-service
```

```linuxコマンド
$ cd code
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
$ git commit -m "Hello Argo CD!!"
```

```gitコマンド
$ git branch -M main
```

```gitコマンド
$ git branch -M main
```

```linuxコマンド
$ curl http://${LB_EXIP}
```

```linuxコマンド
$ curl http://${LB_EXIP}
```

```linuxコマンド
$ curl http://${LB_EXIP}
```