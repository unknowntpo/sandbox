long utilfunc(long i, long j, long k)
{
    return i;
}

long myfunc(long a, long b, long c, long d, long e, long f, long g, long h)
{
    long xx = a * b * c * d * e * f * g * h;
    long yy = a + b + c + d + e + f + g + h;
    long zz = utilfunc(xx, yy, xx % yy);
    return zz + 20;
}
