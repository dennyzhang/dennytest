Table of Contents
=================
   * [Use http to send log](#use-http-to-send-log)
   * [How To Test](#how-to-test)
   * [Generate log data](#generate-log-data)
   * [Trouble shooting](#trouble-shooting)
   * [Useful Links](#useful-links)

# Use http to send log
```
# export API_TOKEN="XXX" # Change this

curl -X POST \
  https://data.mgmt.cloud.vmware.com/le-mans/v1/streams/ingestion-pipeline-stream \
  -H "Authorization:Bearer $API_TOKEN" \
  -H 'Content-Type:application/json' \
  -H 'structure:default' \
  -d '[{
       "text": "Thu, 01 Mar 2018 20:41:42 GMT Test Payload-test",
       "source": "myhost.vmware.com"
   }]'
```

# How To Test
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

service apache2 start

service apache2 status
```

- Configure /etc/td-agent/td-agent.conf

```
# Overwrite td-agent.conf with td-agent.conf.tmpl
# Change Authorization Bearer <...> to real token

# restart td-agent
/etc/init.d/td-agent stop
/etc/init.d/td-agent start
```

# Generate log data
- generate log
```
# hit td agent
curl -X POST -d 'json={"json":"This is a test message from Fluentd HTTP"}' http://localhost:8888/debug.test

# hit apache
curl http://localhost

# hit via http

# export API_TOKEN="XXX" # Change this

curl -X POST \
  https://data.mgmt.cloud.vmware.com/le-mans/v1/streams/ingestion-pipeline-stream \
  -H "Authorization:Bearer $API_TOKEN" \
  -H 'Content-Type:application/json' \
  -H 'structure:default' \
  -d '[{
       "text": "Thu, 01 Mar 2018 20:41:42 GMT MyTest Payload-test",
       "source": "myhost.vmware.com"
   }]'
```

- search log intelligence

```
search by "test", for td agent log

search by "GET", for apache log

search by "MyTest", for curl log
```

# Trouble shooting

```
/var/log/td-agent/td-agent.log

ls -lth /var/log/apache2/*.log
```

# Useful Links

https://cloud.vmware.com/community/2018/07/10/using-fluentd-send-logs-cloud-vmware-log-intelligence/

https://docs.vmware.com/en/VMware-Log-Intelligence/services/User-Guide/GUID-48C6CA73-FA99-42DE-851D-0A7930D08324.html#GUID-48C6CA73-FA99-42DE-851D-0A7930D08324__section_86D3FCA44F6A4D3CA93ACB081DD7B76A

https://fluentbit.io/documentation/0.13/output/http.html
