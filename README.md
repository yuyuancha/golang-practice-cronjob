# go-practice-cronjob

## 簡述

透過 `cron` 套件，來實作排程功能。

## 步驟及說明

* 紀錄程式開始時間。
* 設定排程每 2 秒、每 5 秒執行一次打印，並顯示時間。
* 為了執行排程，程式暫停一分鐘才結束。

## 安裝 docker 環境及執行程式

* clone GitHub repository
```
$ get clone https://github.com/yuyuancha/golang-practice-goroutine.git
```

* 透過 docker-compose 開啟 docker 容器
```
$ docker-compose up -d
```

* 執行 main.go
```
$ docker-compose exec app go run main.go
```

* 關閉 docker 容器
```
docker-compose down
```