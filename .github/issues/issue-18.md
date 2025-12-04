---
id: 18
database_id: 458843847
node_id: MDU6SXNzdWU0NTg4NDM4NDc=
status: closed
title: "add possibility to merge other env vars"
labels: ["type:task"]
url: https://github.com/kamilsk/lift/issues/18
created_at: 2019-06-20T19:54:14Z
updated_at: 2020-09-05T16:23:22Z
---

# add possibility to merge other env vars

```bash
$ lift --append|-a jira-123 env
merge(dep.vars, engine.vars, env_vars, envs.local.env_vars, envs.jira-123.env_vars)
```
