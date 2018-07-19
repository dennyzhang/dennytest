**Use fluentd to create a log forwarder for vmware log intelligence product.**

By default, log intelligence accept http protocal, instead of syslog protocal

https://cloud.vmware.com/community/2018/07/10/using-fluentd-send-logs-cloud-vmware-log-intelligence/

This post support you sending log to log intelligence via syslog protocal.

Use fluentd as an adapter from syslog to http
- input: syslog plugin
- output: log intelligence http endpoint

Note: this workflow won't give your buffering flush benefit. Thus it won't help on performance side. In vmware world, we will have a more official support from vmware directly.

# How To Test

- Update log intelligence token in fluent.conf

Search and change `    Authorization Bearer CHANGETHIS` to real token.

```
# build image to install fluentd http-ext output plugin
docker-compose build

# Start env & Generate log
docker-compose up -d
docker-compose ps

## ,----------- Sample Output
## | bash-3.2$ docker-compose ps
## |       Name                     Command               State                       Ports                    
## | ----------------------------------------------------------------------------------------------------------
## | fluentd-syslog-li   /bin/entrypoint.sh /bin/sh ...   Up      24224/tcp, 0.0.0.0:40012->40012/tcp, 5140/tcp
## `-----------
```

- Generate log from laptop using rfc5424
```
echo "<14>1 2017-02-28T12:00:00.009Z 192.168.0.1 denny - - - Hello." | nc localhost 40012
```

- Check syslog output
```
docker-compose logs -f fluentd-syslog-li
```

- Go to log intelligence dashboard, then search pattern of `denny`

![images/log_intelligence.png](images/log_intelligence.png)
