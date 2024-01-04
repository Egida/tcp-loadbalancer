# `pow.c` - Proof of Work Component

## Overview
The `pow.c` file is part of an implementation of the Client Puzzle Protocol, a cryptographic protocol designed to mitigate Denial of Service (DoS) attacks. The protocol requires a client to solve a computationally-intensive problem (puzzle) before the server allocates significant resources to it. This proof of work (PoW) approach ensures that only clients that have expended a certain amount of computational effort can make requests, thereby protecting the server against resource-exhaustion attacks.

## Client Puzzle Protocol
The Client Puzzle Protocol operates in the following manner:

1. The client requests access to a server.
2. The server responds with a puzzle that is difficult to solve but easy to verify.
3. The client solves the puzzle, which requires computational work.
4. The client sends the solution back to the server.
5. If the solution is correct, the server grants access to the client.

This protocol is effective because it imposes a cost on the client, which makes it expensive for attackers to flood the server with requests. Legitimate clients, however, will only notice a slight delay.

## `solve` Function
The `solve` function is the core of the Client Puzzle Protocol within the `pow.c` file. It is responsible for finding the solution to the puzzle provided by the server. While the `solve` function is not explicitly detailed within the provided code snippets, its typical implementation would involve the following steps:

1. Receive the puzzle from the server, usually in the form of a partially computed hash value and a difficulty level.
2. Perform a brute-force search or other heuristic to find an input value that, when hashed, matches the given puzzle's criteria (e.g., a hash with a certain number of leading zeroes).
3. Return the solution to the server for verification.

you can use solve function in `pow.c` like this:
```c
    char *hash1 = "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48b"; // removed one character ('b') from the end 
    char *hash2 = "da3811154d59c4267077ddd8bb768fa9b06399c486e1fc00485116b57c9872f5";

    char *result = solve(hash1, hash2);
    printf("%s\n", result); // should print "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"
```
Note: The actual implementation details of the `solve` function would depend on the specific puzzle and hashing algorithm used by the protocol. In this case, the `solve` function is responsible for finding the solution to the puzzle provided by the server based on `sha256` hashing algorithm.

Hash function is cloned from [this repo](https://github.com/amosnier/sha-256). It's compatible with `sha256` hashing algorithm and can be used in `low-price` embedded systems.


## Conclusion
The `pow.c` file contributes to security by implementing the proof of work component of the Client Puzzle Protocol. This protocol plays a vital role in protecting servers against DoS attacks by offloading the cost of proof to the clients.
