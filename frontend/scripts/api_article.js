const url_params = new URLSearchParams(window.location.search);
const title = document.querySelector("#pebble_title")
const categories = document.querySelector("#pebble_categorie")
const breed = document.querySelector("#pebble_breed")
const reference = document.querySelector("#pebble_reference")
const price = document.querySelector("#pebble_price")
const description = document.querySelector("#pebble_description")

async function get_pebble_data(id) {
    const url = `http://localhost:8080/pebble/${id}`
    fetch(url)
	    .then(response => response.json())
        .then(res => {
            console.log(res)    // debug
            apply_pebble(res)
        })
}

function apply_pebble(pebble) {
    title.innerHTML = pebble.title
    document.title = "PRGC - " + pebble.title

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
}

url_params.forEach( async (value, key) => {
    if (key == "id") {
        await get_pebble_data(value)
    }
});
