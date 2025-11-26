# safeguard-webin2

OISG-RemoteApp-Launcher から呼び出され、YAML または JSON ファイルから Asset を探し、あらかじめ定義された DOM エレメントを操作する。

---

## 目次

- [概要](#概要)
- [前提条件](#前提条件)
- [セットアップ](#セットアップ)
- [実行方法](#実行方法)
- [実行時の引数](#実行時の引数)
- [設定ファイル](#設定ファイル)

---

## 概要

OISG-RemoteApp-Launcher から呼び出され、YAML または JSONL ファイルから Asset を探し、あらかじめ定義された DOM エレメントを操作する。
かつてあった、Quest Japan から提供された同目的のプログラムを改良した。

---

## 前提条件

- Golang
- OISG-RemoteApp-Launcher

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
プログラム実行時には、後述の引数を指定する必要があります。

---

## 実行時の引数

`webin2.exe` は以下のコマンドライン引数を受け取ります。

- `-jsonl <ファイルパス>`: JSON Lines 形式の設定ファイルを指定します。`-yaml` と同時に指定することはできません。
- `-yaml <ファイルパス>`: YAML 形式の設定ファイルを指定します。`-jsonl` と同時に指定することはできません。
- `-asset <アセット名>`: 設定ファイルの中から、ここで指定されたアセット名を持つ設定を探して処理を実行します。OISG-RemoteApp-Launcher から渡される想定です。
- `-account <アカウントID>`: ログインに使用するアカウントIDです。OISG-RemoteApp-Launcher から渡される想定です。
- `-password <パスワード>`: ログインに使用するパスワードです。OISG-RemoteApp-Launcher から渡される想定です。

---

## 設定ファイル

設定ファイルは `YAML` または `JSON Lines` 形式で記述します。1つのファイルに複数の設定（アセット）を含めることができます。

### 設定ファイルの構造 (YAML)

各アセットの設定は `---` で区切ります。

```yaml
# このファイルは Webin2 プロジェクトの一部です。
asset: PaloAltoNetworksHUB
lang: ja
use_edge: false
secret: true
cert_validation: true
actions:
  - type: navigate
    target: https://apps.paloaltonetworks.com/hub
  - type: account
    target: input[name="identifier"]
  - type: click
    target: input.button.button-primary
  - type: password
    target: input[name="credentials.passcode"]
  - type: click
    target: input.button.button-primary
---
asset: CATOSolutionsLab
use_edge: false
secret: true
cert_validation: true
actions:
  - type: navigate
    target: https://sol-lab.auth.catonetworks.com
  - type: account
    target: input#username
  - type: password
    target: input#password
  - type: click
    target: input.btn-submit
```

### パラメータ

- `asset` (string, 必須): アセット名を指定します。`-asset` 引数で渡された値と一致するものが使われます。
- `lang` (string, 任意): ブラウザの表示言語を指定します (例: `en`, `ja`)。省略した場合は、システムのデフォルト言語が使用されます。
- `use_edge` (boolean, 任意): `true` に設定すると、Google Chrome の代わりに Microsoft Edge を使用してブラウザ操作を行います。デフォルトは `false` です。
- `secret` (boolean, 任意): `true` に設定すると、ログファイルにパスワードなどの機密情報が出力されなくなります。デフォルトは `true` です。
- `cert_validation` (boolean, 任意): `true` に設定すると、サイトのSSL/TLS証明書の検証を行います。デフォルトは `true` です。
- `actions` (array, 必須): ブラウザで実行する操作のリストを定義します。

### Actions

`actions` 配列には、実行したい操作を順番に記述します。

- `type`: 実行する操作の種類を指定します。
  - `navigate`: `target` で指定された URL に遷移します。
  - `account`: `target` で指定された入力フィールド（CSSセレクタで指定）に `-account` 引数の値を入力します。
  - `password`: `target` で指定された入力フィールド（CSSセレクタで指定）に `-password` 引数の値を入力します。
  - `click`: `target` で指定された要素（CSSセレクタで指定）をクリックします。
- `target`: 操作の対象を指定します。`type` が `navigate` の場合は URL を、それ以外の場合は CSS セレクタを記述します。
