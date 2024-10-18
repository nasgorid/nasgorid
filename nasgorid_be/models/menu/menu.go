
package menu



import (

    "go.mongodb.org/mongo-driver/bson/primitive"

)



// Menu represents a menu item

type Menu struct {

    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

    Name        string             `bson:"name" json:"name"`

    Description string             `bson:"description" json:"description"`

    Price       float64            `bson:"price" json:"price"`

    CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`

    UpdatedAt   primitive.DateTime `bson:"updated_at" json:"updated_at"`

}
