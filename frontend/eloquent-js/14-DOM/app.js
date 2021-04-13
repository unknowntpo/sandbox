function talksAbout(node, string) {
    if (node.nodeType == Node.ELEMENT_NODE) {
        for (child of node.childNodes) {
            if (talksAbout(child, string)) {
                return true;
            }
        }
        // what going on at here ?
        // All childNodes of node does not contain 'string', so return false
        return false;
    } else if (node.nodeType == Node.TEXT_NODE){
        console.log(node.nodeValue)
        return node.nodeValue.indexOf(string) > -1;
    }
}

console.log(talksAbout(document.body, "book"))
