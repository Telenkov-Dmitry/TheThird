function invoke() {
  var rt = document.getElementById("root");
  var TheList = document.createElement("ul");
  rt.appendChild(TheList);
  var addingBut = document.createElement("BUTTON");
  var inp = document.createElement("input");
  addingBut.id = "add_task";
  addingBut.innerText = "Добавить";
  addingBut.addEventListener("click", function() {
    var newLi = document.createElement("li");
    var but = document.createElement("BUTTON");
    but.innerText = "Удалить";
    but.addEventListener("click", function() {
        TheList.removeChild(newLi);
    });
    newLi.innerHTML = "<span>" + inp.value + "</span>";
    newLi.appendChild(but);
    TheList.appendChild(newLi);
  });
  var firstLi = document.createElement("li");
  var firstbut = document.createElement("BUTTON");
  firstbut.innerText = "Удалить";
  firstbut.addEventListener("click", function() {
      TheList.removeChild(firstLi);
  });
  firstLi.innerHTML = "<span>Сделать задание #3 по web-программированию</span>";
  firstLi.appendChild(firstbut);
  TheList.appendChild(firstLi);
  inp.id = "add_task_input";
  rt.appendChild(inp);
  rt.appendChild(addingBut);
}
