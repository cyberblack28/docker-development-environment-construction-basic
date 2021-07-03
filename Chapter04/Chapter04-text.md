# 第4章 コンテナオーケストレーション

## 4.4 Kubernetes 環境のセットアップ

### 4.4.1 GKE のセットアップ

#### 【1】GCPプロジェクトの確認

コマンド
```
gcloud projects list
```
コマンド結果
```
PROJECT_ID              NAME              PROJECT_NUMBER
mercurial-shape-278704  My First Project  xxxxxxxxxxxx
```

#### 【2】プロジェクトIDの設定

コマンド
```
gcloud config set project <YOUR-PROJECT-ID>
```
コマンド結果
```
Updated property [core/project].
```

コマンド
```
export PROJECT_ID=$(gcloud config get-value project); echo $PROJECT_ID
```
コマンド結果
```
Your active configuration is: [cloudshell-3891]
mercurial-shape-278704
```

#### 【3】Zoneの設定

コマンド
```
gcloud config set compute/zone asia-northeast1-a
```
コマンド結果
```
Updated property [compute/zone].
```

コマンド
```
export COMPUTE_ZONE=$(gcloud config get-value compute/zone); echo $COMPUTE_ZONE
```
コマンド結果
```
Your active configuration is: [cloudshell-3891]
asia-northeast1-a
```

#### 【4】APIの有効化

コマンド
```
gcloud services enable cloudapis.googleapis.com container.googleapis.com
```
コマンド結果
```
Operation "operations/acf.p2-540200718496-4784cf88-330f-40cb-be37-f9e842b7a148" finished successfully.
```

#### 【5】Kubernetesクラスタの作成

コマンド
```
gcloud container get-server-config
```
コマンド結果
```
Fetching server config for asia-northeast1-a
channels:
- channel: RAPID
  defaultVersion: 1.19.9-gke.1900
  validVersions:
  - 1.20.5-gke.2000
  - 1.19.9-gke.1900
- channel: REGULAR
  defaultVersion: 1.18.17-gke.100
  validVersions:
  - 1.19.9-gke.1400
  - 1.19.8-gke.1600
  - 1.18.17-gke.700
  - 1.18.17-gke.100
- channel: STABLE
  defaultVersion: 1.18.17-gke.100
  validVersions:
  - 1.18.17-gke.100
  - 1.17.17-gke.5400
  - 1.17.17-gke.4900
defaultClusterVersion: 1.18.17-gke.100
defaultImageType: COS
validImageTypes:
- WINDOWS_LTSC
- COS
- UBUNTU
- COS_CONTAINERD
- UBUNTU_CONTAINERD
- WINDOWS_SAC
validMasterVersions:
- 1.19.9-gke.1900
- 1.19.9-gke.1400
- 1.19.8-gke.1600
- 1.18.17-gke.1900
- 1.18.17-gke.1200
- 1.18.17-gke.700
- 1.18.17-gke.100
- 1.18.16-gke.2100
- 1.18.16-gke.1201
- 1.18.16-gke.502
- 1.18.16-gke.302
- 1.18.16-gke.300
- 1.18.15-gke.1502
- 1.18.15-gke.1501
- 1.17.17-gke.7200
- 1.17.17-gke.6700
- 1.17.17-gke.6000
- 1.17.17-gke.5400
- 1.17.17-gke.4900
- 1.17.17-gke.4400
- 1.17.17-gke.3700
validNodeVersions:
- 1.19.9-gke.1900
- 1.19.9-gke.1400
- 1.19.8-gke.1600
- 1.18.17-gke.1900
- 1.18.17-gke.1200
- 1.18.17-gke.700
- 1.18.17-gke.100
- 1.18.16-gke.2100
- 1.18.16-gke.1201
- 1.18.16-gke.1200
- 1.18.16-gke.502
- 1.18.16-gke.500
- 1.18.16-gke.302
- 1.18.16-gke.300
- 1.18.15-gke.2500
- 1.18.15-gke.1502
- 1.18.15-gke.1501
- 1.18.15-gke.1500
- 1.18.15-gke.1102
- 1.18.15-gke.1100
- 1.18.15-gke.800
- 1.18.14-gke.1600
- 1.18.14-gke.1200
- 1.18.12-gke.1210
- 1.18.12-gke.1206
- 1.18.12-gke.1205
- 1.18.12-gke.1201
- 1.18.12-gke.1200
- 1.18.12-gke.300
- 1.18.10-gke.2701
- 1.18.10-gke.2101
- 1.18.10-gke.1500
- 1.18.10-gke.601
- 1.18.9-gke.2501
- 1.18.9-gke.1501
- 1.18.9-gke.801
- 1.18.6-gke.4801
- 1.18.6-gke.3504
- 1.18.6-gke.3503
- 1.17.17-gke.7200
- 1.17.17-gke.6700
- 1.17.17-gke.6000
- 1.17.17-gke.5400
- 1.17.17-gke.4900
- 1.17.17-gke.4400
- 1.17.17-gke.3700
- 1.17.17-gke.3000
- 1.17.17-gke.2800
- 1.17.17-gke.2500
- 1.17.17-gke.1500
- 1.17.17-gke.1101
- 1.17.17-gke.1100
- 1.17.17-gke.600
- 1.17.16-gke.1600
- 1.17.16-gke.1300
- 1.17.15-gke.800
- 1.17.15-gke.300
- 1.17.14-gke.1600
- 1.17.14-gke.1200
- 1.17.14-gke.400
- 1.17.13-gke.2600
- 1.17.13-gke.2001
- 1.17.13-gke.1401
- 1.17.13-gke.1400
- 1.17.13-gke.600
- 1.17.12-gke.2502
- 1.17.12-gke.1504
- 1.17.12-gke.1501
- 1.17.12-gke.500
- 1.17.9-gke.6300
- 1.17.9-gke.1504
- 1.16.15-gke.14800
- 1.16.15-gke.12500
- 1.16.15-gke.11800
- 1.16.15-gke.10600
- 1.16.15-gke.7801
- 1.16.15-gke.7800
- 1.16.15-gke.7300
- 1.16.15-gke.6900
- 1.16.15-gke.6000
- 1.16.15-gke.5500
- 1.16.15-gke.4901
- 1.16.15-gke.4301
- 1.16.15-gke.4300
- 1.16.15-gke.3500
- 1.16.15-gke.2601
- 1.16.15-gke.1600
- 1.16.15-gke.500
- 1.16.13-gke.404
- 1.16.13-gke.403
- 1.16.13-gke.401
- 1.16.13-gke.400
- 1.16.13-gke.1
- 1.16.11-gke.5
- 1.16.10-gke.8
- 1.16.9-gke.6
- 1.16.9-gke.2
- 1.16.8-gke.15
- 1.16.8-gke.12
- 1.16.8-gke.9
- 1.15.12-gke.6002
- 1.15.12-gke.6001
- 1.15.12-gke.5000
- 1.15.12-gke.4002
- 1.15.12-gke.4000
- 1.15.12-gke.20
- 1.15.12-gke.17
- 1.15.12-gke.16
- 1.15.12-gke.13
- 1.15.12-gke.9
- 1.15.12-gke.6
- 1.15.12-gke.3
- 1.15.12-gke.2
- 1.15.11-gke.17
- 1.15.11-gke.15
- 1.15.11-gke.13
- 1.15.11-gke.12
- 1.15.11-gke.11
- 1.15.11-gke.9
- 1.15.11-gke.5
- 1.15.11-gke.3
- 1.15.11-gke.1
- 1.15.9-gke.26
- 1.15.9-gke.24
- 1.15.9-gke.22
- 1.15.9-gke.12
- 1.15.9-gke.9
- 1.15.9-gke.8
- 1.15.8-gke.3
- 1.15.8-gke.2
- 1.15.7-gke.23
- 1.15.7-gke.2
- 1.15.4-gke.22
- 1.14.10-gke.1504
- 1.14.10-gke.902
- 1.14.10-gke.50
- 1.14.10-gke.46
- 1.14.10-gke.45
- 1.14.10-gke.42
- 1.14.10-gke.41
- 1.14.10-gke.40
- 1.14.10-gke.37
- 1.14.10-gke.36
- 1.14.10-gke.34
- 1.14.10-gke.32
- 1.14.10-gke.31
- 1.14.10-gke.27
- 1.14.10-gke.24
- 1.14.10-gke.22
- 1.14.10-gke.21
- 1.14.10-gke.17
- 1.14.10-gke.0
- 1.14.9-gke.23
- 1.14.9-gke.2
- 1.14.9-gke.0
- 1.14.8-gke.33
- 1.14.8-gke.21
- 1.14.8-gke.18
- 1.14.8-gke.17
- 1.14.8-gke.14
- 1.14.8-gke.12
- 1.14.8-gke.7
- 1.14.8-gke.2
- 1.14.7-gke.40
- 1.14.7-gke.25
- 1.14.7-gke.23
- 1.14.7-gke.17
- 1.14.7-gke.14
- 1.14.7-gke.10
- 1.14.6-gke.13
- 1.14.6-gke.2
- 1.14.6-gke.1
- 1.14.3-gke.11
- 1.14.3-gke.10
- 1.14.3-gke.9
- 1.14.2-gke.9
- 1.14.1-gke.5
- 1.13.12-gke.30
- 1.13.12-gke.25
- 1.13.12-gke.17
- 1.13.12-gke.16
- 1.13.12-gke.14
- 1.13.12-gke.13
- 1.13.12-gke.10
- 1.13.12-gke.8
- 1.13.12-gke.4
- 1.13.12-gke.2
- 1.13.11-gke.23
- 1.13.11-gke.15
- 1.13.11-gke.14
- 1.13.11-gke.12
- 1.13.11-gke.11
- 1.13.11-gke.9
- 1.13.11-gke.5
- 1.13.10-gke.7
- 1.13.10-gke.0
- 1.13.9-gke.11
- 1.13.9-gke.3
- 1.13.7-gke.24
- 1.13.7-gke.19
- 1.13.7-gke.15
- 1.13.7-gke.8
- 1.13.7-gke.0
- 1.13.6-gke.13
- 1.13.6-gke.6
- 1.13.6-gke.5
- 1.13.6-gke.0
- 1.13.5-gke.10
- 1.12.10-gke.22
- 1.12.10-gke.20
- 1.12.10-gke.19
- 1.12.10-gke.18
- 1.12.10-gke.17
- 1.12.10-gke.15
- 1.12.10-gke.13
- 1.12.10-gke.11
- 1.12.10-gke.5
- 1.12.9-gke.16
- 1.12.9-gke.15
- 1.12.9-gke.13
- 1.12.9-gke.10
- 1.12.9-gke.7
- 1.12.9-gke.3
- 1.12.8-gke.12
- 1.12.8-gke.10
- 1.12.8-gke.7
- 1.12.8-gke.6
- 1.12.7-gke.26
- 1.12.7-gke.25
- 1.12.7-gke.24
- 1.12.7-gke.22
- 1.12.7-gke.21
- 1.12.7-gke.17
- 1.12.7-gke.10
- 1.12.7-gke.7
- 1.12.6-gke.11
- 1.12.6-gke.10
- 1.12.6-gke.7
- 1.12.5-gke.10
- 1.12.5-gke.5
- 1.11.10-gke.6
- 1.11.10-gke.5
```

コマンド
```
gcloud container clusters create k8s-cluster --zone $COMPUTE_ZONE --cluster-version=1.18.17-gke.100 --async
```
コマンド結果
```
WARNING: Currently VPC-native is not the default mode during cluster creation. In the future, this will become the default mode and can be disabled using `--no-enable-ip-alias` flag. Use `--[no-]enable-ip-alias` flag to suppress this warning.
WARNING: Starting with version 1.18, clusters will have shielded GKE nodes by default.
WARNING: Your Pod address range (`--cluster-ipv4-cidr`) can accommodate at most 1008 node(s).
WARNING: Starting with version 1.19, newly created clusters and node-pools will have COS_CONTAINERD as the default node image when no image type is specified.
NAME         TYPE  LOCATION           TARGET  STATUS_MESSAGE  STATUS        START_TIME  END_TIME
k8s-cluster        asia-northeast1-a                          PROVISIONING
```

#### 【6】Kubernetesクラスタの起動確認

コマンド
```
gcloud container clusters list
```
コマンド結果
```
NAME         LOCATION           MASTER_VERSION   MASTER_IP      MACHINE_TYPE  NODE_VERSION     NUM_NODES  STATUS
k8s-cluster  asia-northeast1-a  1.18.17-gke.100  35.200.13.121  e2-medium     1.18.17-gke.100  3          RUNNING
```

#### 【7】Credentialの取得

コマンド
```
gcloud container clusters get-credentials k8s-cluster --zone $COMPUTE_ZONE --project $PROJECT_ID
```
コマンド結果
```
Fetching cluster endpoint and auth data.
kubeconfig entry generated for k8s-cluster.
```

#### 【8】Node状況の確認

コマンド
```
kubectl get nodes
```
コマンド結果
```
NAME                                         STATUS   ROLES    AGE     VERSION
gke-k8s-cluster-default-pool-18e132c4-5sg2   Ready    <none>   4m47s   v1.18.17-gke.100
gke-k8s-cluster-default-pool-18e132c4-fpr2   Ready    <none>   4m47s   v1.18.17-gke.100
gke-k8s-cluster-default-pool-18e132c4-vnmj   Ready    <none>   4m46s   v1.18.17-gke.100
```

#### 【9】Kubernetesクラスタの削除

コマンド
```
gcloud container clusters delete k8s-cluster --zone $COMPUTE_ZONE --async
```
コマンド結果
```
The following clusters will be deleted.
 - [k8s-cluster] in [asia-northeast1-a]

Do you want to continue (Y/n)?  Y
```