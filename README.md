# サーバー起動

`go run main.go`

# Get Started

## DB の接続情報を書いた.env ファイルを以下の内容で準備

```
CONNECT=user:password@tcp(localhost:3306)/database_name?parseTime=true
```

## Hello World

```terminal:Shell
$ curl http://localhost:3000/

> Hi!
```

# API Reference

## API Information

#### `GET` `api/v1`

```terminal:Shell
$ curl http://localhost:3000/api/v1

> Magiot Game API Version 1.0
```

## Devices

### 全デバイスの取得

#### `GET` `api/v1/devices`

Sample Code

---

Shell

```terminal
$ curl http://localhost:3000/api/v1/devices
```

Response

```json
[
  {
    "id": 1
    "name": "sample_device"
  },
  {
    "id": 2
    "name": "sample_device2"
  }
]
```

### デバイスの追加

#### `POST` `api/v1/devices/{device_name}`

Sample Code

---

Shell

```terminal
$ curl http://localhost:3000/api/v1/devices/sample_device
```

Response

```json
{
  "id": 1
  "name": "sample_device"
}
```

### 指定した id のデバイスの取得

#### `GET` `api/v1/devices/{id}`

Sample Code

---

Shell

```terminal
$ curl http://localhost:3000/api/v1/devices/1
```

Response

```json
{
  "id": 1
  "name": "sample_device"
}
```

### 指定した id のデバイスの削除

#### `DELETE` `api/v1/devices/{id}`

Sample Code

---

Shell

```terminal
$ curl -X DELETE http://localhost:3000/api/v1/devices/1
```

Response

```json

```
