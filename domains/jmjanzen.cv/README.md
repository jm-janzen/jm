# JM's online portfolio I guess?

Anyway, the resume is served by this dumb router. Eventually would like to do something more interesting with this domain/module, but for now I just did this because I saw how cheap the `cv` TLD was with my registrar.

### Make the html résumé

I mean it's just an html file, but I use the incredible [pdf2htmlEX](https://github.com/pdf2htmlEX/pdf2htmlEX) to turn my PDF résumé into an HTML copy.

```shell
pdf2htmlEX assets/resume.pdf ./assets/resume.html --zoom 1.5
```
