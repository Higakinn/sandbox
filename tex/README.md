# 論文環境構築(tex)

VS Code Remote Developmentの機能を用いて、LaTeX環境を容易に構築するサンプルコードです。

[![](https://images.microbadger.com/badges/image/korosuke613/ubuntu-texlive-ja-devcontainer.svg)](https://microbadger.com/images/korosuke613/ubuntu-texlive-ja-devcontainer "Get your own image badge on microbadger.com")


Dockerイメージは、[ubuntu-texlive-ja](https://hub.docker.com/r/korosuke613/ubuntu-texlive-ja)を利用しています。

(**Dockerが必要です**)

## install

- Docker
  - https://docs.docker.com/install/
- VS Code
  - 拡張機能で[Remote-Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)を入れておく

## setup

1. VS Codeで`texlive-ja-devcontainer-template-master`を開く。
2. 左下の`><`アイコンを押して、`Rebuild Container` or `Remote-Containers: Reopen in Container`を実行する。
3. 待つ。
4. latex環境が構築される。(texの拡張機能などがインストールされた環境が立ち上がる。)
5. ファイルを保存する度にコンパイルされ,pdfが更新される。

## build

以下の方法でもコンパイル可能。

```bash
latexmk main.tex
```

`main.pdf`という実行ファイルができているはず。

## reference

- 参考：https://korosuke613.hatenablog.com/entry/2019/06/24/171246