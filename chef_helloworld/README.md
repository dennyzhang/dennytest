- Get chef image for binaries and libraries
docker pull chef/chef

- Start chef container
```
export docker_image="chef/chef"
docker stop my-test; docker rm my-test
docker run -t -d --privileged -h mytest --name my-test --entrypoint=/bin/bash "$docker_image"

docker exec -it my-test bash

which curl
```

- start helloworld

https://www.morethanseven.net/2010/10/30/Chef-hello-world/

example cookbook will only install curl package
```
sudo chef-solo -c config/solo.rb -j config/node.json

which curl
```
