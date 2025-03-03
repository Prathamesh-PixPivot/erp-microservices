package migrations

import (
	"go.uber.org/zap"
	"hrms/internal/domain"

	"gorm.io/gorm"
)

// Migrate runs database migrations
func Migrate(db *gorm.DB, logger *zap.Logger) {
	err := db.AutoMigrate(
		&domain.Organization{},
		&domain.Employee{},
		&domain.Shift{},
		&domain.WorkHistory{},
		&domain.EmployeeDocument{},
		&domain.EmployeeExit{},
		&domain.LoanAdvance{},
		&domain.Department{},
		&domain.Designation{},
		&domain.Attendance{},
		&domain.Leave{},
		&domain.LeaveBalance{},
		&domain.PublicHoliday{},
		&domain.LeavePolicy{},
		&domain.Payroll{},
		&domain.SalaryStructure{},
		&domain.Performance{},
		&domain.Bonus{},
		&domain.EmployeeBenefits{},
		&domain.EmployeePerk{},
		&domain.Expense{},
		&domain.PerformanceReview{},
		&domain.SkillDevelopment{},
		&domain.PerformanceKPI{},
	)
	if err != nil {
		logger.Fatal("❌ Migration failed", zap.Error(err))
	}
	logger.Info("✅ Database migration completed successfully")
}
