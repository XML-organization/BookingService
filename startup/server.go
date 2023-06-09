package startup

import (
	"booking-service/handler"
	"booking-service/repository"
	"booking-service/service"
	"booking-service/startup/config"
	"fmt"
	"log"
	"net"

	booking "github.com/XML-organization/common/proto/booking_service"
	saga "github.com/XML-organization/common/saga/messaging"
	"github.com/XML-organization/common/saga/messaging/nats"
	"github.com/neo4j/neo4j-go-driver/neo4j"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "booking_service"
)

func (server *Server) Start() {
	postgresClient := server.initPostgresClient()
	neo4jDriver := server.initNeo4jDriver()
	bookingNeo4jRepo := server.initBookingNeo4jRepository(neo4jDriver)
	bookingRepo := server.initBookingRepository(postgresClient)
	bookingService := server.initBookingService(bookingRepo, bookingNeo4jRepo)

	bookingHandler := server.initBookingHandler(bookingService)

	server.startGrpcServer(bookingHandler)
}

func (server *Server) initNeo4jDriver() *neo4j.Driver {
	driver, err := repository.GetNeo4jClient("bolt://accommodation_recommendation_db:7687", "neo4j", "password")

	if err != nil {
		return nil
	}

	return &driver
}

func (server *Server) initBookingNeo4jRepository(driver *neo4j.Driver) *repository.BookingNeo4jRepository {
	return repository.NewBookingNeo4jRepository(*driver)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initPostgresClient() *gorm.DB {
	client, err := repository.GetClient(
		server.config.BookingDBHost, server.config.BookingDBUser,
		server.config.BookingDBPass, server.config.BookingDBName,
		server.config.BookingDBPort)
	if err != nil {
		log.Fatal("INIT POSTGRES CLIENT PUCAAAA")
	}
	println("PRESKOCIO IF")
	return client
}

func (server *Server) initBookingRepository(client *gorm.DB) *repository.BookingRepository {
	return repository.NewBookingRepository(client)
}

func (server *Server) initBookingService(repo *repository.BookingRepository, neo4jRepo *repository.BookingNeo4jRepository) *service.BookingService {
	return service.NewBookingService(repo, neo4jRepo)
}

func (server *Server) initBookingHandler(service *service.BookingService) *handler.BookingHandler {
	return handler.NewBookingHandler(service)
}

func (server *Server) startGrpcServer(bookingHandler *handler.BookingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	booking.RegisterBookingServiceServer(grpcServer, bookingHandler)
	reflection.Register(grpcServer)
	log.Println("GRPC BOOKING SERVER USPJESNO NAPRAVLJEN")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
		println("GRPC BOOKING SERVER NIJE USPJESNO NAPRAVLJEN")
	}
}
