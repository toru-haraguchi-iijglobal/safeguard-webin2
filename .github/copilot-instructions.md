## 目的

このリポジトリは OISG-RemoteApp-Launcher から呼ばれ、YAML または JSON Lines の設定から「Asset」を検索してブラウザ操作（ログインなど）を自動化する小さな Go ツールです。

以下はこのコードベースに素早く寄与するときに、AI コーディングアシスタント（Copilot 等）が知っておくべき具体的なポイントと実例です。

## ビッグピクチャ（アーキテクチャ）
- 実行形式: `webin2`（Windows 想定） — `webin2.exe` は CLI 引数で `-jsonl` または `-yaml`（いずれか一方）と `-asset` / `-account` / `-password` を受け取る。
- 設定読み取り: `search_jsonlines` (webin2jsonl.go) は JSONL を 1 行ずつ decode して `Asset` を探索。 `search_yaml` (webin2yaml.go) は複数ドキュメントの YAML を順に decode して探索。
- 実行: `run` (webin2chromedp.go) が `Definition` → `Actions` を解析して chromedp の操作列に組み立て実行する。
- ロギング: `logger.go` に共通のログ関数 (Debug/Info/Warn/Error/Fatal) があり、実行の各段階で呼ばれている。

## 重要なファイル（参照先）
- `webin2.go` — エントリポイント、引数処理、ログファイル生成、処理フローの起点。
- `webin2structs.go` — `Definition` と `Action` の型定義（構造がプロジェクト全体で重要）。
- `webin2jsonl.go` / `webin2yaml.go` — 設定ファイルのパースと Asset 検索ロジック。
- `webin2chromedp.go` — chromedp を使って実際のブラウザ操作（navigate/click/account/password/sleep 等）を組み立てる。
- `logger.go` — すべてのコンポーネントで使われるログ関数。`LogArgs` はパスワードを赤acted 表示する（パスワードはログに表示されない実装）

## コーディング、変更時の具体的ルール
- 修正は `Definition` / `Action` 型に合わせる（`webin2structs.go` を必ず確認）。新しい action type を入れる場合、`webin2chromedp.go` の `switch` に追加し、`LogActionStart` / `LogActionComplete` を呼ぶ。
- 機密情報: `logger.go` の `LogArgs` は password を赤acted しているため、絶対に plain-text を logger に直接出さないでください。
- エラーハンドリング: 既存コードは致命的エラーで `Fatal()` を呼んで終了する傾向がある。安全に継続できる処理は `Error`/`Warn` を用いる。
- プラットフォーム: Windows 上で使われることを想定している（Edge の固定パスなど）。変更時は `webin2chromedp.go` の ExecPath を確認。

## 典型的な開発ワークフロー（手順）
1. ビルド
```powershell
go build
```
2. 実行（例: サンプル YAML を使ってローカルで動かす）
```powershell
.\webin2.exe -yaml sample_webin2.yaml -asset PaloAltoNetworksHUB -account myuser -password mypass
```
3. ログ: 実行ディレクトリに `webin2_<PID>.log` が出力される。実行成功時はログファイルは削除される（参考: `webin2.go`）。

## よく使うパターン（具体例）
- 設定フォーマット
  - YAML: `sample_webin2.yaml` に複数ドキュメント（---）で複数 asset を定義する。各ドキュメントは `Definition` に map される。
  - JSONL: 1 行=1 asset オブジェクト（`webin2jsonl.go` を参照）。
- action の対応
  - `navigate` → chromedp.Navigate(url)
  - `click` → chromedp.Click(selector, chromedp.ByQuery, chromedp.NodeVisible)
  - `account` → chromedp.SendKeys(selector, account, ...)
  - `password` → chromedp.SendKeys(selector, password, ...)
  - `sleep` → chromedp.Sleep(time.Millisecond * Value)

## テスト / 検証
- 現時点でユニットテストは含まれていません。変更を加えたら `go build` を行い、ローカルで `sample_webin2.yaml` を使った手動検証で挙動を確認してください。

## 変更やレビュー時のヒント
- 変更を加える場所は影響範囲が小さく保たれている（構成読み込み → Definition → Actions → chromedp 実行）ため、1点修正しやすいが、chromedp の実行はブラウザ挙動に依存するため手動確認を必ず行う。
- ログは細かく出る設計なので、デバッグ時は `logger.go` の `SetLogLevel` を使って詳細出力に変更する（内部で使用される箇所を参照）。

---
If anything here is unclear or you want it expanded (more examples, testing steps, or safety notes), tell me which section to update and I'll iterate. ✅
