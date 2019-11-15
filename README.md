download protobuf for micro:

# macOS
brew install protobuf

# Ubuntu/Debian

apt/snap install protobuf

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro

compile the proto file inventory.proto:

cd [$GOPATH]/inventory

make proto

## Getting Started
- [Inventory](#inventory.md)
- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)


# inventory Service


1. First understand this service and ask questions , add issues to the repo when needed.

2. Run a mongo Docker image  (DO NOT INSTALL MONGODB IN YOUR SERVER DEV OR ANY DATABASE !)

3. clone and run https://github.com/mrvautin/adminMongo

4. Add new database "inventory-services" to mongo database

5. Add "inventory" as a collection to the inventory database

6. go run main.go , this will use default setup

7. execute Micro command line   $ micro > list services

8. You should find "go.micro.srv.inventory" , then we are good to go



# Running Micro Web

Run micro > micro web
You should see the service is registered and running
We can also use this web interface to test all methods for our micro-service, this is a mix of JSON to gPRC

# Running client

We have created client to test our services when called by other services:

cd /client

go run client.go

# Running the gateway

We also have a gateway HTTP server,  the same as client functionalty,  but use the micro-way of exposing

services to http world.

cd /gateway

go run gateway.go


```
Read :

curl -d '{"data":"value", "data2":"value2"}' -H "Content-Type: application/json" -X POST http://localhost:8000/read

Add:

curl -d '{"data":"value", "data2":"value2"}' -H "Content-Type: application/json" -X POST http://localhost:8000/add

Event:

curl -d '{"data":"value", "data2":"value2"}' -H "Content-Type: application/json" -X POST http://localhost:8000/event

```


Clone this repository to new folder that is targeting your domain subject
git clone gitlab/inventory.git <yourfolder>

Cd  <yourfolder>

Git init

Notes: you can't use this or push code its a inventory only

# Concept
Using go micro will generate a stub code that we will used as the base of
our micro services journey with some additional design patterns, in order to have fast roadmap, I have made deep research that we can partially depends on go-micro for creating our service oriented platform  .

![inventory](inventory.png)

This is the Inventory service is generates as follows

```

micro new inventory

Creating service go.micro.srv.inventory in [$GOPATH]/go/src/inventory

.

├── main.go

├── plugin.go

├── client

│   └── client.go       -> our native gPRC client not http or JSON

├── datastore

│   └── datastore.go    -> our data store adaptor

├── gateway

│   └── gateway.go      -> our JSON to gPRC gatway

├── handler

│   └── inventory.go    -> for gPRC client interface

├── subscriber

│   └── inventory.go    -> for message broker not used for now

├── proto/inventory

│   └── inventory.proto  -> stub file alwasys generates

├── Dockerfile

├── Makefile

└── README.md

```

Client , datastore , gateway is not part of generated codes , added based on our design choices ,

each will be explained and followed by the developer.

# Client.go

This code is a standard gPRC client it was design to give the micro-service package native entry point using gPRC protocol

, when developing the interface this cost must be used to test the interface developed to access the core service , later this will be the base for any client want to consume the services using gPRC , its the developer responsibility to get enough code when possible that allow the interface to be tested

Example :

// get instant of default micro client

client := inventory.CreateInventoryClient("go.micro.inventory.srv",micro.defualtclient)
// call the services read from the proto buff generated interace
req.id = "100"
rsp := client.Read(req)
fmt.Println("Reading from inventory client %v",rsp)
// do the same with all other defined interfaces


# Datastore.go

Datastore is an adaptor , designed to allow protobuf protocol interface a have persistent storage i.e database,
the design allow any database system to be used , in the example bellow we use, MongoDB but we can use Mysql , Postrgress , Oracle or any database

// We need to create instant of client to MongoDB

func NewMongoDB(uri string ,db string) *mongodb.client {
...

	return &DataService{
		c:  client,
		db: dbc,
	}, nil
}


// The core concept we defined a data services like this :

```

type  DataDervice struct {
     DS *mongodb.cleint   // we can have Mysql driver
     dbc string          // point to physical database name i.e "inventory"
}


// Read a physical record from the database using

func (d *DataService) Read(req interface{}) (rsp interface{}, error) {

    f:= BSON.fileter("id":req.id)

    d.ds.FindOne(...)

}

... all other CURD methods

thats all

In the main program using Micro-Serices we need to modify how it works as follows

````

// New datastore

...

uri:= defaultHost + port)

// initialise datastore to work with handler

ds, err:= datastore.NewDataService(uri,database)

if (err !=nil) {

    log.Fatal(err)

}

// create new handler in our example inventory handler that used by Micro Services

    bpHdl := &handler.Inventory{DS:ds}
    // Register Handler
    inventory.RegisterInventoryHandler(service.Server(), custHdl)

```

## Micro Configuration

- FQDN: go.micro.srv.inventory
- Type: srv
- Alias: inventory

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./inventory-srv
```

Build a docker image
```
make docker
```

