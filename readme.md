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
