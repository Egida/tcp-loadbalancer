#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

// Define a client structure
typedef struct {
    char* host;
    int port;
    int conn;
} Client;

// Function to create a new client
Client* NewClient(char* host, int port) {
    Client* client = (Client*)malloc(sizeof(Client));
    if (client == NULL) {
        perror("Memory allocation failed");
        exit(EXIT_FAILURE);
    }
    
    client->host = host;
    client->port = port;
    client->conn = -1;

    return client;
}

// Function to connect to the server
void Connect(Client* client) {
    struct sockaddr_in serverAddr;
    memset(&serverAddr, 0, sizeof(serverAddr));
    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(client->port);

    if (inet_pton(AF_INET, client->host, &(serverAddr.sin_addr)) <= 0) {
        perror("Address conversion failed");
        exit(EXIT_FAILURE);
    }

    client->conn = socket(AF_INET, SOCK_STREAM, 0);
    if (client->conn == -1) {
        perror("Socket creation failed");
        exit(EXIT_FAILURE);
    }

    if (connect(client->conn, (struct sockaddr*)&serverAddr, sizeof(serverAddr)) == -1) {
        perror("Connection failed");
        exit(EXIT_FAILURE);
    }
}

// Function to send a message and receive a reply
void SendMessage(Client* client, const char* msg) {
    if (client->conn == -1) {
        Connect(client);
    }

    char buffer[1024];
    ssize_t bytesSent = send(client->conn, msg, strlen(msg), 0);
    if (bytesSent == -1) {
        perror("Failed to send TCP message");
        exit(EXIT_FAILURE);
    }

    ssize_t bytesRead = recv(client->conn, buffer, sizeof(buffer), 0);
    if (bytesRead == -1) {
        perror("Failed to receive TCP message");
        exit(EXIT_FAILURE);
    }
}