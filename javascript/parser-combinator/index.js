const updateParserState = (state, index, result) => ({
    ...state,
    index,
    result
});

const updateParserResult = (state,  result) => ({
    ...state,
    result
});


const updateParserError = (state, errMsg) => ({
    ...state,
    isError: true,
    error: errMsg
});

const str = s => parserState => {
    const {
        targetString,
        index,
        isError
    } = parserState;

    if (isError) {
        return parserState;
    }

    const slicedTarget = targetString.slice(index)

    if (slicedTarget.length === 0) {
        return updateParserError(parserState, `str: Tried to match "${s}", but got unexpected end of input`);
    }

    if (targetString.slice(index)
        .startsWith(s)) {
        // Success

        return updateParserState(parserState, index + s.length, s);
    }

    return updateParserError(
        parserState,
        `Tried to match ${s}, but got "${targetString.slice(index, index +10)}"`
    );
}

const sequenceOf = parsers => parserState => {
    if (parserState.isError) {
        return parserState;
    }

    const results = [];
    let nextState = parserState;

    for (let p of parsers) {
        nextState = p(nextState);
        results.push(nextState.result);
    }

    return updateParserResult(nextState, results);
}

const run = (parser, targetString) => {
    const initialState = {
        targetString,
        index: 0,
        result: null,
        isError: false,
        error: null
    };
    return parser(initialState);
}


const parser = sequenceOf([
    str('hello there!'),
    str('goodbye there!')
]);


console.log(
    run(parser, 'hello there!goodbye there!')
)
