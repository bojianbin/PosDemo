package main

import (
	"errors"
	"fmt"
)

type AllContent struct {
	//cashier out

	//scan good
	Quantity       float32 `json:"quantity"`
	Quantity_type  int     `json:"quantity_type"`
	Item_id        string  `json:"item_id"`
	Title          string  `json:"title"`
	Sku_code       string  `json:"sku_code"`
	Unit_price     string  `json:"unit_price"`
	Total_quantity float32 `json:"total_quantity"`
	Index          int     `json:"index"`
	Member_title   string  `json:"member_title"`
	Member_no      string  `json:"member_no"`
	Member_level   string  `json:"member_level"`

	//pay
	Amount     string   `json:"amount"`
	Pay_amount []string `json:"pay_amount"`
	Time       string   `json:"time"`
	Discount   string   `json:"discount"`
	Change     string   `json:"change"`
	List_no    string   `json:"list_no"`
	Pay_type   []string `json:"pay_type"`
	Pay        string   `json:"pay"`

	//authorize
	Authorizer string `json:"authorizer"`
	Event_id   string `json:"event_id"`
	ActionID   int    `json:"actionID"`
	Action     string `json:"action"`

	//create order
	//Member_no    string `json:"member_no"`
	//Member_level string `json:"member_level"`
	Order_id string `json:"order_id"`

	//delete whole order
	Duty_code string `json:"duty_code"`
	//Order_id  string `json:"order_id"`

	//delete good
	//Quantity      int    `json:"quantity"`
	//Quantity_type int    `json:"quantity_type"`
	//Title         string `json:"title"`
	//Sku_code      string `json:"sku_code"`
	//Unit_price    string `json:"unit_price"`
	//Order_id      string `json:"order_id"`

	//postpone order
	//Duty_code string `json:"duty_code"`
	//Order_id  string `json:"order_id"`

	//continue order
	//Duty_code string `json:"duty_code"`
	//Order_id  string `json:"order_id"`

	//cashier login

	//reduce good
	//Quantity       int    `json:"quantity"`
	//Quantity_type  int    `json:"quantity_type"`
	//Title          string `json:"title"`
	//Sku_code       string `json:"sku_code"`
	//Unit_price     string `json:"unit_price"`
	//Order_id       string `json:"order_id"`
	//Total_quantity int    `json:"total_quantity"`

	//return good
	//Quantity       int    `json:"quantity"`
	//Quantity_type  int    `json:"quantity_type"`
	//Title          string `json:"title"`
	//Sku_code       string `json:"sku_code"`
	//Unit_price     string `json:"unit_price"`
	//Order_id       string `json:"order_id"`
	//Total_quantity int    `json:"total_quantity"`
}

func (r *OneMessage) print() error {
	name := []string{"", "收银员退出", "扫描商品", "支付找零", "授权", "创立订单", "整单删除", "挂单", "商品删除",
		"解挂", "收银员登陆", "商品删减", "商品退货"}
	if r.Type <= 0 || r.Type >= len(name) {
		return errors.New("[type] value error")
	}

	fmt.Println("*****************start********************************")
	fmt.Println(name[r.Type])
	fmt.Printf("cashier    \t\t\t%+v\n", r.Cashier)
	fmt.Printf("cashierID  \t\t\t%+v\n", r.CashierID)
	fmt.Printf("posCode    \t\t\t%+v\n", r.PosCode)
	fmt.Printf("shopId     \t\t\t%+v\n", r.ShopId)
	fmt.Printf("shopName   \t\t\t%+v\n", r.ShopName)
	fmt.Printf("timeStamp  \t\t\t%+v\n", r.TimeStamp)

	switch r.Type {
	case 1:
	case 2:
		fmt.Printf("quantity      \t\t\t%+v\n", r.Content.Quantity)
		fmt.Printf("quantity_type \t\t\t%+v\n", r.Content.Quantity_type)
		fmt.Printf("item_id       \t\t\t%+v\n", r.Content.Item_id)
		fmt.Printf("title         \t\t\t%+v\n", r.Content.Title)
		fmt.Printf("sku_code      \t\t\t%+v\n", r.Content.Sku_code)
		fmt.Printf("unit_price    \t\t\t%+v\n", r.Content.Unit_price)
		fmt.Printf("total_quantity\t\t\t%+v\n", r.Content.Total_quantity)
		fmt.Printf("index         \t\t\t%+v\n", r.Content.Index)
		fmt.Printf("member_level  \t\t\t%+v\n", r.Content.Member_level)
		fmt.Printf("member_no     \t\t\t%+v\n", r.Content.Member_no)
		fmt.Printf("member_title  \t\t\t%+v\n", r.Content.Member_title)
	case 3:
		fmt.Printf("amount        \t\t\t%+v\n", r.Content.Amount)
		fmt.Printf("pay_amount    \t\t\t%+v\n", r.Content.Pay_amount)
		fmt.Printf("time          \t\t\t%+v\n", r.Content.Time)
		fmt.Printf("discount      \t\t\t%+v\n", r.Content.Discount)
		fmt.Printf("change        \t\t\t%+v\n", r.Content.Change)
		fmt.Printf("list_no       \t\t\t%+v\n", r.Content.List_no)
		fmt.Printf("pay_type      \t\t\t%+v\n", r.Content.Pay_type)
		fmt.Printf("pay           \t\t\t%+v\n", r.Content.Pay)
	case 4:
		fmt.Printf("authorizer    \t\t\t%+v\n", r.Content.Authorizer)
		fmt.Printf("event_id      \t\t\t%+v\n", r.Content.Event_id)
		fmt.Printf("actionID      \t\t\t%+v\n", r.Content.ActionID)
		fmt.Printf("action        \t\t\t%+v\n", r.Content.Action)
	case 5:
		fmt.Printf("member_level  \t\t\t%+v\n", r.Content.Member_level)
		fmt.Printf("member_no     \t\t\t%+v\n", r.Content.Member_no)
		fmt.Printf("order_id      \t\t\t%+v\n", r.Content.Order_id)
	case 6:
		fmt.Printf("order_id      \t\t\t%+v\n", r.Content.Order_id)
		fmt.Printf("dduty_code    \t\t\t%+v\n", r.Content.Duty_code)
	case 7:
		//fmt.Printf("挂单")
		fmt.Printf("order_id      \t\t\t%+v\n", r.Content.Order_id)
		fmt.Printf("duty_code     \t\t\t%+v\n", r.Content.Duty_code)
	case 8:
		//fmt.Printf("商品删除")
		fmt.Printf("quantity      \t\t\t%+v\n", r.Content.Quantity)
		fmt.Printf("quantity_type \t\t\t%+v\n", r.Content.Quantity_type)
		fmt.Printf("title         \t\t\t%+v\n", r.Content.Title)
		fmt.Printf("sku_code      \t\t\t%+v\n", r.Content.Sku_code)
		fmt.Printf("unit_price    \t\t\t%+v\n", r.Content.Unit_price)
		fmt.Printf("order_id      \t\t\t%+v\n", r.Content.Order_id)
	case 9:
		//fmt.Printf("解挂")
		fmt.Printf("order_id      \t\t\t%+v\n", r.Content.Order_id)
		fmt.Printf("duty_code     \t\t\t%+v\n", r.Content.Duty_code)
	case 10:
		//fmt.Printf("收银员登陆")
	case 11:
		//fmt.Printf("商品删减")
		fallthrough
	case 12:
		//fmt.Printf("商品退货")
		fmt.Printf("quantity      \t\t\t%+v\n", r.Content.Quantity)
		fmt.Printf("quantity_type \t\t\t%+v\n", r.Content.Quantity_type)
		fmt.Printf("title         \t\t\t%+v\n", r.Content.Title)
		fmt.Printf("sku_code      \t\t\t%+v\n", r.Content.Sku_code)
		fmt.Printf("unit_price    \t\t\t%+v\n", r.Content.Unit_price)
		fmt.Printf("order_id      \t\t\t%+v\n", r.Content.Order_id)
		fmt.Printf("total_quantity\t\t\t%+v\n", r.Content.Total_quantity)
	}
	fmt.Println("******************end*********************************")

	return nil
}

type OneMessage struct {
	Type      int    `json:"type"`
	Cashier   string `json:"cashier"`
	CashierID string `json:"cashierID"`
	PosCode   string `json:"posCode"`
	ShopId    string `json:"shopId"`
	ShopName  string `json:"shopName"`
	TimeStamp string `json:"timeStamp"`

	Content AllContent `json:"content"`
}
