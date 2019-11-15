package stockpool

import (
	"strconv"

	"inventory-service/db/mgodb"

	"github.com/micro/go-micro/util/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// StockPool sub handler
	StockPool struct {
		DB *mgodb.MogoDB
	}
)

// New inventory
func New(db *mgodb.MogoDB) *StockPool {
	return &StockPool{DB: db}
}

// Create StockPool
// This is just example how to guide you using layer database and split package
func (sp *StockPool) Create(payload interface{}) error {
	log.Logf("Handler....", payload)
	inserted, err := sp.DB.InsertOne(payload)
	if err != nil {
		log.Debug("insert StockPool got error:", err)
		return err
	}
	log.Logf("inserted:", inserted)
	return nil
}

// Update StockPool
func (sp *StockPool) Update(id int32, param interface{}) error {
	log.Logf("Handler....", param)
	filter := bson.M{"poolid": id}
	res, err := sp.DB.UpdateOne(filter, bson.M{"$set": param})
	if err != nil {
		log.Debug(err)
		return err
	}
	if res != nil {
		log.Debugf("Updated a document: ", id)
	}
	return nil
}

// FindOne StockPool
func (sp *StockPool) FindOne(cond bson.M) *mongo.SingleResult {
	return sp.DB.FindOne(cond)
}

// Find StockPool
func (sp *StockPool) Find(cond bson.M) *mgodb.MogoDB {
	return sp.DB.Find(cond)
}
// FindAll without condition supplier
func (sp *StockPool) FindAll() *mgodb.MogoDB {
	return sp.DB.Find(bson.M{})
}

func (sp *StockPool) Delete(idstr string) error {
	id, err := strconv.ParseInt(idstr, 10,32)
	if err != nil {
		log.Debugf("Error when parsing the value: ", err)
	}
	filter := bson.M{"poolid": id}
	res, err := sp.DB.DeleteOne(filter)
	if err != nil {
		return err
	}
	if res != nil {
		log.Debugf("Deleted a document: ", id)
	}
	return nil
}
