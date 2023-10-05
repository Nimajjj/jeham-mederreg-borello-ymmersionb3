// get pebble id
const url_params = new URLSearchParams(window.location.search);

// get elements
const title = document.querySelector("#pebble_title")
const categories = document.querySelector("#pebble_categorie")
const breed = document.querySelector("#pebble_breed")
const reference = document.querySelector("#pebble_reference")
const price = document.querySelector("#pebble_price")
const description = document.querySelector("#pebble_description")

const img1 = document.querySelector("#img1")
const imgs2 = document.querySelectorAll(".img2")
const imgs3 = document.querySelectorAll(".img3")


// call api
async function get_pebble_data(id) {
    const url = `http://localhost:8080/pebble/${id}`
    fetch(url)
	    .then(response => response.json())
        .then(res => {
            console.log(res)    // debug
            apply_pebble(res)
        })
}

// apply style
function apply_pebble(pebble) {
    document.title = "PRGC - " + pebble.title
    title.innerHTML = pebble.title

    let categoriesStr = ""
    pebble.categorie.forEach((element, i) => {
        if (i != 0) {
            categoriesStr += " - "
        }
        categoriesStr += element.toUpperCase()
    });
    categories.innerHTML = "Cat. " + categoriesStr

    breed.innerHTML = pebble.breed
    reference.innerHTML = "Art." + pebble.ID
    price.innerHTML = pebble.price + "€"
    description.innerHTML = pebble.description

    pebble.image.forEach((element, i) => {
        if (i == 0) {
            img1.src = "../img/img-json/" + element
        }
        if (i == 1) {
            imgs2[0].src = "../img/img-json/" + element
            imgs2[1].src = "../img/img-json/" + element
        }
        if (i == 2) {
            imgs3[0].src = "../img/img-json/" + element
            imgs3[1].src = "../img/img-json/" + element
        }
    })

    if (pebble.image.length < 3) {
        imgs3[0].remove()
        imgs3[1].remove()
    }
}

// call api caller
url_params.forEach( async (value, key) => {
    if (key == "id") {
        await get_pebble_data(value)
    }
});
