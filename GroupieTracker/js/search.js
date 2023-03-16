const cssStyle = document.getElementById("yiajstyle");
var activeBtn = "name";
let divs;

/*************************************** SMUSH API from GREENSOCKS ***************************************/
var controller = new ScrollMagic.Controller();

/* Définit l'animation de la citation */
$(document).ready(function () {
  var smushTween = new TimelineMax()
    .from('.smush-text', 1, {
      fontSize: '20px',
      opacity: 0,
      lineHeight: '20px',
      ease: Power2.easeOut
    });

  var smushScene = new ScrollMagic.Scene({
    triggerElement: '.smush-text',
    triggerHook: 'onEnter',
    duration: '100%'
  })
    .setTween(smushTween)
    .addTo(controller);
});

/*************************************** FILTRES CONCERTS ***************************************/
/* Gère le css des filtres et empêche la page de retourner au top lors du clic */
$(document).ready(function () {
  // Récupère tous les boutons filtres
  const filterBtns = document.querySelectorAll('.filterBtn');

  // Boucle dans la liste des boutons
  filterBtns.forEach(function (button) {
    // Pour chaque clic, vérifie dans la liste le bouton "actif"
    button.addEventListener('click', function (event) {
      filterBtns.forEach(function (btn) {
        btn.classList.remove('active');
      })

      document.getElementById('filterConcerts').value = '';

      button.classList.add('active');
      activeBtn = event.target.getAttribute('value');
      
      // Empêche le top
      event.preventDefault();
      event.stopPropagation();
    });
  });
});

/* Gère le footer selon le nombre de résultats de la recherche */
$(document).ready(function () {
  // Récupère les lignes de chaque concert
  divs = document.querySelectorAll('div.concertDetail');
  
  function showNoResults(count) {
    // S'il n'y a aucun résultat
    if (count === 0) {
      document.querySelector('.noResults').style.display = 'flex';
      document.querySelector('#concerts > .subtitle').innerHTML = 'Vérifie que tu te sois bien servi de tes doigts !';
      document.getElementById('endSearch').style.display = 'none';
      // S'il y a plus d'un résultat
    } else {
      document.querySelector('.noResults').style.display = 'none';
      document.querySelector('#concerts > .subtitle').innerHTML = 'On ne va pas en inventer pour toi !';
      document.getElementById('endSearch').style.display = '';
      document.getElementById('endSearch').innerHTML = '...that\'s all, folks!';
    }
  };

  /* Filtre les résultats pour chaque caractère entré/frappe de touche */
  $("#filterConcerts").keyup(function () {

    // Récupère la valeur de l'input et met le compteur à zéro
    var filter = $(this).val(), count = 0;
    // Boucle dans la liste des divs correspondant à "divToFilter"
    var divToFilter = "div#" + activeBtn + "Concert";
    $(divToFilter).each(function () {
      
      // Si la div ne contient pas l'input, le résultat disparaît
      if ($(this).text().search(new RegExp(filter, "i")) < 0) {
        $(this).parent().fadeOut();
        // Si la div contient l'input, le résultat apparaît
      } else {
        $(this).parent().show();
        count++;
        document.getElementById('endSearch').innerHTML = '...that\'s all, folks!';
      }
      // Affiche le message selon le nombre de résultats
      showNoResults(count);
    });

  });
});

/*************************************** FILTRES GROUPES INDEX ***************************************/
/* Gère le css des résultats */
$(document).ready(function () {
  $("#filter").keyup(function () {

    // Récupère la valeur de l'input et met le compteur à zéro
    var filter = $(this).val(), count = 0;

    // Boucle dans la liste des divs "itemband"
    $("div.itemband").each(function () {

      // Si la div ne contient pas l'input, le résultat disparaît
      if ($(this).text().search(new RegExp(filter, "i")) < 0) {
        $(this).fadeOut();

        // Si la div contient l'input, le résultat apparaît
      } else {
        $(this).show();
        count++;
        document.getElementById('endSearch').innerHTML = '...that\'s all, folks!';
      }
    });

    // Change l'affichage du pluriel si plus de 1 résultat
    var textItem = " BANDS";
    var textItems = " BAND";
    if (count > 1) {
      $("#filter-count").text(count + textItem);
    } else {
      $("#filter-count").text(count + textItems);
    }

    // Si le filtre est remis à zéro, retire le CSS modifié
    if (filter == "") {
      removeCSS();
      document.getElementById('endSearch').innerHTML = '';

      $("#filter-count").text('');
      $("#filter-count").css({
        "display": "none"
      });

      // Sinon, change le CSS
    } else {
      changeCSS();
      $("div.itembandResult").css({
        "display": "none"
      })
    }

  });

});

/*************************************** LIVE SEARCH CSS SWAP ***************************************/
/* Change le CSS des divs contenant le nom/lien du groupe et sa photo après la recherche (plus d'alternance des couleurs/positions) */
function changeCSS() {
  $("div.itemband").filter(function () {
    return $(this).css("display") !== "none";
  }).addClass("even");

  $("div.listband").css({
    "min-height": "100vh",
    "justify-content": "flex-start",
    "align-items": "center"
  });

  $("div.itemband.even").css({
    "flex-direction": "row-reverse",
    "background-position": "right",
    "border-color": "#222"
  });

  $("div.itemband.even").css({
    "background-color": "#111",
    "flex-direction": "row-reverse",
    "background-position": "right",
    "border-color": "#222"
  });

  $("div.itemband.even .name").css({
    "text-align": "left"
  });
}

/* Réinstaure le CSS d'origine des divs contenant le nom/lien du groupe et sa photo si l'input est vide */
function removeCSS() {
  $("div.itemband").filter(function () {
    return $(this).css("display") !== "none";
  }).removeClass("even");

  $("div.itemband").filter(function () {
    return $(this).css("display") !== "none";
  }).removeAttr("style");
  $("div.name").removeAttr("style");

  $("div.itembandResult").css({
    "display": "flex"
  })
}

/*************************************** BUTTON TO TOP ***************************************/
/* Gère un bouton qui scroll up jusqu'en haut du body */
$(document).ready(function () {
  // Show/hide the button based on scroll position
  $(window).scroll(function () {
    if ($(this).scrollTop() > 100) {
      $('.btn-top').fadeIn();
      $('.btn-top').css({
        "display": "flex",
        "justify-content": "center",
        "align-items": "center"
      });
    } else {
      $('.btn-top').fadeOut();
    }
  });

  // Scroll to top when button is clicked
  $('.btn-top').click(function () {
    $('html, body').animate({
      scrollTop: 0
    }, 800);
    return false;
  });
});