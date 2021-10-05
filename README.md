# サーバー起動

`go run main.go`

# Get Started

## Hello World

```terminal:Shell
$ curl http://localhost:3000/

> Hi!
```

# API Reference

## API Information

#### `GET` `api/v1/info`

```terminal:Shell
$ curl http://localhost:3000/api/v1

> Magiot Game API Version 1.0
```

## Devises

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
