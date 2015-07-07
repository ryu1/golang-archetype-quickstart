# Learning-Go

## Installation

1. gobrewをgithubから取得する

	Goのバージョン管理には、cryptojuice/gobrewを使用します。
	
	```
	$ git clone git://github.com/cryptojuice/gobrew.git ~/.gobrew
	```

2. .zshrcとかに以下を追加する

	```
	export PATH="$HOME/.gobrew/bin:$PATH"
	eval "$(gobrew init -)"
	```

3. 任意のバージョンのGoをインストールする

	```
	$ gobrew install 1.4.2
	$ gobrew use 1.4.2
	$ gobrew rehash
	$ go version
	```

4. gomの_vendorディレクトリを$GOPATHに設定する

	```
	$ cd _vendor
	$ gobrew workspace set
	$ echo $GOPATH
	```

5. 開発をサポートするツールをインストール

	```
	$ go get -u gopkg.in/godo.v1/cmd/godo
	$ go get github.com/mitchellh/gox
	$ go get github.com/mattn/gom
	$ go get -u github.com/golang/lint/golint
	```

	goxだけ、ツールチェーンのビルドを別途行う。
	
	```
	$ gox -build-toolchain
	```

## Unit Test

```
$ godo test --watch
```

## Build

1. 各プラットフォームの実行形式を生成する

```
$ godo build
```
