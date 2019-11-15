package main
//********************************************************************************************
// By Emran A. Hamdan : 10/10/2019
// gateway.go is the JSON implementation when connect to micro-services                         *
// each service must have this type of client for testing the internal connectivity          *
// of all APIs using JSON, this is the same as client.go , the only different is that service *
// can be consumed from normal HTTP REST interface                                           *
// This is a backup and may not use in the ETW system but it's important for our connectivity *
// Testing  client-to-client is gPRC only
//********************************************************************************************


import (
	"context"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/google/uuid"

	inventory "inventory-service/proto/inventory"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/web"
	"github.com/micro/go-micro/broker"

)

// Gateway is the handelr that wrap a client for http testing
// export the interface go-micro inventory interface to the http world
type Gateway struct {
	bgw inventory.InventoryService
}

func (g *Gateway) ReadData (w http.ResponseWriter, r *http.Request ) {
	// set header for caller
	w.Header().Set("Content-Type", "application/json")
	log.Debug("http gateway read call")
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
	}

	fmt.Printf("%s", reqBody)

	//
	 type reqData struct {
		Id string
	 }
	var rd reqData
	err = json.Unmarshal(reqBody,&rd)
	type ErrRsp struct {
		id string
		msg string
	}
	// make fake respsone that can change as per gateway requirement

	rsp := ErrRsp{id:"100",msg:"Error"}
	var req = inventory.Request{Id: rd.Id}

	// prepare data for rsp
	var bd *inventory.Supplier
	bd, err = g.bgw.ReadSupplier(context.Background(), &req)

	if err != nil {
		fmt.Fprintf(w, "%q",rsp)
		return
	}
	// the easy way to send JSON
	json.NewEncoder(w).Encode(bd)

}
// AddData to inventory service
func (g *Gateway) AddData (w http.ResponseWriter, r *http.Request ) {
	// set header for caller
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", reqBody)

	var bd inventory.Supplier

	// change to our struct

	err = json.Unmarshal(reqBody,&bd)

	// fake error respone
	type ErrRsp struct {
		id string
		msg string
	}
	rsp := ErrRsp{id:"100",msg:"Error"}

	// save will call handler.save , then will call datastore.store , spreation of concept

	rspd, err:= g.bgw.SaveSupplier(context.Background(), &bd)


	if err != nil {
		fmt.Fprintf(w, "%q",rsp)

	}

	json.NewEncoder(w).Encode(rspd)

}
// UpdateData to inventory service
func (g *Gateway) UpdateData (w http.ResponseWriter, r *http.Request ) {
	// set header for caller
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", reqBody)

	var bd inventory.Supplier

	// change to our struct

	err = json.Unmarshal(reqBody,&bd)

	// fake error respone
	type ErrRsp struct {
		id string
		msg string
	}
	rsp := ErrRsp{id:"100",msg:"Error"}

	// save will call handler.save , then will call datastore.store , spreation of concept

	rspd, err:= g.bgw.UpdateSupplier(context.Background(), &bd)


	if err != nil {
		fmt.Fprintf(w, "%q",rsp)

	}

	json.NewEncoder(w).Encode(rspd)

}

// UpdateData to inventory service
func (g *Gateway) DeleteData (w http.ResponseWriter, r *http.Request ) {
	// set header for caller
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", reqBody)

	var bd inventory.Request

	// change to our struct

	err = json.Unmarshal(reqBody,&bd)

	// fake error respone
	type ErrRsp struct {
		id string
		msg string
	}
	rsp := ErrRsp{id:"100",msg:"Error"}

	// save will call handler.save , then will call datastore.store , spreation of concept

	rspd, err:= g.bgw.DeleteSupplier(context.Background(), &bd)


	if err != nil {
		fmt.Fprintf(w, "%q",rsp)

	}

	json.NewEncoder(w).Encode(rspd)

}

// Event publish to a topic
func (g *Gateway) Event (w http.ResponseWriter, r *http.Request ) {
	// set header for caller
	w.Header().Set("Content-Type", "application/json")
	// We can unmarshal to pre-define JSON object
	 reqBody, err := ioutil.ReadAll(r.Body)
     if err != nil {
        log.Fatal(err)
     }
     fmt.Printf("%s", reqBody)

	// var bd inventory.Supplier

	// // change to our struct

	// err = json.Unmarshal(reqBody,&bd)

	// fake error respone
	type RspMsg struct {
		id string
		msg string
	}
	rsp := RspMsg{id:"100",msg:"Error"}

	// save will call handler.save , then it will call datastore.store ,
	id := uuid.New().String()
	msg := &broker.Message{
		Header: map[string]string{
			"Msgid": fmt.Sprintf("%s",id ),
			"RemoteAddr": fmt.Sprintf("%s",r.RemoteAddr ),
		},
		// just send the body as received for now
		Body:  reqBody,
	}


	topic := "go.micro.srv.inventory"

	if err := broker.Publish(topic, msg); err != nil {
		fmt.Printf("[pub] failed: %v", err)
	} else {
		fmt.Println("[pub] message:", string(msg.Body))
	}

	// do we need to save pub Supplier?
	// we can do here also datastore save
	//rspd, err:= g.bgw.Save(context.Background(), &bd)


	if err != nil {
		fmt.Fprintf(w, "%q",rsp)
	} else {
		rsp := RspMsg{id:"200",msg:"Sucess "}
		fmt.Fprintf(w, "%q",rsp)
		return
	}


	//json.NewEncoder(w).Encode(rspd)
	fmt.Fprintf(w, "%s","event message published")

}



func main() {

	// New Service
	service := web.NewService(
		web.Address("localhost:8000"),
		web.Name("inventory.gateway.api"),
	)

	// create instant of the gatway
	gw := new(Gateway)
	// get a default client as used in client.go pattren
	client := inventory.NewInventoryService("go.micro.srv.inventory", microclient.DefaultClient)
	gw.bgw = client

	// start broker
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	// process normal go http Mux handler , we can use any REST framework in go like go-rest

	service.HandleFunc("/read", gw.ReadData)
	service.HandleFunc("/add", gw.AddData)
	service.HandleFunc("/event", gw.Event)



	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

