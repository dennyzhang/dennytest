
# Use wavefront pythont client

Here we show how to verify wavefront token is valid

https://github.com/wavefrontHQ/python-client

## Start python test env

```
docker run -t -d --privileged -h mytest --name my-test -v $PWD/client_test.py:/root/client_test.py:rw --entrypoint=/bin/sh "python:2.7"

docker exec -it my-test pip install wavefront-api-client
```

## Update credentials

Change client_test.py for base_url and api_key. Something like below:

```
# TODO: change base_url and api_key
base_url = 'https://try.wavefront.com'
api_key = 'XXX'
```

## Run test

In client_test.py, we have dummy code query wavefront server.
```
docker exec -it my-test python /root/client_test.py | head

# If we see output like this, our wavefront token is valid
## ,----------- Sample Output
## |    /Users/zdenny  docker exec -it my-test python /root/client_test.py | head                                                                                                         ✔ 0
## | {'response': {'cursor': '10.85.93.13',
## |               'items': [{'created_epoch_millis': None,
## |                          'creator_id': None,
## |                          'description': None,
## |                          'hidden': False,
## |                          'id': '(none)',
## |                          'marked_new_epoch_millis': None,
## |                          'source_name': '(none)',
## |                          'tags': {u'~status.ok': True},
## |                          'updated_epoch_millis': None,
## `-----------
```
