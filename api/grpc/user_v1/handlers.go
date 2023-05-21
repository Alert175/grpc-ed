package user_v1

import (
	"context"
	grpc_user_v1 "grpc-ed/pkg/grpc/user_v1/api/grpc/user_v1"
	"io"
	"log"
)

// Get Обработка стандартного вызова
func (s *GrpcUserV1Server) Get(ctx context.Context, req *grpc_user_v1.GetRequest) (res *grpc_user_v1.GetResponse, err error) {
	return &grpc_user_v1.GetResponse{Info: &grpc_user_v1.UserInfo{Id: req.GetId()}}, nil
}

// InputStream Обработка входящего потока
func (s *GrpcUserV1Server) InputStream(stream grpc_user_v1.UserV1_InputStreamServer) error {
	for {
		record, errRec := stream.Recv()
		if errRec != nil {
			if errRec == io.EOF {
				break
			}
			return errRec
		}
		log.Printf("new id: %d", record.GetId())
	}
	return nil
}

// OutputStream обработка иходящего потока
func (s *GrpcUserV1Server) OutputStream(req *grpc_user_v1.GetRequest, stream grpc_user_v1.UserV1_OutputStreamServer) (err error) {
	info := grpc_user_v1.UserInfo{}
	info.Id = req.GetId()
	res := grpc_user_v1.GetResponse{}
	res.Info = &info
	err = stream.Send(&res)
	err = stream.Send(&res)
	err = stream.Send(&res)
	err = stream.Send(&res)
	return err
}

// BidirectionalStream Обработка двунапрвленного потока
func (s *GrpcUserV1Server) BidirectionalStream(stream grpc_user_v1.UserV1_BidirectionalStreamServer) (err error) {
	for {
		record, errRec := stream.Recv()
		if errRec != nil {
			if errRec == io.EOF {
				break
			}
			return errRec
		}
		log.Printf("new id: %d", record.GetId())
	}
	info := grpc_user_v1.UserInfo{}
	info.Id = 123
	res := grpc_user_v1.GetResponse{}
	res.Info = &info
	err = stream.Send(&res)
	return nil
}
