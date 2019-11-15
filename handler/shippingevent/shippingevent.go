package shippingevent

import (
	"strconv"

	"inventory-service/db/mgodb"

	"github.com/micro/go-micro/util/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// ShippingEvent sub handler
	ShippingEvent struct {
		DB *mgodb.MogoDB
	}
)

// New inventory
func New(db *mgodb.MogoDB) *ShippingEvent {
	return &ShippingEvent{DB: db}
}

// Create ShippingEvent
// This is just example how to guide you using layer database and split package
func (sp *ShippingEvent) Create(payload interface{}) error {
	log.Logf("Handler....", payload)
	inserted, err := sp.DB.InsertOne(payload)
	if err != nil {
		log.Debug("insert ShippingEvent got error:", err)
		return err
	}
	log.Logf("inserted:", inserted)
	return nil
}

// Update ShippingEvent
func (sp *ShippingEvent) Update(id int32, param interface{}) error {
	log.Logf("Handler....", param)
	filter := bson.M{"orderid": id}
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

// FindOne ShippingEvent
func (sp *ShippingEvent) FindOne(cond bson.M) *mongo.SingleResult {
	return sp.DB.FindOne(cond)
}

// Find ShippingEvent
func (sp *ShippingEvent) Find(cond bson.M) *mgodb.MogoDB {
	return sp.DB.Find(cond)
}
// FindAll without condition ShippingEvent
func (sp *ShippingEvent) FindAll() *mgodb.MogoDB {
	return sp.DB.Find(bson.M{})
}

func (sp *ShippingEvent) Delete(idstr string) error {
	id, err := strconv.ParseInt(idstr, 10,32)
	if err != nil {
		log.Debugf("Error when parsing the value: ", err)
	}
	filter := bson.M{"orderid": id}
	res, err := sp.DB.DeleteOne(filter)
	if err != nil {
		return err
	}
	if res != nil {
		log.Debugf("Deleted a document: ", id)
	}
	return nil
}
