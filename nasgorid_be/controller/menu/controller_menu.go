package menu

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"
    
    "nasgorid_be/config"
    "nasgorid_be/models/menu"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllMenu fetches all menus from the database
func GetAllMenu(w http.ResponseWriter, r *http.Request) {
    var menus []menu.Menu
    collection := config.ConnectDB().Collection("menu")

    cursor, err := collection.Find(context.TODO(), bson.D{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var menu menu.Menu
        cursor.Decode(&menu)
        menus = append(menus, menu)
    }

    json.NewEncoder(w).Encode(menus)
}

// InsertMenu inserts a new menu item into the "menu" collection
func InsertMenu(w http.ResponseWriter, r *http.Request) {
    var newMenu menu.Menu
    err := json.NewDecoder(r.Body).Decode(&newMenu)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    newMenu.ID = primitive.NewObjectID()
    newMenu.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
    newMenu.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

    menuCollection := config.ConnectDB().Collection("menu")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = menuCollection.InsertOne(ctx, newMenu)
    if err != nil {
        log.Printf("Error inserting menu: %v", err)
        http.Error(w, "Failed to insert menu", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newMenu)
}

// UpdateMenu updates a menu item based on its ID
func UpdateMenu(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(vars["id"])
    if err != nil {
        http.Error(w, "Invalid menu ID", http.StatusBadRequest)
        return
    }

    var updatedMenu menu.Menu
    err = json.NewDecoder(r.Body).Decode(&updatedMenu)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    updatedMenu.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
    menuCollection := config.ConnectDB().Collection("menu")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"_id": id}
    update := bson.M{"$set": updatedMenu}

    result, err := menuCollection.UpdateOne(ctx, filter, update)
    if err != nil || result.MatchedCount == 0 {
        log.Printf("Error updating menu: %v", err)
        http.Error(w, "Menu not found or failed to update", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Menu berhasil diperbarui")
}

// DeleteMenu deletes a menu item based on its ID
func DeleteMenu(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(vars["id"])
    if err != nil {
        http.Error(w, "Invalid menu ID", http.StatusBadRequest)
        return
    }

    menuCollection := config.ConnectDB().Collection("menu")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err = menuCollection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        log.Printf("Error deleting menu: %v", err)
        http.Error(w, "Menu not found or failed to delete", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Menu berhasil dihapus")
}
