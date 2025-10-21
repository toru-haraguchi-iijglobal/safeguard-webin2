# ログ出力改善ドキュメント

## 概要

webin2アプリケーションのログ出力を改善し、以下の機能を追加しました：

1. **構造化ログ**: タイムスタンプとログレベルを含む一貫したフォーマット
2. **ログレベル**: DEBUG, INFO, WARN, ERROR, FATAL の5段階
3. **セキュリティ**: パスワードは `[REDACTED]` として出力
4. **詳細なトレース**: 各アクション、ファイル操作、ブラウザ設定の詳細なログ

## ログフォーマット

```
[YYYY-MM-DD HH:MM:SS.mmm] [LEVEL] メッセージ
```

例:
```
[2025-10-21 15:07:43.123] [INFO] === Webin2 Application Started ===
[2025-10-21 15:07:43.124] [INFO] Process ID: 12345
[2025-10-21 15:07:43.125] [INFO] Command line arguments:
[2025-10-21 15:07:43.126] [INFO]   - jsonl: config.jsonl
[2025-10-21 15:07:43.127] [INFO]   - password: [REDACTED]
```

## ログレベル

### DEBUG
- 詳細な技術情報
- ブラウザ選択（Chrome/Edge）
- ファイル操作の詳細
- 各アクションの完了状態
- chromedpの内部ログ

### INFO
- アプリケーションの起動・終了
- コマンドライン引数
- Asset検索の開始・結果
- ブラウザ設定
- 各アクションの開始
- 実行サマリー

### WARN
- Asset が見つからない
- 不明なアクションタイプ
- 非致命的な問題

### ERROR
- ファイルオープン失敗（続行可能な場合）
- JSON/YAMLパースエラー
- chromedp実行エラー
- 設定エラー

### FATAL
- 致命的なエラー（アプリケーション終了）
- ファイルが開けない
- パース失敗
- chromedp致命的エラー

## 主な改善点

### 1. アプリケーションライフサイクル

起動時と終了時に明確なマーカーを出力：

```
[INFO] === Webin2 Application Started ===
[INFO] Process ID: 12345
[INFO] Log file: C:\path\to\webin2_12345.log
...
[INFO] === Webin2 Application Completed Successfully ===
```

### 2. セキュリティ

パスワードは常にマスクされます：

```
[INFO] Command line arguments:
[INFO]   - account: user@example.com
[INFO]   - password: [REDACTED]
```

### 3. 詳細なアクションログ

各chromedpアクションの詳細を記録：

```
[INFO] Action 0: type=navigate, target=https://example.com, value=0
[INFO] Action 1: type=click, target=#login-button, value=0
[INFO] Action 2: Sending account credentials
[INFO] Action 3: Sending password credentials
[DEBUG] Action 3 completed: password
```

### 4. ブラウザ設定のトレース

```
[INFO] Browser configuration:
[INFO]   - Use Edge: true
[INFO]   - Secret mode: true
[INFO]   - Certificate validation: false
[DEBUG] Browser: Microsoft Edge
```

### 5. ファイル操作とAsset検索

```
[INFO] Searching for asset 'MyWebsite' in file 'config.jsonl'
[DEBUG] File operation: Opening - config.jsonl
[DEBUG] Line 1: Asset 'OtherWebsite' does not match, continuing search
[DEBUG] Line 2: Asset 'MyWebsite' does not match, continuing search
[INFO] Asset 'MyWebsite' found successfully
[DEBUG] Asset found at line 3 in config.jsonl
```

### 6. エラーの詳細情報

```
[ERROR] JSON Line 5 contains error. Last correct asset: Website1
[FATAL] Failed to parse JSON Line: unexpected end of JSON input
```

## 使用方法

### ログレベルの変更（将来の拡張用）

現在はINFOレベルがデフォルトですが、logger.goの `SetLogLevel()` 関数を使用してレベルを変更できます：

```go
SetLogLevel(DEBUG)  // すべてのログを出力
SetLogLevel(ERROR)  // エラーとFATALのみ
```

### カスタムログメッセージ

新しいログ関数を使用：

```go
Debug("詳細な情報: %v", details)
Info("処理完了: %s", result)
Warn("警告: %s", warning)
Error("エラー発生: %v", err)
Fatal("致命的エラー: %v", err)  // プログラム終了
```

### 専用ログ関数

よく使用されるログパターンには専用関数を使用：

```go
LogStartup(pid, logFilename)
LogArgs(jsonl, yaml, asset, account)
LogAssetSearch(asset, filename)
LogAssetFound(asset)
LogBrowserConfig(useEdge, secret, certValidation)
LogChromedpStart(actionCount)
LogActionStart(index, actionType, target, value)
```

## ログファイル

- ファイル名: `webin2_<PID>.log`
- 場所: 実行ファイルと同じディレクトリ
- 成功時: 自動削除
- 失敗時: 保持（トラブルシューティング用）

## トラブルシューティング

### ログファイルが見つからない

正常終了した場合、ログファイルは自動的に削除されます。問題がある場合のみログファイルが残ります。

### ログが出力されない

1. ファイルへの書き込み権限を確認
2. ディスク容量を確認
3. アンチウイルスソフトがファイルをブロックしていないか確認

### ログが多すぎる

将来のバージョンでは、環境変数やコマンドラインオプションでログレベルを制御できるようになります。

## 今後の拡張予定

1. ログレベルのコマンドラインオプション追加
2. ログのローテーション機能
3. JSON形式でのログ出力オプション
4. 外部ログシステムへの送信機能

## 変更されたファイル

- `logger.go`: 新規作成（構造化ログモジュール）
- `webin2.go`: ログ関数を更新
- `webin2chromedp.go`: 詳細なアクションログを追加
- `webin2jsonl.go`: ファイル操作とAsset検索のログを改善
- `webin2yaml.go`: ファイル操作とAsset検索のログを改善
