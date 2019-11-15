package main

//********************************************************************************************
// By Emran A. Hamdan : 30/9/2019
// Client.go is the proto buffer implementation when connect to a micro-services                    *
// each service must have this type of client for testing the internal connectivity          *
// of all APIs , later these will be used "As Is" to allow other service to act as client        *
// please notice how we defined metrics data                                                  *
//********************************************************************************************

// Defined your own name , comments here

import (
	"context"
	"fmt"
	"strconv"

	// What service we want consume today
	inventory "inventory-service/proto/inventory"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
)

var (
	topic = "go.micro.srv.inventory"
)

func main() {
	ClientTestSuppliers()
	fmt.Printf("----------------------------------")
	ClientTestPurchaseOrder()
	fmt.Printf("----------------------------------")
	ClientTestStockPool()
	fmt.Printf("----------------------------------")
	ClientTestMasterStock()
	fmt.Printf("----------------------------------")
	ClientTestDeliveryNote()
	fmt.Printf("----------------------------------")
	ClientTestShippingEvent()
}
func ClientTestSuppliers() {
	// get instant of default client which define the interface we need to consume from gPRC services
	client := inventory.NewInventoryService("go.micro.srv.inventory", microclient.DefaultClient)

	// Create request
	var strs []string
	for i := 0; i < 2; i++ {
		strs = append(strs, "abc")
	}
	supplier := &inventory.Supplier{
		SupplierId:    1001,
		NameEn   :"test",
		NameAr              : "test",
		CreateAt             :"",
		UpdateAt             :"",
		CreateBy             :"",
		UpdateBy             :"",
		ApprovedBy           :"",
		ApproveDate          :"",
		Class               :"",
		OEM                 :"",
		Exclusive           :"",
		Payments             :"",
		Days                 :7,
		Brands               :strs,
		AgreementType        :"",
		StartDate            :"",
		ExpiryDate           :"",
		Renewal              :"",
		AgreementPDF         :"",
		Contacts            : []*inventory.Supplier_Contacts{
			&inventory.Supplier_Contacts{
				Name                 :"test",
				Address              :"",
				Tel                    :"",
				Mobile                 :"",
				Email                  :"",
				Web                    :"",
				Designations  :"",
			},
		},
	}
	rspSave, errSave := client.SaveSupplier(context.Background(), supplier)
if errSave!=nil{
	log.Debug("Some error in Supplier services")
	fmt.Println("Error Save customer response : ", errSave)
}
fmt.Printf("Client Save Supplier object success : %+v\n",rspSave.Msg)

//Read request
	// point to Request protobuf struct to pass parameters
	var req = inventory.Request{Id: strconv.Itoa(int(supplier.SupplierId))}
	// point to inventory Supplier struct for returned respone
	var rsp *inventory.Supplier
	var err error
	// Test Read of Supplier handler interface
	rsp, err = client.ReadSupplier(context.Background(), &req)
	if err != nil {
		log.Debug("Some error in Supplier services")
		fmt.Println("Error Read customer response : ", err)
	}
	fmt.Printf("Client reading Supplier object : %+v\n", rsp)

//Update
supplier.NameEn="updated name"
rspSave, errSave = client.UpdateSupplier(context.Background(), supplier)
if errSave!=nil{
	log.Debug("Some error in Supplier services")
	fmt.Println("Error update customer response : ", err)
}
fmt.Printf("Client Update Supplier object success : %+v\n",rspSave.Msg)

//delete
rspSave, errSave = client.DeleteSupplier(context.Background(), &req)
if errSave!=nil{
	log.Debug("Some error in Supplier services")
	fmt.Println("Error Delete customer response : ", err)
}
fmt.Printf("Client Delete Supplier object success : %+v\n",rspSave.Msg)

}

//PurchaseOrder
func ClientTestPurchaseOrder(){
		// get instant of default client which define the interface we need to consume from gPRC services
		client := inventory.NewInventoryService("go.micro.srv.inventory", microclient.DefaultClient)
	// Create request
	purchaseOrder := &inventory.PurchaseOrder{
		Po:    201,
		CreateAt:"",
       UpdateAt :"",
	CreateBy :"",
	 UpdateBy :"",
	 Supplier :"1001",
	 BatchID :7,
	 BatchDate :"",
	 Desc :"",
	 Items : []*inventory.PurchaseOrder_Items{
		&inventory.PurchaseOrder_Items{
			ProductID                 :123,
			Desc              :"",
			Qty                    :10,
			Price                 :10,
		},
	},
	}
	rspSave, errSave := client.SavePurchaseOrder(context.Background(), purchaseOrder)
if errSave!=nil{
	log.Debug("Some error in PurchaseOrder services")
	fmt.Println("Error Save customer response : ", errSave)
}
fmt.Printf("Client Save PurchaseOrder object success : %+v\n",rspSave.Msg)

//Read request
	var req = inventory.Request{Id: strconv.Itoa(int(purchaseOrder.Po))}
	var rsp *inventory.PurchaseOrder
	var err error
	rsp, err = client.ReadPurchaseOrder(context.Background(), &req)
	if err != nil {
		log.Debug("Some error in PurchaseOrder services")
		fmt.Println("Error Read customer response : ", err)
	}
	fmt.Printf("Client reading PurchaseOrder object : %+v\n", rsp)

//Update
purchaseOrder.Desc="updated desc"
rspSave, err = client.UpdatePurchaseOrder(context.Background(), purchaseOrder)
if err!=nil{
	log.Debug("Some error in PurchaseOrder services")
	fmt.Println("Error update customer response : ", err)
}
fmt.Printf("Client Update PurchaseOrder object success : %+v\n",rspSave.Msg)

//delete
rspSave, err = client.DeletePurchaseOrder(context.Background(), &req)
if err!=nil{
	log.Debug("Some error in PurchaseOrder services")
	fmt.Println("Error Delete customer response : ", err)
}
fmt.Printf("Client Delete PurchaseOrder object success : %+v\n",rspSave.Msg)

}

//StockPool
	func ClientTestStockPool(){
				// get instant of default client which define the interface we need to consume from gPRC services
				client := inventory.NewInventoryService("go.micro.srv.inventory", microclient.DefaultClient)
				// Create request
				stockPool := &inventory.StockPool{
					 PoolId :301,
	 CreateAt :"",
	 UpdateAt :"",
	 CreateBy :"",
	 UpdateBy :"",
	 StartSerial : 1,
	 EndSerial : 10,
	 Status:"",
	 Total :10,
	 TotalSold : 1,
	 Remaining :9,
				}
				rspSave, errSave := client.SaveStockPool(context.Background(), stockPool)
			if errSave!=nil{
				log.Debug("Some error in StockPool services")
				fmt.Println("Error Save customer response : ", errSave)
			}
			fmt.Printf("Client Save StockPool object success : %+v\n",rspSave.Msg)

				//Read request
	var req = inventory.Request{Id: strconv.Itoa(int(stockPool.PoolId))}
	var rsp *inventory.StockPool
	var err error
	rsp, err = client.ReadStockPool(context.Background(), &req)
	if err != nil {
		log.Debug("Some error in StockPool services")
		fmt.Println("Error Read customer response : ", err)
	}
	fmt.Printf("Client reading StockPool object : %+v\n", rsp)

//Update
stockPool.Total=100
rspSave, err = client.UpdateStockPool(context.Background(), stockPool)
if err!=nil{
	log.Debug("Some error in StockPool services")
	fmt.Println("Error update customer response : ", err)
}
fmt.Printf("Client Update StockPool object success : %+v\n",rspSave.Msg)

//delete
rspSave, err = client.DeleteStockPool(context.Background(), &req)
if err!=nil{
	log.Debug("Some error in StockPool services")
	fmt.Println("Error Delete customer response : ", err)
}
fmt.Printf("Client Delete StockPool object success : %+v\n",rspSave.Msg)
	}

	//MasterStock
	func ClientTestMasterStock(){
			// get instant of default client which define the interface we need to consume from gPRC services
			client := inventory.NewInventoryService("go.micro.srv.inventory", microclient.DefaultClient)

			// Create request
			masterStock := &inventory.MasterStock{
				 Id :1,
	 Po :201,
	 DeliveryNote :401,
	 Total :4,
	 Productid :1,
	 SupplierId :1001,
			 Items : []*inventory.MasterStock_Items{
				&inventory.MasterStock_Items{
					Serial                 :123,
					Pin              :"143243543",
					ExpiryDate                    :"",
					Status                 :"fresh",
				},
			},
			}
			rspSave, errSave := client.SaveMasterStock(context.Background(), masterStock)
		if errSave!=nil{
			log.Debug("Some error in MasterStock services")
			fmt.Println("Error Save customer response : ", errSave)
		}
		fmt.Printf("Client Save MasterStock object success : %+v\n",rspSave.Msg)

		//Read request
	var req = inventory.Request{Id: strconv.Itoa(int(masterStock.Id))}
	var rsp *inventory.MasterStock
	var err error
	rsp, err = client.ReadMasterStock(context.Background(), &req)
	if err != nil {
		log.Debug("Some error in MasterStock services")
		fmt.Println("Error Read customer response : ", err)
	}
	fmt.Printf("Client reading MasterStock object : %+v\n", rsp)

//Update
masterStock.Total=100
rspSave, err = client.UpdateMasterStock(context.Background(), masterStock)
if err!=nil{
	log.Debug("Some error in MasterStock services")
	fmt.Println("Error update customer response : ", err)
}
fmt.Printf("Client Update MasterStock object success : %+v\n",rspSave.Msg)

//delete
rspSave, err = client.DeleteMasterStock(context.Background(), &req)
if err!=nil{
	log.Debug("Some error in MasterStock services")
	fmt.Println("Error Delete customer response : ", err)
}
fmt.Printf("Client Delete MasterStock object success : %+v\n",rspSave.Msg)
	}

	// Delivery Note test
	func ClientTestDeliveryNote(){
					// get instant of default client which define the interface we need to consume from gPRC services
					client := inventory.NewInventoryService("go.micro.srv.inventory", microclient.DefaultClient)

					// Create request
					deliveryNote := &inventory.DeliveryNote{
						  DeliveryId :401,
						  RefPO :201,
						  CreatedAt :"",
						  UpdateAt :"",
						  CreatedBy :"",
						  UpdateBy:"",
						  Note :"",
					 Items : []*inventory.DeliveryNote_Items{
						&inventory.DeliveryNote_Items{
							ProductId :1,
						    Desc :"",
						    OrderedQty :1,
						    RecviedQty :1,
						},
					},
					}
					rspSave, errSave := client.SaveDeliveryNote(context.Background(), deliveryNote)
				if errSave!=nil{
					log.Debug("Some error in DeliveryNote services")
					fmt.Println("Error Save customer response : ", errSave)
				}
				fmt.Printf("Client Save DeliveryNote object success : %+v\n",rspSave.Msg)

				//Read request
	var req = inventory.Request{Id: strconv.Itoa(int(deliveryNote.DeliveryId))}
	var rsp *inventory.DeliveryNote
	var err error
	rsp, err = client.ReadDeliveryNote(context.Background(), &req)
	if err != nil {
		log.Debug("Some error in DeliveryNote services")
		fmt.Println("Error Read customer response : ", err)
	}
	fmt.Printf("Client reading DeliveryNote object : %+v\n", rsp)

//Update
deliveryNote.Note="client test"
rspSave, err = client.UpdateDeliveryNote(context.Background(), deliveryNote)
if err!=nil{
	log.Debug("Some error in DeliveryNote services")
	fmt.Println("Error update customer response : ", err)
}
fmt.Printf("Client Update DeliveryNote object success : %+v\n",rspSave.Msg)

//delete
rspSave, err = client.DeleteDeliveryNote(context.Background(), &req)
if err!=nil{
	log.Debug("Some error in DeliveryNote services")
	fmt.Println("Error Delete customer response : ", err)
}
fmt.Printf("Client Delete DeliveryNote object success : %+v\n",rspSave.Msg)
	}


	// Shipping Event
	func ClientTestShippingEvent(){
		// get instant of default client which define the interface we need to consume from gPRC services
		client := inventory.NewInventoryService("go.micro.srv.inventory", microclient.DefaultClient)
		// Create request
		shippingEvent := &inventory.ShippingEvent{
			OrderId :501,
			  CreateAt :"",
			  ItemType :"",
			  ItemsSerial :4,
			  Events :"",
		}
		rspSave, errSave := client.SaveShippingEvent(context.Background(), shippingEvent)
	if errSave!=nil{
		log.Debug("Some error in ShippingEvent services")
		fmt.Println("Error Save customer response : ", errSave)
	}
	fmt.Printf("Client Save ShippingEvent object success : %+v\n",rspSave.Msg)


				//Read request
				var req = inventory.Request{Id: strconv.Itoa(int(shippingEvent.OrderId))}
				var rsp *inventory.ShippingEvent
				var err error
				rsp, err = client.ReadShippingEvent(context.Background(), &req)
				if err != nil {
					log.Debug("Some error in ShippingEvent services")
					fmt.Println("Error Read customer response : ", err)
				}
				fmt.Printf("Client reading ShippingEvent object : %+v\n", rsp)

			//Update
			shippingEvent.ItemType="test"
			rspSave, err = client.UpdateShippingEvent(context.Background(), shippingEvent)
			if err!=nil{
				log.Debug("Some error in ShippingEvent services")
				fmt.Println("Error update customer response : ", err)
			}
			fmt.Printf("Client Update ShippingEvent object success : %+v\n",rspSave.Msg)

			//delete
			rspSave, err = client.DeleteShippingEvent(context.Background(), &req)
			if err!=nil{
				log.Debug("Some error in ShippingEvent services")
				fmt.Println("Error Delete customer response : ", err)
			}
			fmt.Printf("Client Delete ShippingEvent object success : %+v\n",rspSave.Msg)
	}
