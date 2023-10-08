const url_params = new URLSearchParams(window.location.search);
const container = document.querySelector(".container_products")

const filter1 = document.querySelector("#filter1")
const filter2 = document.querySelector("#filter2")
const filter3 = document.querySelector("#filter3")

const sortFilter = document.querySelector("#sort")

const searchBar = document.querySelector("#searchBar")
const searchBt = document.querySelector("#submitSearch")


let CAT = []
let categories = "[]" 
let sort = "nil"
let keywords = "nil"


function create_pebbles(pebble) {
    const div1 = document.createElement("div")
    div1.classList.add("rock_product")
    div1.style.backgroundImage = `url("../img/img-json/${pebble.image[0]}")`

    div1.addEventListener("mouseenter", () => div1.style.backgroundImage = `url("../img/img-json/${pebble.image[1]}")`)
    div1.addEventListener("mouseleave", () => div1.style.backgroundImage = `url("../img/img-json/${pebble.image[0]}")`)

    div1.addEventListener("click", () => {
        window.location.href = `article.html?id=${pebble.ID}`;
    });

    container.appendChild(div1)
}


function display_filters(catList) {
    catList.forEach(element => {
        option = document.createElement("option")
        option.value = element.Title
        option.innerText = element.Title
        filter1.appendChild(option.cloneNode(true))
        filter2.appendChild(option.cloneNode(true))
        filter3.appendChild(option.cloneNode(true))
    });

    console.log(CAT)

    CAT.forEach((cat, i) => {
        if (i == 0) {c = Array.from(filter1.children)}
        if (i == 1) {c = Array.from(filter2.children)}
        if (i == 2) {c = Array.from(filter3.children)}
        c.forEach((child, i) => {
            if (child.innerText == cat) {
                child.selected = true
            }
        })
    })

    if (categories != "[]") {
    }
}

function reloadPage() {
    let url = `categories.html`;

    CAT.forEach((cat, i) => {
        if (cat == "all") {return}
        if (i == 0) {
            url += `?categories=${cat}`
        } else {
            url += `,${cat}`
        }
    })

    sortValue = sortFilter.value
    if (sortValue != "default") {
        if (!url.includes("?")) {url+="?"}
        url += `&sort=${sortValue}`
    }

    if (searchBar.value != "") {
        if (!url.includes("?")) {url+="?"}
        url += `&keywords=${searchBar.value}`
    }

    window.location.href = url
}

sortFilter.addEventListener('change', () => reloadPage())
submitSearch.addEventListener('click', () => reloadPage())


filter1.addEventListener('change', function() {
    if (CAT.length >= 1) {
        CAT[0] = this.value
    }
    else {
        CAT.push(this.value)
    }

    reloadPage()
});

filter2.addEventListener('change', function() {
    if (CAT.length >= 2) {
        CAT[1] = this.value
    }
    else {
        CAT.push(this.value)
    }

    reloadPage()
});

filter3.addEventListener('change', function() {
    if (CAT.length >= 3) {
        CAT[2] = this.value
    }
    else {
        CAT.push(this.value)
    }

    reloadPage()
});

async function call_search_api() {
    if (categories != "[]") {
        let catArr = categories.split(",")

        categories = "["
        catArr.forEach((e, i) => {
            console.log(CAT)
            CAT.push(e)
            categories += `"${e}"`
            if (i < catArr.length - 1) {
                categories += ","
            }
        })
        categories += "]"
    }
    
    fetch("http://localhost:8080/categories/")
	    .then(response => response.json())
        .then(res => {
            display_filters(res)
        })


    const url = `http://localhost:8080/pebbles/${categories}/${sort}/${keywords}`
    console.log(url)

    fetch(url)
	    .then(response => response.json())
        .then(res => {
            res.forEach(e => create_pebbles(e))
        })
}

url_params.forEach( async (value, key) => {
    if (key == "categories") {
        categories = value
        cat = value
    }
    if (key == "sort") {
        sort = value
        let c = Array.from(sortFilter.children)
        c.forEach((child, i) => {
            if (child.value == sort) {
                child.selected = true
            }
        })
        
    }
    if (key == "keywords") {
        keywords = value
    }
})
call_search_api()


