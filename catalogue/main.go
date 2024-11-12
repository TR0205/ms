package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	// protocol buffersの略
	pb "gihyo/catalogue/proto/book"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Book struct {
	Id     int
	Title  string
	Author string
	Price  int
}

var (
	book1 = Book{
		Id:     1,
		Title:  "The Awaking",
		Author: "Kate Chopin",
		Price:  100,
	}
	book2 = Book{
		Id:     2,
		Title:  "City ob Glass",
		Author: "Paul Auster",
		Price:  2000,
	}
	books = []Book{book1, book2}
)

func getBook(i int32) Book {
	return books[i-1]
}

type server struct {
	pb.UnimplementedCatalogueServer
}

func (s *server) GetBook(ctx context.Context, in *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	book := getBook(in.Id)

	protoBook := &pb.Book{
		Id:     int32(book.Id),
		Title:  book.Title,
		Author: book.Author,
		Price:  int32(book.Price),
	}

	return &pb.GetBookResponse{Book: protoBook}, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCatalogueServer(s, &server{})
	// リフレクションサービスの登録(gRPCurlを使用する際、gRPCサーバーがどのようなサービスやメソッドを公開しているか知ることができる)
	// 例：$ grpcurl -plaintext localhost:50051 list
	// book.Catalogue
	// grpc.reflection.v1.ServerReflection
	// grpc.reflection.v1alpha.ServerReflection
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
