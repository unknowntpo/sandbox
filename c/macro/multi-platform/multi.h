#ifndef MULTI_H
#define MULTI_H

#if defined(__linux__)
#error "linux"
#elif defined(__APPLE__)
#error "apple"
#elif defined(_WIN32)
#error "win32"
#else
#error "unknown platform"
#endif

#include <stdio.h>
#include <stdlib.h>

#endif
