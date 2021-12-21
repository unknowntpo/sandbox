function Person(name, age) {
    let person = {};
    person.name = 'Leo';
    person.age = 20;

    person.eat = function() {
        console.log(`${this.name} is eating`);
    }

    person.sleep = function() {
        console.log(`${this.name} is sleeping`);
    }

}

const person1 = Person('Mike', 23);
const person2 = Person('Alis', 34);

console.log(person2.sleep() === person2.sleep());