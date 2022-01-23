// REf:  [What is duck typing in JavaScript?](https://medium.com/@eamonocallaghan/what-is-duck-typing-in-javascript-f3eb10853361#:~:text=Duck%20typing%20is%20an%20informal,x%20%2C%20they%20must%20be%20ducks.)

class Duck {
    type = () => console.log('gua'.repeat(3));
}

class Cat {
    // Comment out this line to make it fail.
    type = () => console.log('gua'.repeat(3));
}

let makeDucksType = (possibleDuckImposer) => possibleDuckImposer.type();

let duck = new Duck();
let cat = new Cat();

//[duck, cat].forEach(obj => makeDucksType(obj));
makeDucksType(duck);
makeDucksType(cat);

console.log(makeDucksType, duck, cat)
