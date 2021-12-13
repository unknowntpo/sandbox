// for 
for (let i = 1; i < 10; i++) {
    console.log(i);
}

console.log("----------");

// while
{
    let i = 1;
    while (i < 10) {
        console.log(i);
        i++;
    }
}

console.log("----------");

// do while
{
    let i = 1;
    do {
        console.log(i);
        i++;
    } while (i < 10);
}

console.log("----------");

// for each

let a = [1, 2, 3, 4, 5, 6, 7, 8, 9];
a.forEach(function(v) {
    console.log(v);
})

console.log("----------");

// for each: arrow function
a.forEach((v) => console.log(v));

console.log("----------");