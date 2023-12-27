#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

typedef struct {
    int activeConnections;
    char* host;
    int port;
} server;

typedef int (*TCPHandlerFunction)(int conn, const char* message);

server* NewServer(const char* host, int port) {
    server* s = (server*)malloc(sizeof(server));
    if (s == NULL) {
        perror("Error in malloc");
        exit(EXIT_FAILURE);
    }

    s->activeConnections = 0;
    s->host = strdup(host);
    s->port = port;

    return s;
}

void Listen(server* s, TCPHandlerFunction handlerFn) {
    int serverSocket = socket(AF_INET, SOCK_STREAM, 0);
    if (serverSocket == -1) {
        perror("Error in socket");
        exit(EXIT_FAILURE);
    }

    struct sockaddr_in serverAddr;
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_addr.s_addr = inet_addr(s->host);
    serverAddr.sin_port = htons(s->port);

    if (bind(serverSocket, (struct sockaddr*)&serverAddr, sizeof(serverAddr)) == -1) {
        perror("Error in bind");
        exit(EXIT_FAILURE);
    }

    if (listen(serverSocket, 10) == -1) {
        perror("Error in listen");
        exit(EXIT_FAILURE);
    }

    printf("Start listening on [%s:%d]\n", s->host, s->port);

    while (1) {
        int clientSocket = accept(serverSocket, NULL, NULL);
        if (clientSocket == -1) {
            perror("Error in accept");
            exit(EXIT_FAILURE);
        }

        // Assuming DefaultHandler returns 0 on success
        if (fork() == 0) {
            close(serverSocket);
            handlerFn(clientSocket);
            close(clientSocket);
            exit(EXIT_SUCCESS);
        } else {
            close(clientSocket);
        }
    }

    close(serverSocket);
}

void DefaultHandler(int clientSocket) {
    printf("Serving client\n");
    // Your DefaultHandler logic goes here
}

void HandleMessage(int clientSocket, const char* msg, TCPHandlerFunction handlerFn) {
    // Your HandleMessage logic goes here
}

void ResponseString(int clientSocket, const char* str) {
    ResponseByte(clientSocket, (const unsigned char*)str, strlen(str));
}

void ResponseInt(int clientSocket, int i) {
    char str[20];
    sprintf(str, "%d", i);
    ResponseString(clientSocket, str);
}

void ResponseByte(int clientSocket, const unsigned char* b, size_t len) {
    write(clientSocket, b, len);
}