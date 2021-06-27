const A = require('arcsecond')

const stringParser = A.many(A.str('hello')).map(
    results => results.map(
        x => x.toUpperCase()
    ).join('')
)

console.log(
    stringParser.run('hellohellohellohellohellohello')
)