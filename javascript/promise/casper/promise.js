let mingRunPromise = (someone) => {
    let ran = parseInt(Math.random() * 2);
    console.log(`${someone} start running`);
    return new Promise((resolve, reject) => {
        if (ran) {
            setTimeout(function() {
                resolve(`${someone} has run for 3 secs (fulfilled)`);
            }, 3000);
        } else {
            reject(new Error(`${someone} failed to run (rejected)`));
        }
    });
}

// TODO: What is the argument of then func ?
mingRunPromise('ming').then((data) => {
    console.log(data);
}).catch((err) => {
    console.log(err);
});
