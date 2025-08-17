Say you have an image named `asimov`, you would simply run:

```shell
docker save asimov:latest | gzip | ssh user@remote docker load
```

Then you can use the image of that name:tag as if it was built on your remote. For instance, by doing something like:

```shell
ssh user@remote docker run -p 80:8080 --name the.mule -d asimov:latest
```

And Bob's your uncle!
