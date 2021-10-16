
"use strict";

let nav_items = document.getElementsByClassName("nav-menu-item")

for (let i = 0; i < nav_items.length; i++) {
    nav_items[i].children[0].addEventListener("mouseover", e => {
        let ui_text = nav_items[i].children[0]
        ui_text.classList.add("ui-text-active")
    })

    nav_items[i].children[0].addEventListener("mouseout", e => {
        let ui_text = nav_items[i].children[0]
        ui_text.classList.remove("ui-text-active")
    })
}