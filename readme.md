# Task

Design and implement “Word of Wisdom” tcp server.  
• TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.  
• The choice of the POW algorithm should be explained.  
• After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.  
• Docker file should be provided both for the server and for the client that solves the POW challenge.  

# POW algorithm selection

As described in `wikipedia` there are few algorithms (functions) that can be used for POW. In this project I will use `hashcash` algorithm. 

## Hashcash pros

### Easy to Verify 
Although creating a valid Hashcash stamp requires significant computational resources, verifying the stamp is relatively fast and easy. This asymmetry is critical in defending against DDoS attacks.

### Customizable Difficulty 
Hashcash allows you to adjust how much work is needed to create a valid hash by changing the number of leading zero bits required. This enables you to strike a balance between protecting the server and maintaining usability for genuine users.

### CPU-Bound Task 
Hashcash relies on a CPU-bound task, meaning that it can’t easily be accelerated by throwing more hardware at it. This makes it impractical for an attacker to mass generate requests.

### Stateless 
The server does not need to keep track of every request that has been made (as in a CAPTCHA system), reducing the memory footprint on the server side.

### Standard and Reusable 
Hashcash is a well-accepted and standard PoW algorithm which is used in various fields including spam reduction, preventing dictionary attacks, and in cryptocurrency (Bitcoin uses a variant). Its design can be used in many PoW applications with minor modifications.

# Application structure

Application is designed as a single binary with two entry points: server and client (analogy with Geth). Both server and client are implemented as separate packages.

## How to run

make, docker (with compose plugin) should be installed.

execute the following command:
```
make run
```

This will start both server and client in docker containers and will run client to solve POW challenge. After quote is received client will be stopped, but next time you can run only the client with the following command:
```
make run-client
```

# Additional information

Dockerfile used for both client and server could be found in infrastructure/local folder.

# Further improvements

- Add unit tests
- Add integration tests
- Introduce difficulty level for POW and measure number of requests per second
- Integrate with some online REST quote service
- Deploy to DO





* местами нет проверки ошибок (скорее всего случайность)


* защиты от reply атаки нет в принципе, всегда можно слать одно и то же решение

