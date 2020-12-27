# prime
This is a monorepo of prime project

- [What's inside](#whats-inside)
- [Handle services](#Handle-services)
- [Access services](#Access-services)
- [Algorithm](#Algorithm)


## What's inside
The repository contains all source code, and document of prime project.

- `dockers`: contains docker-compose file to run the project locally.
- `prime-service`: the `Golang` project provides calculation service.
- `prime-web-client`: the `ReactJS + TypeScript` project provides web UI.

## Handle services
cd to the `dockers` folder and follow these commands:

- `docker-compose up --build`: command to build dockers then start prime service & web client.
- `docker-compose down`: command to stop service & remove docker stuff.

## Access services
- `prime-service`:
This service provides calculation API:

    - host: `http://localhost:8000`
    - API to find the largest prime number less than a given natural number N: `/prime/{N}`

- `prime-web-client`:
This service provides web UI:

    - host: `http://localhost:3000`
    - A textbox to enter N.
    - A button to fire calculation command.
    - A label to display the result prime number.

## Algorithm
This section explains the algorithms used as well as their mathematical background.

- Trial Division: this algorithm provides a simple and effective method to find the largest prime number less than a given number N.
    - Anatomy:
        - Loop through the numbers beginning at (N - 1) all the way to 2 (smallest prime number), and check if which one is prime.
        - For each candidate, check if it can be divided evenly by a number in range [2 - square root of N]. If not, the target number is prime.
    - Time complexity: UPDATING.

- More effective algorithms: UPDATING.
