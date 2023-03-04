package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"net"
	proto "team00/api/proto"
	"time"
)

var (
	minMean = -10
	maxMean = 10

	minSTD = 0.3
	maxSTD = 1.5
)

type Server struct {
	proto.UnimplementedDeviceServiceServer
}

func main() {
	port := flag.Int("p", 8080, "server port")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterDeviceServiceServer(s, &Server{})
	log.Printf("Server is listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

func (s *Server) DeviceInfo(ctx *proto.RequestEmpty, stream proto.DeviceService_DeviceInfoServer) error {
	// Создание уникального ключа для  пользователя
	uuid := uuid.New().String()
	log.Println("UUID generated", uuid)

	// Рандомайзер среднего и
	randMean := float64(minMean) + rand.Float64()*(float64(maxMean)-float64(minMean))
	randStd := minSTD + rand.Float64()*(maxSTD-minSTD)
	log.Println("Random mean value:", randMean, "Random std value: ", randStd)

	// Цикл для отправки всех данных с определенным промежутком времени
	for {
		// Создание переменной с текущем временем в формате UTC и перевод в строку
		t := time.Now().UTC()
		timeUtc := timestamppb.New(t)
		log.Println("Current timestamp in UTC", timeUtc)
		log.Println("Current session", uuid)

		// Создание новой структуры для ответа
		var resp = &proto.ResponseMess{
			SessionID: uuid,
			Frequency: rand.NormFloat64()*randStd + randMean,
			Utc:       timeUtc,
		}
		// Отправляем ответ
		err := stream.Send(resp)
		fmt.Printf("\nresp: %v\n", resp)
		if err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(10 * time.Millisecond)
	}
	return nil
}
