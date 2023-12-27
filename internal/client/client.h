// client.h

#ifndef CLIENT_H
#define CLIENT_H

typedef struct {
    char* host;
    int port;
    int conn;
} Client;

Client* NewClient(char* host, int port);
void Connect(Client* client);
void SendMessage(Client* client, const char* msg);

#endif // CLIENT_H