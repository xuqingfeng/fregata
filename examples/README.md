## slack

```bash
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/slack \
  -H 'content-type: application/json' \
  -d '{"text":"test", "username": "fregata", "channel": "@qingfeng", "icon_emoji": ":medal:"}'
```

![slack](./slack.png)

## wechat

```bash
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/wechat \
  -H 'content-type: application/json' \
  -d '{"text": "😃", "to": "filehelper"}'
```

<img src="https://raw.githubusercontent.com/xuqingfeng/fregata/master/examples/wechat.png" alt="wechat" style="max-width: 50%;"/>

## telegram

```bash
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/telegram \
  -H 'content-type: application/json' \
  -d '{"text": "*fregata*\n```go\nfmt.Println(\"Hello World.\")\n```"}'
```

<img src="https://raw.githubusercontent.com/xuqingfeng/fregata/master/examples/telegram.png" alt="telegram" style="max-width: 50%;"/>
