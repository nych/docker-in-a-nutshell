# Docker Cheat Sheet
An incomplete cheatsheet on how to interact with Docker.
## Interact with the Registry
### Search
If you want to search for a certain image you can do this by using the *Docker CLI* or search for it on [hub.docker.com](hub.docker.com).

`docker search hello-world`
### Pull
If you want to use an *Image* on your host machine, you have to download it firstly. This can be achieved with the following command.

`docker pull hello-world`

## Interact with Containers
### Create/Run containers
A container can be created with the following command.

`docker create hello-world`

This just creates the container but does not start it.

`docker start hello-world`

There is a shortcut combining the previous two commands. In practice the `run` command is seen more often. Something that applies to both `create` and `run` is that they pull they attempt to pull the image if it is not yet available on the system.

`docker run hello-world`

### List containers
To list all running containers use the following command.

`docker ps`

This will not show containers in a state other than *Up*. To list containers in other states, set the *--all* flag.

`docker ps --all`

### Stop container
To stop a container you can use either the *stop* command or the *kill* command. They differ in the way they terminate the container. Stop is the equivalent to a *SIGTERM* whereas kill is the equivalent to a *SIGKILL*. However both commands can specify the signal to send with the *--signal* option.

`docker stop our-hello-world-container`

`docker kill our-hello-world-container`

### Remove containers
Containers which are not in the *Running* state can be removed using the `rm` command.

`docker rm our-hello-world-container`

When running a container you can set the `--rm` flag to automatically delete the container when it exits.

### Open a shell in a running container
Sometimes it is necessary to have a shell in a running container. This can be achieved by using the `exec` command. Exec allows us to execute a command in a running container which in combination with the `--interactive` and `--tty` flag can be used to get a terminal.

`docker exec -it our-ubuntu-container sh`

## Interact with Images
### List local images
To get a list of images that are available on your host machine, the following commands can be used

`docker image ls`

or

`docker images`

### Build Image
Docker needs two things to build an image. The path to a *Dockerfile* containing the instructions how to build the image and a path to the *Build Context*. The *Build Context* often causes confusion when new to docker. Let's examine the following command to build an image from the *Dockerfile* in this directory.

`docker build --file ./Dockerfile --tag our-first-image:latest .`

The `--file` property is not necessary in this case as the `build` command defaults to `PATH/Dockerfile` where path refers to the *Build Context* which is passed as an argument to the build command. In the example shown the *Build Context* is the current directory indicated by the *.* at the end. It is vital to know that everything you refer in your *Dockerfile* must be present in the *Build Context*. In this case the *Dockerfile* refers the files `requirements.txt` and `main.py` in its `COPY` instructions.
If there are files and/or directories you want to exclude from the *Build Context* you can do so by adding a `.dockerignore` file. Its contents are very similar to a `.gitignore` file. See the *extended-devcontainer* demo which uses a `.dockerignore` file to exclude the `build/` directory from the *Build Context*.
With the `--tag` property it is possible name and tag the resulting image. The property name and its help message are a bit missleading as it states *Name and optionally a tag (format: "name:tag")*. How to name/tag images depends on whether it is to be pushed on the public registry or a private one. See [docs.docker.com](https://docs.docker.com/engine/reference/commandline/tag/) for more information about naming conventions and format.

### Remove image
An image can be removed with the following commands

`docker image rm our-first-image:latest`

or

`docker rmi our-first-image:latest`
