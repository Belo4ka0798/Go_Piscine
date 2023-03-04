package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	// "github.com/google/uuid"
	"log"
	"team00/DBRequests"
	"team00/anomaly"
	proto "team00/api/proto"

	"google.golang.org/grpc"
	// "time"
)

func main() {
	port := flag.Int("p", 8080, "server port")
	k := flag.Float64("k", 3.0, "anomaly deviation (k > 0)")
	flag.Parse()
	if *k < 0.0 {
		log.Fatal("'K' be greater than zero")
	}

	// grpc.WithInsecure() is deprecated
	//conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), grpc.WithInsecure())
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewDeviceServiceClient(conn)

	// Использование ручки (end-point)
	stream, err := client.DeviceInfo(context.Background(), &proto.RequestEmpty{})
	if err != nil {
		log.Fatalf("Error calling DeviceInfo: %v", err)
	}
	a := anomaly.Init(*k)

	// Подключение базы данных
	db := DBRequests.CreateConnection(
		":5432",
		"",
		"",
		"postgres")
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// Создание таблицы в базе данных
	err = DBRequests.CreateTable(db)
	defer db.Close()

	// Цикл обработки входящего потока
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving response from stream: %v", err)
		}

		// Проверка явления аномалий
		isAnomaly := a.Do(response.Frequency)

		// Запись аномалий в базу данных
		if isAnomaly == true {
			err = DBRequests.AddEntry(db, response.SessionID, response.Frequency, response.Utc.AsTime())
			if err != nil {
				fmt.Println("err:", err)
				return
			}
		}
		// fmt.Printf("Received message: SessionID=%s, Frequency=%f, UTC=%v\n", response.SessionID, response.Frequency, response.Utc)
	}
}
