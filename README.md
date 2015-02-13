※まだ実用未満

# tmm

マインドマップ風の木構造メモ

## Description

SQLite形式のDBファイルにメモの内容を保持する。
マインドマップ風に一つのノードに対して任意の数のノードを追加して木構造のメモを作る。

## Demo

## VS. 

txtファイル： 編集は簡単、インデントで木構造にして行くと読みにくい
マインドマップ： 見た目でたどるのは容易、レイアウトや見栄えに気を配ると面倒
当ツール： ビジュアルはあくまでテキスト、レイアウトができない分お手軽

## Requirement

go version go1.4

github.com/jinzhu/gorm
github.com/mattn/go-sqlite3
github.com/mitchellh/cli

## Usage

tmm start 

## Install

go get github.com/goldbook/tmm

## Contribution

## Licence



## Author

[goldbook](https://github.com/goldbook)
