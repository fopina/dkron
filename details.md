`maxBufSize = 256000` in the shell handler, 100 executions per job, 25M to download when listing executions.

Add parameter to truncate execution output and use it in the UI (not by default to avoid API breaking changes)

## test

Create job with shell command `seq 1000000` (`seq 1000000 | wc -c` = `6888894`)

Run it 100+ times:

```
for i in $(seq 100); do curl 'http://localhost:8080/v1/jobs/test' -X POST; done
```
