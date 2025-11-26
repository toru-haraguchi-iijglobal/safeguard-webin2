# safeguard-webin2

OISG-RemoteApp-Launcher から呼び出され、YAML または JSON ファイルから Asset を探し、あらかじめ定義された DOM エレメントを操作する。

---

## 目次

- 概要
- 前提条件
- セットアップ
- 実行方法

---

## 概要

OISG-RemoteApp-Launcher から呼び出され、YAML または JSONL ファイルから Asset を探し、あらかじめ定義された DOM エレメントを操作する。
かつてあった、Quest Japan から提供された同目的のプログラムを改良した。

---

## 前提条件

Golang
OISG-RemoteApp-Launcher

---

## セットアップ

```bash
git clone https://github.com/toru-haraguchi-iijglobal/safeguard-webin2
cd safeguard-webin2
go build
```

---

## 実行方法

One Identity Safeguard for Priveledged Passwords/Sessions のマニュアルに則って OISG-RemoteApp-Launcher から呼び出されるように、RDS サーバーにインストールする。
実行時は CLI 引数で設定ファイルと Asset を指定して呼び出します。主な引数は以下の通りです。

- `-jsonl <ファイルパス>`: JSON Lines 形式の設定ファイルを指定します。`-yaml` と同時に指定することはできません。
- `-yaml <ファイルパス>`: YAML 形式の設定ファイルを指定します。`-jsonl` と同時に指定することはできません。
- `-asset <アセット名>`: 設定ファイルの中から、ここで指定されたアセット名を持つ設定を探して処理を実行します。
- `-account <アカウントID>`: ログインに使用するアカウントIDです。
- `-password <パスワード>`: ログインに使用するパスワードです。
- `-lang <言語>`: ブラウザ言語を指定します（例: `en`, `ja`）。CLI で指定した場合は設定ファイル内の `lang` を上書きします。

実行例:

```powershell
.\webin2.exe -yaml sample_webin2.yaml -asset PaloAltoNetworksHUB -account myuser -password mypass -lang ja
```
