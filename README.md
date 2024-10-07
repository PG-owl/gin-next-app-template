# アプリケーション名
**gin-next-app-template**

## 概要
ginとnext.jsを使ったアプリ用のテンプレートリポジトリです

## 環境構築手順

### 前提条件
- Gitがインストールされていること
- Docker Desktopがインストールされていること（[公式サイト](https://www.docker.com/ja-jp/)を参照）

### 手順
1. 任意のディレクトリ内にGitHubリポジトリをクローンします。
	```shell
	$ git clone https://github.com/PG-owl/gin-next-app-template.git
	```

2.  以下のコマンドを実行し、Dockerコンテナをビルドして起動します。
	``` shell
	$ cd gin-next-app-template
	$ docker-compose up --build
	```

3. ブラウザで<http://localhost:8080/>にアクセスし、*Hello, Poke Quick!* と表示されれば環境構築は完了です。

* (*補足*) 
	``` shell
	// 全Dockerコンテナを停止する場合のコマンド
	$ docker stop $(docker ps -q)
	```

	``` shell
	// 停止した全Dockerコンテナを起動する場合のコマンド
	$ docker stop $(docker ps -a -q)
	```