package main

import (
	"go-iris/common"
	"go-iris/datamodels"
)

func main() {
	data := map[string]string{
		"ID":           "1",
		"productName":  "112233",
		"productNum":   "333",
		"productImage": "444",
		"productUrl":   "555",
	}
	product:=&datamodels.Product{}
	common.DataToStructByTagSql(data,product)
}
