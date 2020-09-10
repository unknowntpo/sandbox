<?php

define("MAXSIZE", 6);
/* TODO: how to init an array with explicitly specify the size */



/* Test */


/* 
 * Main control flow
 */
main();

/* 
 * Function Definition
 */
function gen_array()
{
    $a = array(3, 2, 1, 4, 2 , 6, -2);
    return $a;
}

/* FIXME: unexpected output when showing temp[] = [0, 0, 0, 0] (!next doesn't work) */
/* Read the array key value tutorial and try again */
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
    if ($h1 == $h2) return $a;
    $temp = array_fill(0, MAXSIZE, 0);
    $idx_temp = 0;  // index of temp[]
    $idxh = $h1;  // index of first segment
    $idxt = $h2;  // index of second segment
    
    show_array($temp);
    /* compare, put smaller one into temp */
    while ($idxh <= $t1 && $idxt <= $t2) {
        if ($a[$idxh] <= $a[$idxt])
            $temp[$idx_temp++] = $a[$idxh++];
        else
            $temp[$idx_temp++] = $a[$idxt++];
    }
    

    /* Update: copy rest of ele to temp */
    while($idxt <= $t2)
        $temp[$idx_temp++] = $a[$idxt++];
    while($idxh <= $t1)
        $temp[$idx_temp++] = $a[$idxh++];
    
    /* Copy ele from temp to a */
    for ($i = $h1, $idx_temp = 0; $i <= $t2;)
        $a[$i++] = $temp[$idx_temp++];

    return $a;
}

function main()
{
    /* Main control flow */
    $a = gen_array();
    $size = count($a);
    show_array($a);
    $a = merge_sort($a, 0, $size - 1);
    show_array($a);
}


?>
