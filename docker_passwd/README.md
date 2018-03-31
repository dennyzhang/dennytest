## Requirement

How I can create a Pod with a specific root password?

Even when the pod get restarted or recreated, the root password persist?

## How To Test

docker build -f Dockerfile -t denny/image:v1 --rm=true .
