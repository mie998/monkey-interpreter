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
- 関数リテラルが fn() {} という指定方法のみだが、JS の arrow function を追加してみる (p100) と、その前に無名関数を定義できるようにしないと！(p171)
- 関数呼び出しにおいて、LPAREN が中置演算子とみなして解析を行っているわけだが、他の方法はないのだろうか。あらゆる処理系はこのように処理されているのか。(p107)
- Rust で内容を再実装してみても面白いかも。golang ではインターフェースで頑張っているところをジェネリクスを使えば簡潔に書けて可読性が上がることが期待される。
- eval では Integer と言ったら int64 にしかならないので、int8 とかも実装してみる (p125)
- GT や LT などの実装は Integer しか用意していないが、string などの他の型に対しても実装してみる (p137)
- return の処理を理解しきっていないかも。(p149)
- Error の拡張。最低限ファイル名と行番号は調べられるようにする。できればスタックトレースも実装する(p150)
- NULL を消そう。考えることがどれくらいあるのか知りたい.
- `let add_curry_3 = fn(x) { add(x, 3) };` みたいなことをできるようにしたい。現在: `parser error happen!!!: 
        expected next token to be =, got INT instead
        no prefix parse function for = found` これは `let newAdder = fn(x) { fn(y) {x + y }};` みたいにすればいいかも。要検討
- repl で補完が効かないのが気に食わないので補完機能を追加する。
- loop なくない？        
- GC を使い回しているらしいので GO の GC について理解するのは礼儀でしょうか。

