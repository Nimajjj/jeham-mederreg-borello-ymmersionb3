// get pebble id
const url_params = new URLSearchParams(window.location.search);

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
        const nameEl = document.createElement('p');
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

        // TODO(nmj): add href to script to increase item from cart
        const addBt = document.createElement("a")
        addBt.innerHTML = "(+)"
        addBt.href = "#"
        div.appendChild(addBt);
        
        // TODO(nmj): add href to script to decrease item from cart
        const lessBt = document.createElement("a")
        lessBt.innerHTML = "(-)<br>"
        lessBt.href = "#"
        div.appendChild(lessBt);

        const deleteBt = document.createElement("a")
        deleteBt.classList.add("delete")
        deleteBt.innerHTML = "Supprimer"
        // TODO(nmj): add href to script to remove item from cart
        div.appendChild(deleteBt);

        // Add to container
        ensembleProduit.appendChild(div);
        container.appendChild(ensembleProduit);
    })

    // TODO(nmj): remove laurie's script for total
    document.querySelector("#somme-total").innerHTML = summ + " €"
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
        await get_cart_data(value)
    }
});
