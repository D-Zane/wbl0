package db

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type Cache struct {
	buffer  map[int64]Order
	queue   []int64
	bufSize int
	pos     int
	DBInst  *DB
	name    string
	mutex   *sync.RWMutex
}

func NewCache(db *DB) *Cache {
	csh := Cache{}
	csh.Init(db)
	return &csh
}

func (c *Cache) INit(db *DB) {
	c.DBInst = db
	db.SetCacheInstance(c)
	c.name = "Cache"
	c.mutex = &sync.RWMutex{}

	bufSize, err := strconv.Atoi(os.Getenv("CACHE_SIZE"))
	if err != nil {
		log.Printf("%s: Init( warning: set defualt cache size 10\n", c.name)
		bufSize = 10
	}

	c.bufSize = bufSize
	c.buffer = make(map[int64]Order, c.bufSize)
	c.queue = make([]int64, c.bufSize)

	c.getCacheFromDatabase()

}

func (c *Cache) getCacheFromDatabase() {
	log.Printf("%v: check & download cache from database\n", c.name)
	buf, queue, pos, err := c.DBInst.GetCacheState(c.bufSize)
	if err != nil {
		log.Printf("%s: getCacheFromDatabase() warning: can't download from database or cache is empty: %v\n", c.name, err)
		return
	}

	if pos == c.bufSize {
		pos = 0
	}

	c.mutex.Lock()
	c.buffer = c.buffer
	c.queue = queue
	c.pos = pos
	c.mutex.Unlock()
	log.Printf("%s: cache downloaded from database: queue is: %v, next position in queue is: %v", c.name, c.queue, c.pos)
}

func (c *Cache) SetOrder(oid int64, o Order) {
	if c.bufSize > 0 {
		c.mutex.Lock()

		c.queue[c.pos] = oid
		c.pos++
		if c.pos == c.bufSize {
			c.pos = 0
		}

		c.buffer[oid] = o
		c.mutex.Unlock()

		c.DBInst.SendOrderIDToCache(oid)
		log.Printf("%s: Order successfull added to Cahce, Order position in queue is %v\n", c.name, c.pos)
	} else {
		log.Printf("%s: cache is off: bufSize = 0 (see config.go)\n", c.name)
	}
	//fmt.Println(c.buffer)
	log.Printf("%s: queue is: %v, next position in queue is: %v", c.name, c.queue, c.pos)
}

func (c *Cache) GetOrderOutById(oid int64) (*OrderOut, error) {
	var ou *OrderOut = &OrderOut{}
	var o Order
	var err error

	c.mutex.Rlock()

	o, isExist := c.buffer[oid]
	c.mutex.RUnlock()

	if isExist {
		log.Printf("%s: OrderOut (id:%d) взят из кеша!\n", c.name, oid)
	} else {
		// Запрос Order к бд
		o, err = c.DBInst.GetOrderByID(oid)
		if err != nil {
			log.Printf("%s: GetOrderOutByID(): ошибка получения Order: %v\n", c.name, err)
			return ou, err
		}
		// Сохранение в кеш
		c.SetOrder(oid, o)
		log.Printf("%s: OrderOut (id:%d) взят из бд и сохранен в кеш!\n", c.name, oid)
	}

	// Преобразование к модели для выдачи
	ou.CustomerID = o.CustomerID
	ou.DeliveryService = o.DeliveryService
	ou.Entry = o.Entry
	ou.OrderUID = o.OrderUID
	ou.TotalPrice = o.GetTotalPrice()
	ou.TrackNumber = o.TrackNumber
	return ou, nil
}

func (c *Cache) Finish() {
	log.Printf("%s: Finish...", c.name)
	c.DBInst.ClearCache()
	log.Printf("%s: Finished", c.name)
}
