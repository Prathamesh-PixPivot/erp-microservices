package models

import "gorm.io/gorm"
import pb "organization-service/grpc/userpb"

type Organization struct {
    gorm.Model
    GSTIn      string  `json:"gstIn"`
    Name       string  `json:"name"`
    Phone      string  `json:"phone"`
    Email      string  `json:"email"`
    Address    string  `json:"address"`
    TemplateID string  `json:"template_id"`
    Website    string  `json:"website"`
    City       string  `json:"city"`
    Country    string  `json:"country"`
    State      string  `json:"state"`
    ModuleID   string  `json:"module_id"`
    Zipcode    string  `json:"zipcode"`
    Industry   string  `json:"industry"`
    // Users will be fetched from user-service
    Users      []*pb.User  `json:"users" gorm:"-"`
}
