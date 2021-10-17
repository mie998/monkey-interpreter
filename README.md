# monkey-interpreter
Go言語で作るインタプリタ 実装

## TODO
- Unicode への対応(UTF-8 と UTF-16 を処理する, p7)
- 3文字以上必要な演算子への対応 (p19 とか)
- トップダウン式構文解析とボトムアップ構文解析との違いを理解 (p30)
- statement を増やす (p30 ~)
  - 今は let と return のみ
- if を statement でなく expression にしてみる (p93)
- ++ や -- のように、前置も後置もできる演算子を定義してみる (p83)
- switch expression の実装 
- ast とか parser function とか一か所に固まりすぎて見ずらいのでいい感じにファイル分けを検討する
- monkey 用の formatter とか作れたら良さそう。テストもはかどるかも。
- if-else しかないので、elif を追加する。
- 関数リテラルが fn() {} という指定方法のみだが、JS の arrow function を追加してみる (p100)
- 関数呼び出しにおいて、LPAREN が中置演算子とみなして解析を行っているわけだが、他の方法はないのだろうか。あらゆる処理系はこのように処理されているのか。(p107)
