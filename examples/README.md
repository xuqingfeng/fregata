## slack

```bash
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/slack \
  -H 'content-type: application/json' \
  -d '{"message":"test", "username": "fregata", "channel": "@qingfeng", "icon_emoji": ":medal:"}'
```

![slack](./slack.png)

## wechat

```bash
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/wechat \
  -H 'content-type: application/json' \
  -d '{"message": "😃", "to": "filehelper"}'
```

<img src="https://raw.githubusercontent.com/xuqingfeng/fregata/master/examples/wechat.png" alt="wechat" width=300/>

## telegram

```bash
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/telegram \
  -H 'content-type: application/json' \
  -d '{"message": "*fregata*\n```go\nfmt.Println(\"Hello World.\")\n```"}'
```

<img src="https://raw.githubusercontent.com/xuqingfeng/fregata/master/examples/telegram.png" alt="telegram" width=300/>
