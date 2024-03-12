# Docker Lab
This lab contains different parts which must be solved. The resulting artifacts (Dockerfiles, Scripts, Sourcefiles, etc.) can be stored in this directory.
## Part 1
**Task Objective:** Utilize the Docker Command Line Interface (CLI) to manage a Python environment.
1) **Pull Python Image:** Use the Docker CLI to pull a Python image from Docker Hub (hub.docker.com). You are free to choose any version of the Python image you prefer.
2) **Identify Image Details:** After pulling the image, use the Docker CLI to find the image's ID and size.
As a bonus challenge, identify the image's ENTRYPOINT and CMD instructions. These specify the default executable and arguments for containers based on this image.
3) **Run a Python Script:** Write a Python script on your host machine that prints "Hello World". Save this 
Run this script using the Docker image you previously pulled, ensuring not to write the script within the container but to execute it from your host's filesystem. (Hint check out the `-v` option)

**Instructions Summary:**
* Pull a specific version of the Python Docker image.
* Use the Docker CLI to find the image's ID, size, and optionally, its ENTRYPOINT and CMD.
* Write a "Hello World" Python script on your host and run it using the pulled Docker image.

## Part 2
**Task Objective:** Create a Docker image that includes and runs a Python script, with the ability to pass arguments to the script at runtime.
1) **Prepare Your Python Script:** Modify your hello_world.py script to accept command-line arguments. Use a default value that ensures the script prints something meaningful even when no arguments are provided.
Hint: Consider using the sys.argv from the sys module for parsing command-line arguments.
2) **Write a Dockerfile:** Create a Dockerfile that specifies the Python image you used in Part 1 as the base. Ensure your Dockerfile copies the modified hello_world.py script into the image. Define the command to run your script, making sure it can receive arguments passed during container creation.
3) **Build and Run Your Docker Image:** Build your Docker image using the Docker CLI. Tag it with a meaningful name and version. Run a container from your image, passing an argument to your script. Verify the output matches the argument. Test the script without passing arguments to ensure it still prints a default message.

**Instructions Summary:**
* Modify the Python script to accept and print command-line arguments, with a default behavior when no arguments are given.
* Write a Dockerfile that builds an image with your script, ready to accept runtime arguments.
* Build the Docker image, then run containers from it to test with and without an argument.

## Part 3 (Bonus)
**Task Objective:** Learn to use the Docker Command Line Interface (CLI) to remove Docker artifacts from your system, including images, containers, and the build cache.
1) **Remove Docker Containers:** First, you'll need to remove any running or stopped containers. Containers must be stopped before they can be removed.
2) **Remove Docker Images:** With all containers stopped and removed, you can now remove the Docker images. This will free up space on your system.
**Caution:** Removing Docker images cannot be undone and might result in loss of data if you have not pushed your images to a registry or have local-only images.
3) **Clean the Docker Build Cache (Bonus):** Cleaning the build cache is an optional but recommended step to reclaim disk space and ensure your Docker builds start from a clean state.

**Instructions Summary:**
* Use Docker CLI commands to stop and remove all containers from your system.
* Remove all Docker images, being mindful of the potential for data loss.
* As a bonus challenge, clear the Docker build cache to free up additional space and maintain a clean build environment.