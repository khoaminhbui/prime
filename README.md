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

    - link: `http://localhost:3000`
    - To test prime number calculation, open the link and all necessary items come right on the home page:
        - A textbox to enter N.
        - A button to fire calculation command.
        - A label to display the result.

## Algorithm
This section explains the algorithms used as well as their mathematical background.

Primality test can be categorized basically into 2 types: deterministic and probabilistic. Deterministic test provides result with absolute certainty but slow while probabilistic test run faster but with a trade off that potentially (though with very small probability) faslely identify composite number as prime.

In this scope, we only consider deterministic primality test. Including `Trial Division` (the easiest one) and `Elliptic Curve Primality Proving` (one of the quickiest and widely used for general purpose test). 
- `Trial Division`: this algorithm provides a simple and effective method for small N.
    - Anatomy:
        - Loop through the numbers beginning at (N - 1) all the way to 2 (smallest prime number), and check if which one is prime.
        - For each candidate, perform a primality test by checking if it can be divided evenly by a number in range [2 - sqrt(N)]. If not, the target number is prime.
    - Time complexity: O(n * sqrt(n)).
    - Improvement: 
        - Number theorem suggests that there is always a prime number in range [n/2, n], so the loop can stop at n/2 rather than 2. (In fact, this is an observation that refects the nature of the algorithm and does not contribute to algorithm performance).
        - The candidates can be further filtered out by rules as only test number of type (2k + 1), (6k + 1, 5), (30k + 1, 7, 11, 13, 17, 19, 23, 29).
        - Instead of divide by all number in range [2, sqrt(N)], we can find a way to deal with the prime number in range only. But this is an open question, since we dont have those prime numbers at hand.

- `Elliptic Curve Primality Proving`: 
    - It's too complicated and out of scope to discuss in details and implement this algorithm.
    - We can also combine probabilistic method with deterministic method where the the probabilistic method can quickly filter out the composite number, then continue the deterministic method for the pseudoprime (number pass probabilistic test).
