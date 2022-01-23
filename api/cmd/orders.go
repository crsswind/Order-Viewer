package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Order struct {
	OrderName       string    `json:"order_name"`
	CustomerCompany string    `json:"customer_company"`
	CustomerName    string    `json:"customer_name"`
	OrderDate       time.Time `json:"order_date"`
	DeliveredAmount float64   `json:"delivered_amount"`
	TotalAmount     float64   `json:"total_amount"`
}

const (
	DB_QUERY = `SELECT 
    	o.order_name, cc.company_name as customer_company, c.Name as customer_name, o.created_at as order_date, COALESCE(d.delivered_amount,0), COALESCE(oi.total_amount,0) 
	FROM orders AS o 
	JOIN customers AS c ON o.customer_id = c.user_id
	JOIN customer_companies AS cc ON c.company_id = cc.company_id
	LEFT JOIN LATERAL (
		SELECT SUM(oi.quantity * COALESCE(oi.price_per_unit,0)) AS total_amount
		FROM order_items oi
		WHERE oi.order_id=o.id AND oi.product LIKE $1
	) oi
	ON true
	LEFT JOIN LATERAL (
		SELECT SUM(d.delivered_quantity * COALESCE(oi.price_per_unit,0)) AS delivered_amount
		FROM deliveries d
		JOIN order_items oi ON d.order_item_id=oi.id
		WHERE oi.order_id=o.id AND oi.product LIKE $1
	) d 
	ON true 
	WHERE EXISTS(SELECT 1 FROM order_items oii WHERE oii.order_id=o.id AND oii.product LIKE $1) %s
	ORDER BY o.id ASC
	LIMIT $2 OFFSET $3`
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"), os.Getenv("APP_DB_USER"), os.Getenv("APP_DB_PASS"), os.Getenv("APP_DB_NAME"))
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		panic(err)
	}

	return db
}

func GetOrders(c *gin.Context) {
	searchText := c.DefaultQuery("searchText", "")
	pageIndex, err := strconv.Atoi(c.DefaultQuery("pageIndex", "0"))
	fromDateStr := c.Query("fromDate")
	toDateStr := c.Query("toDate")

	if err != nil {
		pageIndex = 0
	}

	pageSize := 5
	recordsToSkip := pageSize * pageIndex
	searchText = "%" + searchText + "%"
	timeFilter := ""

	arguments := []interface{}{searchText, pageSize, recordsToSkip}

	var fromDate time.Time
	if fromDateStr != "" {
		fromDate, err = time.Parse(time.RFC3339, fromDateStr+"T00:00:00Z")
		if err == nil {
			timeFilter = "AND created_at >= $4 "
			arguments = append(arguments, fromDate)
		}
	}

	var toDate time.Time
	if toDateStr != "" {
		toDate, err = time.Parse(time.RFC3339, toDateStr+"T00:00:00Z")
		if err == nil {
			timeFilter += "AND created_at <= $5 "
			arguments = append(arguments, toDate)
		}
	}

	db := setupDB()

	statement, err := db.Prepare(fmt.Sprintf(DB_QUERY, timeFilter))
	if err != nil {
		c.AbortWithStatus(500)
	}

	rows, err := statement.Query(arguments...)

	if err != nil {
		c.AbortWithStatus(500)
	}

	var orders []Order

	for rows.Next() {
		var order_name string
		var customer_company string
		var customer_name string
		var order_date time.Time
		var delivered_amount float64
		var total_amount float64

		rows.Scan(&order_name, &customer_company, &customer_name, &order_date, &delivered_amount, &total_amount)
		orders = append(orders, Order{OrderName: order_name, CustomerCompany: customer_company, CustomerName: customer_name, OrderDate: order_date, DeliveredAmount: delivered_amount, TotalAmount: total_amount})
	}

	if orders == nil {
		orders = make([]Order, 0)
	}

	c.JSON(http.StatusOK, orders)
}
