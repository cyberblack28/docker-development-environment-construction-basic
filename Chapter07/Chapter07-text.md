# 第7章 コンテナアプリケーション開発におけるCI/CD

## 7.4 CI環境構築

### 7.4.1 GitHubとGitHub Actionsの環境構築

#### Codeリポジトリの作成

コマンド
```
cd
```

コマンド
```
cd docker-development-environment-construction-basic/Chapter07/7-4-1-01
```

コマンド
```
cat code/app/main.go
```
コマンド結果
```
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

コマンド
```
cat code/app/Dockerfile
```
コマンド結果
```
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

コマンド
```
export GITCONFIG_EMAIL="you@example.com"
```

コマンド
```
export GITCONFIG_USER_NAME="User Name"
```

コマンド
```
cd code
```

コマンド
```
git init
```
コマンド結果
```
Initialized empty Git repository in /home/iyutaka2021/docker-development-environment-construction-basic/Chapter07/7-4-1-01/code/.git/
```

コマンド
```
git config --global user.email "${GITCONFIG_EMAIL}"
```

コマンド
```
git config --global user.name "${GITCONFIG_USER_NAME}"
```

コマンド
```
git add .
```

コマンド
```
git commit -m "first commit"
```
コマンド結果
```
[master (root-commit) 2f0ec0b] first commit
 2 files changed, 25 insertions(+)
 create mode 100644 app/Dockerfile
 create mode 100644 app/main.go
```

コマンド
```
git branch -M main
```

コマンド
```
git remote add origin https://github.com/${GITCONFIG_USER_NAME}/code.git
```

コマンド
```
git push -u origin main
```
コマンド結果
```
Username for 'https://github.com': cyberblack28
Password for 'https://cyberblack28@github.com':
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 4 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (5/5), 615 bytes | 615.00 KiB/s, done.
Total 5 (delta 0), reused 0 (delta 0)
To https://github.com/cyberblack28/code.git
 * [new branch]      main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
```

#### Configリポジトリの作成

コマンド
```
cd ../
```

コマンド
```
mkdir config
```

コマンド
```
cd config
```

コマンド
```
helm create gitops-helm
```
コマンド結果
```
Creating gitops-helm
```

コマンド
```
rm -rf gitops-helm/templates/*
```

コマンド
```
cp -p ../helm-yaml/gitops-deployment.yaml gitops-helm/templates
```

コマンド
```
cat  gitops-helm/templates/gitops-deployment.yaml
```
コマンド結果
```
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

コマンド
```
cp -p ../helm-yaml/gitops-service.yaml gitops-helm/templates
```

コマンド
```
cat  gitops-helm/templates/gitops-service.yaml
```
コマンド結果
```
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

コマンド
```
cp -p ../helm-yaml/values.yaml gitops-helm
```

コマンド
```
vim gitops-helm/values.yaml
```
コマンド結果
```
#Common
label: gitops

#Deployment
replicas: 3
image:
  repository: docker.io/DOCKERHUB_REPO_NAME/gitops-go-app # 変更箇所 DOCKERHUB_REPO_NAME
  tag: 0

imageConfig:
  pullPolicy: IfNotPresent

#Service
service_type: LoadBalancer
```

コマンド
```
git init
```
コマンド結果
```
Initialized empty Git repository in /home/iyutaka2021/docker-development-environment-construction-basic/Chapter07/7-4-1-01/config/.git/
```

コマンド
```
git add .
```

コマンド
```
git commit -m "first commit"
```
コマンド結果
```
[master (root-commit) 5395a9a] first commit
 5 files changed, 92 insertions(+)
 create mode 100644 gitops-helm/.helmignore
 create mode 100644 gitops-helm/Chart.yaml
 create mode 100644 gitops-helm/templates/gitops-deployment.yaml
 create mode 100644 gitops-helm/templates/gitops-service.yaml
 create mode 100644 gitops-helm/values.yaml
```

コマンド
```
git branch -M main
```

コマンド
```
git remote add origin https://github.com/${GITCONFIG_USER_NAME}/config.git
```

コマンド
```
git push -u origin main
```
コマンド結果
```
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

コマンド
```
cd ../code
```

コマンド
```
git pull
```
コマンド結果
```
remote: Enumerating objects: 6, done.
remote: Counting objects: 100% (6/6), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
Unpacking objects: 100% (5/5), done.
From https://github.com/cyberblack28/code
   2f0ec0b..311c1ac  main       -> origin/main
Updating 2f0ec0b..311c1ac
Fast-forward
 .github/workflows/main.yml | 36 ++++++++++++++++++++++++++++++++++++
 1 file changed, 36 insertions(+)
 create mode 100644 .github/workflows/main.yml
```

コマンド
```
ls -la
```
コマンド結果
```
total 20
drwxr-xr-x 5 iyutaka2021 iyutaka2021 4096 May 30 14:10 .
drwxr-xr-x 6 iyutaka2021 iyutaka2021 4096 May 30 14:05 ..
drwxr-xr-x 2 iyutaka2021 iyutaka2021 4096 May 30 14:00 app
drwxr-xr-x 8 iyutaka2021 iyutaka2021 4096 May 30 14:10 .git
drwxr-xr-x 3 iyutaka2021 iyutaka2021 4096 May 30 14:10 .github
```

### 7.4.2 main.ymlの作成手順

#### main.ymlの作成

コマンド
```
cp -p ../github-actions/main.yml .github/workflows/
```

コマンド
```
cat .github/workflows/main.yml
```
コマンド結果
```
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
          DOCKER_BUILDKIT=1 docker image build . -f app/Dockerfile --tag ${{ secrets.USERNAME }}/gitops-go-app:${{ github.run_number }}

        # Trivyによるイメージスキャン
      - name: Run Trivy vulnerability scanner
        uses: aquasecurity/trivy-action@master
        with:
          image-ref: '${{ secrets.USERNAME }}/gitops-go-app:${{ github.run_number }}'
          format: 'table'
          exit-code: '1'
          ignore-unfixed: true
          severity: 'CRITICAL,HIGH'

        # コンテナイメージをDocker Hubにプッシュ
      - name: Push image to Docker Hub
        run: |
          # Docker Hub ログイン
          docker login docker.io --username ${{ secrets.USERNAME }} --password ${{ secrets.DOCKER_PASSWORD }}
          # イメージプッシュ
          docker image push ${{ secrets.USERNAME }}/gitops-go-app:${{ github.run_number }}

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

コマンド
```
git add .
```

コマンド
```
git commit -m "create main.yml"
```
コマンド結果
```
[main 842cd35] create main.yml
 1 file changed, 62 insertions(+), 36 deletions(-)
 rewrite .github/workflows/main.yml (90%)
```

コマンド
```
git branch -M main
```

コマンド
```
git push -u origin main
```
コマンド結果
```
Enumerating objects: 9, done.
Counting objects: 100% (9/9), done.
Delta compression using up to 4 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (5/5), 1.38 KiB | 1.38 MiB/s, done.
Total 5 (delta 0), reused 0 (delta 0)
To https://github.com/cyberblack28/code.git
   311c1ac..842cd35  main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
```

## 7.5 CD環境構築

### 7.5.2 Argo CDのインストール

#### Argo CDサーバのインストール

コマンド
```
helm repo add argo https://argoproj.github.io/argo-helm
```
コマンド結果
```
"argo" has been added to your repositories
```

コマンド
```
helm repo update
```
コマンド結果
```
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "argo" chart repository
...Successfully got an update from the "bitnami" chart repository
Update Complete. ⎈Happy Helming!⎈
```

コマンド
```
helm search repo argocd
```
コマンド結果
```
NAME                            CHART VERSION   APP VERSION     DESCRIPTION
argo/argocd-applicationset      0.1.5           v0.1.0          A Helm chart for installing ArgoCD ApplicationSet
argo/argocd-notifications       1.3.2           1.1.1           A Helm chart for ArgoCD notifications, an add-o...
argo/argo-cd                    3.6.4           2.0.3           A Helm chart for ArgoCD, a declarative, GitOps ...
```

コマンド
```
kubectl create namespace argocd
```
コマンド結果
```
namespace/argocd created
```

コマンド
```
kubectl get namespace argocd
```
コマンド結果
```
NAME     STATUS   AGE
argocd   Active   20s
```

コマンド
```
helm install argo-cd -n argocd argo/argo-cd --version 3.6.4
```
コマンド結果
```
manifest_sorter.go:192: info: skipping unknown hook: "crd-install"
manifest_sorter.go:192: info: skipping unknown hook: "crd-install"
NAME: argo-cd
LAST DEPLOYED: Sun May 30 14:26:31 2021
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

コマンド
```
kubectl get pods,services -n argocd
```
コマンド結果
```
NAME                                                         READY   STATUS    RESTARTS   AGE
pod/argo-cd-argocd-application-controller-676599564b-gt2bz   1/1     Running   0          54s
pod/argo-cd-argocd-dex-server-7b7d8d89b7-kmxf5               1/1     Running   0          54s
pod/argo-cd-argocd-redis-685b9cf4b9-bjw84                    1/1     Running   0          54s
pod/argo-cd-argocd-repo-server-5dd97dfcfd-ct5r2              1/1     Running   0          54s
pod/argo-cd-argocd-server-796fcf56d4-xs74h                   1/1     Running   0          54s

NAME                                            TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)             AGE
service/argo-cd-argocd-application-controller   ClusterIP   10.48.13.200   <none>        8082/TCP            54s
service/argo-cd-argocd-dex-server               ClusterIP   10.48.0.207    <none>        5556/TCP,5557/TCP   54s
service/argo-cd-argocd-redis                    ClusterIP   10.48.9.140    <none>        6379/TCP            54s
service/argo-cd-argocd-repo-server              ClusterIP   10.48.8.132    <none>        8081/TCP            54s
service/argo-cd-argocd-server                   ClusterIP   10.48.10.158   <none>        80/TCP,443/TCP      54s
```

#### Argo CD GUIの設定

コマンド
```
kubectl patch service argo-cd-argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```
コマンド結果
```
service/argo-cd-argocd-server patched
```

コマンド
```
kubectl get service -n argocd
```
コマンド結果
```
NAME                                    TYPE           CLUSTER-IP     EXTERNAL-IP    PORT(S)                      AGE
argo-cd-argocd-application-controller   ClusterIP      10.48.13.200   <none>         8082/TCP                     2m55s
argo-cd-argocd-dex-server               ClusterIP      10.48.0.207    <none>         5556/TCP,5557/TCP            2m55s
argo-cd-argocd-redis                    ClusterIP      10.48.9.140    <none>         6379/TCP                     2m55s
argo-cd-argocd-repo-server              ClusterIP      10.48.8.132    <none>         8081/TCP                     2m55s
argo-cd-argocd-server                   LoadBalancer   10.48.10.158   34.xx.xx.xx   80:30177/TCP,443:32249/TCP   2m55s
```

コマンド
```
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```
コマンド結果
```
qAGYKIRoLiMsIEby
```

#### Argo CD Clientのインストール

コマンド
```
sudo curl -sSL -o /usr/local/bin/argocd https://github.com/argoproj/argo-cd/releases/download/$VERSION/argocd-linux-amd64
```

コマンド
```
sudo chmod +x /usr/local/bin/argocd
```

コマンド
```
argocd version
```
コマンド結果
```
argocd: v2.0.3+8d2b13d
  BuildDate: 2021-05-27T17:38:37Z
  GitCommit: 8d2b13d733e1dff7d1ad2c110ed31be4804406e2
  GitTreeState: clean
  GoVersion: go1.16
  Compiler: gc
  Platform: linux/amd64
FATA[0000] Failed to establish connection to 34.84.196.132:443: dial tcp 34.84.196.132:443: connect: connection refused
```

コマンド
```
argocd --insecure login 34.xx.xx.xx --username admin
```
コマンド結果
```
Password:
'admin:login' logged in successfully
Context '34.85.78.201' updated
```

コマンド
```
argocd  account update-password --account admin
```
コマンド結果
```
*** Enter current password:
*** Enter new password:
*** Confirm new password:
Password updated
Context '34.xx.xx.xx' updated
```

## 7.6 GitOpsの実行

コマンド
```
LB_EXIP=$(kubectl get service gitops-service -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

コマンド
```
curl http://${LB_EXIP}
```
コマンド結果
```
Hello GitOps!!
```

コマンド
```
kubectl get service gitops-service
```
コマンド結果
```
NAME             TYPE           CLUSTER-IP    EXTERNAL-IP     PORT(S)        AGE
gitops-service   LoadBalancer   10.48.13.12   34.84.113.170   80:32558/TCP   9m45s
```

コマンド
```
vim app/main.go
```
コマンド結果
```
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

コマンド
```
$ git add .
```

コマンド
```
git commit -m "Hello Argo CD"
```
コマンド結果
```
[main ce79239] Hello Argo CD
 1 file changed, 2 insertions(+), 2 deletions(-)
```

コマンド
```
git branch -M main
```

コマンド
```
git push -u origin main
```
コマンド結果
```
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 4 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 393 bytes | 393.00 KiB/s, done.
Total 4 (delta 1), reused 0 (delta 0)
remote: Resolving deltas: 100% (1/1), completed with 1 local object.
To https://github.com/cyberblack28/code.git
   842cd35..ce79239  main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.
```

コマンド
```
curl http://${LB_EXIP}
```
コマンド結果
```
Hello Argo CD!!
```

コマンド
```
curl http://${LB_EXIP}
```
コマンド結果
```
Hello GitOps!!
```

コマンド
```
curl http://${LB_EXIP}
```
コマンド結果
```
Hello Argo CD!!
```