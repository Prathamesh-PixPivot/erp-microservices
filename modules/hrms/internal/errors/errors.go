package errors

import "errors"

var (
	ErrEmployeeNotFound      = errors.New("employee not found")
	ErrEmployeeAlreadyExists = errors.New("employee with this email already exists")
)

var (
	ErrOrganizationNotFound         = errors.New("organization not found")
	ErrOrganizationAlreadyExists    = errors.New("organization with this email or name already exists")
	ErrInvalidEmailFormat           = errors.New("invalid email format")
	ErrInvalidPhoneFormat           = errors.New("invalid phone number format")
	ErrCannotDeleteOrgWithEmployees = errors.New("cannot delete organization with existing employees")
)

var (
	ErrShiftNotFound      = errors.New("shift not found")
	ErrShiftAlreadyExists = errors.New("shift with this name already exists")
)

var (
	ErrWorkHistoryNotFound = errors.New("work history record not found")
)

var (
	ErrEmployeeDocumentNotFound = errors.New("employee document not found")
)

var (
	ErrEmployeeExitNotFound = errors.New("employee exit record not found")
	ErrClearancePending     = errors.New("employee clearance is still pending")
)

var (
	ErrLoanNotFound         = errors.New("loan/advance request not found")
	ErrLoanAlreadyProcessed = errors.New("loan/advance request has already been processed")
	ErrInsufficientDetails  = errors.New("insufficient details provided for loan/advance processing")
)

var (
	ErrDepartmentNotFound      = errors.New("department not found")
	ErrDepartmentAlreadyExists = errors.New("department with the same name already exists in the organization")
	ErrInvalidOrganization     = errors.New("invalid organization ID")
)

var (
	ErrDesignationNotFound      = errors.New("designation not found")
	ErrDesignationAlreadyExists = errors.New("designation with the same title and level already exists in the department")
	ErrInvalidDepartment        = errors.New("invalid department ID")
)

var (
	ErrAttendanceNotFound = errors.New("attendance record not found")
	ErrDuplicateCheckIn   = errors.New("employee has already checked in for the day")
	ErrInvalidCheckOut    = errors.New("cannot check out without a prior check-in")
	ErrInvalidAttendance  = errors.New("invalid attendance data")
)

var (
	ErrLeaveNotFound      = errors.New("leave request not found")
	ErrLeaveConflict      = errors.New("leave dates overlap with an existing approved leave")
	ErrInvalidLeaveDates  = errors.New("leave end date must be after start date")
	ErrUnauthorizedAction = errors.New("unauthorized action on leave request")
	ErrInvalidLeaveStatus = errors.New("invalid leave status")
	ErrInvalidLeaveID     = errors.New("invalid leave ID")
)

var (
	ErrLeaveBalanceNotFound = errors.New("leave balance record not found")
	ErrInsufficientLeave    = errors.New("insufficient leave balance")
	ErrInvalidLeaveType     = errors.New("invalid leave type")
)

var (
	ErrPublicHolidayExists   = errors.New("public holiday already exists on this date")
	ErrPublicHolidayNotFound = errors.New("public holiday not found")
)

var (
	ErrLeavePolicyExists   = errors.New("leave policy already exists for this leave type in the organization")
	ErrLeavePolicyNotFound = errors.New("leave policy not found")
)

var (
	ErrPayrollAlreadyProcessed = errors.New("payroll already processed for this employee in the given month")
	ErrPayrollNotFound         = errors.New("payroll record not found")
	ErrInvalidPayrollData      = errors.New("invalid payroll data")
)

var (
	ErrSalaryStructureExists   = errors.New("salary structure already exists for this designation")
	ErrSalaryStructureNotFound = errors.New("salary structure not found")
	ErrInvalidSalaryData       = errors.New("invalid salary structure data")
)

var (
	ErrInvalidPerformance  = errors.New("invalid performance data")
)

var (
	ErrBonusNotFound = errors.New("bonus record not found")
	ErrInvalidBonus  = errors.New("invalid bonus data")
)

var (
	ErrBenefitsNotFound = errors.New("employee benefits record not found")
	ErrInvalidBenefits  = errors.New("invalid employee benefits data")
)

var (
	ErrPerkNotFound  = errors.New("employee perk not found")
	ErrInvalidPerk   = errors.New("invalid perk data")
	ErrDuplicatePerk = errors.New("employee perk already exists")
)

var (
	ErrExpenseNotFound = errors.New("expense record not found")
	ErrInvalidExpense  = errors.New("invalid expense data")
)

var (
	ErrInvalidReviewData   = errors.New("invalid review data")
	ErrUnauthorizedAccess  = errors.New("unauthorized access to review")
)

var (
	ErrSkillAlreadyExists   = errors.New("skill development entry already exists")
	ErrSkillCreationFailed  = errors.New("failed to create skill development entry")
	ErrSkillUpdateFailed    = errors.New("failed to update skill development entry")
	ErrSkillDeletionFailed  = errors.New("failed to delete skill development entry")
)

var (
	ErrKPIAlreadyExists   = errors.New("performance KPI already exists for this review")
	ErrKPICreationFailed  = errors.New("failed to create performance KPI entry")
	ErrKPIUpdateFailed    = errors.New("failed to update performance KPI entry")
	ErrKPIDeletionFailed  = errors.New("failed to delete performance KPI entry")
)

var (
	ErrPerformanceReviewNotFound = errors.New("performance review not found")
	ErrReviewAlreadyExists       = errors.New("review already exists for this period")
	ErrReviewCreationFailed      = errors.New("failed to create performance review")
	ErrReviewUpdateFailed        = errors.New("failed to update performance review")
	ErrReviewDeletionFailed      = errors.New("failed to delete performance review")
)
var (
	// common
	ErrKPINotFound               = errors.New("KPI not found")
	ErrSkillNotFound             = errors.New("skill development record not found")
	ErrPerformanceNotFound = errors.New("performance review not found")
	ErrUnauthorized = errors.New("unauthorized action")
)

