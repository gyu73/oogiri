## 大喜利画像一覧画面
GET("/", oogiri.Get)
GET("/oogiris", oogiri.Get)

## 大喜利画像投稿ページ
GET("/oogiri/new", oogiri.New)

## 大喜利画像投稿
POST("/oogiri", oogiri.Create)

## 大喜利画像詳細ページ(コメントも表示)
GET("/oogiri/:id", oogiri.Show)

## 大喜利コメント投稿
POST("/oogiri/answer", answer.Create)

## 大喜利削除機能
POST("/oogiri/answer", answer.Create)

## ユーザー詳細ページ
GET("/user/:id", user.Show)

## 使ったライブラリ
"github.com/labstack/echo"
"github.com/jinzhu/gorm"