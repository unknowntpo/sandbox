fn main() {
    let (mut str1, str2) = two_words();
    join_words(&mut str1, str2);
    println!("concaticated string is {:?}", str1);
}

fn two_words() -> (String, String) {
    (format!("fellow"), format!("Rustaceans"))
}

// Use mutable borrow to concatinate on to prefix
fn join_words(prefix: &mut String, suffix: String) {
    prefix.push(' ');
    for ch in suffix.chars() {
        prefix.push(ch);
    }
}

/*
 fn join_words(mut prefix: String, suffix: String) -> String {
    prefix.push(' ');
    for ch in suffix.chars() {
        prefix.push(ch);
    }
    prefix
}
*/
