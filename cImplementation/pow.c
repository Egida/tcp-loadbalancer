#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sha256.c"
#include <stdbool.h>

#define MAX_COMMAND_LENGTH 256
#define MAX_HASH_LENGTH 64

char *alphanumerics = "0123456789abcdefghijklmnopqrstuvwxyz";

// max : returns the maximum of two numbers
size_t max(size_t a, size_t b)
{
    return a > b ? a : b;
}

// hash_to_string : converts a hash to a string
static void hash_to_string(char string[65], const uint8_t hash[32])
{
    size_t i;
    for (i = 0; i < 32; i++)
    {
        string += sprintf(string, "%02x", hash[i]);
    }
}

// appendChars : appends two strings together. str1 + str2
char *appendChars(char *str1, char *str2)
{
    char *result = malloc(sizeof(char) * (strlen(str1) + strlen(str2) + 1));
    if (result == NULL)
    {
        perror("Error allocating memory");
        exit(EXIT_FAILURE);
    }
    strcpy(result, str1);
    strcat(result, str2);
    return result;
}

// appendChar : appends a char to a string. str + ch
char *appendChar(char *str, char ch)
{
    char *result = malloc(sizeof(char) * (strlen(str) + 2));
    if (result == NULL)
    {
        perror("Error allocating memory");
        exit(EXIT_FAILURE);
    }
    strcpy(result, str);
    result[strlen(str)] = ch;
    result[strlen(str) + 1] = '\0';
    return result;
}

// calculate_sha256_hash : calculates the sha256 hash of a string
char *calculate_sha256_hash(const char *input)
{
    char command[MAX_COMMAND_LENGTH];
    char hash[MAX_HASH_LENGTH];
    FILE *fp;

    snprintf(command, MAX_COMMAND_LENGTH, "echo  \"%s\" | shasum -a 256", input);

    fp = popen(command, "r");
    if (fp == NULL)
    {
        perror("Error opening pipe");
        exit(EXIT_FAILURE);
    }

    if (fgets(hash, MAX_HASH_LENGTH + 2, fp) == NULL)
    {
        perror("Error reading hash");
        exit(EXIT_FAILURE);
    }

    if (pclose(fp) == -1)
    {
        perror("Error closing pipe");
        exit(EXIT_FAILURE);
    }

    hash[strlen(hash) - 1] = '\0';

    char *result = strdup(hash);
    if (result == NULL)
    {
        perror("Error allocating memory");
        exit(EXIT_FAILURE);
    }
    return result;
}

// isOk : checks if the hash of a string is equal to another hash
bool isOk(char *hash1, char *hash2)
{
    int lenHash1 = (int)strlen(hash1);
    uint8_t hash[32];
    char hash_string[65];

    calc_sha_256(hash, hash1, lenHash1);
    hash_to_string(hash_string, hash);
    if (strcmp(hash_string, hash2) == 0)
    {
        return true;
    }
    return false;
}

// backtrack : backtracks the hash1 string to match the hash2 string
char *backtrack(char *hash1, char *hash2)
{

    int lenHash1 = (int)strlen(hash1);
    int lenHash2 = (int)strlen(hash2);
    if (lenHash1 == lenHash2)
    {
        return hash1;
    }
    else
    {
        for (int i = 0; i < 36; i++)
        {
            char ch = alphanumerics[i];
            char *result = appendChar(hash1, ch);

            if (isOk(appendChar(hash1, ch), hash2))
            {
                return result;
            }
        }
    }
    return NULL;
}

// solve : solves the problem
char *solve(char *hash1, char *hash2)
{
    size_t lh1 = strlen(hash1);
    size_t lh2 = strlen(hash2);
    int diff = (int)lh2 - lh1;
    int max_length = (int)max(lh1, lh2);
    char *result = malloc(sizeof(char) * max_length + 1);
    if (result == NULL)
    {
        perror("Error allocating memory");
        exit(EXIT_FAILURE);
    }
    for (int i = 0; i < max_length; i++)
    {
        if (i < max_length - diff)
        {
            result[i] = hash1[i];
        }
    }
    char *path = "";
    result = backtrack(hash1, hash2);
    return result;
}

int main()
{
    // main str := "a"
    // sample hash 1 : ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb
    // sample hash 2 : da3811154d59c4267077ddd8bb768fa9b06399c486e1fc00485116b57c9872f5

    char *hash1 = "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48b"; // removed two chars from the end
    char *hash2 = "da3811154d59c4267077ddd8bb768fa9b06399c486e1fc00485116b57c9872f5";

    char *result = solve(hash1, hash2);
    printf("%s\n", result); // should print "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"
    return 0;
}