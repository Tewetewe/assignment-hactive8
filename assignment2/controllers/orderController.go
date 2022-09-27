package controllers

import (
	"fmt"
	"net/http"
	"time"

	"assignment2/database"
	"assignment2/models"

	"github.com/gin-gonic/gin"
)

type Order struct {
	OrderID      int       `json:"orderId"`
	CustomerName string    `json:"customerName"`
	Item         []Item    `json:"items"`
	OrderedAt    time.Time `json:"orderedAt"`
}

type Item struct {
	ItemID      int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"orderId"`
}

var (
	db         = database.StartDB()
	OrderModel = models.Order{}
	ItemModel  = models.Item{}
)

func CreateOrder(ctx *gin.Context) {
	var OrderModel = models.Order{}
	var ItemModel = models.Item{}
	var newOrder Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement1 := `INSERT INTO orders (customer_name, ordered_at) VALUES ($1, $2) Returning *`
	sqlStatement2 := `INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4) Returning *`

	err1 := db.QueryRow(sqlStatement1, newOrder.CustomerName, newOrder.OrderedAt).
		Scan(&OrderModel.OrderID, &OrderModel.CustomerName, &OrderModel.OrderedAt)

	if err1 != nil {
		panic(err1)
	}

	err2 := db.QueryRow(sqlStatement2, newOrder.Item[0].ItemCode, newOrder.Item[0].Description, newOrder.Item[0].Quantity, OrderModel.OrderID).
		Scan(&ItemModel.ItemID, &ItemModel.ItemCode, &ItemModel.Description, &ItemModel.Quantity, &ItemModel.OrderID)

	if err2 != nil {
		panic(err2)
	}

	OrderModel.Item = append(OrderModel.Item, ItemModel)

	ctx.JSON(http.StatusCreated, gin.H{
		"order": OrderModel,
		"item":  ItemModel,
	})

}

func GetOrders(ctx *gin.Context) {
	var orders = []models.Order{}

	sqlStatement := `SELECT * FROM orders JOIN items ON orders.order_id = items.order_id`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var OrderModel = models.Order{}
		var ItemModel = models.Item{}
		err = rows.Scan(&OrderModel.OrderID, &OrderModel.CustomerName, &OrderModel.OrderedAt, &ItemModel.ItemID, &ItemModel.ItemCode, &ItemModel.Description, &ItemModel.Quantity, &ItemModel.OrderID)
		if err != nil {
			panic(err)
		}

		OrderModel.Item = append(OrderModel.Item, ItemModel)

		orders = append(orders, OrderModel)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}

func UpdateOrder(ctx *gin.Context) {

	orderId := ctx.Param("orderID")

	var updatedOrder Order

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement1 := `
	UPDATE orders
	SET customer_name = $2, ordered_at = $3
	WHERE order_id = $1;`

	sqlStatement2 := `
	UPDATE items
	SET item_code = $2, description = $3, quantity = $4
	WHERE order_id = $1;`

	res, err := db.Exec(sqlStatement1, orderId, updatedOrder.CustomerName, updatedOrder.OrderedAt)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	_, err1 := db.Exec(sqlStatement2, orderId, updatedOrder.Item[0].ItemCode, updatedOrder.Item[0].Description, updatedOrder.Item[0].Quantity)
	if err1 != nil {
		panic(err1)
	}

	if count == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("order with id %v not found", orderId),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %v has been successfully updated", orderId),
	})

}

func DeleteOrder(ctx *gin.Context) {

	orderId := ctx.Param("orderID")

	sqlStatement1 := `
	DELETE from items
	WHERE order_id = $1`

	sqlStatement2 := `
	DELETE from orders
	WHERE order_id = $1`

	_, err := db.Exec(sqlStatement1, orderId)
	if err != nil {
		panic(err)
	}

	res, err1 := db.Exec(sqlStatement2, orderId)
	if err1 != nil {
		panic(err1)
	}
	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	if count == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("order with id %v not found", orderId),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %v has been successfully deleted", orderId),
	})
}
