let header = document.querySelector("header");
let menu = document.querySelectorAll("a");

document.addEventListener("scroll", () => {
    let scrollPosition = window.scrollY;
    if (scrollPosition > 0) {
        header.style.backgroundColor = "#F2F2F4";
        header.style.borderBottom = "1px solid #000000";
        header.style.color = "#000000";
        menu.forEach((item) => {
            item.style.color = "#000000";
        });
        header.style.transition = "background-color 1s ease-in-out, color 1s ease-in-out";
    } else {
        header.style.backgroundColor = "transparent";
        header.style.color = "#FFFFFF";
        header.style.borderBottom = "none";
        header.style.transition = "background-color 1s ease-in-out, color 1s ease-in-out";
    }
});
