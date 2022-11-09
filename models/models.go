package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id             primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name     *string            `json:"first_name" validate:"required,min=3,max=30"`
	Last_Name      *string            `json:"last_name"  validate:"required,min=3,max=30"`
	Password       *string            `json:"password"   validate:"required,min=5"`
	Email          *string            `json:"email"      validate:"email,required"`
	Phone          *string            `json:"phone"      validate:"required"`
	Token          *string            `json:"token"`
	Refresh_Token  *string            `json:"refresh_token"`
	Created_At     time.Time          `json:"created_at"`
	Updated_At     time.Time          `json:"updated_at"`
	User_Id        string             `json:"user_id"`
	UsersCart      []ProductUser      `json:"users_cart" bson:"users"`
	Address_Detail []Address          `json:"address_detail" bson:"address"`
	Order_Status   []Order            `json:"order_status" bson:"orders"`
}
type Product struct {
	Product_Id   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}
type ProductUser struct {
	Product_Id   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}
type Address struct {
	Address_iD primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house" bson:"house"`
	Street     *string            `json:"street" bson:"street"`
	City       *string            `json:"city" bson:"city"`
	PinCode    *string            `json:"pin_code" bson:"pin_code"`
}
type Order struct {
	Order_Id       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_list"bson:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
