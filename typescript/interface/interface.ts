function getFullName(person: Person) {
    return `${person.firstName} ${person.lastName}`;
}

interface Person {
    firstName: string;
    lastName: string;
}

let person = {
    firstName: 'John',
    lastName: 'Doe'
};

console.log(getFullName(person));

let jane = {
   firstname: 'Jane',
   middleName: 'K.',
   lastName: 'Doe',
   age: 22
};

console.log(getFullName(jane));
