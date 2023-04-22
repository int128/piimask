# piimask [![go](https://github.com/int128/piimask/actions/workflows/go.yaml/badge.svg)](https://github.com/int128/piimask/actions/workflows/go.yaml)

This is a tiny tool to mask Personally Identifiable Information (PII) in database.

Status: Proof of Concept

## Example

```sh
docker-compose up -d
DATABASE_URL=postgresql://app:example@localhost/app go run .
```

```sql
UPDATE users SET
first_name = 'REDACTED' /* character varying */,
last_name = 'REDACTED' /* character varying */,
email = 'REDACTED' /* character varying */,
phone = 'REDACTED' /* character varying */
;
UPDATE messages SET
sender_id = 'REDACTED' /* character varying */,
recipient_id = 'REDACTED' /* character varying */,
title = 'REDACTED' /* text */,
body = 'REDACTED' /* text */
;
```
