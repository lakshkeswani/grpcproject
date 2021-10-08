package controller

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"grpcproject/project/resources"

	"gorm.io/gorm"
	"grpcproject/project/projectpb"
	"grpcproject/project/server/model"
	"grpcproject/project/server/service"
)

func (*Server) CreateBooking(ctx context.Context, req *projectpb.CreateBookingRequest) (*projectpb.CreateBookingResponse, error) {
	resources.Log.Info(" CreateBooking Called  ")

	var Booking model.Booking
	Booking.ID = uint(req.GetBooking().GetBookingId())
	Booking.TripId = int(req.GetBooking().GetTripId())
	PickupTime, err := ptypes.Timestamp(req.GetBooking().GetPickupTime())
	if err != nil {
	}
	DropOffTime, err := ptypes.Timestamp(req.GetBooking().GetDropOffTime())
	if err != nil {
	}
	Booking.PickupTime = PickupTime
	Booking.DropOffTime = DropOffTime
	Booking.Fare = int(req.GetBooking().GetFare())
	Booking.PickupStopId = int(req.GetBooking().GetPickupStopId())
	Booking.DropOffStopId = int(req.GetBooking().GetDropOffStopId())
	Booking.UserId = int(req.GetBooking().GetUserId())
	Booking.Status = req.GetBooking().GetStatus()

	_, err = service.CreateBooking(Booking)
	if err != nil {
		return nil, err
	}
	return &projectpb.CreateBookingResponse{Result: "Booking created"}, nil

}

func (*Server) ReadBooking(ctx context.Context, req *projectpb.ReadBookingRequest) (*projectpb.ReadBookingResponse, error) {
	resources.Log.Info(" ReadBooking Called  ")

	//ok,err:=tokenwork(ctx, 0)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if ok {
	//}

	var Booking model.Booking
	Booking, err := service.ReadBooking(int(req.GetBookingId()))
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}
	trip, err := bookingToBookingpb(Booking)
	return &projectpb.ReadBookingResponse{Booking: trip}, nil
}
func bookingToBookingpb(Booking model.Booking) (*projectpb.Booking, error) {
	PickupTime, err := ptypes.TimestampProto(Booking.PickupTime)
	if err != nil {
		return nil, err
	}
	DropOffTime, err := ptypes.TimestampProto(Booking.DropOffTime)
	if err != nil {
		return nil, err
	}
	return &projectpb.Booking{
		BookingId:     int64(Booking.ID),
		UserId:        int64(Booking.UserId),
		TripId:        int64(Booking.TripId),
		Fare:          -int64(Booking.Fare),
		PickupTime:    PickupTime,
		DropOffTime:   DropOffTime,
		PickupStopId:  int64(Booking.PickupStopId),
		DropOffStopId: int64(Booking.DropOffStopId),
		Status:        Booking.Status,
	}, nil
}
func (*Server) UpdateBooking(ctx context.Context, req *projectpb.UpdateBookingRequest) (*projectpb.UpdateBookingResponse, error) {
	var Booking model.Booking

	resources.Log.Info("service UpdateBooking called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}

	Booking.ID = uint(req.GetBooking().GetBookingId())
	Booking.TripId = int(req.GetBooking().GetTripId())
	PickupTime, err := ptypes.Timestamp(req.GetBooking().GetPickupTime())
	if err != nil {
	}
	DropOffTime, err := ptypes.Timestamp(req.GetBooking().GetDropOffTime())
	if err != nil {
	}
	Booking.PickupTime = PickupTime
	Booking.DropOffTime = DropOffTime
	Booking.Fare = int(req.GetBooking().GetFare())
	Booking.PickupStopId = int(req.GetBooking().GetPickupStopId())
	Booking.DropOffStopId = int(req.GetBooking().GetDropOffStopId())
	Booking.UserId = int(req.GetBooking().GetUserId())
	Booking.Status = req.GetBooking().GetStatus()
	_, err = service.UpdateBooking(int(req.GetBooking().GetBookingId()), Booking)
	if err != nil {
		return nil, err
	}
	trip, err := bookingToBookingpb(Booking)
	return &projectpb.UpdateBookingResponse{Booking: trip}, nil

}
func (*Server) DeleteBooking(ctx context.Context, req *projectpb.DeleteBookingRequest) (*projectpb.DeleteBookingResponse, error) {
	resources.Log.Info("service DeleteBooking called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}
	if err != nil {
		return nil, err
	}
	err = service.DeleteBooking(req.GetBookingId())
	return &projectpb.DeleteBookingResponse{Result: "sucessfully deleted"}, nil
}
func (*Server) GetAllBookingByUser(ctx context.Context, Req *projectpb.GetAllBookingByUserRequest) (*projectpb.GetAllBookingByUserResponse, error) {
	resources.Log.Info("service DeleteBooking called")

	/*	_,err:=tokenwork(ctx, int(Req.GetUserId()))
		if err != nil {
			return nil, err
		}*/
	var bookings []model.Booking
	var bookingResponse []*projectpb.Booking
	service.GetAllBookingsByUser(int(Req.GetUserId()), &bookings)
	for _, booking := range bookings {
		bookingpb, _ := bookingToBookingpb(booking)
		bookingResponse = append(bookingResponse, bookingpb)
	}
	return &projectpb.GetAllBookingByUserResponse{Bookings: bookingResponse}, nil
}
