const locationIQKey = 'pk.987f2a86bff0fbbbbc3f6a5d5cee7b18';
const input = document.getElementById("myInput")
const locationSelect = document.querySelector("#location-completion")


input.oninput = async () => {
    const url = `https://api.locationiq.com/v1/autocomplete?key=${locationIQKey}&q=${input.value}&limit=5&dedupe=1`;
    const response = await getAutocompleteResults(url)
    console.log(response)

    let oldOption = document.querySelectorAll("option")
    oldOption.forEach(element => element.remove())

    response.forEach(element => {
        option = document.createElement("option")
        option.value = element.place_id
        option.innerText = element.display_name
        locationSelect.appendChild(option)
    });
}

async function getAutocompleteResults(url) {
  const response = await fetch(url);
  return response.json();
}
