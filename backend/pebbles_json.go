package main

import (
    "encoding/json"
    "io/ioutil"
    "log"

    "prgc/model"
)


func LoadJSONPebbles() []model.Pebble {
    var pebbles []model.Pebble

    content, err := ioutil.ReadFile("./Pebbles.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }
        // Now let's unmarshall the data into `payload`
    var payload map[string][]map[string]interface{}
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

    for i, pebble := range payload["Pebbles"] {
        var new_pebble model.Pebble

        new_pebble.ID = int(i) 
        new_pebble.Title = pebble["title"].(string)
        new_pebble.Description = pebble["description"].(string)
        new_pebble.Price = pebble["price"].(float64)
        new_pebble.Breed = pebble["breed"].(string)
        new_pebble.Quantity = int(pebble["quantity"].(float64))
        new_pebble.Weight = pebble["weight"].(float64)
        new_pebble.Categories = []string{}

        for _, cat := range pebble["categorie"].([]interface{}) {
            new_pebble.Categories = append(new_pebble.Categories, cat.(string))
        }

        pebbles = append(pebbles, new_pebble)
    }
 
    return pebbles
}
