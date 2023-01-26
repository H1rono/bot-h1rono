# bot-h1rono:@BOT_H1rono:

## `@BOT_H1rono (help|:question:|:hatena:)`

これを表示します。

## `@BOT_H1rono :oisu-:`または`@BOT_H1rono :oisu-1::oisu-2::oisu-3::oisu-4yoko:`

:@BOT_H1rono:がチャンネルに参加します。スタンプにエフェクトはつけられません。

## `@BOT_H1rono :wave:`

:@BOT_H1rono:がチャンネルから退出します。

## `@BOT_H1rono :[pingを含むスタンプ][お好みでエフェクト]:`

ping

## `(@BOT_H1rono) :[スタンプのパターン][お好みでエフェクト]:`

:@BOT_H1rono:が参加しているチャンネルであればメンションは不要です。パターンにマッチするスタンプ全てを返信してくれます。エフェクト(`.`以下)がついていた場合は返信するスタンプ全てにエフェクトが入ります。エフェクトにパターンは使えません。パターンで使える文字は以下の通りです。

- `?`: 1文字にマッチします
- `*`: 0文字以上にマッチします
- `+`: 1文字以上にマッチします

ただし、ここでマッチする文字はスタンプで使える文字(アルファベット大文字/小文字、数字、`_`と`-`)のみです。

例:

- `:oisu-?:`: :oisu-4::oisu-2::oisu-1::oisu-3:
- `:oisu*:`: :oisu-4::oisu-iteraton_mark::oisu-tu::oisu-ten::oisu-ri::oisu-cat1::oisu-bo::oisu-ko::oisu-yo::oisu-kyurun::oisu-4yoko::oisu-he::oisu-ra::oisu-cat3::oisu-ho::oisu-dakuten::oisu-si::oisu-cat2::oisu-2::oisu-1::oisu-::oisu-3:
- `:oisu-+:`: :oisu-4::oisu-iteraton_mark::oisu-tu::oisu-ten::oisu-ri::oisu-cat1::oisu-bo::oisu-ko::oisu-yo::oisu-kyurun::oisu-4yoko::oisu-he::oisu-ra::oisu-cat3::oisu-ho::oisu-dakuten::oisu-si::oisu-cat2::oisu-2::oisu-1::oisu-3:

## else

[リポジトリ](https://git.trap.jp/H1rono_K/bot-h1rono)を見てください。Web APIがあったりします。
