module.exports = {
  // モード値を production に設定すると最適化された状態で、
  // development に設定するとソースマップ有効でJSファイルが出力される
  mode: 'development',
 
  // メインとなるJavaScriptファイル（エントリーポイント）
  entry: './src/main.ts',
  // ファイルの出力設定
  output: {
    //  出力ファイルのディレクトリ名
    path: `${__dirname}/../public/js`,
    // 出力ファイル名
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        // 拡張子 .ts の場合
        test: /\.ts(x?)$/,
        // TypeScript をコンパイルする
        use: 'ts-loader'
      },
    ]
  },
  // import 文で .ts ファイルを解決するため
  resolve: {
    extensions: ['.ts','','.tsx'],
    // Webpackで利用するときの設定
    alias: {
      vue: 'vue/dist/vue.js'
    }
  }
};