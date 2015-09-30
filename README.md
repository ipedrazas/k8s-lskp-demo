# Kubernetes Landscape Demo


Running applications in containers is easy. Running applications in Kubernetes is, once you know how the system works, relatively easily.
Managing applications is always hard regardless of where and how you run them.

This repo contains a few examples of code, containers and kubernetes artifacts.

### Step 1 - Simple go app

We have written our app. It's a microservice that says `ping!`. The service is exposed in port `8080`. You can run the app in your machine, in a container, and in a container that runs in Kubernetes.

* Run the app in your machine: assuming you have set up correctly your golang env, you just have to do `go run main.go`, then you can request the service `curl localhost:8080` and you should see the `ping!` result.

* Container: first you have to build your container, but before that, you have to build your service with all the dependencies (build.sh will do that for you). Then you build your container `docker build -t ipedrazas/ping` and finally you can run it: `docker run -it --rm -p 8080:8080 ipedrazas/ping`, then you can request the service `curl localhost:8080` and you should see the `ping!` result.


