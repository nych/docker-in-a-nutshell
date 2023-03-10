# Docker Compose Cheat Sheet
An incomplete cheatsheet on how to interact with Docker Compose.

## Start an application
To start up an application declared in the compose file, you can use the `up` command. When resources are missing compose takes care of that too. It creates networks and volumes as well as it pulls or builds required docker images.

`docker compose up`

To start the composition in detachted mode (in the background), set the `--detach` flag.

## Shutdown an application
To shutdown an application declared in the compose file, you can use the `down` command. This will terminate all services/containers and remove other resources like networks.

`docker compose down`

## Compose commands vs. Docker commands
Compose is using docker and thus it is possible to use the regular docker commands to interact with artifacts created and managed by compose. In practice this is neither wrong nor right but you should be aware that certain things may not behave the way you would expect it. Or to put it in other words, it is easy to mess up those commands and end up getting an unexpected result. Let's consider the following example. You have started the demo service with the `compose up --detach` command. Now you want to see how your application behaves if the redis service dead. Using dockers `kill` command you terminate the redis service and do your tests. Satisfied with the result you are curious how your application behaves if the redis service becomes available again. Without thinking about the consequences you are typing `docker run -d redis` and are surprised, that your application still fails. What happened? There is a redis container but as you didn't explicitly set its name with the `--name` property, the name was generated randomly. The app service however specifies that the redis host (container name which can be resolved) is redis. There are two options to solve the problem. Either you explicitly set the name `docker run --name redis -d redis` or you use the `compose up` command. Compose will figure out that the redis service (represented by a container named redis) is not up and running and start the redis service.