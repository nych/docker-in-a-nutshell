# Docker Basics
Describes the most common artifacts, commands and concepts in the context of Docker.
## Terminology
### Docker Image
A *Docker Image* is a blueprint to instantiate *Docker Containers* from. A more detailed but abstract definition can be found on [docs.docker.com](https://docs.docker.com/glossary/#image). If you have worked with virtual machines before, think of it as a snapshot. If you are a programmer and have worked with a language that supports the object oriented paradigm, then think of it as a *Class*. The most important thing to understand in this context is the immutability of *Docker Images*.

### Dockerfile
A *Dockerfile* contains a set of instructions divided with newlines that the tell the *Docker Builder* how to create/build a *Docker Image*. A comprehensive documentation can be found on [docs.docker.com](https://docs.docker.com/engine/reference/builder).

### Docker Container
A *Docker Container# is a a runnable instance of a *Docker Image*. If you have worked with virtual machines before, think of it as a virtual machine. Unlike a virtual machine in the traditional sense, a container has no need for virtualized hardware but is just a simple process with a higher degreee of isolation. If you are a programmer think of it as an *Object* instanciated from a *Class*. For further information check out the slides or the following resources:
- [atlassion.com](https://www.atlassian.com/microservices/cloud-computing/containers-vs-vms)
- [sites.google.com](https://sites.google.com/site/mytechnicalcollection/cloud-computing/docker/container-vs-process).

### Docker Registry
A *Docker Registry* manages *Docker Images* in a centralized way. The official registry hosted by Docker can be found on [hub.docker.com](https://hub.docker.com). There are other public available registries like *Microsofts* artifact registry.

