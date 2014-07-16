gyowitter
=========

一日に何度も執拗に Yo を送ってくるような人がいてお困りではありませんか？ このソフトは相手から Yo が届くと Yo を返しつつ、 Yo が来たことを Twitter にさらし上げしてくれます。

## 利用方法

### Yo

- Yo の API Token が必要です
  - http://dev.justyo.co

### Twitter

- Twitter の OAuth Consumer Key / Secret が必要です 
- Twitter の OAuth Access Key / Secret が必要です

それぞれ https://dev.twitter.com で取得して下さい。

### 設置方法

バイナリをビルドしてテキトーなところに置きます。お好きなポート（デフォルトだと 8080 ）で起動して、 Nginx などで dev.justyo.co で設定した callback URL への Yo の hook を localhost:8080 などにプロキシしてあげます。誰かがあなたのアカウントに Yo すると Yo を返しつつ Twitter に post します。
