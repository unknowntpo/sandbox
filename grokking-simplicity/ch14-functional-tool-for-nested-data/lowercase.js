var user = {
    firstName: "Joe",
    lastName: "Nash",
    email: "JOE@EXAMPLE.COM"
}


let update = (user, key, f) => {
    field = user[key]
    user[key] = f(field)
    return user
}

const lowercase = (field) => field.toLowerCase()

user = update(user, 'email', lowercase)
console.log(user)
