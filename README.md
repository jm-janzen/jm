# JM Online

Source of my canonical online presences.

|source|url|desc|
|---|---|---|
|[jmjanzen.com](./domains/jmjanzen.com)|https://jmjanzen.com|my main website|
|[api.jmjanzen.com](./domains/api.jmjanzen.com)|https://api.jmjanzen.com|my main website|
|[jmjanzen.cv](./domains/jmjanzen.cv)|https://jmjanzen.cv|my résumé|

---

## Run jmjanzen.com locally

### Launch manually

```shell
go run cmd/jmjanzen.com/main.go
```

### Watch file changes and live reload

Using [air](github.com/air-verse/air) (might need to proxy port -- see official docs).

```shell
air -c ./config/.air.jmjanzen.com.toml
```

## Run api.jmjanzen.com locally

```shell
go run cmd/api.jmjanzen.com
```

For local development I recommend the excellent [wgo](https://github.com/bokwoon95/wgo).

## Run jmjanzen.cv locally

```
go run cmd/jmjanzen.cv/main.go
```

### Make the html résumé

I mean it's just an html file, but I use the incredible [pdf2htmlEX](https://github.com/pdf2htmlEX/pdf2htmlEX), like:

```shell
pdf2htmlEX assets/resume.pdf ./assets/resume.html --zoom 1.5
```
