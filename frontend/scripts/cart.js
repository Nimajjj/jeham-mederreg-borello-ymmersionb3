document.addEventListener("DOMContentLoaded", function () {
    return
    // deso :(

    var elements = document.querySelectorAll(".price-product");


    var sommeTotale = 0;


    elements.forEach(function (element) {
        var prix = parseFloat(element.textContent);
        sommeTotale += prix;
    });


    var sommeTotaleElement = document.getElementById("somme-totale");
    sommeTotaleElement.textContent = "Somme totale : " + sommeTotale.toFixed(2);
});
