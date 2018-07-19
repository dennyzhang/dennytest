Use fluentd to create a log forward:
- input: syslog
- output: log intelligence http endpoint

# How To Test

```
docker-compose up -d
docker-compose ps

docker-compose logs fluentd-syslog-li
```

- Unknown output plugin 'http_ext'
```
## ,-----------
## | zdenny-a02:fluentd_syslog_to_logintelligence zdenny$ docker-compose logs fluentd-syslog-li
## | Attaching to fluentd-syslog-li
## | fluentd-syslog-li    | 2018-07-19 18:30:21 +0000 [info]: parsing config file is succeeded path="/fluentd/etc/fluent.conf"
## | fluentd-syslog-li    | 2018-07-19 18:30:21 +0000 [error]: config error file="/fluentd/etc/fluent.conf" error_class=Fluent::ConfigError error="Unknown output plugin 'http_ext'. Run 'gem search -rd fluent-plugin' to find plugins"
## `-----------
```

# More Debugging
```
docker exec -it fluentd-syslog-li sh
```
