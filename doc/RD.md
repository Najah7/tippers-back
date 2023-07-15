# 要件定義

## 機能要件
- 送金系の機能（アシスト or ダイレクトに遷移）
    - paypayIDを表示することでPayPay送金を容易にする
    - 
- ユーザ機能
    - ユーザの認証
        - restorent_id
        - email
        - password
        - ユーザ名
        - プロフィール画像
        - dream
        - is_student
        - major
        - is_employed
        - period_of_workings
    - プロフィールページ
        - プロフィール画像
        - ユーザ名
        - dream
        - major
- 一覧機能
    - 店員一覧（チップ払う側への機能。位置情報を使う予定）
        - 店員のプロフィール画像
        - 店員の名前
        - one_line_message
        - is_student
        - period_of_workings
    - お店一覧（チップを払う側への機能。位置情報を使う予定）
        - お店のプロフィール画像
        - お店の名前
        - one_line_description
        - num_using_app_employee
    - チップの履歴を確認する機能（デザインは上記と同じにするイメージなのでここに配置）
        - チップを払ったお店の情報
        - Thanksの表示
        - 金額（支払ったユーザからのみ見れるイメージ）
- お店登録機能（チップをもらう側の機能。使う人にはユーザ登録と自分のプロフィールの編集、お店の登録をしてもらう。）
- Thanks機能 or チャット機能
    - 後でありがとうを伝えるための機能。チャットにすると、めんどくさや、トラブルの原因になる可能性もあるので、シンプルなThanks機能にするのもありかも。（悩み中）


## 非機能要件
- token認証
- スマホでの利用を想定（レスポンシブ対応）
- ディバイスにより、ページを振り分ける（PCの時はコンソールページ一ページのアプリを考案している）
- 画像のアップロードなどの特殊な場合を除き、どんなに長くても1秒以内に画面遷移したい
- もし、みんながやる気で技術的な挑戦をしたいなら、SSRを使ってみるのもありかもしれない。特に、プロフィールやお店の詳細ページなどに使うのはアリかも、一覧ページはSSRを使わない方がbetterな気がする。