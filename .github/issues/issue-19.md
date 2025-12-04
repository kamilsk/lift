---
id: 19
database_id: 460805054
node_id: MDU6SXNzdWU0NjA4MDUwNTQ=
status: closed
title: "full support defaults of go.avito.ru/gl/redis"
labels: ["type:task"]
url: https://github.com/kamilsk/lift/issues/19
created_at: 2019-06-26T07:25:05Z
updated_at: 2020-09-05T16:22:10Z
---

# full support defaults of go.avito.ru/gl/redis

```
"REDIS_HOST"
"REDIS_PORT"
"REDIS_DATABASE"
"REDIS_MAX_ACTIVE"
"REDIS_MAX_IDLE"
"REDIS_CONNECT_TIMEOUT"
"REDIS_IDLE_TIMEOUT"
"REDIS_READ_TIMEOUT"
"REDIS_WRITE_TIMEOUT"

"REDIS_SHARDED_HOST"
"REDIS_SHARDED_PORT"
"REDIS_SHARDED_MAX_ACTIVE"
"REDIS_SHARDED_MAX_IDLE"
"REDIS_SHARDED_CONNECT_TIMEOUT"
"REDIS_SHARDED_IDLE_TIMEOUT"
"REDIS_SHARDED_READ_TIMEOUT"
"REDIS_SHARDED_WRITE_TIMEOUT"
```

```go
type Options struct {
	network        string
	host           string
	port           int
	db             int
	maxActive      int
	maxIdle        int
	connectTimeout time.Duration
	idleTimeout    time.Duration
	readTimeout    time.Duration
	writeTimeout   time.Duration
}

var defaultOptions = Options{
	network:        "tcp",
	db:             -1,
	connectTimeout: defaultConnectionTimeout,
	port:           6379,
}
```
