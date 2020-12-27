# prime
This is a monorepo of prime project

- [What's inside](#whats-inside)
- [Start services](#Start-services)
- [Algorithm](#Algorithm)


## What's inside
The repository contains all source code, and document of prime project.

- `dockers`: contains docker-compose file to run the project locally.
- `prime-service`: the Golang project provides calculation service.
- `prime-web-client`: the ReactJS project provides web UI.

## Start services
cd to the `dockers` folder and follow these comamnds:

- `docker-compose up --build`: command to build dockers then start prime service & web client.
- `docker-compose down`: command to stop service & remove docker stuff.

## Algorithm
This section explains the algorithms used as well as its mathematical background.

- Trial Division: this algorithm provide a simple and effective method to find the largest prime number less than a given number N. 
    - Anatomy:
        - Loop through the numbers beginning at (N - 1) all the way to 2 (smallest prime number), and check if which one is prime.
        - For each candidate, check if it can be divided evenly by a number in range [2 - square root of N]. If not, the target number is prime.
    - Time complexity: UPDATING.

- More effective algorithm: UPDATING.
