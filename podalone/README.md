# Podalone
## Pods, replication controllers and services

"Can you have a pod without replication controller and service?"

Yes, you can... but!

Replication controllers are in charge of making sure that the expected pods are running. Services allow pods and external services to talk to each other.

Imagine a pod that does a task and then it dies... Do we want to have a pod?  a RC? or a service?

If our pod doesn't need to be reachable by any other service we don't need the Service. Workers are a kind of "Pods that might not need services".

Do we want an RC? well, it depends. Imagine we have a pod that has a container that increments 10 times a value. Do we want to have that pod always up? What happens once the container has finished its task?

The `podalone-pod.yaml` contains a pod with a container that does that: it goes to a redis service and increments the value every 30 seconds and then it dies. Then we have the same pod inside an RC with 1 replicas.
