const url_params = new URLSearchParams(window.location.search);
console.log(url_params)

let categories = "[]" 
let sort = "nil"
let keywords = "nil"

async function call_search_api() {
    let catArr = categories.split(",")
    categories = "["
    catArr.forEach((e, i) => {
        categories += `"${e}"`
        if (i < catArr.length - 1) {
            categories += ","
        }
    })
    categories += "]"
    

    const url = `http://localhost:8080/pebbles/${categories}/${sort}/${keywords}`
    console.log(url)

    fetch(url)
	    .then(response => response.json())
        .then(res => {
            console.log(res)    // debug
        })
}

url_params.forEach( async (value, key) => {
    if (key == "categories") {
        categories = value
    }
    if (key == "sort") {
        sort = value
    }
    if (key == "keywords") {
        keywords = value
    }
})
call_search_api()
