#define <stdbool.h>
#define NULL_GUARD(q)                                                          \
  do {                                                                         \
    if (!q)                                                                    \
      return 0;                                                                \
  } while (0)

int main() {
  char *p = NULL;
  NULL_GUARD(p);
  return 0;
}
