package migrations

import (
	"go.uber.org/zap"
	"hrms/internal/domain"

	"gorm.io/gorm"
)

// Migrate runs database migrations
func Migrate(db *gorm.DB, logger *zap.Logger) {
	err := db.AutoMigrate(
		// Core entities
		&domain.Organization{},
		&domain.Department{},
		&domain.Designation{},
		&domain.Shift{},
	
		// Employee-related tables
		&domain.Employee{},
		&domain.WorkHistory{},
		&domain.EmployeeDocument{},
		&domain.EmployeeExit{},
		&domain.LoanAdvance{},
		&domain.Payroll{},
	
		// Attendance & Leave Management
		&domain.Attendance{},
		&domain.Leave{},
		&domain.LeaveBalance{},
		&domain.PublicHoliday{},
		&domain.LeavePolicy{},
	
		// Compensation & Benefits
		&domain.SalaryStructure{},
		&domain.Bonus{},
		&domain.EmployeeBenefits{},
		&domain.EmployeePerk{},
		&domain.Expense{},
	
		// Performance & Development
		&domain.PerformanceReview{},
		&domain.Performance{},
		&domain.PerformanceKPI{},
		&domain.SkillDevelopment{},
	)
	
	if err != nil {
		logger.Fatal("❌ Migration failed", zap.Error(err))
	}
	logger.Info("✅ Database migration completed successfully")
}
