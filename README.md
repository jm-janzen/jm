# JM Online

Source of my canonical online presences.

|source|url|desc|
|---|---|---|
|[jmjanzen.com](./domains/jmjanzen.com)|https://jmjanzen.com|my main website|

---

## Run jmjanzen.com locally

### Launch manually

```shell
go run cmd/jmjanzen.com/main.go
```

### Watch file changes

Using [air](github.com/air-verse/air) (might need to proxy port -- see official docs).

```shell
air -c ./config/.air.jmjanzen.com.toml
```
