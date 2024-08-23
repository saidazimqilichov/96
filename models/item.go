package models

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var items = make(map[int]Item)
var nextID = 1

func GetAllItems() map[int]Item {
    return items
}

func CreateItem(item Item) Item {
    item.ID = nextID
    items[nextID] = item
    nextID++
    return item
}

func GetItem(id int) (Item, bool) {
    item, found := items[id]
    return item, found
}

func UpdateItem(item Item) Item {
    items[item.ID] = item
    return item
}

func DeleteItem(id int) {
    delete(items, id)
}