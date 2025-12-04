---
id: 5
database_id: 443747151
node_id: MDU6SXNzdWU0NDM3NDcxNTE=
status: closed
title: "forward mapping"
labels: ["type:refactoring"]
url: https://github.com/kamilsk/lift/issues/5
created_at: 2019-05-14T07:12:54Z
updated_at: 2019-05-30T15:45:11Z
---

# forward mapping

ports are defined as one-to-one, for example, 6379 -> 6379, but it's not so flexible. some ports may be busy or not available. need possibility to define mapping for port forwarding.

`REDIS_PORT = 16379 (real is 6379)`
