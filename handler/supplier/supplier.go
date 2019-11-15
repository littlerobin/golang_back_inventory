package supplier

import (
	"strconv"

	"inventory-service/db/mgodb"

	"github.com/micro/go-micro/util/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type (
	// Supplier sub handler
	Supplier struct {
		DB *mgodb.MogoDB
	}
)

// New inventory
func New(db *mgodb.MogoDB) *Supplier {
	return &Supplier{DB: db}
}

// Create supplier
// This is just example how to guide you using layer database and split package
func (sp *Supplier) Create(payload interface{}) error {
	log.Logf("Handler....", payload)
	inserted, err := sp.DB.InsertOne(payload)
	if err != nil {
		log.Debug("insert supplier got error:", err)
		return err
	}
	log.Logf("inserted:", inserted)
	return nil
}

// Update supplier
func (sp *Supplier) Update(id int32, param interface{}) error {
	log.Logf("Handler....", param)
	filter := bson.M{"supplierid": id}
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

// FindOne supplier
func (sp *Supplier) FindOne(cond bson.M) *mongo.SingleResult {
	log.Log("Inventory.Read called")
	res := sp.DB.FindOne(cond)
	return res
}

// Find supplier
func (sp *Supplier) Find(cond bson.M) *mgodb.MogoDB {
	return sp.DB.Find(cond)
}

// FindAll without condition supplier
func (sp *Supplier) FindAll() *mgodb.MogoDB {
	return sp.DB.Find(bson.M{})
}

func (sp *Supplier) Delete(idstr string) error {
	id, err := strconv.ParseInt(idstr, 10,32)
	if err != nil {
		log.Debugf("Error when parsing the value: ", err)
	}
	filter := bson.M{"supplierid": id}
	res, err := sp.DB.DeleteOne(filter)
	if err != nil {
		return err
	}
	if res != nil {
		log.Debugf("Deleted a document: ", id)
	}
	return nil
}
