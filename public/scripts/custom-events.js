const drawer = document.getElementById("right-drawer");

document.body.addEventListener("open-right-drawer", function (evt) {
    drawer.open = true;
});

document.body.addEventListener("close-right-drawer", function (evt) {
    drawer.open = false;
});