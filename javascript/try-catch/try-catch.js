try {
    lala;
} catch (error ){
    console.log("something goes wrong")
    console.log(`name: ${error.name}\n message: ${error.message}\n, ${error.stack}`)
}
