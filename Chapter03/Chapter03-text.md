# 第3章 コンテナアプリケーション開発ライフサイクルBuild・Ship・Run

## 3.1 Build / コンテナビルド

### 3.1.1 Dockerfile作成からのビルドの実行

#### Dockerfileの作成

```gitコマンド
# git clone https://github.com/cyberblack28/container-develop-environment-construction-guide
```

```unixコマンド
# cd container-develop-environment-construction-guide/Chapter03/3-1-1-01
```

```linuxコマンド
# cat Dockerfile
#CentOS7のベースイメージをPull
FROM centos:7

#yumでepel-releaseをインストール
RUN yum -y install epel-release

#yumでnginxをインストール
RUN yum -y install nginx

#ホスト側のindex.htmlをコピー
COPY index.html /usr/share/nginx/html

#コンテナ起動時に実行するコマンド
ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
```

#### index.htmlファイルの作成

```linuxコマンド
# cat index.html
<!DOCTYPE html>
<html>
<head>
<title>First Docker Build</title>
</head>
<body>
<p>Hello, Container !!</p>
</body>
```

#### コンテナイメージの作成

```dockerコマンド
# docker image build -t cyberblack28/sample-nginx .
Sending build context to Docker daemon  3.072kB
Step 1/5 : FROM centos:7
7: Pulling from library/centos
2d473b07cdd5: Pull complete
Digest: sha256:0f4ec88e21daf75124b8a9e5ca03c37a5e937e0e108a255d890492430789b60e
Status: Downloaded newer image for centos:7
 ---> 8652b9f0cb4c
Step 2/5 : RUN yum -y install epel-release
 ---> Running in ae5353dc60f3
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: ftp-srv2.kddilabs.jp
 * extras: ftp-srv2.kddilabs.jp
 * updates: ty1.mirror.newmediaexpress.com
Resolving Dependencies
--> Running transaction check
---> Package epel-release.noarch 0:7-11 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                Arch             Version         Repository        Size
================================================================================
Installing:
 epel-release           noarch           7-11            extras            15 k

Transaction Summary
================================================================================
Install  1 Package

Total download size: 15 k
Installed size: 24 k
Downloading packages:
warning: /var/cache/yum/x86_64/7/extras/packages/epel-release-7-11.noarch.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for epel-release-7-11.noarch.rpm is not installed
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-9.2009.0.el7.centos.x86_64 (@CentOS)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : epel-release-7-11.noarch                                     1/1
  Verifying  : epel-release-7-11.noarch                                     1/1

Installed:
  epel-release.noarch 0:7-11

Complete!
Removing intermediate container ae5353dc60f3
 ---> 5558949c41b5
Step 3/5 : RUN yum -y install nginx
 ---> Running in 8040ae6650f1
Loaded plugins: fastestmirror, ovl
Loading mirror speeds from cached hostfile
 * base: ftp-srv2.kddilabs.jp
 * epel: d2lzkl7pfhq30w.cloudfront.net
 * extras: ftp-srv2.kddilabs.jp
 * updates: ty1.mirror.newmediaexpress.com
Resolving Dependencies
--> Running transaction check
---> Package nginx.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: nginx-all-modules = 1:1.16.1-3.el7 for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: nginx-filesystem = 1:1.16.1-3.el7 for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libcrypto.so.1.1(OPENSSL_1_1_0)(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libssl.so.1.1(OPENSSL_1_1_0)(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libssl.so.1.1(OPENSSL_1_1_1)(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: nginx-filesystem for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: openssl for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: redhat-indexhtml for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: system-logos for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libcrypto.so.1.1()(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libprofiler.so.0()(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libssl.so.1.1()(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Running transaction check
---> Package centos-indexhtml.noarch 0:7-9.el7.centos will be installed
---> Package centos-logos.noarch 0:70.0.6-3.el7.centos will be installed
---> Package gperftools-libs.x86_64 0:2.6.1-1.el7 will be installed
---> Package nginx-all-modules.noarch 1:1.16.1-3.el7 will be installed
--> Processing Dependency: nginx-mod-http-image-filter = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-http-perl = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-http-xslt-filter = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-mail = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-stream = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
---> Package nginx-filesystem.noarch 1:1.16.1-3.el7 will be installed
---> Package openssl.x86_64 1:1.0.2k-21.el7_9 will be installed
--> Processing Dependency: openssl-libs(x86-64) = 1:1.0.2k-21.el7_9 for package: 1:openssl-1.0.2k-21.el7_9.x86_64
--> Processing Dependency: make for package: 1:openssl-1.0.2k-21.el7_9.x86_64
---> Package openssl11-libs.x86_64 1:1.1.1g-3.el7 will be installed
--> Running transaction check
---> Package make.x86_64 1:3.82-24.el7 will be installed
---> Package nginx-mod-http-image-filter.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: gd for package: 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libgd.so.2()(64bit) for package: 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64
---> Package nginx-mod-http-perl.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: perl >= 5.006001 for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(:MODULE_COMPAT_5.16.3) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(Exporter) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(XSLoader) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(constant) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(strict) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(warnings) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: libperl.so()(64bit) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
---> Package nginx-mod-http-xslt-filter.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.11)(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.18)(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libexslt.so.0()(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libxslt.so.1()(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
---> Package nginx-mod-mail.x86_64 1:1.16.1-3.el7 will be installed
---> Package nginx-mod-stream.x86_64 1:1.16.1-3.el7 will be installed
---> Package openssl-libs.x86_64 1:1.0.2k-19.el7 will be updated
---> Package openssl-libs.x86_64 1:1.0.2k-21.el7_9 will be an update
--> Running transaction check
---> Package gd.x86_64 0:2.0.35-27.el7_9 will be installed
--> Processing Dependency: libpng15.so.15(PNG15_0)(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libjpeg.so.62(LIBJPEG_6.2)(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libpng15.so.15()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libjpeg.so.62()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libfreetype.so.6()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libfontconfig.so.1()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libXpm.so.4()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libX11.so.6()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
---> Package libxslt.x86_64 0:1.1.28-6.el7 will be installed
---> Package perl.x86_64 4:5.16.3-299.el7_9 will be installed
--> Processing Dependency: perl(Socket) >= 1.3 for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Scalar::Util) >= 1.10 for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl-macros for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(threads::shared) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(threads) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Time::Local) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Time::HiRes) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Storable) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Socket) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Scalar::Util) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Pod::Simple::XHTML) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Pod::Simple::Search) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Getopt::Long) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Filter::Util::Call) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Temp) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Spec::Unix) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Spec::Functions) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Spec) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Path) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Cwd) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Carp) for package: 4:perl-5.16.3-299.el7_9.x86_64
---> Package perl-Exporter.noarch 0:5.68-3.el7 will be installed
---> Package perl-constant.noarch 0:1.27-2.el7 will be installed
---> Package perl-libs.x86_64 4:5.16.3-299.el7_9 will be installed
--> Running transaction check
---> Package fontconfig.x86_64 0:2.13.0-4.3.el7 will be installed
--> Processing Dependency: fontpackages-filesystem for package: fontconfig-2.13.0-4.3.el7.x86_64
--> Processing Dependency: dejavu-sans-fonts for package: fontconfig-2.13.0-4.3.el7.x86_64
---> Package freetype.x86_64 0:2.8-14.el7_9.1 will be installed
---> Package libX11.x86_64 0:1.6.7-3.el7_9 will be installed
--> Processing Dependency: libX11-common >= 1.6.7-3.el7_9 for package: libX11-1.6.7-3.el7_9.x86_64
--> Processing Dependency: libxcb.so.1()(64bit) for package: libX11-1.6.7-3.el7_9.x86_64
---> Package libXpm.x86_64 0:3.5.12-1.el7 will be installed
---> Package libjpeg-turbo.x86_64 0:1.2.90-8.el7 will be installed
---> Package libpng.x86_64 2:1.5.13-8.el7 will be installed
---> Package perl-Carp.noarch 0:1.26-244.el7 will be installed
---> Package perl-File-Path.noarch 0:2.09-2.el7 will be installed
---> Package perl-File-Temp.noarch 0:0.23.01-3.el7 will be installed
---> Package perl-Filter.x86_64 0:1.49-3.el7 will be installed
---> Package perl-Getopt-Long.noarch 0:2.40-3.el7 will be installed
--> Processing Dependency: perl(Pod::Usage) >= 1.14 for package: perl-Getopt-Long-2.40-3.el7.noarch
--> Processing Dependency: perl(Text::ParseWords) for package: perl-Getopt-Long-2.40-3.el7.noarch
---> Package perl-PathTools.x86_64 0:3.40-5.el7 will be installed
---> Package perl-Pod-Simple.noarch 1:3.28-4.el7 will be installed
--> Processing Dependency: perl(Pod::Escapes) >= 1.04 for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
--> Processing Dependency: perl(Encode) for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
---> Package perl-Scalar-List-Utils.x86_64 0:1.27-248.el7 will be installed
---> Package perl-Socket.x86_64 0:2.010-5.el7 will be installed
---> Package perl-Storable.x86_64 0:2.45-3.el7 will be installed
---> Package perl-Time-HiRes.x86_64 4:1.9725-3.el7 will be installed
---> Package perl-Time-Local.noarch 0:1.2300-2.el7 will be installed
---> Package perl-macros.x86_64 4:5.16.3-299.el7_9 will be installed
---> Package perl-threads.x86_64 0:1.87-4.el7 will be installed
---> Package perl-threads-shared.x86_64 0:1.43-6.el7 will be installed
--> Running transaction check
---> Package dejavu-sans-fonts.noarch 0:2.33-6.el7 will be installed
--> Processing Dependency: dejavu-fonts-common = 2.33-6.el7 for package: dejavu-sans-fonts-2.33-6.el7.noarch
---> Package fontpackages-filesystem.noarch 0:1.44-8.el7 will be installed
---> Package libX11-common.noarch 0:1.6.7-3.el7_9 will be installed
---> Package libxcb.x86_64 0:1.13-1.el7 will be installed
--> Processing Dependency: libXau.so.6()(64bit) for package: libxcb-1.13-1.el7.x86_64
---> Package perl-Encode.x86_64 0:2.51-7.el7 will be installed
---> Package perl-Pod-Escapes.noarch 1:1.04-299.el7_9 will be installed
---> Package perl-Pod-Usage.noarch 0:1.63-3.el7 will be installed
--> Processing Dependency: perl(Pod::Text) >= 3.15 for package: perl-Pod-Usage-1.63-3.el7.noarch
--> Processing Dependency: perl-Pod-Perldoc for package: perl-Pod-Usage-1.63-3.el7.noarch
---> Package perl-Text-ParseWords.noarch 0:3.29-4.el7 will be installed
--> Running transaction check
---> Package dejavu-fonts-common.noarch 0:2.33-6.el7 will be installed
---> Package libXau.x86_64 0:1.0.8-2.1.el7 will be installed
---> Package perl-Pod-Perldoc.noarch 0:3.20-4.el7 will be installed
--> Processing Dependency: perl(parent) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
--> Processing Dependency: perl(HTTP::Tiny) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
--> Processing Dependency: groff-base for package: perl-Pod-Perldoc-3.20-4.el7.noarch
---> Package perl-podlators.noarch 0:2.5.1-3.el7 will be installed
--> Running transaction check
---> Package groff-base.x86_64 0:1.22.2-8.el7 will be installed
---> Package perl-HTTP-Tiny.noarch 0:0.033-3.el7 will be installed
---> Package perl-parent.noarch 1:0.225-244.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                       Arch     Version                 Repository
                                                                           Size
================================================================================
Installing:
 nginx                         x86_64   1:1.16.1-3.el7          epel      563 k
Installing for dependencies:
 centos-indexhtml              noarch   7-9.el7.centos          base       92 k
 centos-logos                  noarch   70.0.6-3.el7.centos     base       21 M
 dejavu-fonts-common           noarch   2.33-6.el7              base       64 k
 dejavu-sans-fonts             noarch   2.33-6.el7              base      1.4 M
 fontconfig                    x86_64   2.13.0-4.3.el7          base      254 k
 fontpackages-filesystem       noarch   1.44-8.el7              base      9.9 k
 freetype                      x86_64   2.8-14.el7_9.1          updates   380 k
 gd                            x86_64   2.0.35-27.el7_9         updates   146 k
 gperftools-libs               x86_64   2.6.1-1.el7             base      272 k
 groff-base                    x86_64   1.22.2-8.el7            base      942 k
 libX11                        x86_64   1.6.7-3.el7_9           updates   607 k
 libX11-common                 noarch   1.6.7-3.el7_9           updates   164 k
 libXau                        x86_64   1.0.8-2.1.el7           base       29 k
 libXpm                        x86_64   3.5.12-1.el7            base       55 k
 libjpeg-turbo                 x86_64   1.2.90-8.el7            base      135 k
 libpng                        x86_64   2:1.5.13-8.el7          base      213 k
 libxcb                        x86_64   1.13-1.el7              base      214 k
 libxslt                       x86_64   1.1.28-6.el7            base      242 k
 make                          x86_64   1:3.82-24.el7           base      421 k
 nginx-all-modules             noarch   1:1.16.1-3.el7          epel       20 k
 nginx-filesystem              noarch   1:1.16.1-3.el7          epel       21 k
 nginx-mod-http-image-filter   x86_64   1:1.16.1-3.el7          epel       30 k
 nginx-mod-http-perl           x86_64   1:1.16.1-3.el7          epel       39 k
 nginx-mod-http-xslt-filter    x86_64   1:1.16.1-3.el7          epel       29 k
 nginx-mod-mail                x86_64   1:1.16.1-3.el7          epel       57 k
 nginx-mod-stream              x86_64   1:1.16.1-3.el7          epel       85 k
 openssl                       x86_64   1:1.0.2k-21.el7_9       updates   493 k
 openssl11-libs                x86_64   1:1.1.1g-3.el7          epel      1.5 M
 perl                          x86_64   4:5.16.3-299.el7_9      updates   8.0 M
 perl-Carp                     noarch   1.26-244.el7            base       19 k
 perl-Encode                   x86_64   2.51-7.el7              base      1.5 M
 perl-Exporter                 noarch   5.68-3.el7              base       28 k
 perl-File-Path                noarch   2.09-2.el7              base       26 k
 perl-File-Temp                noarch   0.23.01-3.el7           base       56 k
 perl-Filter                   x86_64   1.49-3.el7              base       76 k
 perl-Getopt-Long              noarch   2.40-3.el7              base       56 k
 perl-HTTP-Tiny                noarch   0.033-3.el7             base       38 k
 perl-PathTools                x86_64   3.40-5.el7              base       82 k
 perl-Pod-Escapes              noarch   1:1.04-299.el7_9        updates    52 k
 perl-Pod-Perldoc              noarch   3.20-4.el7              base       87 k
 perl-Pod-Simple               noarch   1:3.28-4.el7            base      216 k
 perl-Pod-Usage                noarch   1.63-3.el7              base       27 k
 perl-Scalar-List-Utils        x86_64   1.27-248.el7            base       36 k
 perl-Socket                   x86_64   2.010-5.el7             base       49 k
 perl-Storable                 x86_64   2.45-3.el7              base       77 k
 perl-Text-ParseWords          noarch   3.29-4.el7              base       14 k
 perl-Time-HiRes               x86_64   4:1.9725-3.el7          base       45 k
 perl-Time-Local               noarch   1.2300-2.el7            base       24 k
 perl-constant                 noarch   1.27-2.el7              base       19 k
 perl-libs                     x86_64   4:5.16.3-299.el7_9      updates   690 k
 perl-macros                   x86_64   4:5.16.3-299.el7_9      updates    44 k
 perl-parent                   noarch   1:0.225-244.el7         base       12 k
 perl-podlators                noarch   2.5.1-3.el7             base      112 k
 perl-threads                  x86_64   1.87-4.el7              base       49 k
 perl-threads-shared           x86_64   1.43-6.el7              base       39 k
Updating for dependencies:
 openssl-libs                  x86_64   1:1.0.2k-21.el7_9       updates   1.2 M

Transaction Summary
================================================================================
Install  1 Package  (+55 Dependent packages)
Upgrade             (  1 Dependent package)

Total download size: 42 M
Downloading packages:
Delta RPMs disabled because /usr/bin/applydeltarpm not installed.
warning: /var/cache/yum/x86_64/7/epel/packages/nginx-1.16.1-3.el7.x86_64.rpm: Header V4 RSA/SHA256 Signature, key ID 352c64e5: NOKEY
Public key for nginx-1.16.1-3.el7.x86_64.rpm is not installed
--------------------------------------------------------------------------------
Total                                               13 MB/s |  42 MB  00:03
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-7
Importing GPG key 0x352C64E5:
 Userid     : "Fedora EPEL (7) <epel@fedoraproject.org>"
 Fingerprint: 91e9 7d7c 4a5e 96f1 7f3e 888f 6a2f aea2 352c 64e5
 Package    : epel-release-7-11.noarch (@extras)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 2:libpng-1.5.13-8.el7.x86_64                                1/58
  Installing : freetype-2.8-14.el7_9.1.x86_64                              2/58
  Installing : fontpackages-filesystem-1.44-8.el7.noarch                   3/58
  Installing : dejavu-fonts-common-2.33-6.el7.noarch                       4/58
  Installing : dejavu-sans-fonts-2.33-6.el7.noarch                         5/58
  Installing : fontconfig-2.13.0-4.3.el7.x86_64                            6/58
  Installing : libXau-1.0.8-2.1.el7.x86_64                                 7/58
  Installing : libxcb-1.13-1.el7.x86_64                                    8/58
  Installing : 1:openssl11-libs-1.1.1g-3.el7.x86_64                        9/58
  Installing : 1:nginx-filesystem-1.16.1-3.el7.noarch                     10/58
  Installing : libxslt-1.1.28-6.el7.x86_64                                11/58
  Updating   : 1:openssl-libs-1.0.2k-21.el7_9.x86_64                      12/58
  Installing : libX11-common-1.6.7-3.el7_9.noarch                         13/58
  Installing : libX11-1.6.7-3.el7_9.x86_64                                14/58
  Installing : libXpm-3.5.12-1.el7.x86_64                                 15/58
  Installing : libjpeg-turbo-1.2.90-8.el7.x86_64                          16/58
  Installing : gd-2.0.35-27.el7_9.x86_64                                  17/58
  Installing : 1:make-3.82-24.el7.x86_64                                  18/58
  Installing : 1:openssl-1.0.2k-21.el7_9.x86_64                           19/58
  Installing : centos-indexhtml-7-9.el7.centos.noarch                     20/58
  Installing : centos-logos-70.0.6-3.el7.centos.noarch                    21/58
  Installing : groff-base-1.22.2-8.el7.x86_64                             22/58
  Installing : 1:perl-parent-0.225-244.el7.noarch                         23/58
  Installing : perl-HTTP-Tiny-0.033-3.el7.noarch                          24/58
  Installing : perl-podlators-2.5.1-3.el7.noarch                          25/58
  Installing : perl-Pod-Perldoc-3.20-4.el7.noarch                         26/58
  Installing : 1:perl-Pod-Escapes-1.04-299.el7_9.noarch                   27/58
  Installing : perl-Encode-2.51-7.el7.x86_64                              28/58
  Installing : perl-Text-ParseWords-3.29-4.el7.noarch                     29/58
  Installing : perl-Pod-Usage-1.63-3.el7.noarch                           30/58
  Installing : perl-threads-1.87-4.el7.x86_64                             31/58
  Installing : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                      32/58
  Installing : perl-Exporter-5.68-3.el7.noarch                            33/58
  Installing : perl-constant-1.27-2.el7.noarch                            34/58
  Installing : perl-Socket-2.010-5.el7.x86_64                             35/58
  Installing : perl-Filter-1.49-3.el7.x86_64                              36/58
  Installing : perl-Time-Local-1.2300-2.el7.noarch                        37/58
  Installing : perl-Carp-1.26-244.el7.noarch                              38/58
  Installing : 4:perl-macros-5.16.3-299.el7_9.x86_64                      39/58
  Installing : perl-Storable-2.45-3.el7.x86_64                            40/58
  Installing : perl-PathTools-3.40-5.el7.x86_64                           41/58
  Installing : perl-threads-shared-1.43-6.el7.x86_64                      42/58
  Installing : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 43/58
  Installing : 1:perl-Pod-Simple-3.28-4.el7.noarch                        44/58
  Installing : perl-File-Temp-0.23.01-3.el7.noarch                        45/58
  Installing : perl-File-Path-2.09-2.el7.noarch                           46/58
  Installing : 4:perl-libs-5.16.3-299.el7_9.x86_64                        47/58
  Installing : perl-Getopt-Long-2.40-3.el7.noarch                         48/58
  Installing : 4:perl-5.16.3-299.el7_9.x86_64                             49/58
  Installing : gperftools-libs-2.6.1-1.el7.x86_64                         50/58
  Installing : 1:nginx-mod-mail-1.16.1-3.el7.x86_64                       51/58
  Installing : 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64           52/58
  Installing : 1:nginx-mod-stream-1.16.1-3.el7.x86_64                     53/58
  Installing : 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64          54/58
  Installing : 1:nginx-1.16.1-3.el7.x86_64                                55/58
  Installing : 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64                  56/58
  Installing : 1:nginx-all-modules-1.16.1-3.el7.noarch                    57/58
  Cleanup    : 1:openssl-libs-1.0.2k-19.el7.x86_64                        58/58
  Verifying  : perl-HTTP-Tiny-0.033-3.el7.noarch                           1/58
  Verifying  : fontconfig-2.13.0-4.3.el7.x86_64                            2/58
  Verifying  : 1:nginx-mod-mail-1.16.1-3.el7.x86_64                        3/58
  Verifying  : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                       4/58
  Verifying  : perl-threads-1.87-4.el7.x86_64                              5/58
  Verifying  : perl-Exporter-5.68-3.el7.noarch                             6/58
  Verifying  : perl-constant-1.27-2.el7.noarch                             7/58
  Verifying  : perl-PathTools-3.40-5.el7.x86_64                            8/58
  Verifying  : gperftools-libs-2.6.1-1.el7.x86_64                          9/58
  Verifying  : perl-Socket-2.010-5.el7.x86_64                             10/58
  Verifying  : fontpackages-filesystem-1.44-8.el7.noarch                  11/58
  Verifying  : groff-base-1.22.2-8.el7.x86_64                             12/58
  Verifying  : centos-logos-70.0.6-3.el7.centos.noarch                    13/58
  Verifying  : 1:perl-parent-0.225-244.el7.noarch                         14/58
  Verifying  : gd-2.0.35-27.el7_9.x86_64                                  15/58
  Verifying  : centos-indexhtml-7-9.el7.centos.noarch                     16/58
  Verifying  : perl-Filter-1.49-3.el7.x86_64                              17/58
  Verifying  : perl-File-Temp-0.23.01-3.el7.noarch                        18/58
  Verifying  : 1:perl-Pod-Simple-3.28-4.el7.noarch                        19/58
  Verifying  : 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64                  20/58
  Verifying  : perl-Time-Local-1.2300-2.el7.noarch                        21/58
  Verifying  : libxcb-1.13-1.el7.x86_64                                   22/58
  Verifying  : 1:make-3.82-24.el7.x86_64                                  23/58
  Verifying  : 1:perl-Pod-Escapes-1.04-299.el7_9.noarch                   24/58
  Verifying  : perl-Pod-Perldoc-3.20-4.el7.noarch                         25/58
  Verifying  : 1:openssl-1.0.2k-21.el7_9.x86_64                           26/58
  Verifying  : libXpm-3.5.12-1.el7.x86_64                                 27/58
  Verifying  : libjpeg-turbo-1.2.90-8.el7.x86_64                          28/58
  Verifying  : perl-Carp-1.26-244.el7.noarch                              29/58
  Verifying  : perl-threads-shared-1.43-6.el7.x86_64                      30/58
  Verifying  : libX11-common-1.6.7-3.el7_9.noarch                         31/58
  Verifying  : libX11-1.6.7-3.el7_9.x86_64                                32/58
  Verifying  : 4:perl-macros-5.16.3-299.el7_9.x86_64                      33/58
  Verifying  : perl-Storable-2.45-3.el7.x86_64                            34/58
  Verifying  : 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64           35/58
  Verifying  : dejavu-sans-fonts-2.33-6.el7.noarch                        36/58
  Verifying  : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 37/58
  Verifying  : 2:libpng-1.5.13-8.el7.x86_64                               38/58
  Verifying  : 1:nginx-mod-stream-1.16.1-3.el7.x86_64                     39/58
  Verifying  : 1:openssl-libs-1.0.2k-21.el7_9.x86_64                      40/58
  Verifying  : freetype-2.8-14.el7_9.1.x86_64                             41/58
  Verifying  : perl-Encode-2.51-7.el7.x86_64                              42/58
  Verifying  : perl-Pod-Usage-1.63-3.el7.noarch                           43/58
  Verifying  : dejavu-fonts-common-2.33-6.el7.noarch                      44/58
  Verifying  : perl-podlators-2.5.1-3.el7.noarch                          45/58
  Verifying  : 4:perl-5.16.3-299.el7_9.x86_64                             46/58
  Verifying  : perl-File-Path-2.09-2.el7.noarch                           47/58
  Verifying  : libxslt-1.1.28-6.el7.x86_64                                48/58
  Verifying  : 1:nginx-filesystem-1.16.1-3.el7.noarch                     49/58
  Verifying  : 1:nginx-1.16.1-3.el7.x86_64                                50/58
  Verifying  : 1:openssl11-libs-1.1.1g-3.el7.x86_64                       51/58
  Verifying  : libXau-1.0.8-2.1.el7.x86_64                                52/58
  Verifying  : 1:nginx-all-modules-1.16.1-3.el7.noarch                    53/58
  Verifying  : perl-Getopt-Long-2.40-3.el7.noarch                         54/58
  Verifying  : perl-Text-ParseWords-3.29-4.el7.noarch                     55/58
  Verifying  : 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64          56/58
  Verifying  : 4:perl-libs-5.16.3-299.el7_9.x86_64                        57/58
  Verifying  : 1:openssl-libs-1.0.2k-19.el7.x86_64                        58/58

Installed:
  nginx.x86_64 1:1.16.1-3.el7

Dependency Installed:
  centos-indexhtml.noarch 0:7-9.el7.centos
  centos-logos.noarch 0:70.0.6-3.el7.centos
  dejavu-fonts-common.noarch 0:2.33-6.el7
  dejavu-sans-fonts.noarch 0:2.33-6.el7
  fontconfig.x86_64 0:2.13.0-4.3.el7
  fontpackages-filesystem.noarch 0:1.44-8.el7
  freetype.x86_64 0:2.8-14.el7_9.1
  gd.x86_64 0:2.0.35-27.el7_9
  gperftools-libs.x86_64 0:2.6.1-1.el7
  groff-base.x86_64 0:1.22.2-8.el7
  libX11.x86_64 0:1.6.7-3.el7_9
  libX11-common.noarch 0:1.6.7-3.el7_9
  libXau.x86_64 0:1.0.8-2.1.el7
  libXpm.x86_64 0:3.5.12-1.el7
  libjpeg-turbo.x86_64 0:1.2.90-8.el7
  libpng.x86_64 2:1.5.13-8.el7
  libxcb.x86_64 0:1.13-1.el7
  libxslt.x86_64 0:1.1.28-6.el7
  make.x86_64 1:3.82-24.el7
  nginx-all-modules.noarch 1:1.16.1-3.el7
  nginx-filesystem.noarch 1:1.16.1-3.el7
  nginx-mod-http-image-filter.x86_64 1:1.16.1-3.el7
  nginx-mod-http-perl.x86_64 1:1.16.1-3.el7
  nginx-mod-http-xslt-filter.x86_64 1:1.16.1-3.el7
  nginx-mod-mail.x86_64 1:1.16.1-3.el7
  nginx-mod-stream.x86_64 1:1.16.1-3.el7
  openssl.x86_64 1:1.0.2k-21.el7_9
  openssl11-libs.x86_64 1:1.1.1g-3.el7
  perl.x86_64 4:5.16.3-299.el7_9
  perl-Carp.noarch 0:1.26-244.el7
  perl-Encode.x86_64 0:2.51-7.el7
  perl-Exporter.noarch 0:5.68-3.el7
  perl-File-Path.noarch 0:2.09-2.el7
  perl-File-Temp.noarch 0:0.23.01-3.el7
  perl-Filter.x86_64 0:1.49-3.el7
  perl-Getopt-Long.noarch 0:2.40-3.el7
  perl-HTTP-Tiny.noarch 0:0.033-3.el7
  perl-PathTools.x86_64 0:3.40-5.el7
  perl-Pod-Escapes.noarch 1:1.04-299.el7_9
  perl-Pod-Perldoc.noarch 0:3.20-4.el7
  perl-Pod-Simple.noarch 1:3.28-4.el7
  perl-Pod-Usage.noarch 0:1.63-3.el7
  perl-Scalar-List-Utils.x86_64 0:1.27-248.el7
  perl-Socket.x86_64 0:2.010-5.el7
  perl-Storable.x86_64 0:2.45-3.el7
  perl-Text-ParseWords.noarch 0:3.29-4.el7
  perl-Time-HiRes.x86_64 4:1.9725-3.el7
  perl-Time-Local.noarch 0:1.2300-2.el7
  perl-constant.noarch 0:1.27-2.el7
  perl-libs.x86_64 4:5.16.3-299.el7_9
  perl-macros.x86_64 4:5.16.3-299.el7_9
  perl-parent.noarch 1:0.225-244.el7
  perl-podlators.noarch 0:2.5.1-3.el7
  perl-threads.x86_64 0:1.87-4.el7
  perl-threads-shared.x86_64 0:1.43-6.el7

Dependency Updated:
  openssl-libs.x86_64 1:1.0.2k-21.el7_9

Complete!
Removing intermediate container 8040ae6650f1
 ---> c39fdd92f3fc
Step 4/5 : COPY index.html /usr/share/nginx/html
 ---> 2fd5a19d6053
Step 5/5 : ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
 ---> Running in f0b63c663e9a
Removing intermediate container f0b63c663e9a
 ---> 59cb9dc6b0e2
Successfully built 59cb9dc6b0e2
Successfully tagged cyberblack28/sample-nginx:latest
```

#### コンテナイメージの確認

```dockerコマンド
# docker image ls
REPOSITORY                  TAG       IMAGE ID       CREATED          SIZE
cyberblack28/sample-nginx   latest    dda2f2973802   10 minutes ago   558MB
centos                      7         8652b9f0cb4c   5 months ago     204MB
```

#### レイヤの確認

```dockerコマンド
# docker image history cyberblack28/sample-nginx
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
59cb9dc6b0e2   12 minutes ago   /bin/sh -c #(nop)  ENTRYPOINT ["/usr/sbin/ng…   0B
2fd5a19d6053   12 minutes ago   /bin/sh -c #(nop) COPY file:3e884108a93ee9b1…   126B
c39fdd92f3fc   12 minutes ago   /bin/sh -c yum -y install nginx                 238MB
5558949c41b5   12 minutes ago   /bin/sh -c yum -y install epel-release          116MB
8652b9f0cb4c   5 months ago     /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B
<missing>      5 months ago     /bin/sh -c #(nop)  LABEL org.label-schema.sc…   0B
<missing>      5 months ago     /bin/sh -c #(nop) ADD file:b3ebbe8bd304723d4…   204MB
```

### 3.1.2 イメージの軽量化

* distroless:
[https://github.com/GoogleContainerTools/distroless](https://github.com/GoogleContainerTools/distroless)

### 3.1.3 マルチステージビルド

#### Go 言語のサンプルアプリケーションの作成

```linuxコマンド
# cd ../3-1-3-01
```

```linuxコマンド
# cat main.go
package main

import "fmt"

func main() {
        fmt.Println("Let's start multi-stage builds !!")
}
```

```linuxコマンド
# cat Dockerfile-msb
#Stage1
#Goのビルドを実行するベースイメージをPull
FROM golang:1.16.4-alpine3.13 as builder
#ローカルのmain.goをイメージ内にコピー
COPY ./main.go ./
#main.goをソースからビルド（コンパイル）して、msbというビルド成果物を生成
RUN go build -o /msb ./main.go

#Stage2
#alpineベースイメージをPull
FROM alpine:3.13
#ビルドイメージからビルド成果物のmsbをコピーして配置
COPY --from=builder /msb /usr/local/bin/msb
#msbの実
```

```dockerコマンド
# docker image build -t msb -f Dockerfile-msb .
Sending build context to Docker daemon  3.584kB
Step 1/6 : FROM golang:1.16.4-alpine3.13 as builder
1.16.4-alpine3.13: Pulling from library/golang
540db60ca938: Pull complete
adcc1eea9eea: Pull complete
4c4ab2625f07: Pull complete
c5e7595549f7: Pull complete
3df88182f7ac: Pull complete
Digest: sha256:4dd403b2e7a689adc5b7110ba9cd5da43d216cfcfccfbe2b35680effcf336c7e
Status: Downloaded newer image for golang:1.16.4-alpine3.13
 ---> 722a834ff95b
Step 2/6 : COPY ./main.go ./
 ---> c066d1644250
Step 3/6 : RUN go build -o /msb ./main.go
 ---> Running in 808cdbbbd29a
Removing intermediate container 808cdbbbd29a
 ---> 0b9401537c5c
Step 4/6 : FROM alpine:3.13
3.13: Pulling from library/alpine
540db60ca938: Already exists
Digest: sha256:69e70a79f2d41ab5d637de98c1e0b055206ba40a8145e7bddb55ccc04e13cf8f
Status: Downloaded newer image for alpine:3.13
 ---> 6dbb9cc54074
Step 5/6 : COPY --from=builder /msb /usr/local/bin/msb
 ---> 628d25db0d06
Step 6/6 : ENTRYPOINT ["/usr/local/bin/msb"]
 ---> Running in b67b541189db
Removing intermediate container b67b541189db
 ---> 1ffb1b502170
Successfully built 1ffb1b502170
Successfully tagged msb:latest
```

```dockerコマンド
# docker container run -it --rm msb
Let's start multi-stage builds !!
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
msb                         latest              1ffb1b502170   2 minutes ago    7.55MB
<none>                      <none>              0b9401537c5c   2 minutes ago    303MB
cyberblack28/sample-nginx   latest              59cb9dc6b0e2   16 minutes ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   4 days ago       301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

### 3.1.4 ビルドツール

#### BuildKitを利用したビルド

```dockerコマンド
# DOCKER_BUILDKIT=1 docker image build -t msb -f Dockerfile-msb .
[+] Building 1.7s (11/11) FINISHED
 => [internal] load build definition from Dockerfile-msb                                                                                                                         0.1s
 => => transferring dockerfile: 589B                                                                                                                                             0.0s
 => [internal] load .dockerignore                                                                                                                                                0.1s
 => => transferring context: 2B                                                                                                                                                  0.0s
 => [internal] load metadata for docker.io/library/alpine:3.13                                                                                                                   0.0s
 => [internal] load metadata for docker.io/library/golang:1.16.4-alpine3.13                                                                                                      0.0s
 => [stage-1 1/2] FROM docker.io/library/alpine:3.13                                                                                                                             0.0s
 => [internal] load build context                                                                                                                                                0.1s
 => => transferring context: 134B                                                                                                                                                0.0s
 => [builder 1/3] FROM docker.io/library/golang:1.16.4-alpine3.13                                                                                                                0.1s
 => [builder 2/3] COPY ./main.go ./                                                                                                                                              0.1s
 => [builder 3/3] RUN go build -o /msb ./main.go                                                                                                                                 0.8s
 => [stage-1 2/2] COPY --from=builder /msb /usr/local/bin/msb                                                                                                                    0.1s
 => exporting to image                                                                                                                                                           0.1s
 => => exporting layers                                                                                                                                                          0.0s
 => => writing image sha256:ab4aa580a77fb3fb3ed6053a328f3b59a5a1a2c2ab00022ac085e1537192db9c                                                                                     0.0s
 => => naming to docker.io/library/msb                                                                                                                                           0.0s
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
msb                         latest              ab4aa580a77f   3 minutes ago    7.55MB
<none>                      <none>              1ffb1b502170   7 minutes ago    7.55MB
<none>                      <none>              0b9401537c5c   7 minutes ago    303MB
cyberblack28/sample-nginx   latest              59cb9dc6b0e2   21 minutes ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   4 days ago       301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

#### Trivyのインストール

```dockerコマンド
# curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/master/contrib/install.sh | sh -s -- -b /usr/local/bin
aquasecurity/trivy info checking GitHub for latest tag
aquasecurity/trivy info found version: 0.17.2 for v0.17.2/Linux/64bit
aquasecurity/trivy info installed /usr/local/bin/trivy
```

#### Trivyによる検査

省略箇所はあまりにも量が多く、随時変わるものなのでこちらでも省略とします。

```trivyコマンド
# trivy image cyberblack28/sample-nginx
2021-05-08T14:48:03.203Z        INFO    Detecting RHEL/CentOS vulnerabilities...
2021-05-08T14:48:03.222Z        INFO    Trivy skips scanning programming language libraries because no supported file was detected

cyberblack28/sample-nginx (centos 7.9.2009)
===========================================
Total: 687 (UNKNOWN: 0, LOW: 385, MEDIUM: 297, HIGH: 5, CRITICAL: 0)
：
：<省略>
：
```

#### イメージの検査

省略箇所はあまりにも量が多く、随時変わるものなのでこちらでも省略とします。

```trivyコマンド
# trivy nginx:1.19.2
2020-10-06T13:20:26.558Z        INFO    Need to update DB
2020-10-06T13:20:26.558Z        INFO    Downloading DB...
18.68 MiB / 18.68 MiB [------------------------------------------------------------------------------------------------------------------------------------------------------] 100.00% 4.42 MiB p/s 4s
2020-10-06T13:20:36.486Z        INFO    Detecting Debian vulnerabilities...

nginx:1.19.2 (debian 10.5)
==========================
Total: 151 (UNKNOWN: 0, LOW: 116, MEDIUM: 35, HIGH: 0, CRITICAL: 0)

+------------------+---------------------+----------+---------------------------+-------------------+--------------------------------------------------------------+
|     LIBRARY      |  VULNERABILITY ID   | SEVERITY |     INSTALLED VERSION     |   FIXED VERSION   |                            TITLE                             |
+------------------+---------------------+----------+---------------------------+-------------------+--------------------------------------------------------------+
| apt              | CVE-2011-3374       | LOW      | 1.8.2.1                   |                   | It was found that apt-key                                    |
|                  |                     |          |                           |                   | in apt, all versions, do not                                 |
|                  |                     |          |                           |                   | correctly...                                                 |
+------------------+---------------------+          +---------------------------+-------------------+--------------------------------------------------------------+
| bash             | CVE-2019-18276      |          | 5.0-4                     |                   | bash: when effective UID is                                  |
|                  |                     |          |                           |                   | not equal to its real UID                                    |
|                  |                     |          |                           |                   | the...                                                       |
+                  +---------------------+          +                           +-------------------+--------------------------------------------------------------+
|                  | TEMP-0841856-B18BAF |          |                           |                   |                                                              |
+------------------+---------------------+          +---------------------------+-------------------+--------------------------------------------------------------+
・
・<省略>
・
+------------------+---------------------+          +---------------------------+-------------------+--------------------------------------------------------------+
| sysvinit-utils   | TEMP-0517018-A83CE6 |          | 2.93-8                    |                   |                                                              |
+------------------+---------------------+          +---------------------------+-------------------+--------------------------------------------------------------+
| tar              | CVE-2005-2541       |          | 1.30+dfsg-6               |                   | Tar 1.15.1 does not properly                                 |
|                  |                     |          |                           |                   | warn the user when extracting                                |
|                  |                     |          |                           |                   | setuid or...                                                 |
+                  +---------------------+          +                           +-------------------+--------------------------------------------------------------+
|                  | CVE-2019-9923       |          |                           |                   | tar: null-pointer dereference                                |
|                  |                     |          |                           |                   | in pax_decode_header in                                      |
|                  |                     |          |                           |                   | sparse.c                                                     |
+                  +---------------------+          +                           +-------------------+--------------------------------------------------------------+
|                  | TEMP-0290435-0B57B5 |          |                           |                   |                                                              |
+------------------+---------------------+----------+---------------------------+-------------------+--------------------------------------------------------------+
```

### 3.2.4 イメージのPush/Pullコマンドの実行

#### Docker Hubへのログイン

```dockerコマンド
# docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: cyberblack28
Password:
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
```

#### Pushコマンドの実行

```dockerコマンド
# docker image push cyberblack28/sample-nginx
Using default tag: latest
The push refers to repository [docker.io/cyberblack28/sample-nginx]
09ff5c3f3c16: Pushed
6724fad9c54d: Pushed
8ec59fb6f8fd: Pushed
174f56854903: Mounted from library/centos
latest: digest: sha256:c80b9c9ef042fecee96e7bd4b4e8456d113757133fe575adfe61768a434aaae7 size: 1161
```

#### IMAGE IDの確認

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
msb                         latest              ab4aa580a77f   9 minutes ago    7.55MB
<none>                      <none>              1ffb1b502170   13 minutes ago   7.55MB
<none>                      <none>              0b9401537c5c   13 minutes ago   303MB
cyberblack28/sample-nginx   latest              59cb9dc6b0e2   27 minutes ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   4 days ago       301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

#### イメージの削除と確認

```dockerコマンド
# docker image rm 59cb9dc6b0e2
Untagged: cyberblack28/sample-nginx:latest
Untagged: cyberblack28/sample-nginx@sha256:c80b9c9ef042fecee96e7bd4b4e8456d113757133fe575adfe61768a434aaae7
Deleted: sha256:59cb9dc6b0e253c56c88a3b0ac12a09e5d1e2e30f47b2489a899e3a09504bf1c
Deleted: sha256:2fd5a19d6053511680d37fea7b2ca74698b1fa6b9c0b69f13773a321ec87b8ae
Deleted: sha256:38dd054c1d3f973efd9344aba750f5583b425779818ab59f1ca81a187bf8c7ef
Deleted: sha256:c39fdd92f3fc8f9ee980dd6a594bf9685efa7357fb6439157035faca00fe425f
Deleted: sha256:46929772216af6d6ad2d7b5b1130a033ce09a771e3d40f48616e1cc17d7f5c31
Deleted: sha256:5558949c41b554a1377ae604cc086a964a763405a3ef7121faa022b70e828384
Deleted: sha256:dc254c8b63c3b59e102926bc8214ce1269a57890de864f8e98deb663edf399f7
```

```dockerコマンド
# docker image ls
REPOSITORY   TAG                 IMAGE ID       CREATED          SIZE
msb          latest              ab4aa580a77f   11 minutes ago   7.55MB
<none>       <none>              1ffb1b502170   15 minutes ago   7.55MB
<none>       <none>              0b9401537c5c   15 minutes ago   303MB
golang       1.16.4-alpine3.13   722a834ff95b   4 days ago       301MB
alpine       3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos       7                   8652b9f0cb4c   5 months ago     204MB
```

#### Pullコマンドの実行

```dockerコマンド
# docker image pull cyberblack28/sample-nginx
Using default tag: latest
latest: Pulling from cyberblack28/sample-nginx
2d473b07cdd5: Already exists
8e3bbe9a067d: Pull complete
5bf3670baede: Pull complete
4bc8e8517989: Pull complete
Digest: sha256:c80b9c9ef042fecee96e7bd4b4e8456d113757133fe575adfe61768a434aaae7
Status: Downloaded newer image for cyberblack28/sample-nginx:latest
docker.io/cyberblack28/sample-nginx:latest
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
msb                         latest              ab4aa580a77f   16 minutes ago   7.55MB
<none>                      <none>              1ffb1b502170   19 minutes ago   7.55MB
<none>                      <none>              0b9401537c5c   19 minutes ago   303MB
cyberblack28/sample-nginx   latest              59cb9dc6b0e2   33 minutes ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   4 days ago       301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

#### Image Tag

```linuxコマンド
# cd ../3-2-4-01
```

```dockerコマンド
# cat index.html
<!DOCTYPE html>
<html>
<head>
<title>First Docker Build</title>
</head>
<body>
<p>Happy, Container !!</p>
</body>
</html>
```

```dockerコマンド
# docker image build -t cyberblack28/sample-nginx .
Sending build context to Docker daemon  3.072kB
Step 1/5 : FROM centos:7
 ---> 8652b9f0cb4c
Step 2/5 : RUN yum -y install epel-release
 ---> Running in 6d1ed13029ca
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: ftp-srv2.kddilabs.jp
 * extras: ftp-srv2.kddilabs.jp
 * updates: ty1.mirror.newmediaexpress.com
Resolving Dependencies
--> Running transaction check
---> Package epel-release.noarch 0:7-11 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                Arch             Version         Repository        Size
================================================================================
Installing:
 epel-release           noarch           7-11            extras            15 k

Transaction Summary
================================================================================
Install  1 Package

Total download size: 15 k
Installed size: 24 k
Downloading packages:
warning: /var/cache/yum/x86_64/7/extras/packages/epel-release-7-11.noarch.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for epel-release-7-11.noarch.rpm is not installed
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-9.2009.0.el7.centos.x86_64 (@CentOS)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : epel-release-7-11.noarch                                     1/1
  Verifying  : epel-release-7-11.noarch                                     1/1

Installed:
  epel-release.noarch 0:7-11

Complete!
Removing intermediate container 6d1ed13029ca
 ---> 358d8684562a
Step 3/5 : RUN yum -y install nginx
 ---> Running in c75c12a6a037
Loaded plugins: fastestmirror, ovl
Loading mirror speeds from cached hostfile
 * base: ftp-srv2.kddilabs.jp
 * epel: d2lzkl7pfhq30w.cloudfront.net
 * extras: ftp-srv2.kddilabs.jp
 * updates: ty1.mirror.newmediaexpress.com
Resolving Dependencies
--> Running transaction check
---> Package nginx.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: nginx-all-modules = 1:1.16.1-3.el7 for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: nginx-filesystem = 1:1.16.1-3.el7 for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libcrypto.so.1.1(OPENSSL_1_1_0)(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libssl.so.1.1(OPENSSL_1_1_0)(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libssl.so.1.1(OPENSSL_1_1_1)(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: nginx-filesystem for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: openssl for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: redhat-indexhtml for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: system-logos for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libcrypto.so.1.1()(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libprofiler.so.0()(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Processing Dependency: libssl.so.1.1()(64bit) for package: 1:nginx-1.16.1-3.el7.x86_64
--> Running transaction check
---> Package centos-indexhtml.noarch 0:7-9.el7.centos will be installed
---> Package centos-logos.noarch 0:70.0.6-3.el7.centos will be installed
---> Package gperftools-libs.x86_64 0:2.6.1-1.el7 will be installed
---> Package nginx-all-modules.noarch 1:1.16.1-3.el7 will be installed
--> Processing Dependency: nginx-mod-http-image-filter = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-http-perl = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-http-xslt-filter = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-mail = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
--> Processing Dependency: nginx-mod-stream = 1:1.16.1-3.el7 for package: 1:nginx-all-modules-1.16.1-3.el7.noarch
---> Package nginx-filesystem.noarch 1:1.16.1-3.el7 will be installed
---> Package openssl.x86_64 1:1.0.2k-21.el7_9 will be installed
--> Processing Dependency: openssl-libs(x86-64) = 1:1.0.2k-21.el7_9 for package: 1:openssl-1.0.2k-21.el7_9.x86_64
--> Processing Dependency: make for package: 1:openssl-1.0.2k-21.el7_9.x86_64
---> Package openssl11-libs.x86_64 1:1.1.1g-3.el7 will be installed
--> Running transaction check
---> Package make.x86_64 1:3.82-24.el7 will be installed
---> Package nginx-mod-http-image-filter.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: gd for package: 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libgd.so.2()(64bit) for package: 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64
---> Package nginx-mod-http-perl.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: perl >= 5.006001 for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(:MODULE_COMPAT_5.16.3) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(Exporter) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(XSLoader) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(constant) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(strict) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: perl(warnings) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
--> Processing Dependency: libperl.so()(64bit) for package: 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64
---> Package nginx-mod-http-xslt-filter.x86_64 1:1.16.1-3.el7 will be installed
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.11)(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.18)(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libexslt.so.0()(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
--> Processing Dependency: libxslt.so.1()(64bit) for package: 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64
---> Package nginx-mod-mail.x86_64 1:1.16.1-3.el7 will be installed
---> Package nginx-mod-stream.x86_64 1:1.16.1-3.el7 will be installed
---> Package openssl-libs.x86_64 1:1.0.2k-19.el7 will be updated
---> Package openssl-libs.x86_64 1:1.0.2k-21.el7_9 will be an update
--> Running transaction check
---> Package gd.x86_64 0:2.0.35-27.el7_9 will be installed
--> Processing Dependency: libpng15.so.15(PNG15_0)(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libjpeg.so.62(LIBJPEG_6.2)(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libpng15.so.15()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libjpeg.so.62()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libfreetype.so.6()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libfontconfig.so.1()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libXpm.so.4()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
--> Processing Dependency: libX11.so.6()(64bit) for package: gd-2.0.35-27.el7_9.x86_64
---> Package libxslt.x86_64 0:1.1.28-6.el7 will be installed
---> Package perl.x86_64 4:5.16.3-299.el7_9 will be installed
--> Processing Dependency: perl(Socket) >= 1.3 for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Scalar::Util) >= 1.10 for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl-macros for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(threads::shared) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(threads) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Time::Local) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Time::HiRes) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Storable) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Socket) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Scalar::Util) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Pod::Simple::XHTML) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Pod::Simple::Search) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Getopt::Long) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Filter::Util::Call) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Temp) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Spec::Unix) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Spec::Functions) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Spec) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(File::Path) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Cwd) for package: 4:perl-5.16.3-299.el7_9.x86_64
--> Processing Dependency: perl(Carp) for package: 4:perl-5.16.3-299.el7_9.x86_64
---> Package perl-Exporter.noarch 0:5.68-3.el7 will be installed
---> Package perl-constant.noarch 0:1.27-2.el7 will be installed
---> Package perl-libs.x86_64 4:5.16.3-299.el7_9 will be installed
--> Running transaction check
---> Package fontconfig.x86_64 0:2.13.0-4.3.el7 will be installed
--> Processing Dependency: fontpackages-filesystem for package: fontconfig-2.13.0-4.3.el7.x86_64
--> Processing Dependency: dejavu-sans-fonts for package: fontconfig-2.13.0-4.3.el7.x86_64
---> Package freetype.x86_64 0:2.8-14.el7_9.1 will be installed
---> Package libX11.x86_64 0:1.6.7-3.el7_9 will be installed
--> Processing Dependency: libX11-common >= 1.6.7-3.el7_9 for package: libX11-1.6.7-3.el7_9.x86_64
--> Processing Dependency: libxcb.so.1()(64bit) for package: libX11-1.6.7-3.el7_9.x86_64
---> Package libXpm.x86_64 0:3.5.12-1.el7 will be installed
---> Package libjpeg-turbo.x86_64 0:1.2.90-8.el7 will be installed
---> Package libpng.x86_64 2:1.5.13-8.el7 will be installed
---> Package perl-Carp.noarch 0:1.26-244.el7 will be installed
---> Package perl-File-Path.noarch 0:2.09-2.el7 will be installed
---> Package perl-File-Temp.noarch 0:0.23.01-3.el7 will be installed
---> Package perl-Filter.x86_64 0:1.49-3.el7 will be installed
---> Package perl-Getopt-Long.noarch 0:2.40-3.el7 will be installed
--> Processing Dependency: perl(Pod::Usage) >= 1.14 for package: perl-Getopt-Long-2.40-3.el7.noarch
--> Processing Dependency: perl(Text::ParseWords) for package: perl-Getopt-Long-2.40-3.el7.noarch
---> Package perl-PathTools.x86_64 0:3.40-5.el7 will be installed
---> Package perl-Pod-Simple.noarch 1:3.28-4.el7 will be installed
--> Processing Dependency: perl(Pod::Escapes) >= 1.04 for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
--> Processing Dependency: perl(Encode) for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
---> Package perl-Scalar-List-Utils.x86_64 0:1.27-248.el7 will be installed
---> Package perl-Socket.x86_64 0:2.010-5.el7 will be installed
---> Package perl-Storable.x86_64 0:2.45-3.el7 will be installed
---> Package perl-Time-HiRes.x86_64 4:1.9725-3.el7 will be installed
---> Package perl-Time-Local.noarch 0:1.2300-2.el7 will be installed
---> Package perl-macros.x86_64 4:5.16.3-299.el7_9 will be installed
---> Package perl-threads.x86_64 0:1.87-4.el7 will be installed
---> Package perl-threads-shared.x86_64 0:1.43-6.el7 will be installed
--> Running transaction check
---> Package dejavu-sans-fonts.noarch 0:2.33-6.el7 will be installed
--> Processing Dependency: dejavu-fonts-common = 2.33-6.el7 for package: dejavu-sans-fonts-2.33-6.el7.noarch
---> Package fontpackages-filesystem.noarch 0:1.44-8.el7 will be installed
---> Package libX11-common.noarch 0:1.6.7-3.el7_9 will be installed
---> Package libxcb.x86_64 0:1.13-1.el7 will be installed
--> Processing Dependency: libXau.so.6()(64bit) for package: libxcb-1.13-1.el7.x86_64
---> Package perl-Encode.x86_64 0:2.51-7.el7 will be installed
---> Package perl-Pod-Escapes.noarch 1:1.04-299.el7_9 will be installed
---> Package perl-Pod-Usage.noarch 0:1.63-3.el7 will be installed
--> Processing Dependency: perl(Pod::Text) >= 3.15 for package: perl-Pod-Usage-1.63-3.el7.noarch
--> Processing Dependency: perl-Pod-Perldoc for package: perl-Pod-Usage-1.63-3.el7.noarch
---> Package perl-Text-ParseWords.noarch 0:3.29-4.el7 will be installed
--> Running transaction check
---> Package dejavu-fonts-common.noarch 0:2.33-6.el7 will be installed
---> Package libXau.x86_64 0:1.0.8-2.1.el7 will be installed
---> Package perl-Pod-Perldoc.noarch 0:3.20-4.el7 will be installed
--> Processing Dependency: perl(parent) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
--> Processing Dependency: perl(HTTP::Tiny) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
--> Processing Dependency: groff-base for package: perl-Pod-Perldoc-3.20-4.el7.noarch
---> Package perl-podlators.noarch 0:2.5.1-3.el7 will be installed
--> Running transaction check
---> Package groff-base.x86_64 0:1.22.2-8.el7 will be installed
---> Package perl-HTTP-Tiny.noarch 0:0.033-3.el7 will be installed
---> Package perl-parent.noarch 1:0.225-244.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

================================================================================
 Package                       Arch     Version                 Repository
                                                                           Size
================================================================================
Installing:
 nginx                         x86_64   1:1.16.1-3.el7          epel      563 k
Installing for dependencies:
 centos-indexhtml              noarch   7-9.el7.centos          base       92 k
 centos-logos                  noarch   70.0.6-3.el7.centos     base       21 M
 dejavu-fonts-common           noarch   2.33-6.el7              base       64 k
 dejavu-sans-fonts             noarch   2.33-6.el7              base      1.4 M
 fontconfig                    x86_64   2.13.0-4.3.el7          base      254 k
 fontpackages-filesystem       noarch   1.44-8.el7              base      9.9 k
 freetype                      x86_64   2.8-14.el7_9.1          updates   380 k
 gd                            x86_64   2.0.35-27.el7_9         updates   146 k
 gperftools-libs               x86_64   2.6.1-1.el7             base      272 k
 groff-base                    x86_64   1.22.2-8.el7            base      942 k
 libX11                        x86_64   1.6.7-3.el7_9           updates   607 k
 libX11-common                 noarch   1.6.7-3.el7_9           updates   164 k
 libXau                        x86_64   1.0.8-2.1.el7           base       29 k
 libXpm                        x86_64   3.5.12-1.el7            base       55 k
 libjpeg-turbo                 x86_64   1.2.90-8.el7            base      135 k
 libpng                        x86_64   2:1.5.13-8.el7          base      213 k
 libxcb                        x86_64   1.13-1.el7              base      214 k
 libxslt                       x86_64   1.1.28-6.el7            base      242 k
 make                          x86_64   1:3.82-24.el7           base      421 k
 nginx-all-modules             noarch   1:1.16.1-3.el7          epel       20 k
 nginx-filesystem              noarch   1:1.16.1-3.el7          epel       21 k
 nginx-mod-http-image-filter   x86_64   1:1.16.1-3.el7          epel       30 k
 nginx-mod-http-perl           x86_64   1:1.16.1-3.el7          epel       39 k
 nginx-mod-http-xslt-filter    x86_64   1:1.16.1-3.el7          epel       29 k
 nginx-mod-mail                x86_64   1:1.16.1-3.el7          epel       57 k
 nginx-mod-stream              x86_64   1:1.16.1-3.el7          epel       85 k
 openssl                       x86_64   1:1.0.2k-21.el7_9       updates   493 k
 openssl11-libs                x86_64   1:1.1.1g-3.el7          epel      1.5 M
 perl                          x86_64   4:5.16.3-299.el7_9      updates   8.0 M
 perl-Carp                     noarch   1.26-244.el7            base       19 k
 perl-Encode                   x86_64   2.51-7.el7              base      1.5 M
 perl-Exporter                 noarch   5.68-3.el7              base       28 k
 perl-File-Path                noarch   2.09-2.el7              base       26 k
 perl-File-Temp                noarch   0.23.01-3.el7           base       56 k
 perl-Filter                   x86_64   1.49-3.el7              base       76 k
 perl-Getopt-Long              noarch   2.40-3.el7              base       56 k
 perl-HTTP-Tiny                noarch   0.033-3.el7             base       38 k
 perl-PathTools                x86_64   3.40-5.el7              base       82 k
 perl-Pod-Escapes              noarch   1:1.04-299.el7_9        updates    52 k
 perl-Pod-Perldoc              noarch   3.20-4.el7              base       87 k
 perl-Pod-Simple               noarch   1:3.28-4.el7            base      216 k
 perl-Pod-Usage                noarch   1.63-3.el7              base       27 k
 perl-Scalar-List-Utils        x86_64   1.27-248.el7            base       36 k
 perl-Socket                   x86_64   2.010-5.el7             base       49 k
 perl-Storable                 x86_64   2.45-3.el7              base       77 k
 perl-Text-ParseWords          noarch   3.29-4.el7              base       14 k
 perl-Time-HiRes               x86_64   4:1.9725-3.el7          base       45 k
 perl-Time-Local               noarch   1.2300-2.el7            base       24 k
 perl-constant                 noarch   1.27-2.el7              base       19 k
 perl-libs                     x86_64   4:5.16.3-299.el7_9      updates   690 k
 perl-macros                   x86_64   4:5.16.3-299.el7_9      updates    44 k
 perl-parent                   noarch   1:0.225-244.el7         base       12 k
 perl-podlators                noarch   2.5.1-3.el7             base      112 k
 perl-threads                  x86_64   1.87-4.el7              base       49 k
 perl-threads-shared           x86_64   1.43-6.el7              base       39 k
Updating for dependencies:
 openssl-libs                  x86_64   1:1.0.2k-21.el7_9       updates   1.2 M

Transaction Summary
================================================================================
Install  1 Package  (+55 Dependent packages)
Upgrade             (  1 Dependent package)

Total download size: 42 M
Downloading packages:
Delta RPMs disabled because /usr/bin/applydeltarpm not installed.
warning: /var/cache/yum/x86_64/7/epel/packages/nginx-1.16.1-3.el7.x86_64.rpm: Header V4 RSA/SHA256 Signature, key ID 352c64e5: NOKEY
Public key for nginx-1.16.1-3.el7.x86_64.rpm is not installed
--------------------------------------------------------------------------------
Total                                               13 MB/s |  42 MB  00:03
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-7
Importing GPG key 0x352C64E5:
 Userid     : "Fedora EPEL (7) <epel@fedoraproject.org>"
 Fingerprint: 91e9 7d7c 4a5e 96f1 7f3e 888f 6a2f aea2 352c 64e5
 Package    : epel-release-7-11.noarch (@extras)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-EPEL-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 2:libpng-1.5.13-8.el7.x86_64                                1/58
  Installing : freetype-2.8-14.el7_9.1.x86_64                              2/58
  Installing : fontpackages-filesystem-1.44-8.el7.noarch                   3/58
  Installing : dejavu-fonts-common-2.33-6.el7.noarch                       4/58
  Installing : dejavu-sans-fonts-2.33-6.el7.noarch                         5/58
  Installing : fontconfig-2.13.0-4.3.el7.x86_64                            6/58
  Installing : libXau-1.0.8-2.1.el7.x86_64                                 7/58
  Installing : libxcb-1.13-1.el7.x86_64                                    8/58
  Installing : 1:openssl11-libs-1.1.1g-3.el7.x86_64                        9/58
  Installing : 1:nginx-filesystem-1.16.1-3.el7.noarch                     10/58
  Installing : libxslt-1.1.28-6.el7.x86_64                                11/58
  Updating   : 1:openssl-libs-1.0.2k-21.el7_9.x86_64                      12/58
  Installing : libX11-common-1.6.7-3.el7_9.noarch                         13/58
  Installing : libX11-1.6.7-3.el7_9.x86_64                                14/58
  Installing : libXpm-3.5.12-1.el7.x86_64                                 15/58
  Installing : libjpeg-turbo-1.2.90-8.el7.x86_64                          16/58
  Installing : gd-2.0.35-27.el7_9.x86_64                                  17/58
  Installing : 1:make-3.82-24.el7.x86_64                                  18/58
  Installing : 1:openssl-1.0.2k-21.el7_9.x86_64                           19/58
  Installing : centos-indexhtml-7-9.el7.centos.noarch                     20/58
  Installing : centos-logos-70.0.6-3.el7.centos.noarch                    21/58
  Installing : groff-base-1.22.2-8.el7.x86_64                             22/58
  Installing : 1:perl-parent-0.225-244.el7.noarch                         23/58
  Installing : perl-HTTP-Tiny-0.033-3.el7.noarch                          24/58
  Installing : perl-podlators-2.5.1-3.el7.noarch                          25/58
  Installing : perl-Pod-Perldoc-3.20-4.el7.noarch                         26/58
  Installing : 1:perl-Pod-Escapes-1.04-299.el7_9.noarch                   27/58
  Installing : perl-Encode-2.51-7.el7.x86_64                              28/58
  Installing : perl-Text-ParseWords-3.29-4.el7.noarch                     29/58
  Installing : perl-Pod-Usage-1.63-3.el7.noarch                           30/58
  Installing : perl-threads-1.87-4.el7.x86_64                             31/58
  Installing : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                      32/58
  Installing : perl-Exporter-5.68-3.el7.noarch                            33/58
  Installing : perl-constant-1.27-2.el7.noarch                            34/58
  Installing : perl-Socket-2.010-5.el7.x86_64                             35/58
  Installing : perl-Filter-1.49-3.el7.x86_64                              36/58
  Installing : perl-Time-Local-1.2300-2.el7.noarch                        37/58
  Installing : perl-Carp-1.26-244.el7.noarch                              38/58
  Installing : 4:perl-macros-5.16.3-299.el7_9.x86_64                      39/58
  Installing : perl-Storable-2.45-3.el7.x86_64                            40/58
  Installing : perl-PathTools-3.40-5.el7.x86_64                           41/58
  Installing : perl-threads-shared-1.43-6.el7.x86_64                      42/58
  Installing : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 43/58
  Installing : 1:perl-Pod-Simple-3.28-4.el7.noarch                        44/58
  Installing : perl-File-Temp-0.23.01-3.el7.noarch                        45/58
  Installing : perl-File-Path-2.09-2.el7.noarch                           46/58
  Installing : 4:perl-libs-5.16.3-299.el7_9.x86_64                        47/58
  Installing : perl-Getopt-Long-2.40-3.el7.noarch                         48/58
  Installing : 4:perl-5.16.3-299.el7_9.x86_64                             49/58
  Installing : gperftools-libs-2.6.1-1.el7.x86_64                         50/58
  Installing : 1:nginx-mod-mail-1.16.1-3.el7.x86_64                       51/58
  Installing : 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64           52/58
  Installing : 1:nginx-mod-stream-1.16.1-3.el7.x86_64                     53/58
  Installing : 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64          54/58
  Installing : 1:nginx-1.16.1-3.el7.x86_64                                55/58
  Installing : 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64                  56/58
  Installing : 1:nginx-all-modules-1.16.1-3.el7.noarch                    57/58
  Cleanup    : 1:openssl-libs-1.0.2k-19.el7.x86_64                        58/58
  Verifying  : perl-HTTP-Tiny-0.033-3.el7.noarch                           1/58
  Verifying  : fontconfig-2.13.0-4.3.el7.x86_64                            2/58
  Verifying  : 1:nginx-mod-mail-1.16.1-3.el7.x86_64                        3/58
  Verifying  : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                       4/58
  Verifying  : perl-threads-1.87-4.el7.x86_64                              5/58
  Verifying  : perl-Exporter-5.68-3.el7.noarch                             6/58
  Verifying  : perl-constant-1.27-2.el7.noarch                             7/58
  Verifying  : perl-PathTools-3.40-5.el7.x86_64                            8/58
  Verifying  : gperftools-libs-2.6.1-1.el7.x86_64                          9/58
  Verifying  : perl-Socket-2.010-5.el7.x86_64                             10/58
  Verifying  : fontpackages-filesystem-1.44-8.el7.noarch                  11/58
  Verifying  : groff-base-1.22.2-8.el7.x86_64                             12/58
  Verifying  : centos-logos-70.0.6-3.el7.centos.noarch                    13/58
  Verifying  : 1:perl-parent-0.225-244.el7.noarch                         14/58
  Verifying  : gd-2.0.35-27.el7_9.x86_64                                  15/58
  Verifying  : centos-indexhtml-7-9.el7.centos.noarch                     16/58
  Verifying  : perl-Filter-1.49-3.el7.x86_64                              17/58
  Verifying  : perl-File-Temp-0.23.01-3.el7.noarch                        18/58
  Verifying  : 1:perl-Pod-Simple-3.28-4.el7.noarch                        19/58
  Verifying  : 1:nginx-mod-http-perl-1.16.1-3.el7.x86_64                  20/58
  Verifying  : perl-Time-Local-1.2300-2.el7.noarch                        21/58
  Verifying  : libxcb-1.13-1.el7.x86_64                                   22/58
  Verifying  : 1:make-3.82-24.el7.x86_64                                  23/58
  Verifying  : 1:perl-Pod-Escapes-1.04-299.el7_9.noarch                   24/58
  Verifying  : perl-Pod-Perldoc-3.20-4.el7.noarch                         25/58
  Verifying  : 1:openssl-1.0.2k-21.el7_9.x86_64                           26/58
  Verifying  : libXpm-3.5.12-1.el7.x86_64                                 27/58
  Verifying  : libjpeg-turbo-1.2.90-8.el7.x86_64                          28/58
  Verifying  : perl-Carp-1.26-244.el7.noarch                              29/58
  Verifying  : perl-threads-shared-1.43-6.el7.x86_64                      30/58
  Verifying  : libX11-common-1.6.7-3.el7_9.noarch                         31/58
  Verifying  : libX11-1.6.7-3.el7_9.x86_64                                32/58
  Verifying  : 4:perl-macros-5.16.3-299.el7_9.x86_64                      33/58
  Verifying  : perl-Storable-2.45-3.el7.x86_64                            34/58
  Verifying  : 1:nginx-mod-http-xslt-filter-1.16.1-3.el7.x86_64           35/58
  Verifying  : dejavu-sans-fonts-2.33-6.el7.noarch                        36/58
  Verifying  : perl-Scalar-List-Utils-1.27-248.el7.x86_64                 37/58
  Verifying  : 2:libpng-1.5.13-8.el7.x86_64                               38/58
  Verifying  : 1:nginx-mod-stream-1.16.1-3.el7.x86_64                     39/58
  Verifying  : 1:openssl-libs-1.0.2k-21.el7_9.x86_64                      40/58
  Verifying  : freetype-2.8-14.el7_9.1.x86_64                             41/58
  Verifying  : perl-Encode-2.51-7.el7.x86_64                              42/58
  Verifying  : perl-Pod-Usage-1.63-3.el7.noarch                           43/58
  Verifying  : dejavu-fonts-common-2.33-6.el7.noarch                      44/58
  Verifying  : perl-podlators-2.5.1-3.el7.noarch                          45/58
  Verifying  : 4:perl-5.16.3-299.el7_9.x86_64                             46/58
  Verifying  : perl-File-Path-2.09-2.el7.noarch                           47/58
  Verifying  : libxslt-1.1.28-6.el7.x86_64                                48/58
  Verifying  : 1:nginx-filesystem-1.16.1-3.el7.noarch                     49/58
  Verifying  : 1:nginx-1.16.1-3.el7.x86_64                                50/58
  Verifying  : 1:openssl11-libs-1.1.1g-3.el7.x86_64                       51/58
  Verifying  : libXau-1.0.8-2.1.el7.x86_64                                52/58
  Verifying  : 1:nginx-all-modules-1.16.1-3.el7.noarch                    53/58
  Verifying  : perl-Getopt-Long-2.40-3.el7.noarch                         54/58
  Verifying  : perl-Text-ParseWords-3.29-4.el7.noarch                     55/58
  Verifying  : 1:nginx-mod-http-image-filter-1.16.1-3.el7.x86_64          56/58
  Verifying  : 4:perl-libs-5.16.3-299.el7_9.x86_64                        57/58
  Verifying  : 1:openssl-libs-1.0.2k-19.el7.x86_64                        58/58

Installed:
  nginx.x86_64 1:1.16.1-3.el7

Dependency Installed:
  centos-indexhtml.noarch 0:7-9.el7.centos
  centos-logos.noarch 0:70.0.6-3.el7.centos
  dejavu-fonts-common.noarch 0:2.33-6.el7
  dejavu-sans-fonts.noarch 0:2.33-6.el7
  fontconfig.x86_64 0:2.13.0-4.3.el7
  fontpackages-filesystem.noarch 0:1.44-8.el7
  freetype.x86_64 0:2.8-14.el7_9.1
  gd.x86_64 0:2.0.35-27.el7_9
  gperftools-libs.x86_64 0:2.6.1-1.el7
  groff-base.x86_64 0:1.22.2-8.el7
  libX11.x86_64 0:1.6.7-3.el7_9
  libX11-common.noarch 0:1.6.7-3.el7_9
  libXau.x86_64 0:1.0.8-2.1.el7
  libXpm.x86_64 0:3.5.12-1.el7
  libjpeg-turbo.x86_64 0:1.2.90-8.el7
  libpng.x86_64 2:1.5.13-8.el7
  libxcb.x86_64 0:1.13-1.el7
  libxslt.x86_64 0:1.1.28-6.el7
  make.x86_64 1:3.82-24.el7
  nginx-all-modules.noarch 1:1.16.1-3.el7
  nginx-filesystem.noarch 1:1.16.1-3.el7
  nginx-mod-http-image-filter.x86_64 1:1.16.1-3.el7
  nginx-mod-http-perl.x86_64 1:1.16.1-3.el7
  nginx-mod-http-xslt-filter.x86_64 1:1.16.1-3.el7
  nginx-mod-mail.x86_64 1:1.16.1-3.el7
  nginx-mod-stream.x86_64 1:1.16.1-3.el7
  openssl.x86_64 1:1.0.2k-21.el7_9
  openssl11-libs.x86_64 1:1.1.1g-3.el7
  perl.x86_64 4:5.16.3-299.el7_9
  perl-Carp.noarch 0:1.26-244.el7
  perl-Encode.x86_64 0:2.51-7.el7
  perl-Exporter.noarch 0:5.68-3.el7
  perl-File-Path.noarch 0:2.09-2.el7
  perl-File-Temp.noarch 0:0.23.01-3.el7
  perl-Filter.x86_64 0:1.49-3.el7
  perl-Getopt-Long.noarch 0:2.40-3.el7
  perl-HTTP-Tiny.noarch 0:0.033-3.el7
  perl-PathTools.x86_64 0:3.40-5.el7
  perl-Pod-Escapes.noarch 1:1.04-299.el7_9
  perl-Pod-Perldoc.noarch 0:3.20-4.el7
  perl-Pod-Simple.noarch 1:3.28-4.el7
  perl-Pod-Usage.noarch 0:1.63-3.el7
  perl-Scalar-List-Utils.x86_64 0:1.27-248.el7
  perl-Socket.x86_64 0:2.010-5.el7
  perl-Storable.x86_64 0:2.45-3.el7
  perl-Text-ParseWords.noarch 0:3.29-4.el7
  perl-Time-HiRes.x86_64 4:1.9725-3.el7
  perl-Time-Local.noarch 0:1.2300-2.el7
  perl-constant.noarch 0:1.27-2.el7
  perl-libs.x86_64 4:5.16.3-299.el7_9
  perl-macros.x86_64 4:5.16.3-299.el7_9
  perl-parent.noarch 1:0.225-244.el7
  perl-podlators.noarch 0:2.5.1-3.el7
  perl-threads.x86_64 0:1.87-4.el7
  perl-threads-shared.x86_64 0:1.43-6.el7

Dependency Updated:
  openssl-libs.x86_64 1:1.0.2k-21.el7_9

Complete!
Removing intermediate container c75c12a6a037
 ---> 8bb02bf6f4c5
Step 4/5 : COPY index.html /usr/share/nginx/html
 ---> 65029268118f
Step 5/5 : ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
 ---> Running in a4d13073fda9
Removing intermediate container a4d13073fda9
 ---> d487de1ee98f
Successfully built d487de1ee98f
Successfully tagged cyberblack28/sample-nginx:latest
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
cyberblack28/sample-nginx   latest              d487de1ee98f   7 minutes ago    558MB
msb                         latest              ab4aa580a77f   27 minutes ago   7.55MB
<none>                      <none>              1ffb1b502170   30 minutes ago   7.55MB
<none>                      <none>              0b9401537c5c   30 minutes ago   303MB
cyberblack28/sample-nginx   <none>              59cb9dc6b0e2   44 minutes ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   4 days ago       301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

```dockerコマンド
# docker image tag 59cb9dc6b0e2 cyberblack28/sample-nginx:1.0
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
cyberblack28/sample-nginx   latest              d487de1ee98f   10 minutes ago   558MB
msb                         latest              ab4aa580a77f   29 minutes ago   7.55MB
<none>                      <none>              1ffb1b502170   32 minutes ago   7.55MB
<none>                      <none>              0b9401537c5c   32 minutes ago   303MB
cyberblack28/sample-nginx   1.0                 59cb9dc6b0e2   47 minutes ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   4 days ago       301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

```dockerコマンド
# docker image push cyberblack28/sample-nginx:1.0
The push refers to repository [docker.io/cyberblack28/sample-nginx]
09ff5c3f3c16: Layer already exists
6724fad9c54d: Layer already exists
8ec59fb6f8fd: Layer already exists
174f56854903: Layer already exists
1.0: digest: sha256:c80b9c9ef042fecee96e7bd4b4e8456d113757133fe575adfe61768a434aaae7 size: 1161
```

### 3.3.2 コンテナ起動と操作

#### docker container run

```dockerコマンド
# docker container run --name sample-nginx -d -p 8080:80 cyberblack28/sample-nginx
c04bf6b03914a9e9984fc5f7b7c61d4bc1bd7d3295d06bc6bdf5bdca94b20f91
```

```dockerコマンド
# curl http://localhost:8080
<!DOCTYPE html>
<html>
<head>
<title>First Docker Build</title>
</head>
<body>
<p>Happy, Container !!</p>
</body>
```

```dockerコマンド
# docker container ls
CONTAINER ID   IMAGE                       COMMAND                  CREATED              STATUS              PORTS                                   NAMES
c04bf6b03914   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   About a minute ago   Up About a minute   0.0.0.0:8080->80/tcp, :::8080->80/tcp   sample-nginx
```

#### docker container exec

```dockerコマンド
# docker container exec -it sample-nginx /bin/bash
[root@c04bf6b03914 /]# ls
anaconda-post.log  bin  boot  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
[root@c04bf6b03914 /]# exit
exit
```

#### docker container stop / docker container start

```dockerコマンド
# docker container stop sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE                       COMMAND                  CREATED         STATUS                      PORTS     NAMES
c04bf6b03914   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   4 minutes ago   Exited (0) 15 seconds ago             sample-nginx
```

```dockerコマンド
# docker container start sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE                       COMMAND                  CREATED         STATUS         PORTS                                   NAMES
c04bf6b03914   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   6 minutes ago   Up 8 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   sample-nginx
```

#### docker container cp

```linuxコマンド
# cd ../3-3-2-01
```

```linuxコマンド
# cat copy.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Command Practice</title>
</head>
<body>
<p>Copy Command !!</p>
</body>
</html>
```

```dockerコマンド
# docker container cp copy.html sample-nginx:/usr/share/nginx/html
```

```linuxコマンド
# curl http://localhost:8080/copy.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Command Practice</title>
</head>
<body>
<p>Copy Command !!</p>
</body>
</html>
```

#### docker container stats

```dockerコマンド
# docker container stats sample-nginx
CONTAINER ID   NAME           CPU %     MEM USAGE / LIMIT     MEM %     NET I/O          BLOCK I/O   PIDS
c04bf6b03914   sample-nginx   0.00%     3.777MiB / 3.597GiB   0.10%     1.8kB / 4.25kB   0B / 0B     2
```

#### docker container inspect

```dockerコマンド
# docker container inspect sample-nginx
[
    {
        "Id": "c04bf6b03914a9e9984fc5f7b7c61d4bc1bd7d3295d06bc6bdf5bdca94b20f91",
        "Created": "2021-05-11T14:10:21.422499965Z",
        "Path": "/usr/sbin/nginx",
        "Args": [
            "-g",
            "daemon off;"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 6007,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2021-05-11T14:16:22.027985577Z",
            "FinishedAt": "2021-05-11T14:14:24.906533817Z"
        },
        "Image": "sha256:d487de1ee98f14142bad549aaa2da82abce7956cd9b1df74ca4181bd15a9742b",
        "ResolvConfPath": "/var/lib/docker/containers/c04bf6b03914a9e9984fc5f7b7c61d4bc1bd7d3295d06bc6bdf5bdca94b20f91/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/c04bf6b03914a9e9984fc5f7b7c61d4bc1bd7d3295d06bc6bdf5bdca94b20f91/hostname",
        "HostsPath": "/var/lib/docker/containers/c04bf6b03914a9e9984fc5f7b7c61d4bc1bd7d3295d06bc6bdf5bdca94b20f91/hosts",
        "LogPath": "/var/lib/docker/containers/c04bf6b03914a9e9984fc5f7b7c61d4bc1bd7d3295d06bc6bdf5bdca94b20f91/c04bf6b03914a9e9984fc5f7b7c61d4bc1bd7d3295d06bc6bdf5bdca94b20f91-json.log",
        "Name": "/sample-nginx",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {
                "80/tcp": [
                    {
                        "HostIp": "",
                        "HostPort": "8080"
                    }
                ]
            },
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "CgroupnsMode": "host",
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/ea3cc825606d4cb5bd29df639ee0c15414c642f2e9cf83648bd07081bb69bd52-init/diff:/var/lib/docker/overlay2/b1120b8fdaca705a4a223ea9e87c86a57b271ee85113ca4fd8b4d8686903bd14/diff:/var/lib/docker/overlay2/cc8294ec50e74ed4d7259dbb14d7d5350f11219fef3132fdf7a7defe104c6c49/diff:/var/lib/docker/overlay2/8345f2cffaf35cac263a78620253cda00f0e90cda0254d5f40e01715fd8b500b/diff:/var/lib/docker/overlay2/cb428ed041a5379fc7c3f2d23c07a810e7ba0fcafd7be954ba4533027e44fa2f/diff",
                "MergedDir": "/var/lib/docker/overlay2/ea3cc825606d4cb5bd29df639ee0c15414c642f2e9cf83648bd07081bb69bd52/merged",
                "UpperDir": "/var/lib/docker/overlay2/ea3cc825606d4cb5bd29df639ee0c15414c642f2e9cf83648bd07081bb69bd52/diff",
                "WorkDir": "/var/lib/docker/overlay2/ea3cc825606d4cb5bd29df639ee0c15414c642f2e9cf83648bd07081bb69bd52/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "c04bf6b03914",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "80/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": null,
            "Image": "cyberblack28/sample-nginx",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": [
                "/usr/sbin/nginx",
                "-g",
                "daemon off;"
            ],
            "OnBuild": null,
            "Labels": {
                "org.label-schema.build-date": "20201113",
                "org.label-schema.license": "GPLv2",
                "org.label-schema.name": "CentOS Base Image",
                "org.label-schema.schema-version": "1.0",
                "org.label-schema.vendor": "CentOS",
                "org.opencontainers.image.created": "2020-11-13 00:00:00+00:00",
                "org.opencontainers.image.licenses": "GPL-2.0-only",
                "org.opencontainers.image.title": "CentOS Base Image",
                "org.opencontainers.image.vendor": "CentOS"
            }
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "65d3096cef5d0cad0accbda7b4cd136b1df6812f66d7889313a6208eb0d2700c",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "80/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "8080"
                    },
                    {
                        "HostIp": "::",
                        "HostPort": "8080"
                    }
                ]
            },
            "SandboxKey": "/var/run/docker/netns/65d3096cef5d",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "825c23dece94b84062ee4c38ed5a50a90b6ed9abd072551fc8fd738539dd2f68",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "5fd8d761d73b3b8bbffc09caf632f23f68b487683facfdca0f11a3c8d43e6d98",
                    "EndpointID": "825c23dece94b84062ee4c38ed5a50a90b6ed9abd072551fc8fd738539dd2f68",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
        }
    }
]
```

#### docker container rm

```dockerコマンド
# docker container stop sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS                      PORTS     NAMES
c04bf6b03914   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   15 minutes ago   Exited (0) 11 seconds ago             sample-nginx
```

```dockerコマンド
# docker container rm sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

#### docker container logs

```linuxコマンド
# cd ../3-3-2-02
```

```dockerコマンド
# cat Dockerfile
#CentOS7のベースイメージをPull
FROM centos:7

#yumでepel-releaseをインストール
RUN yum -y install epel-release

#yumでnginxをインストール
RUN yum -y install nginx

#ホスト側のindex.htmlをコピー
COPY index.html /usr/share/nginx/html

# アクセスログとエラーログを標準出力に出力
RUN ln -sf /dev/stdout /var/log/nginx/access.log && ln -sf /dev/stderr /var/log/nginx/error.log

#コンテナ起動時に実行するコマンド
ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
```

```dockerコマンド
# docker image build -t cyberblack28/sample-nginx .
Sending build context to Docker daemon  3.584kB
Step 1/6 : FROM centos:7
 ---> 8652b9f0cb4c
Step 2/6 : RUN yum -y install epel-release
 ---> Using cache
 ---> 358d8684562a
Step 3/6 : RUN yum -y install nginx
 ---> Using cache
 ---> 8bb02bf6f4c5
Step 4/6 : COPY index.html /usr/share/nginx/html
 ---> Using cache
 ---> 65029268118f
Step 5/6 : RUN ln -sf /dev/stdout /var/log/nginx/access.log && ln -sf /dev/stderr /var/log/nginx/error.log
 ---> Running in a83974f3d382
Removing intermediate container a83974f3d382
 ---> be9bcfa9ac93
Step 6/6 : ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
 ---> Running in 13da396949af
Removing intermediate container 13da396949af
 ---> e721ae4e9582
Successfully built e721ae4e9582
Successfully tagged cyberblack28/sample-nginx:latest
```

```dockerコマンド
# docker container run --name sample-nginx -d -p 8080:80 cyberblack28/sample-nginx
089df1d46ec28ddec9f5c14af8b67681eaccdb09235fb5666900ff5aba2d92e0
```

```dockerコマンド
# docker container ls
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS          PORTS                                   NAMES
089df1d46ec2   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   26 seconds ago   Up 25 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   sample-nginx
```

```linuxコマンド
# curl http://localhost:8080
<!DOCTYPE html>
<html>
<head>
<title>First Docker Build</title>
</head>
<body>
<p>Happy, Container !!</p>
</body>
</html>
```

```dockerコマンド
# docker container logs sample-nginx
172.17.0.1 - - [11/May/2021:14:35:45 +0000] "GET / HTTP/1.1" 200 126 "-" "curl/7.68.0" "-"
```

```dockerコマンド
# docker container stop sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE                       COMMAND                  CREATED         STATUS                      PORTS     NAMES
089df1d46ec2   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   3 minutes ago   Exited (0) 13 seconds ago             sample-nginx
```

```dockerコマンド
# docker container rm sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

### 3・3・4 バインドマウント

#### バインドマウントの作成

```linuxコマンド
# cd ../3-3-4-01
```

```dockerコマンド
# cat htdocs/index.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Bind mount</title>
</head>
<body>
<h1>ホストのhtmlファイルをマウントしています！</h1>
</body>
</html>
```

```dockerコマンド
# docker run --name bind-nginx -d -p 8080:80 --mount type=bind,source=/root/container-develop-environment-construction-guide/Chapter03/3-3-4-01/htdocs,target=/usr/share/nginx/html nginx
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
f7ec5a41d630: Pull complete
aa1efa14b3bf: Pull complete
b78b95af9b17: Pull complete
c7d6bca2b8dc: Pull complete
cf16cd8e71e0: Pull complete
0241c68333ef: Pull complete
Digest: sha256:75a55d33ecc73c2a242450a9f1cc858499d468f077ea942867e662c247b5e412
Status: Downloaded newer image for nginx:latest
1d92e100abcfb9301c2f77cc5f707d44820e19bcb970a9f56b96cfaa3be8ee1a
```

```linuxコマンド
# curl http://localhost:8080/index.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Bind mount</title>
</head>
<body>
<h1>ホストのhtmlファイルをマウントしています！</h1>
</body>
```

```dockerコマンド
# docker run --name bind-nginx -d -p 8080:80 -v /root/container-develop-environment-construction-guide/Chapter03/3-3-4-01/htdocs:/usr/share/nginx/html nginx
aaab73a2a5de468200062b5ddcd42bb3e4c4a801053a5341a84a00a13176bf5b
# curl http:////localhost:8080/index.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Bind mount</title>
</head>
<body>
<h1>ホストのhtmlファイルをマウントしています！</h1>
</body>
```

```linuxコマンド
# curl http://localhost:8080/index.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Bind mount</title>
</head>
<body>
<h1>ホストのhtmlファイルをマウントしています！</h1>
</body>
```

#### マウントファイルの変更

```linuxコマンド
# vim htdocs/index.html
```

```dockerコマンド
# curl http://localhost:8080/index.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Bind mount</title>
</head>
<body>
<h1>ホストのhtmlファイルを更新しました！</h1>
</body>
</html>
```

```dockerコマンド
# docker container stop bind-nginx
bind-nginx
```

```dockerコマンド
# docker container rm bind-nginx
bind-nginx
```

### 3・3・5 ボリューム（Volume）

#### ボリュームの作成

```dockerコマンド
# docker volume create htdocs
htdocs
```

```dockerコマンド
# docker volume ls
DRIVER VOLUME NAME
local htdocs
```

```dockerコマンド
# docker container run --name volume-nginx -d -p 8080:80 --mount source=htdocs,target=/usr/share/nginx/html nginx
978821cb0c6c259e3a22245844b86b0f9216f3268af9bf95f5007f19cb975a26
```

```dockerコマンド
# docker container run --name volume-nginx -d -p 8080:80 -v htdocs:/usr/share/nginx/html nginx
67b44f15a9f8495a6fadbeb17bdab87d7a9c68b32978cda4d23269c86cd9751d
```

#### ホスト側ディレクトリの確認

```linuxコマンド
# ls /var/lib/docker/volumes/htdocs/_data/
50x.html  index.html
```

```linuxコマンド
cp -p /root/container-develop-environment-construction-guide/Chapter03/3-3-4-02/volume.html /var/lib/docker/volumes/htdocs/_data/volume.html
```

```linuxコマンド
# curl http://localhost:8080/volume.html
<!DOCTYPE html>
<html>
<head>
<title>Docker Volume</title>
</head>
<body>
<h1>Volumeでhtmlファイルを作成しました！</h1>
</body>
</html>
```

```dockerコマンド
# docker container exec -it volume-nginx /bin/bash
root@978821cb0c6c:/# touch /usr/share/nginx/html/test
root@978821cb0c6c:/# ls /usr/share/nginx/html/
50x.html  index.html  test  volume.html
root@978821cb0c6c:/# exit
exit
```

```dockerコマンド
# ls /var/lib/docker/volumes/htdocs/_data/
50x.html  index.html  test  volume.html
```

```dockerコマンド
# docker container stop volume-nginx
volume-nginx
```

```dockerコマンド
# docker container rm volume-nginx
volume-nginx
```

### 3・3・6 一時ファイルシステムのマウント（tmpfs mount）

#### 一時ファイルシステムの作成

```dockerコマンド
# docker run -itd --name tmpfs-nginx --mount type=tmpfs,destination=/root/tmp,tmpfs-size=10,tmpfs-mode=755 nginx
8708997d938db1c978494621e12a63255b191dc0261a0825eb2c20e82dd1003f
```

```dockerコマンド
# docker container exec -it tmpfs-nginx /bin/bash
root@8708997d938d:/# ls /root
tmp
root@8708997d938d:/# exit
exit
```

```dockerコマンド
# docker container stop tmpfs-nginx
tmpfs-nginx
```

```dockerコマンド
# docker container rm tmpfs-nginx
tmpfs-nginx
```

#### 3・3・7 データボリュームコンテナ（Data Volume Container）

### 作業ディレクトリの作成

```linuxコマンド
# cd /root
```

```linuxコマンド
# mkdir /tmp/data-volume
```

```linuxコマンド
# cd /tmp/data-volume
```

```linuxコマンド
# mkdir share
```

```linuxコマンド
# touch share/share-file.txt
```

```dockerコマンド
# docker run -it -d --name ubuntu -v /tmp/data-volume/share:/tmp/data ubuntu
Unable to find image 'ubuntu:latest' locally
latest: Pulling from library/ubuntu
345e3491a907: Pull complete
57671312ef6f: Pull complete
5e9250ddb7d0: Pull complete
Digest: sha256:cf31af331f38d1d7158470e095b132acd126a7180a54f263d386da88eb681d93
Status: Downloaded newer image for ubuntu:latest
da7bbd309b62f4da57217629ee3fd026475b5515ccd01f8c8ddc65bdc2c2bf49
```

### コンテナの作成

```dockerコマンド
# docker run -it -d --name share01 --volumes-from ubuntu ubuntu
a7f94c95c139799c8be2351d4d3ab1fb106b6c5cac95b10d39b6aa69759a7901
```

```dockerコマンド
# docker run -it -d --name share02 --volumes-from ubuntu ubuntu
92f5643a37cb94febed0d3c902750331f7348fd6fda7a810878bdd62f04f6700
```

```dockerコマンド
# docker container ls
CONTAINER ID   IMAGE     COMMAND       CREATED         STATUS         PORTS     NAMES
92f5643a37cb   ubuntu    "/bin/bash"   2 minutes ago   Up 2 minutes             share02
a7f94c95c139   ubuntu    "/bin/bash"   3 minutes ago   Up 2 minutes             share01
da7bbd309b62   ubuntu    "/bin/bash"   5 minutes ago   Up 5 minutes             ubuntu
```

### データ共有の確認

```dockerコマンド
# docker exec -it ubuntu /bin/bash
root@da7bbd309b62:/# ls /tmp/data/
share-file.txt
root@da7bbd309b62:/# touch /tmp/data/new-sharefile
root@da7bbd309b62:/# ls /tmp/data
new-sharefile  share-file.txt
root@da7bbd309b62:/# exit
exit
```

```dockerコマンド
# docker exec -it share01 /bin/bash
root@a7f94c95c139:/# ls /tmp/data
new-sharefile  share-file.txt
root@a7f94c95c139:/# exit
exit
```

```dockerコマンド
# docker exec -it share02 /bin/bash
root@92f5643a37cb:/# ls /tmp/data
new-sharefile  share-file.txt
root@92f5643a37cb:/# exit
exit
```

```dockerコマンド
# docker container stop ubuntu
ubuntu
```

```dockerコマンド
# docker container stop share01
share01
```

```dockerコマンド
# docker container stop share02
share02
```

```dockerコマンド
# docker container rm ubuntu
ubuntu
```

```dockerコマンド
# docker container rm share01
share01
```

```dockerコマンド
# docker container rm share02
share02
```

### 3・3・8 コンテナのネットワーク

```dockerコマンド
# docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
2572d881e2b2   bridge    bridge    local
10be770dee84   host      host      local
d2180dcb719e   none      null      local
```

```linuxコマンド
# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: ens4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1460 qdisc fq_codel state UP group default qlen 1000
    link/ether 42:01:0a:92:0f:df brd ff:ff:ff:ff:ff:ff
    inet 10.146.15.223/32 scope global dynamic ens4
       valid_lft 3123sec preferred_lft 3123sec
    inet6 fe80::4001:aff:fe92:fdf/64 scope link
       valid_lft forever preferred_lft forever
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 02:42:b8:9d:1c:48 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:b8ff:fe9d:1c48/64 scope link
       valid_lft forever preferred_lft forever
```

### 3・3・9 ブリッジネットワークとアプリケーション

```dockerコマンド
# docker network create wordpress-network
e0d035cf8c1dda9485db8226ecd50b4877e0cf730c0e3a97a93b950ace59f073
```

```dockerコマンド
# docker network ls
NETWORK ID     NAME                DRIVER    SCOPE
f041f9a7a51c   bridge              bridge    local
0db81ad2118c   host                host      local
6d25324cb0e8   none                null      local
e0d035cf8c1d   wordpress-network   bridge    local
```

```dockerコマンド
# docker run -d --name mysql \
--network wordpress-network \
-e MYSQL_ROOT_PASSWORD=wordpress \
-e MYSQL_DATABASE=wordpress \
-e MYSQL_USER=wordpress \
-e MYSQL_PASSWORD=wordpress \
mysql:8.0.25
Unable to find image 'mysql:8.0.25' locally
8.0.25: Pulling from library/mysql
69692152171a: Pull complete
1651b0be3df3: Pull complete
951da7386bc8: Pull complete
0f86c95aa242: Pull complete
37ba2d8bd4fe: Pull complete
6d278bb05e94: Pull complete
497efbd93a3e: Pull complete
f7fddf10c2c2: Pull complete
16415d159dfb: Pull complete
0e530ffc6b73: Pull complete
b0a4a1a77178: Pull complete
cd90f92aa9ef: Pull complete
Digest: sha256:d50098d7fcb25b1fcb24e2d3247cae3fc55815d64fec640dc395840f8fa80969
Status: Downloaded newer image for mysql:8.0.25
35d594a48e41d87f3cb820e8e4e7d7f8afbce65615f5a74aaa64c2fb66336f2a
```

```dockerコマンド
# docker run -d --name wordpress \
--network wordpress-network \
-p 8080:80 \
-e WORDPRESS_DB_HOST=mysql:3306 \
-e WORDPRESS_DB_NAME=wordpress \
-e WORDPRESS_DB_USER=wordpress \
-e WORDPRESS_DB_PASSWORD=wordpress \
wordpress:php7.4-apache
Unable to find image 'wordpress:php7.4-apache' locally
php7.4-apache: Pulling from library/wordpress
69692152171a: Already exists
2040822db325: Pull complete
9b4ca5ae9dfa: Pull complete
ac1fe7c6d966: Pull complete
5b26fc9ce030: Pull complete
3492f4769444: Pull complete
1dec05775a74: Pull complete
77107a42338e: Pull complete
f58e4093c52a: Pull complete
d32715f578d3: Pull complete
7a73fb2558ce: Pull complete
667b573fcff7: Pull complete
75e2da936ffe: Pull complete
759622df3a7b: Pull complete
c2f98ef02756: Pull complete
50e11300b0a6: Pull complete
de37513870b9: Pull complete
f25501789abc: Pull complete
0cf8e3442952: Pull complete
d45ce270a7e6: Pull complete
534cdc5a6ea6: Pull complete
Digest: sha256:e9da0d6c867249f364cd2292ea0dd01d7281e8dfbcc3e4b39b823f9a790b237b
Status: Downloaded newer image for wordpress:php7.4-apache
534c3dabf55661cf2914d21caa1b242e4074a80082884c58b2101122bbbc5f32
```

```dockerコマンド
# docker container ls
CONTAINER ID   IMAGE                     COMMAND                  CREATED          STATUS          PORTS                                   NAMES
534c3dabf556   wordpress:php7.4-apache   "docker-entrypoint.s…"   9 minutes ago    Up 9 minutes    0.0.0.0:8080->80/tcp, :::8080->80/tcp   wordpress
35d594a48e41   mysql:8.0.25              "docker-entrypoint.s…"   10 minutes ago   Up 10 minutes   3306/tcp, 33060/tcp                     mysql
```

```dockerコマンド
# curl http://localhost:8080/wp-admin/install.php
<!DOCTYPE html>
<html lang="en-US" xml:lang="en-US">
<head>
        <meta name="viewport" content="width=device-width" />
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta name="robots" content="noindex,nofollow" />
        <title>WordPress &rsaquo; Installation</title>
        <link rel='stylesheet' id='dashicons-css'  href='http://localhost:8080/wp-includes/css/dashicons.min.css?ver=5.7.2' type='text/css' media='all' />
<link rel='stylesheet' id='buttons-css'  href='http://localhost:8080/wp-includes/css/buttons.min.css?ver=5.7.2' type='text/css' media='all' />
<link rel='stylesheet' id='forms-css'  href='http://localhost:8080/wp-admin/css/forms.min.css?ver=5.7.2' type='text/css' media='all' />
<link rel='stylesheet' id='l10n-css'  href='http://localhost:8080/wp-admin/css/l10n.min.css?ver=5.7.2' type='text/css' media='all' />
<link rel='stylesheet' id='install-css'  href='http://localhost:8080/wp-admin/css/install.min.css?ver=5.7.2' type='text/css' media='all' />
</head>
<body class="wp-core-ui language-chooser">
<p id="logo">WordPress</p>

        <form id="setup" method="post" action="?step=1"><label class='screen-reader-text' for='language'>Select a default language</label>
<select size='14' name='language' id='language'>
<option value="" lang="en" selected="selected" data-continue="Continue" data-installed="1">English (United States)</option>
<option value="af" lang="af" data-continue="Gaan voort">Afrikaans</option>
<option value="ar" lang="ar" data-continue="المتابعة">العربية</option>
<option value="ary" lang="ar" data-continue="المتابعة">العربية المغربية</option>
<option value="as" lang="as" data-continue="Continue">অসমীয়া</option>
<option value="azb" lang="az" data-continue="Continue">گؤنئی آذربایجان</option>
<option value="az" lang="az" data-continue="Davam">Azərbaycan dili</option>
<option value="bel" lang="be" data-continue="Працягнуць">Беларуская мова</option>
<option value="bg_BG" lang="bg" data-continue="Продължение">Български</option>
<option value="bn_BD" lang="bn" data-continue="এগিয়ে চল.">বাংলা</option>
<option value="bo" lang="bo" data-continue="མུ་མཐུད།">བོད་ཡིག</option>
<option value="bs_BA" lang="bs" data-continue="Nastavi">Bosanski</option>
<option value="ca" lang="ca" data-continue="Continua">Català</option>
<option value="ceb" lang="ceb" data-continue="Padayun">Cebuano</option>
<option value="cs_CZ" lang="cs" data-continue="Pokračovat">Čeština</option>
<option value="cy" lang="cy" data-continue="Parhau">Cymraeg</option>
<option value="da_DK" lang="da" data-continue="Forts&#230;t">Dansk</option>
<option value="de_DE" lang="de" data-continue="Fortfahren">Deutsch</option>
<option value="de_DE_formal" lang="de" data-continue="Fortfahren">Deutsch (Sie)</option>
<option value="de_CH" lang="de" data-continue="Fortfahren">Deutsch (Schweiz)</option>
<option value="de_CH_informal" lang="de" data-continue="Weiter">Deutsch (Schweiz, Du)</option>
<option value="de_AT" lang="de" data-continue="Weiter">Deutsch (Österreich)</option>
<option value="dsb" lang="dsb" data-continue="Dalej">Dolnoserbšćina</option>
<option value="dzo" lang="dz" data-continue="Continue">རྫོང་ཁ</option>
<option value="el" lang="el" data-continue="Συνέχεια">Ελληνικά</option>
<option value="en_AU" lang="en" data-continue="Continue">English (Australia)</option>
<option value="en_ZA" lang="en" data-continue="Continue">English (South Africa)</option>
<option value="en_GB" lang="en" data-continue="Continue">English (UK)</option>
<option value="en_CA" lang="en" data-continue="Continue">English (Canada)</option>
<option value="en_NZ" lang="en" data-continue="Continue">English (New Zealand)</option>
<option value="eo" lang="eo" data-continue="Daŭrigi">Esperanto</option>
<option value="es_UY" lang="es" data-continue="Continuar">Español de Uruguay</option>
<option value="es_VE" lang="es" data-continue="Continuar">Español de Venezuela</option>
<option value="es_ES" lang="es" data-continue="Continuar">Español</option>
<option value="es_MX" lang="es" data-continue="Continuar">Español de México</option>
<option value="es_EC" lang="es" data-continue="Continuar">Español de Ecuador</option>
<option value="es_AR" lang="es" data-continue="Continuar">Español de Argentina</option>
<option value="es_CL" lang="es" data-continue="Continuar">Español de Chile</option>
<option value="es_CO" lang="es" data-continue="Continuar">Español de Colombia</option>
<option value="es_CR" lang="es" data-continue="Continuar">Español de Costa Rica</option>
<option value="es_PE" lang="es" data-continue="Continuar">Español de Perú</option>
<option value="es_PR" lang="es" data-continue="Continuar">Español de Puerto Rico</option>
<option value="es_GT" lang="es" data-continue="Continuar">Español de Guatemala</option>
<option value="et" lang="et" data-continue="Jätka">Eesti</option>
<option value="eu" lang="eu" data-continue="Jarraitu">Euskara</option>
<option value="fa_IR" lang="fa" data-continue="ادامه">فارسی</option>
<option value="fa_AF" lang="fa" data-continue="ادامه">(فارسی (افغانستان</option>
<option value="fi" lang="fi" data-continue="Jatka">Suomi</option>
<option value="fr_FR" lang="fr" data-continue="Continuer">Français</option>
<option value="fr_CA" lang="fr" data-continue="Continuer">Français du Canada</option>
<option value="fr_BE" lang="fr" data-continue="Continuer">Français de Belgique</option>
<option value="fur" lang="fur" data-continue="Continue">Friulian</option>
<option value="gd" lang="gd" data-continue="Lean air adhart">Gàidhlig</option>
<option value="gl_ES" lang="gl" data-continue="Continuar">Galego</option>
<option value="gu" lang="gu" data-continue="ચાલુ રાખવું">ગુજરાતી</option>
<option value="haz" lang="haz" data-continue="ادامه">هزاره گی</option>
<option value="he_IL" lang="he" data-continue="להמשיך">עִבְרִית</option>
<option value="hi_IN" lang="hi" data-continue="जारी">हिन्दी</option>
<option value="hr" lang="hr" data-continue="Nastavi">Hrvatski</option>
<option value="hsb" lang="hsb" data-continue="Dale">Hornjoserbšćina</option>
<option value="hu_HU" lang="hu" data-continue="Tovább">Magyar</option>
<option value="hy" lang="hy" data-continue="Շարունակել">Հայերեն</option>
<option value="id_ID" lang="id" data-continue="Lanjutkan">Bahasa Indonesia</option>
<option value="is_IS" lang="is" data-continue="Áfram">Íslenska</option>
<option value="it_IT" lang="it" data-continue="Continua">Italiano</option>
<option value="ja" lang="ja" data-continue="続ける">日本語</option>
<option value="jv_ID" lang="jv" data-continue="Nutugne">Basa Jawa</option>
<option value="ka_GE" lang="ka" data-continue="გაგრძელება">ქართული</option>
<option value="kab" lang="kab" data-continue="Continuer">Taqbaylit</option>
<option value="kk" lang="kk" data-continue="Жалғастыру">Қазақ тілі</option>
<option value="km" lang="km" data-continue="បន្ត">ភាសាខ្មែរ</option>
<option value="kn" lang="kn" data-continue="ಮುಂದುವರೆಸಿ">ಕನ್ನಡ</option>
<option value="ko_KR" lang="ko" data-continue="계속">한국어</option>
<option value="ckb" lang="ku" data-continue="به‌رده‌وام به‌">كوردی‎</option>
<option value="lo" lang="lo" data-continue="ຕໍ່">ພາສາລາວ</option>
<option value="lt_LT" lang="lt" data-continue="Tęsti">Lietuvių kalba</option>
<option value="lv" lang="lv" data-continue="Turpināt">Latviešu valoda</option>
<option value="mk_MK" lang="mk" data-continue="Продолжи">Македонски јазик</option>
<option value="ml_IN" lang="ml" data-continue="തുടരുക">മലയാളം</option>
<option value="mn" lang="mn" data-continue="Үргэлжлүүлэх">Монгол</option>
<option value="mr" lang="mr" data-continue="सुरु ठेवा">मराठी</option>
<option value="ms_MY" lang="ms" data-continue="Teruskan">Bahasa Melayu</option>
<option value="my_MM" lang="my" data-continue="ဆက်လက်လုပ်ေဆာင်ပါ။">ဗမာစာ</option>
<option value="nb_NO" lang="nb" data-continue="Fortsett">Norsk bokmål</option>
<option value="ne_NP" lang="ne" data-continue="जारीराख्नु ">नेपाली</option>
<option value="nl_NL_formal" lang="nl" data-continue="Doorgaan">Nederlands (Formeel)</option>
<option value="nl_BE" lang="nl" data-continue="Doorgaan">Nederlands (België)</option>
<option value="nl_NL" lang="nl" data-continue="Doorgaan">Nederlands</option>
<option value="nn_NO" lang="nn" data-continue="Hald fram">Norsk nynorsk</option>
<option value="oci" lang="oc" data-continue="Contunhar">Occitan</option>
<option value="pa_IN" lang="pa" data-continue="ਜਾਰੀ ਰੱਖੋ">ਪੰਜਾਬੀ</option>
<option value="pl_PL" lang="pl" data-continue="Kontynuuj">Polski</option>
<option value="ps" lang="ps" data-continue="دوام">پښتو</option>
<option value="pt_PT" lang="pt" data-continue="Continuar">Português</option>
<option value="pt_AO" lang="pt" data-continue="Continuar">Português de Angola</option>
<option value="pt_BR" lang="pt" data-continue="Continuar">Português do Brasil</option>
<option value="pt_PT_ao90" lang="pt" data-continue="Continuar">Português (AO90)</option>
<option value="rhg" lang="rhg" data-continue="Continue">Ruáinga</option>
<option value="ro_RO" lang="ro" data-continue="Continuă">Română</option>
<option value="ru_RU" lang="ru" data-continue="Продолжить">Русский</option>
<option value="sah" lang="sah" data-continue="Салҕаа">Сахалыы</option>
<option value="snd" lang="sd" data-continue="اڳتي هلو">سنڌي</option>
<option value="si_LK" lang="si" data-continue="දිගටම කරගෙන යන්න">සිංහල</option>
<option value="sk_SK" lang="sk" data-continue="Pokračovať">Slovenčina</option>
<option value="skr" lang="skr" data-continue="جاری رکھو">سرائیکی</option>
<option value="sl_SI" lang="sl" data-continue="Nadaljujte">Slovenščina</option>
<option value="sq" lang="sq" data-continue="Vazhdo">Shqip</option>
<option value="sr_RS" lang="sr" data-continue="Настави">Српски језик</option>
<option value="sv_SE" lang="sv" data-continue="Fortsätt">Svenska</option>
<option value="sw" lang="sw" data-continue="Endelea">Kiswahili</option>
<option value="szl" lang="szl" data-continue="Kōntynuować">Ślōnskŏ gŏdka</option>
<option value="ta_IN" lang="ta" data-continue="தொடரவும்">தமிழ்</option>
<option value="ta_LK" lang="ta" data-continue="தொடர்க">தமிழ்</option>
<option value="te" lang="te" data-continue="కొనసాగించు">తెలుగు</option>
<option value="th" lang="th" data-continue="ต่อไป">ไทย</option>
<option value="tl" lang="tl" data-continue="Magpatuloy">Tagalog</option>
<option value="tr_TR" lang="tr" data-continue="Devam">Türkçe</option>
<option value="tt_RU" lang="tt" data-continue="дәвам итү">Татар теле</option>
<option value="tah" lang="ty" data-continue="Continue">Reo Tahiti</option>
<option value="ug_CN" lang="ug" data-continue="داۋاملاشتۇرۇش">ئۇيغۇرچە</option>
<option value="uk" lang="uk" data-continue="Продовжити">Українська</option>
<option value="ur" lang="ur" data-continue="جاری رکھیں">اردو</option>
<option value="uz_UZ" lang="uz" data-continue="Продолжить">O‘zbekcha</option>
<option value="vi" lang="vi" data-continue="Tiếp tục">Tiếng Việt</option>
<option value="zh_HK" lang="zh" data-continue="繼續">香港中文版 </option>
<option value="zh_TW" lang="zh" data-continue="繼續">繁體中文</option>
<option value="zh_CN" lang="zh" data-continue="继续">简体中文</option>
</select>
<p class="step"><span class="spinner"></span><input id="language-continue" type="submit" class="button button-primary button-large" value="Continue" /></p></form><script type="text/javascript">var t = document.getElementById('weblog_title'); if (t){ t.focus(); }</script>
        <script type='text/javascript' src='http://localhost:8080/wp-includes/js/jquery/jquery.min.js?ver=3.5.1' id='jquery-core-js'></script>
<script type='text/javascript' src='http://localhost:8080/wp-includes/js/jquery/jquery-migrate.min.js?ver=3.3.2' id='jquery-migrate-js'></script>
<script type='text/javascript' src='http://localhost:8080/wp-admin/js/language-chooser.min.js?ver=5.7.2' id='language-chooser-js'></script>
<script type="text/javascript">
jQuery( function( $ ) {
        $( '.hide-if-no-js' ).removeClass( 'hide-if-no-js' );
} );
</script>
</body>
</html>
```

```dockerコマンド
# docker container stop wordpress
wordpress
```

```dockerコマンド
# docker container stop mysql
mysql
```

## 3.4 コンテナとイメージの一括削除

```dockerコマンド
# docker container prune
WARNING! This will remove all stopped containers.
Are you sure you want to continue? [y/N] y
Deleted Containers:
81b6b2f479493149c99a991a8ada112cda3b71bb29332a3edc85cb33ee583979
2461c5ba5fc414a549174dcc036b335a532a94768bbc539efb1959945306aeb8
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

```dockerコマンド
# docker image prune -a
WARNING! This will remove all images without at least one container associated to them.
Are you sure you want to continue? [y/N] y
Deleted Images:
untagged: golang:1.16.4-alpine3.13
untagged: golang@sha256:4dd403b2e7a689adc5b7110ba9cd5da43d216cfcfccfbe2b35680effcf336c7e
untagged: wordpress:php7.4-apache
untagged: wordpress@sha256:208def35d7fcbbfd76df18997ce6cd5a5221c0256221b7fdaba41c575882d4f0
deleted: sha256:5684f8405e21e6037008f0ab18cd993a3325d12e67bfefe8c8e8423adb6aa041
deleted: sha256:2f10a2cf2aabcf3551f7516c1bec517e6b98f399d2d240dd6a1438dbd22a587f
deleted: sha256:1162800780ba2efe0e0bede36f67f96a155c73c1b5d7bdbf4075882bd4a20d91
deleted: sha256:5d2517fd005a8e9c8551229d29a9eb33c96e0dcfae9276ee4b1404f4f9c86ca6
deleted: sha256:c47f3ea5383b0f937578ffebbd708ed4c8754987b9813c0229682f1d2fd90f88
deleted: sha256:15c528b756d5b8cb9669bd551f89387617330d1c71d82359f41c30e95a199394
deleted: sha256:bd762c71564719c96d44b22d87911bdf725eafd54f67c584a9cfe72b8f2b1d4c
deleted: sha256:5245c88a9af29dd7be42d013919e2c398f068626e75222151fa477a85e58dc34
deleted: sha256:518d55141a5ba8f9b5ec4eb2cbaab7da1ecfae2d8a195c7f60aaed0382f0085b
deleted: sha256:283b010b05aacd1280b14c0ef37f3fa15dea574aa0c37e586b0428402bcd4960
deleted: sha256:a7e3a4c92fbe1da5a7a5233f7d3b8c7c0d2f11ed3607fa9365ca0e40c91ad4c6
deleted: sha256:7fd504e828bdccb43fafc901bc24737dc87f89af7a2d4ca12d514cec3b6d895a
deleted: sha256:883f1bd033776372784e83e55a07e4f7ef16e59bd13ef82b278591c05130a09e
deleted: sha256:24cfcd0051d4c4c70c47e8bc07706088331789f81b1790a18635bcb44b84a117
deleted: sha256:3cd13337872e7120ec64550353c38056a73e8083b06b3324e100d605a75c1bc8
deleted: sha256:fa5ebc51c424d4c89c1afb63e2b5a4d1adb1872caa3f56d2b6d48c953ea4ab0a
deleted: sha256:50b7b271050e44fdef5234756e73472b97efb53025b20e2b5bcea858fce0d3cf
deleted: sha256:711e9ecbeeedaa6a7c5457ed8f01c16e69c5a0acce81ffbe31402233e08a5697
deleted: sha256:0c094d6ea549215b0594dcf11cd851e4704c779a979e2bbfd7621ebfd1af8391
deleted: sha256:4c8ef953bb0bd31d5a7f4b67aade3284b0af4f840cc7dd019940def3fac1aca3
deleted: sha256:018306c9cf4d8c3f63d9ebf5aa23c44a88c19c31956bc570c4136c55a2b95785
untagged: cyberblack28/sample-nginx@sha256:85bfa643cc5fb6cb85544993a216ddf092fc2eb8b4e0c34598ebf86b2f6e1977
deleted: sha256:dd05031e366cc022080df34843ee166f9a707769ab9344545c268c7a8666e0b4
untagged: alpine:3.13
untagged: alpine@sha256:69e70a79f2d41ab5d637de98c1e0b055206ba40a8145e7bddb55ccc04e13cf8f
untagged: nginx:latest
untagged: nginx@sha256:75a55d33ecc73c2a242450a9f1cc858499d468f077ea942867e662c247b5e412
deleted: sha256:62d49f9bab67f7c70ac3395855bf01389eb3175b374e621f6f191bf31b54cd5b
deleted: sha256:3444fb58dc9e8338f6da71c1040e8ff532f25fab497312f95dcee0f756788a84
deleted: sha256:f85cfdc7ca97d8856cd4fa916053084e2e31c7e53ed169577cef5cb1b8169ccb
deleted: sha256:704bf100d7f16255a2bc92e925f7007eef0bd3947af4b860a38aaffc3f992eae
deleted: sha256:d5955c2e658d1432abb023d7d6d1128b0aa12481b976de7cbde4c7a31310f29b
deleted: sha256:11126fda59f7f4bf9bf08b9d24c9ea45a1194f3d61ae2a96af744c97eae71cbf
deleted: sha256:7e718b9c0c8c2e6420fe9c4d1d551088e314fe923dce4b2caf75891d82fb227d
untagged: msb:latest
deleted: sha256:e4292341aafcb3b0caf6c0c491eb47aeaacf06226a90e26ed0d5ab33e29d6702
untagged: cyberblack28/sample-nginx:latest
deleted: sha256:fcd99f4c6ca470456beecad54ec4d86d27e879f5cf636e2385f47798c784e2bf
deleted: sha256:cb9de963da772df8d352541f1259f6a83e90e05338c61925d9162599acff4e6a
deleted: sha256:6a2b94bba6b530d080e4e66a1cf732ffbd151ef1faa7974a86f66ff0eeea75de
deleted: sha256:becc96f8d94ac0087040c143c00ead382429290f51440d4babf8b3ab9f791f0e
deleted: sha256:775c2404df10c39eb979ff6a4ec9b144e6e10a8ea46c9a6ea5e49270aacbebbf
deleted: sha256:b769804e22fe64741e85fc6efaf3debffe805ab8821cfd8b817ad02c66c630bb
deleted: sha256:6bb656a37288230a6e10e0d905aa05c1f7cf89608aff5905f899a9ced85c4c1d
deleted: sha256:aef1753346c0b64695d557bb949e312d564bff15b59741a39783e981b83fbc2e
deleted: sha256:511e90be7dd122ebe340c02fc54af3b7cc9e58e609311173cd514a440ad7a458
deleted: sha256:4c6515bdcb216167a9bd7d9ddff4a70d526409d6f00bf8fcd4cd35104f6b7cdc
deleted: sha256:c16ee41da1f206e505ed8c7368c293c952febbc56aa6dfe294551849edb78931
deleted: sha256:fc75b017ed41c1c5d5e8a80ba0b7ee5d7b273309b63593c9993e9f6af79edbf4
deleted: sha256:6dbb9cc54074106d46d4ccb330f2a40a682d49dda5f4844962b7dce9fe44aaec
untagged: mysql:8.0.22
untagged: mysql@sha256:0fd2898dc1c946b34dceaccc3b80d38b1049285c1dab70df7480de62265d6213
deleted: sha256:d4c3cafb11d573699728f9e7de10d1b976089b01298c0360e03f0afd9a1a8b36
deleted: sha256:3ce5f6a2175f88412c0e8241e9298e721cface0773ea2fd70f1fdaf0e606c7fa
deleted: sha256:f840f44fb69007f9b78a9bb552882753066b8e4a7835aa8f1be83ba50466e346
deleted: sha256:d5c65d80478a6f623a715e67e39c4756ce0d4d09348881d1e86b87033322c70c
deleted: sha256:2a33bb3c3abf34893e689effc0f204c998eb1c8c137050808b9b2725589e3c66
deleted: sha256:537f09811c9ed9704ccbda6a316f3a2bf346530a296753de6545025b2dfde532
deleted: sha256:d08533f1e2acc40ad561a46fc6a76b54c739e6b24f077c183c5709e0a6885312
deleted: sha256:4f9d91a4728e833d1062fb65a792f06e22e425f63824f260c8b5a64b776ddc38
deleted: sha256:20bf4c759d1b0d0e6286d2145453af4e0e1b7ba3d4efa3b8bce46817ad4109de
deleted: sha256:a9371bbdf16ac95cc72555c6ad42f79b9f03a82d964fe89d52bdc5f335a5f42a
deleted: sha256:5b02130e449d94f51e8ff6e5f7d24802246198749ed9eb064631e63833cd8f1d
deleted: sha256:ab74465b38bc1acb16c23091df32c5b7033ed55783386cb57acae8efff9f4b37
deleted: sha256:cb42413394c4059335228c137fe884ff3ab8946a014014309676c25e3ac86864
untagged: cyberblack28/sample-nginx:1.0
untagged: cyberblack28/sample-nginx@sha256:b51f647395095e5364928c0bf53b82a05c73ea15fd757fe641c07f152e4b43d6
deleted: sha256:c0adbbc67c7579f10c0e02d156ae5848b2f6b20cc7e9fbba7c824e0b942d3785
deleted: sha256:5a229a3ebc2f92712884dc511c6b8011ffeb669854fc25607f64f3b33a39ad2a
deleted: sha256:5bfb8527698f016527c8408877281b44d246d6dc3fdcfb52265ec8bb5f2ddee1
deleted: sha256:e11d86dec8982063e784b2356306cbbb532c4f8029f6b12c0950f57b3485ce35
untagged: ubuntu:latest
untagged: ubuntu@sha256:cf31af331f38d1d7158470e095b132acd126a7180a54f263d386da88eb681d93
deleted: sha256:7e0aa2d69a153215c790488ed1fcec162015e973e49962d438e18249d16fa9bd
deleted: sha256:3dd8c8d4fd5b59d543c8f75a67cdfaab30aef5a6d99aea3fe74d8cc69d4e7bf2
deleted: sha256:8d8dceacec7085abcab1f93ac1128765bc6cf0caac334c821e01546bd96eb741
deleted: sha256:ccdbb80308cc5ef43b605ac28fac29c6a597f89f5a169bbedbb8dec29c987439
untagged: centos:7
untagged: centos@sha256:0f4ec88e21daf75124b8a9e5ca03c37a5e937e0e108a255d890492430789b60e
deleted: sha256:8652b9f0cb4c0599575e5a003f5906876e10c1ceb2ab9fe1786712dac14a50cf
deleted: sha256:174f5685490326fc0a1c0f5570b8663732189b327007e47ff13d2ca59673db02
deleted: sha256:fb2a599580ef628da15da5245f1ff065e703505b5539e0761b12b17f53b0e546
deleted: sha256:18077913be34b197239204227e465a1ffea8cfa34c1643fdbd6739560c38baf6
deleted: sha256:2b8a7666257b699a24ea39451d27121652fd560a78e38bdb7bb78ec2c27d4ca1
deleted: sha256:c7cdd6ac183911c469eaf9b2e8d1932d65eb52eb292471c22f80174a304aa281
deleted: sha256:722a834ff95bfd3dac4bc5aae498c245eb76b98108ed9de4293df2596e60cf1a

Total reclaimed space: 2.148GB
```

```dockerコマンド
# docker image ls
REPOSITORY   TAG       IMAGE ID   CREATED   SIZE
```

```linuxコマンド
# exit
logout
$ exit
logout
Connection to 34.84.148.90 closed.
$
```