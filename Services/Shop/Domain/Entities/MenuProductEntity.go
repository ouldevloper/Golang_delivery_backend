package shop_domain_entities

type MenuProductEntity interface {
	GetId() int
	GetMenuId() int
	GetProductId() int
}