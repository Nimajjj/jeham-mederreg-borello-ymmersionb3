const locationIQKey = 'pk.987f2a86bff0fbbbbc3f6a5d5cee7b18';
const input = document.getElementById("myInput")

// https://api.locationiq.com/v1/autocomplete?key=pk.987f2a86bff0fbbbbc3f6a5d5cee7b18&q=30%20rue%20lieutau&limit=5&dedupe=1

input.oninput = async () => {
    const url = `https://api.locationiq.com/v1/autocomplete?key=${locationIQKey}&q=${input.value}&limit=5&dedupe=1`;
    const response = await getAutocompleteResults(url)
    console.log(response)
}


async function getAutocompleteResults(url) {
  const response = await fetch(url);
  return response.json();
}

async function getCaillou(id) {
    const url = `http://localhost:8080/pebble/${id}`
    console.log(url)
    fetch(url)
	    .then(response => response.json())
        .then(res => {
            console.log(res)
        })
}


const urlParams = new URLSearchParams(window.location.search);
urlParams.forEach( async (pebble_id, c) => {
     await getCaillou(pebble_id)
});
