# Kubernetes Introduction


Running applications in containers is easy. Running applications in Kubernetes is, once you know how the system works, relatively easily.
Managing applications is always hard regardless of where and how you run them.

This repo contains a few examples of code, containers and kubernetes artifacts.

* [Ping](ping): a simple http endpoint
* [PingPodRedis](pingpodredis): a simple http endpoint that reads & writes to a Redis database. Both containers are in a Pod.
* [Ping-redis](ping-redis):  a simple http endpoint that reads & writes to a Redis database. It runs on different Pods using an internal service.





Useful Commands

Get cluster info:

    kubectl cluster-info

Get pods

    kubectl get pods

Get services

    kubectl get services

Describe a service: handy when you want to get the public dns

    kubectl describe service pingredis

Scaling up & down

    kubectl scale --replicas=3 rc pingredis

Check your pods are created

    kubectl get pods

Delete one of them to see how a new one is created

    kubectl delete pod pingredis-XXXXX


Debugging

    kubectl get events

You can get the logs of your containers easily too. If your pod runs more than one container `kubectl logs POD_NAME CONTAINER`
    kubectl logs redis-XXXXX

There are ways of interacting directly with your containers. For example, you can `exec` a command in your container. Like when doing a `docker exec...`
    kubectl exec redis-XXXXX -- redis-cli get ping

Or you can do port forwarding between your host and your container

    kubectl port-forward -p redis-XXXXX 6380:6379


