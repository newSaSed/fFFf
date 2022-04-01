function  closeMenu() {
    let opened = document.querySelectorAll("nav li.open");
    for (let k=0; k< opened.length; k++){
        opened[k].classList.remove("open");
    }
}
let menuItem = document.querySelectorAll("nav li");
for (let i=0; i < menuItem.length; i++){
    menuItem[i].onclick = function (ev){
        ev.stopPropagation();
        closeMenu();
        this.classList.add("open");
    }
}

document.body.onclick = function (ev) {
    let element = ev.target;
    if (!element) return;

    if (element.tagName !== "span") {
        closeMenu();
    }
}