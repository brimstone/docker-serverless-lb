docker-serverless-lb
====================

Example serverless http handler.

Architecture
------------

```
+---------+    +---------+           +-----------------+
| Browser |--->| main.go |--+-- / -->| root/Dockerfile |
+---------+    +---------+  |        +-----------------+
                            |
                            |               +-------------------+
                            +-- /assets/ -->| assets/Dockerfile |
                                            +-------------------+
```

`main.go`
---------

This is the main load balancer. It's simply configured to launch a container
and pass it the request from the browser with CGI.

It connects to the Docker engine defined by environment variables. This allows
the loadbalancer to connect to engines local, remote, or swarm managers.


`root/Dockerfile`
-----------------

This returns a simple HTML file for any request.


`assets/Dockerfile`
-------------------

This returns jeff's face for only the `/assets/goldblum.jpg` route. Everything
else gets a 404.


Ideas
-----

- Support accepting a PUT at any endpoint with a compose yaml/json defining the
  container to use
- Index images available on the server and setup endpoints with the labels
  declared in the images
