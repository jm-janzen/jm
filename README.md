# JM Online

Source of my canonical online presences.

|source|url|desc|
|---|---|---|
|[jmjanzen.com](./domains/jmjanzen.com)|https://jmjanzen.com|my main website|
|[api.jmjanzen.com](./domains/api.jmjanzen.com)|https://api.jmjanzen.com|api of me|
|[jmjanzen.cv](./domains/jmjanzen.cv)|https://jmjanzen.cv|my résumé|
|[blog.jmjanzen.com](./domains/blog.jmjanzen.com)|https://blog.jmjanzen.com|the blog|

---

## Running locally

To run any particular domain locally, simply execute the command `go run ./cmd/main.go <DOMAIN>`, for example

```shell
go run ./cmd/main.go api.jmjanzen.com
```

### Running locally (with live reloading)

If you want live reloading of the frontend use [air](github.com/air-verse/air) run the following, and make sure to connect to the `proxy_port` to actually get the live reloading. By default, this will use the base configuration, [here](./.air.toml).

```shell
air -proxy.app_port 8080 -proxy.proxy_port 9090 jmjanzen.com
```

If you just want to run everything, say like on a VPS, run the following commands to get started:

```shell
cp .env.example .env
cp config/db/seed.example.js assets/

docker compose up -d
```
