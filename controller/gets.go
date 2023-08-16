package controller

import "exam/models"

func (c *Controller) GetProduct() map[string]models.Product {
	products, _ := c.Strg.Product().GetList(&models.ProductGetListRequest{Offset: 1, Limit: 1000})
	res := make(map[string]models.Product)

	for _, product := range products.Products {
		res[product.Id] = *product
	}

	return res
}

func (c *Controller) GetUsrName() map[string]string {
	users, _ := c.Strg.User().GetList(&models.UserGetListRequest{Offset: 1, Limit: 1000})
	res := make(map[string]string)
	for _, user := range users.Users {
		res[user.Id] = user.FirstName + " " + user.LastName
	}
	return res
}

func (c *Controller) GetCatName() map[string]string {
	categories, _ := c.Strg.Category().GetList(&models.CategoryGetListRequest{Offset: 1, Limit: 1000})
	res := make(map[string]string)
	for _, cat := range categories.Categorys {
		res[cat.Id] = cat.Name
	}
	return res
}
