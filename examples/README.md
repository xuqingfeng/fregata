## slack

```
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/slack \
  -H 'content-type: application/json' \
  -d '{"text":"test", "username": "fregata", "channel": "@qingfeng", "icon_emoji": ":medal:"}'
```

![slack](./slack.png)

## wechat

```
curl -X POST \
  http://127.0.0.1:2017/fregata/v1/wechat \
  -H 'content-type: application/json' \
  -d '{"text": "😃", "to": "filehelper"}'
```

![wechat](./wechat.png)