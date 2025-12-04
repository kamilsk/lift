---
id: 7
database_id: 444805612
node_id: MDU6SXNzdWU0NDQ4MDU2MTI=
status: closed
title: "exclude duplicates"
labels: ["type:bug"]
url: https://github.com/kamilsk/lift/issues/7
created_at: 2019-05-16T07:51:38Z
updated_at: 2019-05-25T05:55:35Z
---

# exclude duplicates

```bash
$ lift forward -f testdata/app.toml 
forward -- demo-local-mongodb- 27017 demo-local-postgresql- 5432 demo-local-rabbitmq- 5672 5672 demo-local-redis- 6379 demo-local-sphinx- 9306
```

**demo-local-rabbitmq- 5672 5672**
