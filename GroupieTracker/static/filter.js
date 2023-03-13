/*************************************** FILTRE COULEUR ***************************************/
/* À save avec un cookie */

$(document).ready(function(){
var button = document.getElementById("color-swap");

button.addEventListener("click", function() {

  // Vérifie la présence de la red-stylesheet
  if (!document.getElementById("red-stylesheet")) {

    // Crée la balise link de la red-stylesheet
    var link = document.createElement("link");
    link.id = "red-stylesheet";
    link.rel = "stylesheet";
    link.href = "../static/red_style.css";

    // Ajoute la balise à <head>
    document.head.appendChild(link);

    // Si la red-stylesheet existe et est sélectionnée
  } else {

    // Retire la balise <link>
    var link = document.getElementById("red-stylesheet");
    link.parentNode.removeChild(link);
  }
});
});