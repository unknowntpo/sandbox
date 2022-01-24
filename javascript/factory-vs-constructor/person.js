// Ref: [JavaScript Prototype and Prototype Chain explained.](https://chamikakasun.medium.com/javascript-prototype-and-prototype-chain-explained-fdc2ec17dd04)
// Ref: [JavaScript Factory functions vs Constructor functions.](https://chamikakasun.medium.com/javascript-factory-functions-vs-constructor-functions-585919818afe)

// Factory function
function person(firstName, lastName, age) {
    const person = {};
    person.firstName = firstName;
    person.lastName = lastName;
    person.age = age;

    return person;
}

// Constructor function (start by capital letter)
function Person(firstName, lastName, age) {
    // What does 'this' come from ?
    this.firstName = firstName;
    this.lastName = lastName;
    this.age = age;
}

// use console.log() in browser to show the diff between 1, 2, 3, 4 
// 1
const mike = person("Mike", "Kale", 13);
// 2
const andy = new person("Andy", "Kale", 21);
// 3
const eric = Person("Eric", "Blala", 20);
// 4
const dennis = new Person("Dennis", "Blala", 20);
