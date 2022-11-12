#include <stdio.h>
#include <stdlib.h> 

extern char **__environ;

int main (int argc, char **argv) {
    FILE *fp = fopen("/volumes/v1/output.txt", "w");
    for (int i = 0; i < argc; i++) {
        printf("%s\n", argv[i]);
        fprintf(fp, "%s\n", argv[i]);
    }
    printf("\n");
    fprintf(fp, "\n");


    fclose(fp);
    return 42;
}
