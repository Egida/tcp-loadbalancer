#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sha256.c"

#define MAX_COMMAND_LENGTH 256
#define MAX_HASH_LENGTH 64

size_t max(size_t a, size_t b) {
    return a > b ? a : b;
}

char* calculate_sha256_hash(const char* input) {
    char command[MAX_COMMAND_LENGTH];
    char hash[MAX_HASH_LENGTH];
    FILE* fp;

    snprintf(command, MAX_COMMAND_LENGTH, "echo  \"%s\" | shasum -a 256", input);

    fp = popen(command, "r");
    if (fp == NULL) {
        perror("Error opening pipe");
        exit(EXIT_FAILURE);
    }

    if (fgets(hash, MAX_HASH_LENGTH+2, fp) == NULL) {
        perror("Error reading hash");
        exit(EXIT_FAILURE);
    }

    if (pclose(fp) == -1) {
        perror("Error closing pipe");
        exit(EXIT_FAILURE);
    }

    hash[strlen(hash) - 1] = '\0';

    char* result = strdup(hash);
    if (result == NULL) {
        perror("Error allocating memory");
        exit(EXIT_FAILURE);
    }
    return result;
}

char* alphanumerics = "0123456789abcdefghijklmnopqrstuvwxyz";

char* solve(char* hash1 , char* hash2){
    size_t lh1 = strlen(hash1);
    size_t lh2 = strlen(hash2);
    int diff = (int)lh2 - lh1;
    int max_length = (int)max(lh1, lh2);
    char *result = malloc(sizeof(char) * max_length + 1);
    if (result == NULL) {
        perror("Error allocating memory");
        exit(EXIT_FAILURE);
    }
    for (int i = 0; i < max_length; i++) {
        if (i < max_length-diff) {
            result[i] = hash1[i];
        } else {
           result[i] = alphanumerics[rand() % 36];
        }
    }
    printf("hash1: %s\n", hash1);
    printf("hash2: %s\n", hash2);

    printf("result: %s\n", result);
    return result;
}


static void hash_to_string(char string[65], const uint8_t hash[32])
{
	size_t i;
	for (i = 0; i < 32; i++) {
		string += sprintf(string, "%02x", hash[i]);
	}
}

// main str := "a"
// sample hash 1 : ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb
// sample hash 2 : da3811154d59c4267077ddd8bb768fa9b06399c486e1fc00485116b57c9872f5

int main() {
    // uint8_t hash[32];
    // char hash_string[65];

    // calc_sha_256(hash, "a", strlen("a"));
	// hash_to_string(hash_string, hash);
    // printf("hash: %s\n", hash_string);
    char *hash1 = "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48"; // removed two chars from the end 
    char *hash2 = "da3811154d59c4267077ddd8bb768fa9b06399c486e1fc00485116b57c9872f5";
    char *result = solve(hash1,hash2);
    return 0;
}