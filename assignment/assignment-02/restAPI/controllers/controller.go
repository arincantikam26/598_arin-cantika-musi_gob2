package controllers

import (
	"net/http"
	"restAPI/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func (db *InDB) CreateOrders(c *gin.Context) {
	var (
		order  models.Order
		item   models.Item
		result gin.H
	)

	customer_name := c.PostForm("customer_name")
	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	qty := c.PostForm("quantity")

	quantity, _ := strconv.Atoi(qty)

	order.OrderedAt = time.Now()
	order.CustomerName = customer_name
	item.ItemCode = item_code
	item.Description = description
	item.Quantity = quantity

	var data []Data
	getData := Data{
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    quantity,
	}

	data = append(data, getData)

	err := db.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": "Create Data Failed!",
		}
		c.JSON(http.StatusNoContent, result)
	}
	item.OrderID = order.OrderID

	err = db.DB.Create(&item).Error
	if err != nil {
		result = gin.H{
			"result": "Create Data Faile!",
		}
		c.JSON(http.StatusNoContent, result)
	}
	c.JSON(http.StatusOK, gin.H{
		"order_id":      order.OrderID,
		"ordered_at":    order.OrderedAt,
		"customer_name": order.CustomerName,
		"items":         data,
	})

}

func (idb *InDB) GetOrders(c *gin.Context) {

	var (
		orders []models.Order
		result gin.H
	)

	idb.DB.Find(&orders)

	if len(orders) <= 0 {
		result = gin.H{
			"result":  nil,
			"count":   0,
			"message": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
		c.JSON(http.StatusOK, result)
	}
}

// update order by orderId
func (idb *InDB) UpdateOrder(c *gin.Context) {

	// req param :orderId
	id := c.Param("orderId")

	// struct data2
	type Data2 struct {
		LineItemId  int    `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}

	// variable init
	var (
		order    models.Order
		item     models.Item
		newOrder models.Order
		result   gin.H
	)

	// check id order
	err := idb.DB.Where("order_id = ?", id).First(&order).Error
	if err != nil {

		// response data order not found
		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		// check order_id on item
		err = idb.DB.Where("order_id = ?", id).First(&item).Error
		if err != nil {

			// response data item not found
			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)

		} else {

			// save data to new data variable
			orderid, _ := strconv.Atoi(id)
			customer_name := c.PostForm("customer_name")
			newOrder.OrderID = orderid
			newOrder.CustomerName = customer_name
			newOrder.OrderedAt = order.OrderedAt

			// array of struct
			var data []Data2
			dataStruct := Data2{
				LineItemId:  item.ItemID,
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
			}

			// add object to array
			data = append(data, dataStruct)

			// update data
			err = idb.DB.Model(&order).Updates(newOrder).Error
			if err != nil {

				// response update failed
				result = gin.H{
					"result": "Update Failed",
				}
				c.JSON(http.StatusNoContent, result)
			} else {

				// response update success
				result = gin.H{
					"result":        "successfully updated data",
					"orderId":       order.OrderID,
					"orderedAt":     order.OrderedAt,
					"customer_name": order.CustomerName,
					"data":          data,
				}
				c.JSON(http.StatusOK, result)
			}
		}
	}
}

// delete data with {id}
func (idb *InDB) DeleteOrder(c *gin.Context) {

	// init variable
	var (
		order  models.Order
		item   models.Item
		result gin.H
	)

	// param orderId
	orderId := c.Param("orderId")

	// querry check id on order
	err := idb.DB.Where("order_id = ?", orderId).First(&order).Error
	if err != nil {

		// response data order not found
		result = gin.H{
			"result": "data order not found",
		}
		c.JSON(http.StatusNotFound, result)
	} else {

		// querry check order_id on item
		err = idb.DB.Where("order_id = ?", orderId).First(&item).Error
		if err != nil {

			// response data item not found
			result = gin.H{
				"result": "data item not found",
			}
			c.JSON(http.StatusNotFound, result)
		} else {

			// Delete data on db, table item
			err = idb.DB.Delete(&item).Error
			if err != nil {

				// response delete data item not success
				result = gin.H{
					"result": "Delete item Failed",
				}
				c.JSON(http.StatusBadRequest, result)
			} else {

				// Delete data on db, table order
				err = idb.DB.Delete(&order).Error
				if err != nil {

					// response delete data order not success
					result = gin.H{
						"result": "Delete order Failed",
					}
					c.JSON(http.StatusBadRequest, result)
				} else {

					// request delete data success
					result = gin.H{
						"result": "data deleted successfully",
						"data":   order,
					}
					c.JSON(http.StatusOK, result)
				}
			}
		}
	}
}
