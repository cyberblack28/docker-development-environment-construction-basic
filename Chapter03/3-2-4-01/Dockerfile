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