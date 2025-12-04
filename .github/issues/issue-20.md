---
id: 20
database_id: 460807219
node_id: MDU6SXNzdWU0NjA4MDcyMTk=
status: closed
title: "full support defaults of go.avito.ru/gl/psql"
labels: ["type:task"]
url: https://github.com/kamilsk/lift/issues/20
created_at: 2019-06-26T07:30:17Z
updated_at: 2020-09-05T16:21:54Z
---

# full support defaults of go.avito.ru/gl/psql

```
ParseEnvLibpq currently recognizes the following environment variables:
PGHOST
PGPORT
PGDATABASE
PGUSER
PGPASSWORD
PGSSLMODE
PGSSLCERT
PGSSLKEY
PGSSLROOTCERT
PGAPPNAME
PGCONNECT_TIMEOUT
```

```go
type Options struct {
	host                 string
	port                 int
	user                 string
	password             string
	dbName               string
	sslMode              string
	Logger               pgx.Logger
	LogLevel             pgx.LogLevel
	Dial                 pgx.DialFunc
	RuntimeParams        map[string]string                         // Run-time parameters to set on connection as session default values (e.g. search_path or application_name)
	OnNotice             pgx.NoticeHandler                         // Callback function called when a notice response is received.
	CustomConnInfo       func(*pgx.Conn) (*pgtype.ConnInfo, error) // Callback function to implement connection strategies for different backends. crate, pgbouncer, pgpool, etc.
	CustomCancel         func(*pgx.Conn) error                     // Callback function used to override cancellation behavior
	preferSimpleProtocol bool
}

var defaultOptions = Options{
	sslMode:              "disable",
	preferSimpleProtocol: true,
}
```
