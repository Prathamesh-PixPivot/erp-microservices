package handler

import (
	"context"
	"time"

	"itsm/internal/models"
	service "itsm/internal/services"

	pb "itsm/internal/transport/grpc/proto"

	"go.uber.org/zap"
)

// ITSMServer implements the ITSMService gRPC interface.
type ITSMServer struct {
	pb.UnimplementedITSMServiceServer
	ITSMService service.ITSMService
	Logger      *zap.Logger
}

func NewITSMServer(address string, itsmService service.ITSMService, logger *zap.Logger) *ITSMServer {
	return &ITSMServer{
		ITSMService: itsmService,
		Logger:      logger,
	}
}

func convertIncident(i *models.Incident) *pb.Incident {
	return &pb.Incident{
		Id:          i.ID,
		Title:       i.Title,
		Description: i.Description,
		Hostid:      i.HostID,
		Status:      i.Status,
		CreatedAt:   i.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   i.UpdatedAt.Format(time.RFC3339),
	}
}

func (s *ITSMServer) CreateIncident(ctx context.Context, req *pb.CreateIncidentRequest) (*pb.CreateIncidentResponse, error) {
	s.Logger.Info("CreateIncident called", zap.Any("request", req))
	incident := &models.Incident{
		Title:       req.Incident.Title,
		Description: req.Incident.Description,
		HostID:      req.Incident.Hostid,
	}
	created, err := s.ITSMService.CreateIncident(incident)
	if err != nil {
		return nil, err
	}
	return &pb.CreateIncidentResponse{
		Incident: convertIncident(created),
	}, nil
}

// Implement GetIncident, UpdateIncident, and ListIncidents similarly.

func (s *ITSMServer) GetIncident(ctx context.Context, req *pb.GetIncidentRequest) (*pb.GetIncidentResponse, error) {
	s.Logger.Info("GetIncident called", zap.String("id", req.Id))
	incident, err := s.ITSMService.GetIncident(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetIncidentResponse{
		Incident: convertIncident(incident),
	}, nil
}

func (s *ITSMServer) UpdateIncident(ctx context.Context, req *pb.UpdateIncidentRequest) (*pb.UpdateIncidentResponse, error) {
	s.Logger.Info("UpdateIncident called", zap.Any("request", req))
	incident := &models.Incident{
		ID:          req.Incident.Id,
		Title:       req.Incident.Title,
		Description: req.Incident.Description,
		HostID:      req.Incident.Hostid,
		Status:      req.Incident.Status,
	}
	if err := s.ITSMService.UpdateIncident(incident); err != nil {
		return nil, err
	}
	return &pb.UpdateIncidentResponse{
		Incident: convertIncident(incident),
	}, nil
}

func (s *ITSMServer) ListIncidents(ctx context.Context, req *pb.ListIncidentsRequest) (*pb.ListIncidentsResponse, error) {
	s.Logger.Info("ListIncidents called")
	incidents, err := s.ITSMService.ListIncidents()
	if err != nil {
		return nil, err
	}
	var pbIncidents []*pb.Incident
	for _, inc := range incidents {
		pbIncidents = append(pbIncidents, convertIncident(inc))
	}
	return &pb.ListIncidentsResponse{
		Incidents: pbIncidents,
	}, nil
}
