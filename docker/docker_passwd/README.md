## Requirement

How I can create a Pod with a specific root password?

Even when the pod get restarted or recreated, the root password persist?

## How To Test
1.1 Build and start container
```
export docker_image="denny/image:v1"
docker build -f Dockerfile -t "$docker_image" --rm=true .

docker stop my-test; docker rm my-test
docker run -d -t --privileged -p 2222:2222 -h mytest --name my-test "$docker_image"
```

1.2 Update root password
```
docker exec -it my-test bash
passwd # change password from default(root) to something else (password1234)
```

1.3 SSH via root password

Login and logout 3 times. Confirm whether we can use the same ssh root password to login
```
ssh -p 2222 root@127.0.0.1
exit
```
