#!/bin/bash
try() {
    expected="$1"
    input="$2"

    ./9cc "$input" > tmp.s
    gcc -o tmp tmp.s
    ./tmp
    actual="$?"

    if [ "$actual" = "$expected" ]; then
        echo "$input => $actual"
    else
        echo "$input => $expected , get $actual"
        exit 1
    fi
}

# Test case
try 0 0
try 42 42
try 1 1
try 2 2
try 3 3

echo OK
