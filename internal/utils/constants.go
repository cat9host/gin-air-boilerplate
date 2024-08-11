package utils

type Method string
type Link string

const (
	Get  Method = "GET"
	Post Method = "POST"
)

const (
	CreateOrder        Link = "/order/submit"
	RequestOrderStatus Link = "/order/query"
)
