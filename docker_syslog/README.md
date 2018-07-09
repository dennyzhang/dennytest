Customize docker log driver to use syslog

Read link:
- https://docs.docker.com/config/containers/logging/syslog/
- https://docs.docker.com/compose/compose-file/#logging

# How To Test

```
docker-compose build

docker-compose up -d

# dummy-container will generate dummy stdout, which will be sent to syslog-dummy-server automatically
docker-compose ps
# ,----------- Sample Output
# | bash-3.2$ docker-compose ps
# |        Name                Command         State                 Ports               
# | -------------------------------------------------------------------------------------
# | dummy-container       /root/start.sh       Up                                        
# | syslog-dummy-server   /srv/tcp-to-stdout   Up      0.0.0.0:12346->12346/tcp, 8080/tcp
# `-----------

# Check syslog-server, we will see one dummy log
docker logs --tail 10 syslog-dummy-server

# Sample Output:
# ,-----------
# | bash-3.2$ docker logs -f syslog-dummy-server
# | 2018/07/05 19:14:45 listening on addr: :12346
# | <30>Jul  5 19:14:46 5d0a710e1022[1127]: i: 1
# | <30>Jul  5 19:14:47 5d0a710e1022[1127]: i: 2
# | <30>Jul  5 19:14:48 5d0a710e1022[1127]: i: 3
# | <30>Jul  5 19:14:49 5d0a710e1022[1127]: i: 4
# | <30>Jul  5 19:14:50 5d0a710e1022[1127]: i: 5
# `-----------
```
