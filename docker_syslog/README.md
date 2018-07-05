Customize docker log driver to use syslog

Read link:
- https://docs.docker.com/config/containers/logging/syslog/
- https://docs.docker.com/compose/compose-file/

# How To Test

```
docker-compose build

docker-compose up -d

# We will see two containers: syslog-dummy-server and dummy-container
docker-compose ps

# Check syslog-server, we will see one dummy log
docker logs syslog-dummy-server

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
