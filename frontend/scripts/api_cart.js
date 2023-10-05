// get pebble id
const url_params = new URLSearchParams(window.location.search);
let cart_id

const container = document.querySelector("#shopping_cart")

// apply cart
function apply_cart(cart) {
    let summ = 0

    cart.Pebbles.forEach(product => {
        summ += product.price * product.quantity
        const ensembleProduit = document.createElement('div');
        ensembleProduit.classList.add("ensemble-produit")

        
        const potoProd = document.createElement('div');
        potoProd.classList.add("photo-product")
        ensembleProduit.appendChild(potoProd)

        const img = document.createElement("img")
        img.classList.add("photo")
        img.src = "../img/img-json/" + product.image[0]
        potoProd.appendChild(img)

        // Create div for each product
        const div = document.createElement('div');
        div.classList.add('description-product');

        // Add name
        const nameEl = document.createElement('a');
        nameEl.href = "article.html?id=" + product.ID
        nameEl.classList.add('name');
        nameEl.innerText = product.title;
        div.appendChild(nameEl);

        // Add description

        // Add price
        const priceEl = document.createElement('p');
        priceEl.classList.add('price-product');
        priceEl.innerText = product.price + " € - " + product.weight + " g";

        const priceContainer = document.createElement('div');
        priceContainer.classList.add('euros');
        priceContainer.appendChild(priceEl);
        div.appendChild(priceContainer);

        const quantityEl = document.createElement("p")
        quantityEl.innerHTML = "Quantité: " + product.quantity
        div.appendChild(quantityEl);

        const addBt = document.createElement("a")
        addBt.innerHTML = "(+)"
        addBt.href = "#"
        addBt.addEventListener("click",  () => increase_item_qt(cart.ID, product.ID))
        div.appendChild(addBt);
        
        const lessBt = document.createElement("a")
        lessBt.innerHTML = "(-)<br>"
        lessBt.href = "#"
        lessBt.addEventListener("click",  () => decrease_item_qt(cart.ID, product.ID))
        div.appendChild(lessBt);

        const deleteBt = document.createElement("a")
        deleteBt.classList.add("delete")
        deleteBt.href = "#"
        deleteBt.innerHTML = "Supprimer"
        deleteBt.addEventListener("click",  () => remove_item(cart.ID, product.ID))
        div.appendChild(deleteBt);

        // Add to container
        ensembleProduit.appendChild(div);
        container.appendChild(ensembleProduit);
    })

    // TODO(nmj): remove laurie's script for total
    document.querySelector("#somme-totale").innerHTML = summ + " €"
}

async function increase_item_qt(id_user, id_item) {
    const url = `http://localhost:8080/cart/add/${id_user}/${id_item}/1`
    fetch(url).then(() => location.reload())
}

async function decrease_item_qt(id_user, id_item) {
    const url = `http://localhost:8080/cart/add/${id_user}/${id_item}/-1`
    fetch(url).then(() => location.reload())
}

async function remove_item(id_user, id_item) {
    const url = `http://localhost:8080/cart/add/${id_user}/${id_item}/-999999`
    fetch(url).then(() => location.reload())
}

// call api
async function get_cart_data(id) {
    const url = `http://localhost:8080/cart/${id}`
    fetch(url)
	    .then(response => response.json())
        .then(res => {
            console.log(res)    // debug
            apply_cart(res)
        })
}

// call api caller
url_params.forEach( async (value, key) => {
    if (key == "id") {
        cart_id = value
        await get_cart_data(value)
    }
});
