Use fluentd to create a log forwarder:
- input: syslog plugin
- output: log intelligence http endpoint

# How To Test

- Update log intelligence token in fluent.conf

Search and change `    Authorization Bearer CHANGETHIS` to real token.

```
# build image to install fluentd http-ext output plugin
docker-compose build

# Start env & Generate log
docker-compose up -d
docker-compose ps
```

- Generate log from laptop using rfc5424
```
echo "<14>1 2017-02-28T12:00:00.009Z 192.168.0.1 denny - - - Hello." | nc localhost 40012
```

- Check syslog output
docker-compose logs -f fluentd-syslog-li

- Go to log intelligence dashboard, then search pattern of `denny`

![images/log_intelligence.png](images/log_intelligence.png)

# More Debugging
```
docker exec -it fluentd-syslog-li sh
```
