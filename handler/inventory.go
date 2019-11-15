package handler

//***************************************************************
// By go-micro code generator .
// Update by Emran Hamdan  30/9/2019
// handler.go is the wrapper interface for protobuf generated code , its very important to remember that this file
// dose not change each time you modify the protoobuf definition and must be done manually,
// the hander called from main.go will not work if the interface signature is not valid , each time you need to
// implements all methods of the inventory interface services when protobuf changes.
//******************************************************************************************

import (
	"context"
	"strconv"

	"github.com/micro/go-micro/util/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"inventory-service/db"
	"inventory-service/db/mgodb"
	purchaseorderhandler "inventory-service/handler/purchaseorder"
	supplierhandler "inventory-service/handler/supplier"
	stockpoolhandler "inventory-service/handler/stockpool"
	deliverynotehandler "inventory-service/handler/deliverynote"
	masterstockhandler "inventory-service/handler/masterstock"
	shippingeventhandler "inventory-service/handler/shippingevent"
	inventory "inventory-service/proto/inventory"
)

// Inventory is handler services , this is example, you need to change  to domain object like
// customer , order , purchase, invenotry ..etc.
//Inventory will must implements all services that needed to be called from the micro-services
type Inventory struct {
	SupplierHandler      SupplierHandler
	PurchaseOrderHandler PurchaseOrderHandler
	StockPoolHandler StockPoolHandler
	MasterStockHandler MasterStockHandler
	DeliveryNoteHandler DeliveryNoteHandler
	ShippingEventHandler ShippingEventHandler
}

type (
	SupplierHandler interface {
		Create(interface{}) error
		FindOne(bson.M) *mongo.SingleResult
		Update(int32, interface{}) error
		Delete(string) error
		Find(bson.M) *mgodb.MogoDB
		FindAll() *mgodb.MogoDB
	}
	PurchaseOrderHandler interface {
		Create(interface{}) error
		FindOne(bson.M) *mongo.SingleResult
		Update(int32, interface{}) error
		Delete(string) error
		FindAll() *mgodb.MogoDB
		Find(bson.M) *mgodb.MogoDB
	}
	StockPoolHandler interface {
		Create(interface{}) error
		FindOne(bson.M) *mongo.SingleResult
		Update(int32, interface{}) error
		Delete(string) error
		FindAll() *mgodb.MogoDB
		Find(bson.M) *mgodb.MogoDB
	}
	MasterStockHandler interface {
		Create(interface{}) error
		FindOne(bson.M) *mongo.SingleResult
		Update(int32, interface{}) error
		Delete(string) error
		FindAll() *mgodb.MogoDB
		Find(bson.M) *mgodb.MogoDB
	}
	DeliveryNoteHandler interface {
		Create(interface{}) error
		FindOne(bson.M) *mongo.SingleResult
		Update(int32, interface{}) error
		Delete(string) error
		FindAll() *mgodb.MogoDB
		Find(bson.M) *mgodb.MogoDB
	}
	ShippingEventHandler interface {
		Create(interface{}) error
		FindOne(bson.M) *mongo.SingleResult
		Update(int32, interface{}) error
		Delete(string) error
		FindAll() *mgodb.MogoDB
		Find(bson.M) *mgodb.MogoDB
	}
)

func New() *Inventory {
	supColl := mgodb.NewCollectionDB(db.SupplierColl)
	purColl := mgodb.NewCollectionDB(db.PurchaseOrderColl)
	stkColl := mgodb.NewCollectionDB(db.StockPoolColl)
	masColl := mgodb.NewCollectionDB(db.MasterStockColl)
	shpColl := mgodb.NewCollectionDB(db.ShippingEventColl)
	DlrColl := mgodb.NewCollectionDB(db.DeliveryNoteColl)
	return &Inventory{
		SupplierHandler:      supplierhandler.New(supColl),
		PurchaseOrderHandler: purchaseorderhandler.New(purColl),
		StockPoolHandler: stockpoolhandler.New(stkColl),
		MasterStockHandler: masterstockhandler.New(masColl),
		ShippingEventHandler: shippingeventhandler.New(shpColl),
		DeliveryNoteHandler: deliverynotehandler.New(DlrColl),
	}
}

// Read is a single request handler called via client.Read or the generated client code
func (e *Inventory) ReadSupplier(ctx context.Context, req *inventory.Request, rsp *inventory.Supplier) error {
	log.Log("Received Inventory.Read request")
	id, err := strconv.ParseInt(req.Id, 10,32)
	if err != nil {
		log.Log("conversion failed", err)
	}
	var rspData *inventory.Supplier
	efind := e.SupplierHandler.FindOne(bson.M{"supplierid": id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
		return efind.Err()
	}
	*rsp = *rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) SaveSupplier(ctx context.Context, req *inventory.Supplier, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Save request", req.NameEn)
//Before saving check whether id Supplier exists already
var rspData *inventory.Supplier
efind := e.SupplierHandler.FindOne(bson.M{"supplierid":req.SupplierId})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists. Hence create a new record")
		//create a new record
	err := e.SupplierHandler.Create(*req)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}
	rsp.Status = 200
	rsp.Msg = "Data stored sucessully "
	return nil
	}
// if exists update the record
if rspData!=nil{
	err := e.SupplierHandler.Update(req.SupplierId, *req)
	if err != nil {
		log.Log("Inventory.UpdateSupplier() error", err)
		return err
	}
	rsp.Status = 201
	rsp.Msg = "Records updated succesfully"
}
	return nil
}

// Process is a handler called via client.Process or the generated client code
func (e *Inventory) ProcessSupplier(ctx context.Context, req *inventory.Filter, rsp *inventory.SupplierBigData) error {
	log.Log("Received Inventory.Process request with filter , Search ", req.Filter, req.Search)
	var rspData  []*inventory.Supplier
	filter :=bson.M{req.Filter : req.Search }
	e.SupplierHandler.Find(filter).Decode(&rspData)
	rsp.Suppliers = rspData
	return nil
}

func (e *Inventory) GetAllSupplier(ctx context.Context, req *inventory.EmptyRequest, rsp *inventory.SupplierBigData) error {
	log.Log("Received Inventory.GetAllSupplier")
	var rspData  []*inventory.Supplier
	e.SupplierHandler.FindAll().Decode(&rspData)
	rsp.Suppliers = rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) UpdateSupplier(ctx context.Context, req *inventory.Supplier, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Update request", req.SupplierId)
//Before updating check whether id Supplier exists already
var rspData *inventory.Supplier
efind := e.SupplierHandler.FindOne(bson.M{"supplierid":req.SupplierId})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists")
	    rsp.Status = 404
		rsp.Msg = "Record does not exists"
		return nil
	}
	err := e.SupplierHandler.Update(req.SupplierId, *req)
	if err != nil {
		log.Log("Inventory.UpdateSupplier() error", err)
		return err
	}

	rsp.Status = 201
	rsp.Msg = "Data updated sucessully "

	return nil
}

func (e *Inventory) DeleteSupplier(ctx context.Context, req *inventory.Request, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Delete request", req.Id)
	err := e.SupplierHandler.Delete(req.Id)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}

	rsp.Status = 204
	rsp.Msg = "Data Deleted sucessully "

	return nil
}

//Purchase order
func (e *Inventory) ReadPurchaseOrder(ctx context.Context, req *inventory.Request, rsp *inventory.PurchaseOrder) error {
	log.Log("Received Inventory.Read request")
	id, err := strconv.ParseInt(req.Id, 10,32)
	if err != nil {
		log.Log("conversion failed", err)
	}
	var rspData *inventory.PurchaseOrder
	efind := e.PurchaseOrderHandler.FindOne(bson.M{"po": id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
		return efind.Err()
	}
	*rsp = *rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) SavePurchaseOrder(ctx context.Context, req *inventory.PurchaseOrder, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Save request", req.Po)

	//Before saving check whether id PurchaseOrder exists already
var rspData *inventory.PurchaseOrder
efind := e.PurchaseOrderHandler.FindOne(bson.M{"po": req.Po})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
		//Else create a new record
	err := e.PurchaseOrderHandler.Create(*req)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}
	rsp.Status = 200
	rsp.Msg = "Data stored sucessully "
	return nil
	}
// if exists update the record
if rspData!=nil{
	err := e.PurchaseOrderHandler.Update(req.Po, *req)
	if err != nil {
		log.Log("Inventory.UpdatePurchaseOrder() error", err)
		return err
	}
	rsp.Status = 201
	rsp.Msg = "Records updated succesfully"
}
	return nil
}

// Process is a handler called via client.Process or the generated client code
func (e *Inventory) ProcessPurchaseOrder(ctx context.Context, req *inventory.Filter, rsp *inventory.PurchaseOrderBigData) error {
	log.Log("Received Inventory.Process request with filter , Search ", req.Filter, req.Search)
	//e.SupplierHandler.FindOne(bson.M{req.Filter, req.Search}).Decode(&rsp)
	return nil

}

func (e *Inventory) GetAllPurchaseOrder(ctx context.Context, req *inventory.EmptyRequest, rsp *inventory.PurchaseOrderBigData) error {
	log.Log("Received Inventory.GetAllPurchaseOrder")
	var rspData []*inventory.PurchaseOrder
	e.PurchaseOrderHandler.FindAll().Decode(&rspData)
	rsp.PurchaseOrders = rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) UpdatePurchaseOrder(ctx context.Context, req *inventory.PurchaseOrder, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Update request", req.Po)
	//Before saving check whether id PurchaseOrder exists already
var rspData *inventory.PurchaseOrder
efind := e.PurchaseOrderHandler.FindOne(bson.M{"po": req.Po})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
	rsp.Status = 404
	rsp.Msg = " Record does not exists"
	return nil
	}
	err := e.PurchaseOrderHandler.Update(req.Po, *req)
	if err != nil {
		log.Log("Inventory.UpdatePurchaseOrder() error", err)
		return err
	}

	rsp.Status = 201
	rsp.Msg = "Data updated sucessully "

	return nil
}

func (e *Inventory) DeletePurchaseOrder(ctx context.Context, req *inventory.Request, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Delete request", req.Id)
	err := e.SupplierHandler.Delete(req.Id)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}

	rsp.Status = 204
	rsp.Msg = "Data Deleted sucessully "

	return nil
}

//Stock Pool
func (e *Inventory) ReadStockPool(ctx context.Context, req *inventory.Request, rsp *inventory.StockPool) error {
	log.Log("Received Inventory.Read request")
	id, err := strconv.ParseInt(req.Id, 10,32)
	if err != nil {
		log.Log("conversion failed", err)
	}
	var rspData *inventory.StockPool
	efind := e.StockPoolHandler.FindOne(bson.M{"poolid": id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
		return efind.Err()
	}
	*rsp = *rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) SaveStockPool(ctx context.Context, req *inventory.StockPool, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Save request", req.PoolId)
	//Before saving check whether poolid exists already
var rspData *inventory.StockPool
efind := e.StockPoolHandler.FindOne(bson.M{"poolid": req.PoolId})
log.Log(efind)
if efind!=nil{
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists. Hence create a new record")
	err := e.StockPoolHandler.Create(*req)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}
	rsp.Status = 200
	rsp.Msg = "Data stored sucessully "
	return nil
	}
}
// if exists update the record
if rspData!=nil{
	err := e.StockPoolHandler.Update(req.PoolId, *req)
	if err != nil {
		log.Log("Inventory.UpdateStockPool() error", err)
		return err
	}
	rsp.Status = 201
	rsp.Msg = "Records updated succesfully"
}
	return nil

}

// Process is a handler called via client.Process or the generated client code
func (e *Inventory) ProcessStockPool(ctx context.Context, req *inventory.Filter, rsp *inventory.StockPoolBigData) error {
	log.Log("Received Inventory.Process request with filter , Search ", req.Filter, req.Search)
	//e.SupplierHandler.FindOne(bson.M{req.Filter, req.Search}).Decode(&rsp)
	return nil

}

func (e *Inventory) GetAllStockPool(ctx context.Context, req *inventory.EmptyRequest, rsp *inventory.StockPoolBigData) error {
	log.Log("Received Inventory.GetAllStockPool")
	var rspData []*inventory.StockPool
	e.StockPoolHandler.FindAll().Decode(&rspData)
	rsp.StockPools = rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) UpdateStockPool(ctx context.Context, req *inventory.StockPool, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Update request", req.PoolId)
//Before saving check whether poolid exists already
var rspData *inventory.StockPool
efind := e.StockPoolHandler.FindOne(bson.M{"poolid":req.PoolId})
log.Log(efind)
if efind!=nil{
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists. Hence create a new record")
	rsp.Status = 404
	rsp.Msg = "Record does not exists "
	return nil
	}
}
	err := e.StockPoolHandler.Update(req.PoolId, *req)

	if err != nil {
		log.Log("Inventory.UpdateStockPool() error", err)
		return err
	}

	rsp.Status = 201
	rsp.Msg = "Data updated sucessully "

	return nil
}

func (e *Inventory) DeleteStockPool(ctx context.Context, req *inventory.Request, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Delete request", req.Id)
	err := e.StockPoolHandler.Delete(req.Id)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}

	rsp.Status = 204
	rsp.Msg = "Data Deleted sucessully "

	return nil
}
//Master Stock
func (e *Inventory) ReadMasterStock(ctx context.Context, req *inventory.Request, rsp *inventory.MasterStock) error {
	log.Log("Received Inventory.Read request")
	id, err := strconv.ParseInt(req.Id, 10,32)
	if err != nil {
		log.Log("conversion failed", err)
	}
	var rspData *inventory.MasterStock
	efind := e.MasterStockHandler.FindOne(bson.M{"id": id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
		return efind.Err()
	}
	*rsp = *rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) SaveMasterStock(ctx context.Context, req *inventory.MasterStock, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Save request", req.Id)
	//Before saving check whether id Master exists already
var rspData *inventory.MasterStock
efind := e.MasterStockHandler.FindOne(bson.M{"id": req.Id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists. Hence create a new record")
			//else create a new record
	err := e.MasterStockHandler.Create(*req)

	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}

	rsp.Status = 200
	rsp.Msg = "Data stored sucessully "
		return nil
	}
// if exists update the record
if rspData!=nil{
	err := e.MasterStockHandler.Update(req.Id, *req)
	if err != nil {
		log.Log("Inventory.UpdateMasterStock() error", err)
		return err
	}
	rsp.Status = 201
	rsp.Msg = "Records updated succesfully"
}
	return nil

}

// Process is a handler called via client.Process or the generated client code
func (e *Inventory) ProcessMasterStock(ctx context.Context, req *inventory.Filter, rsp *inventory.MasterStockBigData) error {
	log.Log("Received Inventory.Process request with filter , Search ", req.Filter, req.Search)
	//e.SupplierHandler.FindOne(bson.M{req.Filter, req.Search}).Decode(&rsp)
	return nil

}

func (e *Inventory) GetAllMasterStock(ctx context.Context, req *inventory.EmptyRequest, rsp *inventory.MasterStockBigData) error {
	log.Log("Received Inventory.GetAllSupplier")
	var rspData []*inventory.MasterStock
	e.MasterStockHandler.FindAll().Decode(&rspData)
	rsp.MasterStocks = rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) UpdateMasterStock(ctx context.Context, req *inventory.MasterStock, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Update request", req.Id)
//Before saving check whether id Master exists already
var rspData *inventory.MasterStock
efind := e.MasterStockHandler.FindOne(bson.M{"id": req.Id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists. Hence create a new record")
	rsp.Status = 404
	rsp.Msg = "Record does not exists"
		return nil
	}
	err := e.MasterStockHandler.Update(req.Id, *req)

	if err != nil {
		log.Log("Inventory.UpdateMasterStock() error", err)
		return err
	}

	rsp.Status = 201
	rsp.Msg = "Data updated sucessully "

	return nil
}

func (e *Inventory) DeleteMasterStock(ctx context.Context, req *inventory.Request, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Delete request", req.Id)
	err := e.MasterStockHandler.Delete(req.Id)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}

	rsp.Status = 204
	rsp.Msg = "Data Deleted sucessully "

	return nil
}
//Shipping Event
func (e *Inventory) ReadShippingEvent(ctx context.Context, req *inventory.Request, rsp *inventory.ShippingEvent) error {
	log.Log("Received Inventory.Read request")
	id, err := strconv.ParseInt(req.Id, 10,32)
	if err != nil {
		log.Log("conversion failed", err)
	}
	var rspData *inventory.ShippingEvent
	efind := e.ShippingEventHandler.FindOne(bson.M{"orderid": id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
		return efind.Err()
	}
	*rsp = *rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) SaveShippingEvent(ctx context.Context, req *inventory.ShippingEvent, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Save request", req.OrderId)
		//Before saving check whether id ShippingEvent exists already
var rspData *inventory.ShippingEvent
efind := e.ShippingEventHandler.FindOne(bson.M{"orderid": req.OrderId})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists. Hence create a new record")
// create the new record
	err := e.ShippingEventHandler.Create(*req)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}
	rsp.Status = 200
	rsp.Msg = "Data stored sucessully "
		return nil
	}
// if exists update the record
if rspData!=nil{
	err := e.ShippingEventHandler.Update(req.OrderId, *req)
	if err != nil {
		log.Log("Inventory.UpdateShippingEvent() error", err)
		return err
	}
	rsp.Status = 201
	rsp.Msg = "Records updated succesfully"
	return nil
}

	return nil

}

// Process is a handler called via client.Process or the generated client code
func (e *Inventory) ProcessShippingEvent(ctx context.Context, req *inventory.Filter, rsp *inventory.ShippingEventBigData) error {
	log.Log("Received Inventory.Process request with filter , Search ", req.Filter, req.Search)
	//e.SupplierHandler.FindOne(bson.M{req.Filter, req.Search}).Decode(&rsp)
	return nil

}

func (e *Inventory) GetAllShippingEvent(ctx context.Context, req *inventory.EmptyRequest, rsp *inventory.ShippingEventBigData) error {
	log.Log("Received Inventory.GetAllSupplier")
	var rspData []*inventory.ShippingEvent
	e.ShippingEventHandler.FindAll().Decode(&rspData)
	rsp.ShippingEvents = rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) UpdateShippingEvent(ctx context.Context, req *inventory.ShippingEvent, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Update request", req.OrderId)
//Before saving check whether id ShippingEvent exists already
var rspData *inventory.ShippingEvent
efind := e.ShippingEventHandler.FindOne(bson.M{"orderid": req.OrderId})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read")
	rsp.Status = 404
	rsp.Msg = "Record does not exists"
		return nil
	}
	err := e.ShippingEventHandler.Update(req.OrderId, *req)

	if err != nil {
		log.Log("Inventory.UpdateShippingEvent() error", err)
		return err
	}

	rsp.Status = 201
	rsp.Msg = "Data updated sucessully "

	return nil
}

func (e *Inventory) DeleteShippingEvent(ctx context.Context, req *inventory.Request, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Delete request", req.Id)
	err := e.ShippingEventHandler.Delete(req.Id)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}

	rsp.Status = 204
	rsp.Msg = "Data Deleted sucessully "

	return nil
}
//Delivery Notes
func (e *Inventory) ReadDeliveryNote(ctx context.Context, req *inventory.Request, rsp *inventory.DeliveryNote) error {
	log.Log("Received Inventory.Read request")
	id, err := strconv.ParseInt(req.Id, 10,32)
	if err != nil {
		log.Log("conversion failed", err)
	}
	var rspData *inventory.DeliveryNote
	efind := e.DeliveryNoteHandler.FindOne(bson.M{"deliveryid": id})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read error")
		return efind.Err()
	}
	*rsp = *rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) SaveDeliveryNote(ctx context.Context, req *inventory.DeliveryNote, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Save request", req.DeliveryId)
		//Before saving check whether id DeliveryNote exists already
var rspData *inventory.DeliveryNote
efind := e.DeliveryNoteHandler.FindOne(bson.M{"deliveryid": req.DeliveryId})
	efind.Decode(&rspData)
	if efind.Err() != nil {
		log.Log("Inventory.Read Does not exists. Hence create a new record")
// create the new record
	err := e.DeliveryNoteHandler.Create(*req)

	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}
	rsp.Status = 200
	rsp.Msg = "Data stored sucessully "
		return nil
	}
// if exists update the record
if rspData!=nil{
	err := e.DeliveryNoteHandler.Update(req.DeliveryId, *req)
	if err != nil {
		log.Log("Inventory.UpdateDeliveryNote() error", err)
		return err
	}
	rsp.Status = 201
	rsp.Msg = "Records updated succesfully"
	return nil
}


	return nil

}

// Process is a handler called via client.Process or the generated client code
func (e *Inventory) ProcessDeliveryNote(ctx context.Context, req *inventory.Filter, rsp *inventory.DeliveryNoteBigData) error {
	log.Log("Received Inventory.Process request with filter , Search ", req.Filter, req.Search)
	//e.SupplierHandler.FindOne(bson.M{req.Filter, req.Search}).Decode(&rsp)
	return nil

}

func (e *Inventory) GetAllDeliveryNote(ctx context.Context, req *inventory.EmptyRequest, rsp *inventory.DeliveryNoteBigData) error {
	log.Log("Received Inventory.GetDeliveryNote")
	var rspData []*inventory.DeliveryNote
	e.DeliveryNoteHandler.FindAll().Decode(&rspData)
	rsp.DeliveryNotes = rspData
	return nil
}

// Save is a server side handler called via client.Save or the generated client code
func (e *Inventory) UpdateDeliveryNote(ctx context.Context, req *inventory.DeliveryNote, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Update request", req.DeliveryId)
	//Before saving check whether id DeliveryNote exists already
	var rspData *inventory.DeliveryNote
	efind := e.DeliveryNoteHandler.FindOne(bson.M{"deliveryid": req.DeliveryId})
		efind.Decode(&rspData)
		if efind.Err() != nil {
			log.Log("Inventory.Read")
		rsp.Status = 404
		rsp.Msg = "Record does not exists"
			return nil
		}
	err := e.DeliveryNoteHandler.Update(req.DeliveryId, *req)

	if err != nil {
		log.Log("Inventory.UpdateDeliveryNote() error", err)
		return err
	}
	rsp.Status = 201
	rsp.Msg = "Data updated sucessully "

	return nil
}

func (e *Inventory) DeleteDeliveryNote(ctx context.Context, req *inventory.Request, rsp *inventory.Confirm) error {
	log.Logf("Received Inventory.Delete request", req.Id)
	err := e.DeliveryNoteHandler.Delete(req.Id)
	if err != nil {
		log.Log("Inventory.Store() error")
		return err
	}

	rsp.Status = 204
	rsp.Msg = "Data Deleted sucessully "

	return nil
}