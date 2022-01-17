#ifndef MUNIT_H
#define MUNIT_H

#define mu_assert(name, got, want)                        \
    do {                                                  \
        printf("========\n");                             \
        printf("TEST: %s\n", (name));                     \
        if ((got) != (want)) {                            \
            printf("FAIL, want %d, got %d\n", want, got); \
        } else {                                          \
            printf("PASS\n");                             \
        }                                                 \
        printf("========\n");                             \
    } while (0)

#endif /* MUNIT_H */
