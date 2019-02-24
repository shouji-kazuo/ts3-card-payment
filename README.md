
# ts3-card-payment

TS CUBICカード(https://ts3card.com/)にログインして請求金額を取得するスクレイピングツール．

## Requirements
- Golang (>= 1.8.1)
- Google Chrome
- お使いのChromeバージョンに合ったChromeDriver
    - http://chromedriver.chromium.org/downloads

## Usage
1. お使いのChromeに合ったバージョンのchromedriver ダウンロードして，PATH通す
2. ビルド
   ```
   go build
   ```
3. 起動
   ```
   ./ts3-card-payment
   ```
4. ユーザ名とパスワードを聞かれるので入力
   ```
   Enter username: *****
   Enter password: *****
   ```
5. 標準出力に結果が出る
   ```
   前回（2月4日）ご請求分	**,***円
   今回（3月4日）ご請求分	**,***円
   次回（4月2日）ご請求予定分	**,***円
   ```

## Options
```
NAME:
   ts3-card-payment - A new cli application

USAGE:
   ts3-card-payment [global options] command [command options]

VERSION:
   v1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --user value, -u value      set user name to login ts3card.com
   --password value, -p value  set password to to login ts3card.com
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```
## TODO

- chrome以外のブラウザ対応
- 結果の金額が `string` 型でしか得られないのなんとかしたい
