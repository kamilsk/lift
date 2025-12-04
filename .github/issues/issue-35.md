---
id: 35
database_id: 703571169
node_id: MDU6SXNzdWU3MDM1NzExNjk=
status: open
title: "output warnings for unknown fields"
labels: ["type:feature"]
url: https://github.com/kamilsk/lift/issues/35
created_at: 2020-09-17T13:06:41Z
updated_at: 2020-09-17T13:06:41Z
---

# output warnings for unknown fields

to prevent misprints

```
[default.handlers]
errors_codes = [500, 503, 504] // invalid key
```
