# 第2章 コンテナアプリケーション開発に必要なソフトウェア

## 2.2 Docker環境のセットアップ

### 2.2.1 GCPのセットアップ

#### Googleアカウントの作成

* Googleアカウントの作成手順:
[https://support.google.com/accounts/answer/27441?hl=ja](https://support.google.com/accounts/answer/27441?hl=ja)

#### Googleアカウントの作成

* Google Cloud Platform の無料枠:
[https://cloud.google.com/free/](https://cloud.google.com/free/)

### 2.2.2 GCEのセットアップ

#### 仮想マシンの作成

【1】インスタンスの作成

```gcloudコマンド
$ gcloud compute instances create docker --zone asia-northeast1-b --machine-type=n1-standard-1 --image-family=ubuntu-2004-lts --image-project=ubuntu-os-cloud --boot-disk-size=20GB
WARNING: You have selected a disk size of under [200GB]. This may result in poor I/O performance. For more information, see: https://developers.google.com/compute/docs/disks#performance.
Created [https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-b/instances/docker].
WARNING: Some requests generated warnings:
 - Disk size: '20 GB' is larger than image size: '10 GB'. You might need to resize the root repartition manually if the operating system does not support automatic resizing. See https://cloud.google.com/compute/docs/disks/add-persistent-disk#resize_pd for details.

NAME    ZONE               MACHINE_TYPE   PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP    STATUS
docker  asia-northeast1-b  n1-standard-1               10.146.0.9   34.84.111.225  RUNNING
```

```gcloudコマンド
$ gcloud compute firewall-rules create docker --allow tcp
Creating firewall...⠹Created [https://www.googleapis.com/compute/v1/projects/mercurial-shape-278704/global/firewalls/docker].
Creating firewall...done.
NAME    NETWORK  DIRECTION  PRIORITY  ALLOW  DENY  DISABLED
docker  default  INGRESS    1000      tcp          False
```

【2】SSH接続

```gcloudコマンド
$ gcloud projects list
PROJECT_ID              NAME              PROJECT_NUMBER
mercurial-shape-278704  My First Project  540200718496
```

```gcloudコマンド
$ gcloud compute ssh --project mercurial-shape-278704 --zone asia-northeast1-b docker
Warning: Permanently added 'compute.7067885094151172666' (ECDSA) to the list of known hosts.
Welcome to Ubuntu 20.04.2 LTS (GNU/Linux 5.4.0-1042-gcp x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  System information as of Sun Apr 25 06:43:13 UTC 2021

  System load:  0.0               Processes:             102
  Usage of /:   7.9% of 19.21GB   Users logged in:       0
  Memory usage: 5%                IPv4 address for ens4: 10.146.15.214
  Swap usage:   0%

1 update can be installed immediately.
0 of these updates are security updates.
To see these additional updates run: apt list --upgradable


The list of available updates is more than a week old.
To check for new updates run: sudo apt update


The programs included with the Ubuntu system are free software;
the exact distribution terms for each program are described in the
individual files in /usr/share/doc/*/copyright.

Ubuntu comes with ABSOLUTELY NO WARRANTY, to the extent permitted by
applicable law.

$
```

【3】root ユーザ設定

```gcloudコマンド
$ sudo passwd root
New password:dockerpractice2020
Retype new password:dockerpractice2020
passwd: password updated successfully
$ su -
Password:dockerpractice2020
```

【4】仮想マシンの削除

```gcloudコマンド
$ gcloud compute instances delete docker --zone asia-northeast1-b
The following instances will be deleted. Any attached disks configured
 to be auto-deleted will be deleted unless they are attached to any
other instances or the `--keep-disks` flag is given and specifies them
 for keeping. Deleting a disk is irreversible and any data on the disk
 will be lost.
 - [docker] in [asia-northeast1-b]

Do you want to continue (Y/n)?  Y

Deleted [https:////www.googleapis.com/compute/v1/projects/mercurial-shape-278704/zones/asia-northeast1-b/instances/docker].
```

【5】ファイアウォールの削除

```gcloudコマンド
$ gcloud compute firewall-rules delete docker
The following firewalls will be deleted:
 - [docker]

Do you want to continue (Y/n)?  Y

Deleted [https:////www.googleapis.com/compute/v1/projects/mercurial-shape-278704/global/firewalls/docker].
```

### 2.2.3 Dockerのインストール

【1】Dockerインストールスクリプトのダウンロード

```Linuxコマンド
# curl -fsSL get.docker.com -o get-docker.sh
```

【2】Dockerのインストール

```Dockerコマンド
# sh get-docker.sh
# Executing docker install script, commit: 7cae5f8b0decc17d6571f9f52eb840fbc13b2737
+ sh -c apt-get update -qq >/dev/null
+ sh -c DEBIAN_FRONTEND=noninteractive apt-get install -y -qq apt-transport-https ca-certificates curl >/dev/null
+ sh -c curl -fsSL "https://download.docker.com/linux/ubuntu/gpg" | apt-key add -qq - >/dev/null
Warning: apt-key output should not be parsed (stdout is not a terminal)
+ sh -c echo "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable" > /etc/apt/sources.list.d/docker.list
+ sh -c apt-get update -qq >/dev/null
+ [ -n  ]
+ sh -c apt-get install -y -qq --no-install-recommends docker-ce >/dev/null
+ [ -n 1 ]
+ sh -c DEBIAN_FRONTEND=noninteractive apt-get install -y -qq docker-ce-rootless-extras >/dev/null
+ sh -c docker version
Client: Docker Engine - Community
 Version:           20.10.6
 API version:       1.41
 Go version:        go1.13.15
 Git commit:        370c289
 Built:             Fri Apr  9 22:47:17 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.6
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.13.15
  Git commit:       8728dd2
  Built:            Fri Apr  9 22:45:28 2021
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

================================================================================

To run Docker as a non-privileged user, consider setting up the
Docker daemon in rootless mode for your user:

    dockerd-rootless-setuptool.sh install

Visit https://docs.docker.com/go/rootless/ to learn about rootless mode.


To run the Docker daemon as a fully privileged service, but granting non-root
users access, refer to https://docs.docker.com/go/daemon-access/

WARNING: Access to the remote API on a privileged Docker daemon is equivalent
         to root access on the host. Refer to the 'Docker daemon attack surface'
         documentation for details: https://docs.docker.com/go/attack-surface/

================================================================================
```

【3】Docker コマンドの実行

```Dockerコマンド
# docker system info
Client:
 Context:    default
 Debug Mode: false
 Plugins:
  app: Docker App (Docker Inc., v0.9.1-beta3)
  buildx: Build with BuildKit (Docker Inc., v0.5.1-docker)

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 0
 Server Version: 20.10.6
 Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Native Overlay Diff: true
  userxattr: false
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Cgroup Version: 1
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: inactive
 Runtimes: io.containerd.runc.v2 io.containerd.runtime.v1.linux runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: 05f951a3781f4f2c1911b05e61c160e9c30eaa8e
 runc version: 12644e614e25b05da6fd08a38ffa0cfe1903fdec
 init version: de40ad0
 Security Options:
  apparmor
  seccomp
   Profile: default
 Kernel Version: 5.4.0-1042-gcp
 Operating System: Ubuntu 20.04.2 LTS
 OSType: linux
 Architecture: x86_64
 CPUs: 1
 Total Memory: 3.597GiB
 Name: docker
 ID: HDNY:MNN3:63I2:LYJN:5KA6:7UMS:62VS:SW2F:RPSC:PNKK:XT6K:5B2N
 Docker Root Dir: /var/lib/docker
 Debug Mode: false
 Registry: https://index.docker.io/v1/
 Labels:
 Experimental: false
 Insecure Registries:
  127.0.0.0/8
 Live Restore Enabled: false

WARNING: No swap limit support
```