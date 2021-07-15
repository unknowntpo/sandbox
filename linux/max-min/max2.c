/*
 * max2: Confirm the same data type before comparison.
 * FIXME: what if we pass in _x as arguments of this macro ?
 * e.g. max2(x, _x);
 * Ref:
 * arch/powerpc/boot/types.h](https://github.com/torvalds/linux/blob/master/arch/powerpc/boot/types.h)
 */
#define max2(x, y)           \
    ({                       \
        typeof(x) _x = (x);  \
        typeof(y) _y = (y);  \
        (void) (&_x == &_y); \
        _x > _y ? _x : _y;   \
    })

int main()
{
    int x = 1, _x = 2;
    return max2(x, _x);
}
