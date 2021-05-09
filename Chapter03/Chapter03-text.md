# 第3章 コンテナアプリケーション開発ライフサイクルBuild・Ship・Run

## 3.1 Build / コンテナビルド

### 3.1.1 Dockerfile作成からのビルドの実行

#### Dockerfileの作成

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
 ---> Running in b26dbab58922
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: ty1.mirror.newmediaexpress.com
 * extras: ty1.mirror.newmediaexpress.com
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
Removing intermediate container b26dbab58922
 ---> 1e01263f61e1
Step 3/5 : RUN yum -y install nginx
 ---> Running in fa33902183c4
Loaded plugins: fastestmirror, ovl
Loading mirror speeds from cached hostfile
 * base: ty1.mirror.newmediaexpress.com
 * epel: d2lzkl7pfhq30w.cloudfront.net
 * extras: ty1.mirror.newmediaexpress.com
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
Total                                               12 MB/s |  42 MB  00:03
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
Removing intermediate container fa33902183c4
 ---> 7dfdc7289736
Step 4/5 : COPY index.html /usr/share/nginx/html
 ---> 1404cd12fc17
Step 5/5 : ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
 ---> Running in 97c1bedd6aef
Removing intermediate container 97c1bedd6aef
 ---> c0adbbc67c75
Successfully built c0adbbc67c75
Successfully tagged cyberblack28/sample-nginx:latest
```

#### コンテナイメージの確認

```dockerコマンド
# docker image ls
REPOSITORY                  TAG       IMAGE ID       CREATED         SIZE
cyberblack28/sample-nginx   latest    c0adbbc67c75   8 minutes ago   558MB
centos                      7         8652b9f0cb4c   5 months ago    204MB
```

#### レイヤの確認

```dockerコマンド
# docker image history cyberblack28/sample-nginx
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
c0adbbc67c75   10 minutes ago   /bin/sh -c #(nop)  ENTRYPOINT ["/usr/sbin/ng…   0B
1404cd12fc17   10 minutes ago   /bin/sh -c #(nop) COPY file:3e884108a93ee9b1…   126B
7dfdc7289736   10 minutes ago   /bin/sh -c yum -y install nginx                 238MB
1e01263f61e1   11 minutes ago   /bin/sh -c yum -y install epel-release          116MB
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
# cat main.go
package main

import "fmt"

func main() {
        fmt.Println("Let's start multi-stage builds !!")
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
 ---> 2b8a7666257b
Step 3/6 : RUN go build -o /msb ./main.go
 ---> Running in f4464264c3ab
Removing intermediate container f4464264c3ab
 ---> fb2a599580ef
Step 4/6 : FROM alpine:3.13
3.13: Pulling from library/alpine
540db60ca938: Already exists
Digest: sha256:69e70a79f2d41ab5d637de98c1e0b055206ba40a8145e7bddb55ccc04e13cf8f
Status: Downloaded newer image for alpine:3.13
 ---> 6dbb9cc54074
Step 5/6 : COPY --from=builder /msb /usr/local/bin/msb
 ---> c16ee41da1f2
Step 6/6 : ENTRYPOINT ["/usr/local/bin/msb"]
 ---> Running in 3c95ca6464cb
Removing intermediate container 3c95ca6464cb
 ---> 4c6515bdcb21
Successfully built 4c6515bdcb21
Successfully tagged msb:latest
```

```dockerコマンド
# docker container run -it --rm msb
Let's start multi-stage builds !!
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
msb                         latest              4c6515bdcb21   2 minutes ago    7.55MB
<none>                      <none>              fb2a599580ef   2 minutes ago    303MB
cyberblack28/sample-nginx   latest              c0adbbc67c75   24 minutes ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   33 hours ago     301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

### 3.1.4 ビルドツール

#### BuildKitを利用したビルド

```dockerコマンド
# DOCKER_BUILDKIT=1 docker image build -t msb -f Dockerfile-msb .
[+] Building 2.3s (11/11) FINISHED
 => [internal] load .dockerignore                                                                                                                                                                0.1s
 => => transferring context: 2B                                                                                                                                                                  0.0s
 => [internal] load build definition from Dockerfile-msb                                                                                                                                         0.1s
 => => transferring dockerfile: 589B                                                                                                                                                             0.0s
 => [internal] load metadata for docker.io/library/alpine:3.13                                                                                                                                   0.0s
 => [internal] load metadata for docker.io/library/golang:1.16.4-alpine3.13                                                                                                                      0.0s
 => [stage-1 1/2] FROM docker.io/library/alpine:3.13                                                                                                                                             0.0s
 => [internal] load build context                                                                                                                                                                0.1s
 => => transferring context: 134B                                                                                                                                                                0.0s
 => [builder 1/3] FROM docker.io/library/golang:1.16.4-alpine3.13                                                                                                                                0.1s
 => [builder 2/3] COPY ./main.go ./                                                                                                                                                              0.1s
 => [builder 3/3] RUN go build -o /msb ./main.go                                                                                                                                                 1.4s
 => [stage-1 2/2] COPY --from=builder /msb /usr/local/bin/msb                                                                                                                                    0.1s
 => exporting to image                                                                                                                                                                           0.0s
 => => exporting layers                                                                                                                                                                          0.0s
 => => writing image sha256:e4292341aafcb3b0caf6c0c491eb47aeaacf06226a90e26ed0d5ab33e29d6702                                                                                                     0.0s
 => => naming to docker.io/library/msb                                                                                                                                                           0.0s
```

```dockerコマンド
# docker image ls
msb                         latest              e4292341aafc   About a minute ago   7.55MB
<none>                      <none>              4c6515bdcb21   5 hours ago          7.55MB
<none>                      <none>              fb2a599580ef   5 hours ago          303MB
cyberblack28/sample-nginx   latest              c0adbbc67c75   5 hours ago          558MB
golang                      1.16.4-alpine3.13   722a834ff95b   38 hours ago         301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago          5.61MB
centos                      7                   8652b9f0cb4c   5 months ago         204MB
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
8967fe867ced: Pushed
e1f7893ab52e: Pushed
acfaf1472707: Pushed
174f56854903: Mounted from library/centos
latest: digest: sha256:b51f647395095e5364928c0bf53b82a05c73ea15fd757fe641c07f152e4b43d6 size: 1161
```

#### IMAGE IDの確認

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED        SIZE
msb                         latest              e4292341aafc   17 hours ago   7.55MB
<none>                      <none>              4c6515bdcb21   21 hours ago   7.55MB
<none>                      <none>              fb2a599580ef   21 hours ago   303MB
cyberblack28/sample-nginx   latest              c0adbbc67c75   22 hours ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   2 days ago     301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago    5.61MB
centos                      7                   8652b9f0cb4c   5 months ago   204MB
```

#### イメージの削除と確認

```dockerコマンド
# docker image rm c0adbbc67c75
Untagged: cyberblack28/sample-nginx:latest
Untagged: cyberblack28/sample-nginx@sha256:b51f647395095e5364928c0bf53b82a05c73ea15fd757fe641c07f152e4b43d6
Deleted: sha256:c0adbbc67c7579f10c0e02d156ae5848b2f6b20cc7e9fbba7c824e0b942d3785
Deleted: sha256:1404cd12fc173f744af2fc45304d2bc5e183dfb93ce4b5225d593d54ca44201e
Deleted: sha256:5a229a3ebc2f92712884dc511c6b8011ffeb669854fc25607f64f3b33a39ad2a
Deleted: sha256:7dfdc72897361211e00fabc1ed1684f36d6164c8b9e39f584583e0c0eb778705
Deleted: sha256:5bfb8527698f016527c8408877281b44d246d6dc3fdcfb52265ec8bb5f2ddee1
Deleted: sha256:1e01263f61e1a2f65e9e0fd5f78016ad787d6ed375918196665517f88585f054
Deleted: sha256:e11d86dec8982063e784b2356306cbbb532c4f8029f6b12c0950f57b3485ce35
```

#### Pullコマンドの実行

```dockerコマンド
# docker image pull cyberblack28/sample-nginx
Using default tag: latest
latest: Pulling from cyberblack28/sample-nginx
2d473b07cdd5: Already exists
81e72b5342fa: Pull complete
7a2a3e571b4a: Pull complete
0a4fc4cba375: Pull complete
Digest: sha256:b51f647395095e5364928c0bf53b82a05c73ea15fd757fe641c07f152e4b43d6
Status: Downloaded newer image for cyberblack28/sample-nginx:latest
docker.io/cyberblack28/sample-nginx:latest
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED        SIZE
msb                         latest              e4292341aafc   17 hours ago   7.55MB
<none>                      <none>              4c6515bdcb21   21 hours ago   7.55MB
<none>                      <none>              fb2a599580ef   21 hours ago   303MB
cyberblack28/sample-nginx   latest              c0adbbc67c75   22 hours ago   558MB
golang                      1.16.4-alpine3.13   722a834ff95b   2 days ago     301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago    5.61MB
centos                      7                   8652b9f0cb4c   5 months ago   204MB
```

#### Image Tag

```dockerコマンド
# cd ../3-2-4-01
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
 ---> Running in f724cf22bace
Loaded plugins: fastestmirror, ovl
Determining fastest mirrors
 * base: ty1.mirror.newmediaexpress.com
 * extras: ty1.mirror.newmediaexpress.com
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
Removing intermediate container f724cf22bace
 ---> aef1753346c0
Step 3/5 : RUN yum -y install nginx
 ---> Running in 274e58cb6f8e
Loaded plugins: fastestmirror, ovl
Loading mirror speeds from cached hostfile
 * base: ty1.mirror.newmediaexpress.com
 * epel: d2lzkl7pfhq30w.cloudfront.net
 * extras: ty1.mirror.newmediaexpress.com
 * updates: ty1.mirror.newmediaexpress.com
http://mirrors.kernel.org/fedora-epel/7/x86_64/repodata/5ca9c8933445f59b58a261448d7ea4a487507718559420b519feef8f2cd02dcf-updateinfo.xml.bz2: [Errno 14] HTTP Error 404 - Not Found
Trying other mirror.
To address this issue please refer to the below wiki article 

https://wiki.centos.org/yum-errors

If above article doesn't help to resolve this issue please use https://bugs.centos.org/.

http://mirrors.syringanetworks.net/fedora-epel/7/x86_64/repodata/85e65dec80c2929659de21faf9ad0463ca47ad5bde3db354cd74b7b8bde346ce-primary.sqlite.bz2: [Errno 14] HTTP Error 404 - Not Found
Trying other mirror.
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
Total                                               22 MB/s |  42 MB  00:01
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
Removing intermediate container 274e58cb6f8e
 ---> b769804e22fe
Step 4/5 : COPY index.html /usr/share/nginx/html
 ---> becc96f8d94a
Step 5/5 : ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
 ---> Running in b2e364603488
Removing intermediate container b2e364603488
 ---> dd05031e366c
Successfully built dd05031e366c
Successfully tagged cyberblack28/sample-nginx:latest
```

```dockerコマンド
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED         SIZE
cyberblack28/sample-nginx   latest              dd05031e366c   6 minutes ago   558MB
msb                         latest              e4292341aafc   17 hours ago    7.55MB
<none>                      <none>              4c6515bdcb21   22 hours ago    7.55MB
<none>                      <none>              fb2a599580ef   22 hours ago    303MB
cyberblack28/sample-nginx   <none>              c0adbbc67c75   22 hours ago    558MB
golang                      1.16.4-alpine3.13   722a834ff95b   2 days ago      301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago     5.61MB
centos                      7                   8652b9f0cb4c   5 months ago    204MB
```

```dockerコマンド
# docker image tag c0adbbc67c75 cyberblack28/sample-nginx:1.0
# docker image ls
REPOSITORY                  TAG                 IMAGE ID       CREATED          SIZE
cyberblack28/sample-nginx   latest              dd05031e366c   18 minutes ago   558MB
msb                         latest              e4292341aafc   18 hours ago     7.55MB
<none>                      <none>              4c6515bdcb21   22 hours ago     7.55MB
<none>                      <none>              fb2a599580ef   22 hours ago     303MB
cyberblack28/sample-nginx   1.0                 c0adbbc67c75   23 hours ago     558MB
golang                      1.16.4-alpine3.13   722a834ff95b   2 days ago       301MB
alpine                      3.13                6dbb9cc54074   3 weeks ago      5.61MB
centos                      7                   8652b9f0cb4c   5 months ago     204MB
```

```dockerコマンド
# docker image push cyberblack28/sample-nginx:1.0
The push refers to repository [docker.io/cyberblack28/sample-nginx]
8967fe867ced: Layer already exists
e1f7893ab52e: Layer already exists
acfaf1472707: Layer already exists
174f56854903: Layer already exists
1.0: digest: sha256:b51f647395095e5364928c0bf53b82a05c73ea15fd757fe641c07f152e4b43d6 size: 1161
```

```dockerコマンド
# docker image push cyberblack28/sample-nginx
Using default tag: latest
The push refers to repository [docker.io/cyberblack28/sample-nginx]
03fba5c533a6: Pushed
c68a3ab6b548: Pushed
cdb6104e7e5e: Pushed
174f56854903: Layer already exists
latest: digest: sha256:85bfa643cc5fb6cb85544993a216ddf092fc2eb8b4e0c34598ebf86b2f6e1977 size: 1161
```

### 3.3.2 コンテナ起動と操作

#### docker container run

```dockerコマンド
# docker container run --name sample-nginx -d -p 8080:80 cyberblack28/sample-nginx
fabb50c8e09a75df779e207f48fde429669681a2cda6255355f52b6bb5c0b204
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
CONTAINER ID   IMAGE                       COMMAND                  CREATED         STATUS         PORTS                                   NAMES
fabb50c8e09a   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   7 minutes ago   Up 7 minutes   0.0.0.0:8080->80/tcp, :::8080->80/tcp   sample-nginx
```

#### docker container exec

```dockerコマンド
# docker container exec -it sample-nginx /bin/bash
[root@fabb50c8e09a /]# ls
anaconda-post.log  bin  boot  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
[root@fabb50c8e09a /]# exit
exit
```

#### docker container stop / docker container start

```dockerコマンド
# docker container stop sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS                     PORTS     NAMES
fabb50c8e09a   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   21 minutes ago   Exited (0) 9 seconds ago             sample-nginx
```

```dockerコマンド
# docker container start sample-nginx
sample-nginx
```

```dockerコマンド
# docker container ls -a
CONTAINER ID   IMAGE                       COMMAND                  CREATED          STATUS         PORTS                                   NAMES
fabb50c8e09a   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   22 minutes ago   Up 8 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   sample-nginx
```

#### docker container cp

```linuxコマンド
# cd ../3-3-2-01
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
CONTAINER ID   NAME           CPU %     MEM USAGE / LIMIT     MEM %     NET I/O         BLOCK I/O   PIDS
fabb50c8e09a   sample-nginx   0.00%     3.789MiB / 3.597GiB   0.10%     1.87kB / 715B   0B / 0B     2
```

#### docker container inspect

```dockerコマンド
# docker container inspect sample-nginx
[
    {
        "Id": "fabb50c8e09a75df779e207f48fde429669681a2cda6255355f52b6bb5c0b204",
        "Created": "2021-05-09T07:30:34.756940885Z",
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
            "Pid": 11179,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2021-05-09T07:52:32.284350756Z",
            "FinishedAt": "2021-05-09T07:52:11.19895919Z"
        },
        "Image": "sha256:dd05031e366cc022080df34843ee166f9a707769ab9344545c268c7a8666e0b4",
        "ResolvConfPath": "/var/lib/docker/containers/fabb50c8e09a75df779e207f48fde429669681a2cda6255355f52b6bb5c0b204/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/fabb50c8e09a75df779e207f48fde429669681a2cda6255355f52b6bb5c0b204/hostname",
        "HostsPath": "/var/lib/docker/containers/fabb50c8e09a75df779e207f48fde429669681a2cda6255355f52b6bb5c0b204/hosts",
        "LogPath": "/var/lib/docker/containers/fabb50c8e09a75df779e207f48fde429669681a2cda6255355f52b6bb5c0b204/fabb50c8e09a75df779e207f48fde429669681a2cda6255355f52b6bb5c0b204-json.log",
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
                "LowerDir": "/var/lib/docker/overlay2/cfe423d4e29362968c7dcba3d82f671db6738fb12ef3e7e2ee6ec26ddaa0469b-init/diff:/var/lib/docker/overlay2/d15e58192be2b14a8027b6b1144bf43f5074f565b6b5c2cf9f75353c544a9ce8/diff:/var/lib/docker/overlay2/54980eeec746d3770946b578579d99c2fe9f86c39639db6fa24b503ffce0fcf5/diff:/var/lib/docker/overlay2/a0bb6497d5250d96620ddd030c28b54aaafa6c17063fc19078966c6efe303404/diff:/var/lib/docker/overlay2/a6f0d38fae88c5d22d01fbd9f9c1c2ec6354ccd6f718f0dd3666d726e6bf1427/diff",
                "MergedDir": "/var/lib/docker/overlay2/cfe423d4e29362968c7dcba3d82f671db6738fb12ef3e7e2ee6ec26ddaa0469b/merged",
                "UpperDir": "/var/lib/docker/overlay2/cfe423d4e29362968c7dcba3d82f671db6738fb12ef3e7e2ee6ec26ddaa0469b/diff",
                "WorkDir": "/var/lib/docker/overlay2/cfe423d4e29362968c7dcba3d82f671db6738fb12ef3e7e2ee6ec26ddaa0469b/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "fabb50c8e09a",
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
            "SandboxID": "5f39981d10458797eba9d93b23132720bc166a6dab260f7c5143544279bb5af5",
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
            "SandboxKey": "/var/run/docker/netns/5f39981d1045",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "daf5b7a7fb382f82f994eaec0480ff25688c831b5a09fe3fc732a86a086b667c",
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
                    "NetworkID": "15a685036c1a02014933dda087ed82d5d29b6dadb93893c0bfbd510eb3ba2e88",
                    "EndpointID": "daf5b7a7fb382f82f994eaec0480ff25688c831b5a09fe3fc732a86a086b667c",
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
CONTAINER ID   IMAGE                       COMMAND                  CREATED             STATUS                      PORTS     NAMES
fabb50c8e09a   cyberblack28/sample-nginx   "/usr/sbin/nginx -g …"   About an hour ago   Exited (0) 10 seconds ago             sample-nginx
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

```dockerコマンド
# cd ../3-3-2-02
# cat Dockerfile
```