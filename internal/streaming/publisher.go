package streaming

import (
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
	sc   *stan.Conn
	name string
}

func NewPublisher(conn *stan.Conn) *Publisher {
	return &Publisher{
		name: "Publisher",
		sc:   conn,
	}
}

func (p *Publisher) Publish() {
	// TODO: Доделать ввод
	item1 := db.Items{ChrtID: 1, Price: 10, Rid: "rid1", Name: "T-Shirt-4", Sale: 9, Size: 9, Size: "M", TotalPrice: 13, NmID: 1, Brand: "Adidas"}
	item2 := db.Items{ChrtID: 2, Price: 12, Rid: "rid2", Name: "Jeans", Sale: 11, Size: "S", TotalPrice: 14, NmID: 2, Brand: "Collins"}
	item3 := db.Items{ChrtID: 3, Price: 18, Rid: "rid3", Name: "Sneakers", Sale: 15, Size: "M", TotalPrice: 20, NmID: 1, Brand: "Nike"}
	payment := db.Payment{Transaction: "tran1", Currency: "Rub", Provider: "Provider 1", Amount: 47, PaymentDt: 2, Bank: "VTB", DeliveryCost: 7, GoodsTotal: 3}
	order := db.Order{OrderUID: "order 2", Entry: "2", InternalSignature: "IS 2", Payment: payment, Items: []db.Items(item1, item2, item3),
		Locale: "Ru", CustomerID: "2", TrackNumber: "2", DeliveryService: "DS 2", Shardkey: "SK 2", SmID: 2}
	orderData, err := json.Marshal(order)
	if err != nil {
		log.Printf("%s: json.Marshal error: %v\n", p.name, err)
	}

	ackHandler := func(ackedNuid string, err error) {
		if err != nil {
			log.Printf("%s: error publishing msg id %s: %v\n", p.name, err)
		}
	}
}
