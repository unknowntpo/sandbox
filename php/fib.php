<?php
/* TODO: Use callback function to specify fib method */
function fib($n)
{
    //return fib_tail($n, 0, 1);
    return fib_iter($n, 0, 1);
}
// $a == F(n)
// $b == F(n + 1)
function fib_tail($n, $a, $b)
{
    if ($n == 0) {
        return $a;
    }
    
    $result = $a + $b; // evaluate $F(n + 2)
    $a = $b;  // update F(n)
    $b = $result; // assign F(n + 2) to $b
    
    return fib_tail($n - 1, $a, $result);
}

function fib_iter($n, $a, $b)
{
    if ($n == 0) {
        return 0;
    }
    if ($n == 1) {
        return 1;
    }
    for ($i = 2; $i <= $n; $i++) {
        $result = $a + $b;  // $result == F(n)
        $a = $b;
        $b = $result;
    }
    return $result;
}
// i    n a b result
// 0    2 0 1 0
// 2    2 1 1 1
// 3    2 1 2 2
$max = 1476;
for ($n = 0; $n < $max; $n++) {
    echo "fibonacci $n = ", fib($n), "\n";
    //echo fib($n);
}
