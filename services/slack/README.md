# slack

**API endpoint**

POST `/fregata/v1/slack`

**input params**

```json
{
    "text": "text",
    "channel": "channel",
    "username": "username",
    "icon_emoji": ":medal:",
    "attachments": [
        {
            "fallback": "fallback",
            "color": "warn",
            "pretext": "pretext",
            "author_name": "author name",
            "author_link": "author link",
            "author_icon": "author icon",
            "title": "title",
            "title_link": "title link",
            "fields": [
                {
                    "title": "title",
                    "value": "value",
                    "short": true
                }
            ],
            "image_url": "image url",
            "thumb_url": "thumb url",
            "footer": "footer",
            "footer_icon": "footer icon",
            "ts": 1489465322
        }
    ]
}
```