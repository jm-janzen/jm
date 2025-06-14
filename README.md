# JM Online

Source of my canonical online presences.

|source|url|desc|
|---|---|---|
|[jmjanzen.com](./domains/jmjanzen.com)|https://jmjanzen.com|my main website|
|[api.jmjanzen.com](./domains/api.jmjanzen.com)|https://api.jmjanzen.com|my main website|

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

## Run api.jmjanzen.com locally

```shell
go run cmd/api.jmjanzen.com
```

For local development I recommend the excellent [wgo](https://github.com/bokwoon95/wgo).
