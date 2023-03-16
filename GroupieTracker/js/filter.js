/*************************************** FILTRE COULEUR ***************************************/
/* À save avec un cookie */

/* $(document).ready(function(){
var button = document.getElementById("color-swap");

button.addEventListener("click", function() {

  // Vérifie la présence de la red-stylesheet
  if (!document.getElementById("red-stylesheet")) {

    // Crée la balise link de la red-stylesheet
    var link = document.createElement("link");
    link.id = "red-stylesheet";
    link.rel = "stylesheet";
    link.type = "text/css";
    link.href = "http://localhost:8080/static/css/red_style.css";

    // Ajoute la balise à <head>
    document.head.appendChild(link);

    // Si la red-stylesheet existe et est sélectionnée
  } else {

    // Retire la balise <link>
    var link = document.getElementById("red-stylesheet");
    link.parentNode.removeChild(link);
  }
});
}); */

document.addEventListener('DOMContentLoaded', () => {
  const themeSwitcher = document.getElementById("switch");
  themeSwitcher.checked = false;
  label = document.getElementById('label-switch');

function clickHandler() {
    if (this.checked) {
        document.body.classList.remove("yellow");
        document.body.classList.add("red");
        label.style.color = "#ae8f2a"
        localStorage.setItem("theme", "red");
    } else {
        document.body.classList.add("yellow");
        document.body.classList.remove("red");
        label.style.color = "#c70000"
        localStorage.setItem("theme", "yellow");
    }


}

themeSwitcher.addEventListener("click", clickHandler);
window.onload = checkTheme();

function checkTheme() {
    const localStorageTheme = localStorage.getItem("theme");

    if (localStorageTheme !== null && localStorageTheme === "red") {
        // set the theme of body
        document.body.className = localStorageTheme;

        // adjust the slider position
        const themeSwitch = document.getElementById("switch");
        themeSwitch.checked = true;
    }
}

});