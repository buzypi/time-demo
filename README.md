Build using:

```
docker build -t time-demo .
```

Confirm that the image has only one file:

```
docker image history time-demo
```

Run it:

```
docker run --rm time-demo
docker run --rm -e TZ="America/New_York"  time-demo
docker run --rm -e TZ="UTC+5"  time-demo
```


