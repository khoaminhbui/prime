# prime
This is a monorepo of prime project

- [What's inside](#whats-inside)
- [Start services](#Start-services)
- [Algorithm](#Algorithm)


## What's inside
The repository contains all source code, and document of iCom project.

- `dockers`: containes docker-compose file to run the project locally.
- `prime-service`: the Golang project provides calculation service for web client.
- `prime-web-client`: the ReactJS project provides web UI.

## Start services
- `docker-compose up --build`: command to build dockers & start prime service & web client.
- `docker-compose down`: command to stop & remove docker stuff.

## Algorithm
This section explains the algorithms used as well as its mathematical background.
- Trial Division: this algorithm provide a simple and effective method to find the largest prime number less than a given number N. 
    - Anatomy:
        - Loop through the number beginning at (N - 1) all the way to 2 (smallest prime number), and check if each number is prime.
        - For a candidate number, check if it can be divided evenly by a number in range [2 - square root of N]. If not, the target number is a prime one.
    - Time complexity: UPDATING.

- More effective algorithm: UPDATING.
