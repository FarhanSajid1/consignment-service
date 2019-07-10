package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/farhansajid1/go-micro-project/consignment-service/proto"
	vessel "github.com/farhansajid1/go-micro-project/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"go.mongodb.org/mongo-driver/mongo"
)

type server struct{}

type ConsignmentItem struct {
	// primitive.ObjectID is the id of the bson object
	ID          primitive.ObjectID `bson:"_id,omitempty"` // generated on the fly
	Description string             `bson:"description"`
	Weight      int32              `bson:"weight"`
	VesselId    string             `bson:"vessel_id"`
	Containers   []*pb.Container    `bson:"container"`
}

/*string id = 1;
  string description = 2;
  int32 weight = 3;
  repeated Container containers = 4;
  string vessel_id = 5; */

var vesselClient vessel.VesselService

func createVesselReq() *vessel.Specification {
	ves := &vessel.Specification{
		MaxWeight: 100,
		Capacity:  500,
	}
	return ves
}

func toBSON(req *pb.Consignment, item *pb.Container) *ConsignmentItem {
	// convert req to bson object
	obj := &ConsignmentItem{
		Description: req.GetDescription(),
		Weight:      req.GetWeight(),
		Containers: []*pb.Container{
			&pb.Container{
				Id:         item.GetId(),
				CustomerId: item.GetCustomerId(),
				Origin:     item.GetOrigin(),
				UserId:     item.GetUserId(),
			},
		},
	}
	return obj
}

func (s *server) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	lis := req.GetContainers()
	item := lis[0]

	// Getting vessel id
	ves := createVesselReq()
	vesselresp, err := vesselClient.FindAvailable(ctx, ves)
	if err != nil {
		log.Printf("there was an error %v", err)
	}

	log.Printf("Found a vessel %v\t%v", vesselresp.Vessel.Descriptor, vesselresp.Vessel.Id)

	obj := toBSON(req, item)
	result, err := collection.InsertOne(ctx, obj)
	if err != nil {
		log.Printf("error adding the value %v", err)
	}
	oid, _ := result.InsertedID.(primitive.ObjectID)
	create := &pb.Consignment{
		Id:          oid.Hex(),
		Description: req.GetDescription(),
		Weight:      req.GetWeight(),
		Containers: []*pb.Container{
			&pb.Container{
				Id:         item.GetId(),
				CustomerId: item.GetCustomerId(),
				Origin:     item.GetOrigin(),
				UserId:     item.GetUserId(),
			},
		},
		VesselId: vesselresp.Vessel.Id,
	}
	// b, err := proto.Marshal(create)
	// if err != nil {
	// 	log.Printf("could not write to bytes %v", err)
	// }
	// ioutil.WriteFile("./input", b, 0644)

	resp.Created = true
	resp.Consignment = create
	return nil
}

var collection *mongo.Collection
var err error

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// database
	log.Printf("Connecting to the database")
	collection = connectDB()
	
	service.Init()
	pb.RegisterShippingServiceHandler(service.Server(), new(server))

	// Register the client for the vessel.
	vesselClient = vessel.NewVesselService("go.micro.srv.vessel", service.Client())



	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

func connectDB() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://@localhost:27017"))
	if err != nil {
		log.Fatalf("failed to create new MongoDB client: %#v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect client
	if err = client.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to MongoDB: %#v", err)
	}

	log.Printf("connected successfully")
	collection = client.Database("consignments").Collection("new")
	return collection
}
