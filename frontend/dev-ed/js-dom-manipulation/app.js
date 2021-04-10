const list = document.getElementById("list")
var btn = document.getElementById("btn")

//var btn = document.querySelector("#btn");
console.log(btn)
btn.addEventListener('click', function(e){
    var inputContent = document.getElementById("input_content").value;
    var newContent = document.createTextNode(inputContent)
    var newList = document.createElement("li")
    newList.appendChild(newContent)
    console.log(newList)
    list.appendChild(newList)
});
  