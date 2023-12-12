const drawer = document.getElementById("right-drawer");

document.body.addEventListener("open-right-drawer", function (evt) {
    drawer.open = true;
});

document.body.addEventListener("close-right-drawer", function (evt) {
    drawer.open = false;
});

// document.body.addEventListener('htmx:responseError', function(evt) {
//     console.log("there was an error", evt)
//     window.location.href = "https://http.cat/status/500";
// });