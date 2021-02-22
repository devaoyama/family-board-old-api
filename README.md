# Family Board API

## 言語/フレームワーク
- go 1.14
- echo
- gorm
- mysql
- docker/docker-compose

## ローカル環境セットアップ

ソースコードをクローンする
```
git clone https://github.com/devkeita/family-board-api.git
```

dockerを立ち上げる
```
docker-compose up -d
```

## EC2上で立ち上げる

ソースコードをクローンする
```
git clone https://github.com/devkeita/family-board-api.git
```

dockerイメージをビルドする
```
docker build -t api .
```

 dockerを立ち上げる
 ```
 docker run -d -p 9000:9000 --name api
 ```
