# Go Users API

This is a basic CRUD API for users. The API is built with the following tools:

- Go: The backend language
- MySQL: The database
- Docker: For containerization and managing the application's environment
- Docker Compose: For defining and running multi-container Docker applications

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).

## Running the Project

To run this project, follow these steps:

1. Copy the `.env.example` file to a new file named `.env`:
```
cp .env.example .env
```


3. Open the `.env` file and replace the dummy values with your actual values.

4. Run the following command to start the Docker Compose services:
```
docker-compose up
```


This will start all the services defined in your `docker-compose.yml` file. Your application should now be running and accessible at the configured ports.


## Issue Tracking

I use GitHub issues to track the progress of the API development. You can view the [open issues](https://github.com/glauber-silva/gousers/issues?q=is%3Aopen) and [closed issues](https://github.com/glauber-silva/gousers/issues?q=is%3Aclosed) to see what has been accomplished and what work remains.
