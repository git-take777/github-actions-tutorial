# GitHub CI/CD 実践ガイド

![GitHub Actions](https://via.placeholder.com/800x200?text=GitHub+CI/CD実践ガイド)

## 📘 本書について

「GitHub CI/CD 実践ガイド――持続可能なソフトウェア開発を支える GitHub Actions の設計と運用」は、GitHub Actions を使った効率的な開発ワークフローの構築方法を解説した書籍です。

プログラミング初心者から経験者まで、誰でも GitHub Actions を使いこなせるようになることを目指しています。

## 🚀 本書の特徴

- **基礎から応用まで**: GitHub Actions の基本概念から高度な使い方まで段階的に学べます
- **実践的なサンプル**: すぐに使えるワークフロー例が豊富に用意されています
- **トラブルシューティング**: よくある問題の解決方法も解説しています
- **持続可能な開発**: チーム開発を効率化するノウハウを紹介しています

## 📋 目次

1. **GitHub Actions とは何か**

   - CI/CD の基本概念
   - GitHub Actions の特徴と利点
   - 他の CI/CD ツールとの比較

2. **GitHub Actions の基本**

   - ワークフローファイルの書き方
   - ワーク不 question 実行ログの確認方法
     - `gh run watch`
     - `gh run view`
   - エラーが起こると、run-on にある実行をしても Actions は実行されない
     - ※ YAML は正しくても GitHubAction で使えないエイリアス、アンカーがある。実行されない。
   - 手動実行
     - on キーへ `workflow_dispatchイベントを指定すればワークフロー手動実行が可能`
     - `gh workflow run manual.yml -f greeting=goodbye`

- on キー
  - on キーに`schedule`を指定すると、指定した時間に自動実行される ただし、タイムゾーンに UTC が設定されているため、日本時間で実行する場合は UTC の時間を日本時間に変換する必要がある。
  - 時間ぴったりに起動するとは限らない。 時間にシビアな場合は不向きである。
- ワークフローの際実行
  - 実装ミスがなくても、連携先のシステムの不具合によりワークフローが失敗することがある。そんな時は再実行する。また失敗したジョブに限定した際実行も可能。
- アクションの探し方
  - GitHub Marketplace で検索
  - Actions を選択し、検索ボックスで Linter を入力すると、Linter のアクションが表示される。
  - Verify Creators
  - アクションは便利だがセキュリティリスクがある。 悪意のあるアクションを使うとコードの改ざんや情報漏洩のリスクがある。 そこで活用したいのが Verified Creators。 これは GitHub によって検証されたアカウントを意味する。 「Verified Creator」というバッジが表示されているアクションは、GitHub が検証したものであることを意味します。
    過信は禁物。理由は、GitHub が検証したアカウントでも、悪意のあるアクションを公開する可能性があるため。
- アクション、ジョブ、ステップの理解
- トリガーの設定方法

3. **実践的なワークフロー例**

- コンテキストの一覧

  - github
  - job
  - runner
  - matrix
  - env
  - jobs
  - secrets
  - needs
  - vars
  - steps
  - strategy
  - inputs
  - action

- github コンテキスト

  - プロパティ

    - github.run_id ・・・ ワークフロー実行 ID
    - github.head_ref ・・・ プルリクエストのソースブランチ
    - github.workflow ・・・ ワークフロー名
    - github.repository ・・・ リポジトリ名
    - github.repository_owner ・・・ リポジトリの所有者
    - github.event ・・・ トリガーイベントの詳細
    - github.event_name ・・・ トリガーイベント名
    - github.ref ・・・ ブランチ名またはタグ名

  - プロパティ版外面

    - github.action ・・・ アクション名
    - github.action_path ・・・ アクションのパス
    - github.action_ref ・・・ アクションのバージョン
    - github.action_repository ・・・ アクションのリポジトリ名
    - github.action_repository_owner ・・・ アクションのリポジトリの所有者
    - github.actor ・・・
    - github.actor_id ・・・
    - github.base_ref ・・・ プルリクエストのベースブランチ
    - github.env
    - github.event_path ・・・ トリガーイベントのパス
    - github.event.issue.body ・・・ イシュー本文
    - github.event.issue.title ・・・ イシューのタイトル
    - github.event.issue.number ・・・ イシュー番号
    - github.event.pull_request.head.ref ・・・ プルリクエストのソースブランチ
    - github.graphql_url ・・・ GraphQL API の URL
    - github.head_ref ・・・ プルリクエストのソースブランチ
    - github.job ・・・ ジョブ名
    - github.ref ・・・ ブランチ名またはタグ名
    - github.ref_name ・・・ ブランチ名またはタグ名
    - github.ref_protected
    - github.ref_type ・・・ ブランチまたはタグ
    - github.retention_days ・・・ リポジトリのリテンション期間
    - github.repository ・・・ リポジトリ名
    - github.repository_owner ・・・ リポジトリの所有者
    - github.run_id ・・・ ワークフロー実行 ID
    - github.run_number ・・・ ワークフロー実行番号
    - github.run_attempt ・・・ ワークフロー実行試行回数
    - github.run_number ・・・ ワークフロー実行番号
    - github.secret_source ・・・ シークレットのソース
    - github.server_url ・・・ GitHub サーバーの URL
    - github.sha ・・・ コミット SHA
    - github.token ・・・ トークン
    - github.workflow ・・・ ワークフロー名
    - github.workflow_ref ・・・ ワークフローのバージョン
    - github.workflow_sha ・・・ ワークフローの SHA
    - github.workspace ・・・ ワークスペースのパス
    - github.workspace_path ・・・ ワークスペースのパス

  - runner コンテキスト

    - runner.name ・・・ ランナー名
    - runner.os ・・・ ランナーの OS
    - runner.arch ・・・ CPU のアーキテクチャ (X86・X64・ARM・ARM64)
    - runner.tool_cache ・・・ ツールキャッシュのパス
    - runner.temp ・・・ 一時ディレクトリのパス
    - runner.tool_cache ・・・ ツールキャッシュのパス
    - runner.debug ・・・ デバッグログの有効化

  - 環境変数
    - 環境変数は env ファイルで定義する。
    - ワークフロージョブの各レベルで定義可能 定義場所によって環境変数のスコープは変わる。また環境変数には文字列だけでなく先ほど登場したコンテキストも指定できる。環境変数を定義するには次のように記述します。
      - env:
        BRANCH: main
  - 環境変数の参照
    - 環境変数の参照方法はシェルコマンド以外からも参照可能。 `${{ env.BRANCH }}`
  - GITHUB_ACTIONS ・・・ ワークフローが GitHub Actions で実行されているかどうかを示すフラグ
  - 概要

  - 環境変数のオーバーライド
    - 環境変数はオーバライド可能。 優先されるのはスコープが狭い方。 ステップレベルの優先順位が一番高く、ワークフローレベルの優先順位が一番低い。※オーバーライドすると読みづらくなるので控えよう。
  - デフォルト環境変数
    - CI
    - GITHUB_ACTION
    - GITHUB_ACTION_PATH
    - GITHUB_ACTION_REPOSITORY
    - GITHUB_ACTION_REPOSITORY_OWNER
    - GITHUB_ACTIONS
    - GITHUB_ACTOR
    - GITHUB_ACTOR_ID
    - GITHUB_API_URL
    - GITHUB_BASE_REF
    - GITHUB_ENV
    - GITHUB_EVENT_NAME
    - GITHUB_EVENT_PATH
    - GITHUB_GRAPHQL_URL
    - GITHUB_HEAD_REF
    - GITHUB_JOB
    - GITHUB_PATH
    - GITHUB_REF
    - GITHUB_REF_NAME
    - GITHUB_REF_PROTECTED
    - GITHUB_REF_TYPE
    - GITHUB_REPOSITORY
    - GITHUB_REPOSITORY_OWNER
    - GITHUB_RETENTION_DAYS
    - GITHUB_RUN_ID
    - GITHUB_RUN_NUMBER
    - GITHUB_RUN_ATTEMPT
    - GITHUB_SERVER_URL
    - GITHUB
    - GITHUB_SHA
    - GITHUB_SERVER_URL
    - GITHUB_STEP_SUMMARY
    - GITHUB_WORKFLOW
    - GITHUB_WORKFLOW_REF
    - ※ コンテキストからデフォルト環境変数へは簡単に書き換えできる
  - 中間環境変数
    - コード `Context.yml`ではコンテキストを直接、シェルコマンドへ埋め込む。しかしこの実装はアンチパターン。コンテキストによっては特殊文字が含まれ、シェルコマンドの実行に意図しない影響を受けてしまう。
      この問題を回避するには環境変数を経由でコンテキストを渡し、クオーとする。中間環境変数と呼ばれるテクニック。
    - コンテキストを中間環境変数に変換するには、`echo "::set-env name={変数名}::{値}"`を使う。
      - 環境変数をクオーとするには、参照時に 「""」で囲む。トークン分割やパス名展開が抑止可能。 一般的なシェル仕様である。安全にコンテキストを参照できる。しかし、どのコンテキストにリスクがあるか判断は難しい。そこでおすすめなのが コーディングルールを設けること
      - コンテキストはシェルコマンドへハードコードせず、環境変数経由で渡す
      - 環境変数は全て 「""」で囲む
  - Variables
  - 複数のフローで同じ値を費た場合が[Varibales]。
  - Variables の設定方法は GUI でも設定可能。 - わからなければ、Cluade で聴く。
  - コンテキスト参照時は忘れずに中間環境変数を経由する
  - Secrets
    - Secrets は GitHub のリポジトリに保存される機密情報。 シークレットは、ワークフローの実行中にアクセスできるようになる。
    - Secrets はログ出力時に自動でマスクされる
    - Secrets 時は登録後に値がまったく確認できなくなる
    - パスワードなどの機密情報は Secrets を使おう。 特にログマスクの有無が重要。
    - Secrets も Variables と同様に事前登録が必要。 `Variables`同様にシェルでも登録可能。
    - Secreat: `gh secret set PASSWORD --body 'SuperSecret!`
  - 式
  - 式は、ワークフローの実行中に評価される式。 式は、ワークフローの実行中に評価される式。 式は、ワークフローの実行中に評価される式。式は コンテキスト参照以外にも、リテラル演算子や関数を利用できる。
  - リテラル
    - データタイプ null, boolean, number, string を表す。
    - リテラル値 null, true, JSON がサポートする数値、文字列。
  - 演算子

        - <Table>
               | 演算子 | 説明 |
               | --- | --- |
               | `==` | 等しいかどうかを判定する。 |
               | `!=` | 等しくないかどうかを判定する。 |
               | `<` | より小さいかどうかを判定する。 |
               | `<=` | 以下かどうかを判定する。 |
               | `>` | より大きいかどうかを判定する。 |
               | `>=` | 以上かどうかを判定する。 |
               | `&&` |
               | `||` |
               | `!` | 否定する。 |
               | `()` | グルーピング |
               | `[]` | 配列のインデックスを指定する。 |
               | `.` | オブジェクトのプロパティを指定する。 |
               </Table>
             ※ 比較演算子は注意。GitHUb Actionsでは、異なる型で比較すると勝手に値が変更される。 trueと数値を比較するとバウ具が混入しやすいので異なる型の比較は避けよう
             - 関数
             <Table>
             |種類| |関数| 説明|
             |文字列比較 contains() startWith() endWith() | 文字列が含まれているかどうかを判定する。 |
             |文字列生成 join() format() | 配列の要素を結合する。 |
             |JOIN操作  | toJSON() fromJSON()| 配列の要素を結合する。 |
             | ハッシュ操作 | hasFiles() | 配列を結合する。 |
             </Table>

          3.6.1 文字列の比較

          3.7 条件分岐

          - ステータスチェック

            - ステータスチェック関数を使うと、失敗時のみ実行する処理が実現できる
            - success()・・・手前の処理が成功したら true
            - failure()・・・手前の処理が失敗したら true
            - always()・・・手前の処理が成功しても失敗しても true
            - cancelled()・・・キャンセルされたら true
              ※ failure 関数をよく使う。
          3.8 ネーミング
          3.8.1 ステップ名とジョブ名 - step jobs の配下に`name`をつけるとネーミングをつけることになる
          3.9 ステップ間の値の共有
              ステップ間の受け渡しは少し工夫が必要
              `name: Missing share data
                on: push
                jpbs:
                runs-on: ubuntu-latest
                steps: - run: export RESULT="hello" - run: echo "RESULT is $RESULT"
              `

- Web アプリケーションのテストと自動デプロイ
- モバイルアプリのビルドと配布
- ドキュメントの自動生成と公開

4. **効率化とベストプラクティス**

   - キャッシュを活用した高速化
   - マトリックスビルドによる並列実行
   - 再利用可能なワークフローの作成

5. **セキュリティと管理**
   - シークレットの安全な管理
   - 権限の適切な設定
   - コスト管理と最適化

## 🛠️ サンプルコード

本リポジトリには書籍で紹介しているサンプルコードが含まれています。各章ごとのフォルダにサンプルがあります。

```
samples/
├── chapter1/
├── chapter2/
├── chapter3/
├── chapter4/
└── chapter5/
```

## 🏁 はじめ方

1. このリポジトリをクローンする

   ```bash
   git clone https://github.com/yourusername/github-cicd-guide.git
   ```

2. サンプルコードを試す
   ```bash
   cd github-cicd-guide/samples/chapter1
   # サンプルの手順に従ってください
   ```

## 🤔 よくある質問

<details>
<summary>Q: プログラミング初心者でも理解できますか？</summary>
A: はい！基本的なGitの知識があれば、順を追って理解できるように書かれています。
</details>

<details>
<summary>Q: どのプログラミング言語に対応していますか？</summary>
A: 主にJavaScript/TypeScript、Python、Javaのサンプルがありますが、基本的な考え方はどの言語でも応用できます。
</details>

<details>
<summary>Q: 電子書籍版はありますか？</summary>
A: はい、Kindle、PDFなど各種フォーマットで入手可能です。詳細は公式サイトをご確認ください。
</details>

## 📝 誤字脱字・内容の修正

本書の内容に誤りを見つけた場合は、Issue を作成するか、Pull Request を送ってください。皆さんの貢献に感謝します！

## 📢 更新情報

- 2025 年 3 月 1 日: サンプルコード v1.0.0 公開
- 2025 年 2 月 15 日: 書籍正式発売
- 2025 年 1 月 10 日: 予約開始

## 📞 お問い合わせ

質問やフィードバックがありましたら、以下の方法でご連絡ください：

- GitHub の Issue を作成する
- Twitter: [@githubcicdbook](https://twitter.com/githubcicdbook)
- メール: contact@githubcicdbook.example.com

## 📜 ライセンス

本リポジトリのサンプルコードは MIT ライセンスで提供されています。詳細は[LICENSE](LICENSE)ファイルをご覧ください。

---

Happy Coding with GitHub Actions! 🚀
