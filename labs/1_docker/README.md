# Docker Lab
This lab contains different parts which must be solved. The resulting artifacts (Dockerfiles, Scripts, Sourcefiles, etc.) can be stored in this directory.
## Part 1
Use the Docker CLI to pull the python image from hub.docker.com. You are free to choose what image version you want to use. Once the Image is pulled to your host, try to figure out its id and size using the CLI. Bonus points if you can also figure out its entrypoint and/or cmd that will be exectued if you run a container from it. Write a python script which prints "Hello World" and run this script using the docker image you pulled previously. Do NOT write the script within the container but on your host.
## Part 2
Use the Docker CLI to build an image that contains the script you have implemnted in *Part 1*. The image should run the script when a container is created from it. Change the script and/or your Dockerfile in such way, that the text printed is taken as an argument. If you pass "FizFaz" as parameter, the output on your terminal should be FizFaz. Bonus points if your Image and/or Script can handle a case where no parameter is passed.
## Part 3
Use the docker CLI to remove all artifacts from your system. Bonus points if you can also clean the docker build cache.