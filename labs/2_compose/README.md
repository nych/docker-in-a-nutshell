# Docker Compose Lab
This lab contains different parts which must be solved. The resulting artifacts (Dockerfiles, Scripts, Sourcefiles, etc.) can be stored in this directory. For this excercise use the compose demo you can find at ../demos/compose-demo.
## Part 1
**Task Objective:** Learn to start a multi-container application using Docker Compose without blocking the terminal and subsequently check the logs of each service individually.
1) **Starting the Application in Detached Mode:** Navigate to the directory containing your docker-compose.yml file. Use Docker Compose to start the application in detached mode. This mode allows the services to run in the background, keeping your terminal session free.
2) **Verifying Service Health:** Once the application is up and running in detached mode, you'll want to ensure it's working as expected. The first step is to check the overall status of your services. Next, inspect the logs of each service individually to ensure there are no errors and that each service is functioning correctly. This step is crucial for troubleshooting and understanding how services interact within your application.

**Instructions Summary:**
* Start the compose-demo application in detached mode using the appropriate Docker Compose command.
* Verify the application's health by first checking the status of all services, then inspecting the logs of each service individually.

## Part 2
**Task Objective:** Enhance an application service by adding a new route that displays a list of the service's available routes/endpoints, then deploy the updated service.
1) **Implementing the New Route:** Within the service's code, add a new route ("/") that handles HTTP GET requests. This route should return a JSON response containing a list of all the service's available routes/endpoints. *Make sure to test this new route locally to ensure it behaves as expected, returning a correctly formatted list of the service's routes.*
2) **Update Docker Compose Configuration:** If necessary, update your service's Dockerfile to ensure any new dependencies required by your route implementation are included. Ensure your docker-compose.yml file is updated to build the latest version of your service if it does not already do so automatically.
3) **Deploy Your Changes:** From the directory containing your docker-compose.yml, rebuild and redeploy your service to incorporate the changes. After deployment, verify the new route is functioning correctly by accessing it and checking the response.

**Instructions Summary:**
* Add a new route ("/") to your service that returns a list of the service's available routes/endpoints.
* Test the new route locally to ensure it works as intended.
* Update your Dockerfile and docker-compose.yml as necessary to include any new dependencies.
* Rebuild and redeploy your service using Docker Compose to make your changes live.
* Verify the deployment by accessing the new route and ensuring it returns the correct response.

## Part 3 (Bonus)
**Task Objective:** Initialize the compose-demo service using Docker CLI commands, achieving the same setup as would be done with Docker Compose, without actually using the Docker Compose tool.
1) **Analyze the Docker Compose File:** Carefully read through the docker-compose.yml file for the compose-demo service. Note the services, networks, volumes, and any other configurations like environment variables or port mappings.
2) **Create Networks and Volumes:** Using Docker CLI, manually create any networks and volumes that are defined in the docker-compose.yml file.
3) **Build Images (if applicable):** If the docker-compose.yml file specifies building an image from a Dockerfile, replicate this step using Docker CLI.
4) **Run Containers:** For each service defined in the docker-compose.yml, start a container using the equivalent docker run command. Make sure to replicate the environment variables, network attachments, volume mounts, and port mappings as specified.
5) **Verify the Setup:** After starting all containers, verify that the network, volume, and service configurations work as intended. You can inspect running containers, networks, and volumes to ensure they are correctly set up and interconnected.
6) **Create a Bash Script:** Enhance your understanding and efficiency by encapsulating all the above Docker CLI commands into a single executable Bash script. This script should automate the setup process, simulating the ease of deployment provided by Docker Compose.