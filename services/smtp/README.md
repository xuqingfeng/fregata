# smtp

**API endpoint**

POST `/fregata/v1/smtp`

**input params**

```json
{
  "from": "from@example.com",
  "to": [
    "to@example.com"
  ],
  "cc": [
    "cc@example.com"
  ],
  "subject": "subject",
  "body": "email body"
}
```