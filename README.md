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

To run any particular domain locally, simply execute the command `go run ./cmd/main.go <DOMAIN>`, or if you want live reloading of the frontend use [air](github.com/air-verse/air) with the configuration files provided at `./config/air.<DOMAIN>.toml`.
