var user = {
    firstName: "Joe",
    lastName: "Nash",
    email: "JOE@EXAMPLE.COM"
}

function update(user, key, f) {
    field = user[key]
    user[key] = f(field)
    return user
}

const lowercase = function (field) {
    return field.toLowerCase()
}
user = update(user, 'email', lowercase)
console.log(user)
