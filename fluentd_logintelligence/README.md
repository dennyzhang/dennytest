https://cloud.vmware.com/community/2018/07/10/using-fluentd-send-logs-cloud-vmware-log-intelligence/

- start a container

```
export docker_image="ubuntu:14.04"
docker stop my-test; docker rm my-test
docker run -t -d --privileged -h mytest --name my-test --entrypoint=/bin/sh "$docker_image"
```

- Install facility

```
docker exec -it my-test bash

apt update

apt-get install -y curl

curl -L https://toolbelt.treasuredata.com/sh/install-ubuntu-trusty-td-agent3.sh | sh
/usr/sbin/td-agent-gem install fluent-plugin-out-http-ext
mkdir /tmp/log
chmod -R 777 /tmp/log

# start agent
/etc/init.d/td-agent start

# check status
/etc/init.d/td-agent status
```

- Run dummy process: apache

```
apt-get install apache2

chmod -R 645 /var/log/apache2
```

- Configure /etc/td-agent/td-agent.conf
