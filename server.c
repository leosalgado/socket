#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* nao conhecidos */
#include <unistd.h>

/* sockadd_in */
#include <arpa/inet.h>


#include <sys/socket.h>


int main() {
    int serversocket = 0;
    struct sockaddr_in serveraddr;

    /* prepare server socket */
    if ((serversocket = socket(AF_INET, SOCK_STREAM, 0)) == -1) {
        printf("Cannot create socket\n\n");
        exit(1);
    };
    /* place n zero values in serveraddr */
    bzero((void*)&serveraddr, sizeof(serveraddr));
    /* IPv4 */
    serveraddr.sin_family = AF_INET;
    /* accept any interface */
    serveraddr.sin_addr.s_addr = htonl(INADDR_ANY);
    serveraddr.sin_port = htons(5000);
    bind(serversocket, (struct sockaddr*)&serveraddr, sizeof(serveraddr));
    listen(serversocket, 10);
    
    /* accept loop */
    for (;;) {
        struct sockaddr_in clientaddr;
        int clientaddrsize = sizeof(clientaddr);
        int clientsocket = 0;
        int childpid = 0;
        char buffer[128];
        int nread = 0;
        int nwritten = 0;
        
        bzero((void*)&clientaddr, sizeof(clientaddr));
        printf("Accepting connection on port 5000\n"); 
        if ((clientsocket = accept(serversocket, (struct sockaddr*)&clientaddr, &clientaddrsize)) == -1) {
            printf("Failure accepting connection. Skipping\n");
            continue;
        }
        printf("Incomming connection...\n");

        /* fork to accept more connections */
        childpid = fork();
        if (childpid != 0) {
            /* parent keep accepting connections */
            close(clientsocket);
            continue;
        }

        /* child read data from socket */
        bzero((void*)buffer, sizeof(buffer));
        if ((nread = read(clientsocket, buffer, sizeof(buffer) - 1)) == -1) {
            printf("Failure receiving data. Closing connection");
            close(clientsocket);
            exit(1);
        }
        
        /* echo back data read from socket */
        printf("Echoing back %s\n", buffer);
        nwritten = write(clientsocket, buffer, strlen(buffer));
        if (nwritten == -1) {
            printf("Failure sending data. Closing connection");
            close(clientsocket);
            exit(1);
        }
        close(clientsocket);
        exit(0);
    }
}

/* https://labcpp.com.br/usando-socket-com-c-ansi-parte-1-servidor/ */