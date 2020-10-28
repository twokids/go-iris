package repositories

import (
	"database/sql"
	"go-iris/common"
	"go-iris/datamodels"
	"strconv"
)

type IProduct interface {
	Conn() error
	Insert(*datamodels.Product) int64
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(table string, mysqlConn *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: mysqlConn,
	}
}

func (c *ProductManager) Conn() (err error) {
	if c.mysqlConn == nil {
		mysql, err := common.NewMysqlConn()
		if err != nil {
			return err
		}
		c.mysqlConn = mysql
	}
	if c.table == "" {
		c.table = "product"
	}
	return
}

func (c *ProductManager) Insert(product *datamodels.Product) (productId int64) {
	if err := c.Conn(); err != nil {
		return
	}

	sql := "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	stmt, errSql := c.mysqlConn.Prepare(sql)
	if errSql != nil {
		return 0
	}

	result, errStmt := stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if errStmt != nil {
		return 0
	}
	productId, err = result.LastInsertId()
	return productId
}

func (c *ProductManager) Delete(productId int64) bool {
	if err := c.Conn(); err != nil {
		return false
	}

	sql := "DELETE product WHERE id=?"
	stmt, errSql := c.mysqlConn.Prepare(sql)
	if errSql != nil {
		return false
	}

	result, errStmt := stmt.Exec(productId)
	if errStmt != nil {
		return false
	}
	affectCount, _ := result.RowsAffected()
	if affectCount > 0 {
		return true
	} else {
		return true
	}
}

func (c *ProductManager) Update(product *datamodels.Product) error {
	if err := c.Conn(); err != nil {
		return err
	}

	sql := "UPDATE product SET productName=?,productNum=?,productImage=?,productUrl=? where ID=" + strconv.FormatInt(product.ID, 10)

	stmt, err := c.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl)
	if err != nil {
		return err
	}

	return nil
}

func (c *ProductManager) SelectByKey(productID int64) (productResult *datamodels.Product, err error) {
	if err = c.Conn(); err != nil {
		return nil, err
	}

	sql := "SELECT * FROM " + c.table + " WHERE ID=" + strconv.FormatInt(productID, 10)
	row, errRow := c.mysqlConn.Query(sql)
	if errRow != nil {
		return &datamodels.Product{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	common.DataToStructByTagSql(result, productResult)

	return
}

func (c *ProductManager) SelectAll() (productArray []*datamodels.Product, err error) {
	if err = c.Conn(); err != nil {
		return nil, err
	}
	sql := "SELECT * FROM " + c.table
	row, errRow := c.mysqlConn.Query(sql)
	defer row.Close()
	if errRow != nil {
		return nil, errRow
	}
	result := common.GetResultRows(row)
	if len(result) == 0 {
		return nil, nil
	}
	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		productArray = append(productArray, product)
	}
	return
}
