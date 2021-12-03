package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/lucianohorvath-ml/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Starting calculator server...")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	// Register reflection service on gRPC server (a.k.a. "Server Reflection").
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Received Sum RPC: %v\n", req)

	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	res := &calculatorpb.SumResponse{
		SumResult: firstNumber + secondNumber,
	}

	return res, nil
}

func (s *server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Printf("Received PrimeNumberDecomposition RPC: %v\n", req)

	number := req.GetNumber()
	var divisor = int64(2)

	for number > 1 {
		if number%divisor == 0 {
			log.Printf("[server streaming] send divisor: %v\n", divisor)

			_ = stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			log.Printf("Divisor has increased to %v\n", divisor)
		}
	}

	return nil
}

func (s *server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Received ComputeAverage RPC (stream)\n")

	var sum, count int32 = 0, 0

	for {
		// Leo un request
		req, err := stream.Recv()

		if err == io.EOF {
			// EOF ==> El cliente terminó de enviar requests. Calculemos el promedio y retornemos.
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: average,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sum += req.GetNumber()
		count++
	}
}

func (s *server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("Received FindMaximum RPC (stream)")
	var maximum int32 = 0

	for {
		// Leo el request
		req, err := stream.Recv()
		if err == io.EOF {
			// EOF ==> El cliente terminó de enviar requests.
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		number := req.GetNumber()
		if number > maximum {
			// Nuevo máximo encontrado, lo retorno inmediatamente y sigo escuchando.
			maximum = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				Maximum: maximum,
			})

			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}

func (s *server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	log.Printf("Received SquareRoot RPC: %v\n", req)

	number := req.GetNumber()
	if number < 0 {
		log.Printf("Error: negative number!")

		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}
