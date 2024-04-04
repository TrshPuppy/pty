#define _XOPEN_SOURCE 600
#include <stdlib.h>
#include <stdint.h>

int callGrant(uint64_t fd) {
    int grantStatus;
    grantStatus = grantpt(fd);

    return grantStatus;
}

