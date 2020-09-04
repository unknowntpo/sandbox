<?php

define("MAXSIZE", 3);
/* TODO: how to init an array with explicitly specify the size */
//$a[MAXSIZE] = array_fill(0, MAXSIZE, 0);
$a[3] = ["a", "b"];
var_dump($a);
/* 
 * Main control flow
 */
//main();

/* 
 * Function Definition
 */
function gen_array()
{
    $a = array(3, 2, 1);
    return $a;
}
function show_array($a)
{
    foreach ($a as $ele) {
        printf("%d%s", $ele, (!next($a) ? "" : " "));
        printf("\n");
    }
}

// get array as input
function merge_sort($a, $h, $t)  // $h: head $t: tail 
{
    if ($h == $t) return $a;
    
    $mid = intdiv($h + $t, 2);
    $a = merge_sort($a, $h, $mid); // left
    $a = merge_sort($a, $mid + 1, $t); // right
    $a = merge($a, $h, $mid, $mid + 1, $t);
    return $a;
}


function merge($a, $h1, $t1, $h2, $t2)
{
    if ($h1 == $t1) return $a;
    $temp[MAXSIZE] = array_fill(0, MAXSIZE, 0);
    $idx_temp = 0;  // index of temp[]
    $idxh = $h1;  // index of first segment
    $idxt = $t1;  // index of second segment
    
    /* compare, put smaller one into temp */
    if ($a[$idxh] <= $a[$idxt])
        $temp[$idx_temp++] = $a[$idxh++];
    else
        $temp[$idx_temp++] = $a[$idxt++];

    /* Update: copy from temp to a */
    if ($idxh > $t1)
        while($idxt <= $t2)
            $temp[$idx_temp++] = $a[$idxt++];
    if ($idxt > $t2)
        while($idxh <= $t1)
            $temp[$idx_temp++] = $a[$idxh++];
    
    for ($i = $h1, $idx_temp = 0; $i <= $t2;)
        $a[$i++] = $temp[$idx_temp++];

    return $a;
}

function main()
{
    /* Main control flow */
    $a = gen_array();
    $size = count($a);

    $a = merge_sort($a, 0, $size - 1);

    show_array($a);
}


?>
