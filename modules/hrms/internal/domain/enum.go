package domain

// EmployeeStatus Enum
type EmployeeStatus string

const (
	Active     EmployeeStatus = "Active"
	Resigned   EmployeeStatus = "Resigned"
	Terminated EmployeeStatus = "Terminated"
	Intern     EmployeeStatus = "Intern"
	Probation  EmployeeStatus = "Probation"
	Hired      EmployeeStatus = "Hired"
)

// EmploymentType Enum
type EmploymentType string

const (
	FullTime   EmploymentType = "Full-time"
	PartTime   EmploymentType = "Part-time"
	Contract   EmploymentType = "Contract"
	Internship EmploymentType = "Internship"
)

// ShiftType Enum
type ShiftType string

const (
	Flexible   ShiftType = "Flexible"
	Fixed      ShiftType = "Fixed"
	Rotational ShiftType = "Rotational"
)

// PunchMethod Enum
type PunchMethod string

const (
	Manual      PunchMethod = "Manual"
	Biometric   PunchMethod = "Biometric"
	Geolocation PunchMethod = "Geolocation"
)

// LeaveType Enum
type LeaveType string

const (
	AnnualLeave    LeaveType = "Annual"
	SickLeave      LeaveType = "Sick"
	CasualLeave    LeaveType = "Casual"
	MaternityLeave LeaveType = "Maternity"
	PaternityLeave LeaveType = "Paternity"
	CustomLeave    LeaveType = "Custom"
)

// LeaveStatus Enum
type LeaveStatus string

const (
	Pending  LeaveStatus = "Pending"
	Approved LeaveStatus = "Approved"
	Rejected LeaveStatus = "Rejected"
)

// Payroll Payment Status
type PaymentStatus string

const (
	PendingPayment PaymentStatus = "Pending"
	Processed      PaymentStatus = "Processed"
	Failed         PaymentStatus = "Failed"
)

// BonusType Enum - Defines the categories of bonuses
type BonusType string

const (
	PerformanceBonus BonusType = "Performance" // Based on employee performance
	FestivalBonus    BonusType = "Festival"    // Holiday or special occasion bonuses
	RetentionBonus   BonusType = "Retention"   // Reward for long-term employees
	ReferralBonus    BonusType = "Referral"    // Bonus for referring a candidate
)

// BonusStatus Enum - Tracks the approval workflow of a bonus
type BonusStatus string

const (
	PendingBonus  BonusStatus = "Pending"  // Awaiting approval
	ApprovedBonus BonusStatus = "Approved" // Bonus has been granted
	RejectedBonus BonusStatus = "Rejected" // Bonus request denied
)

// HealthInsurancePlan Enum - Different health coverage plans
type HealthInsurancePlan string

const (
	BasicHealthPlan   HealthInsurancePlan = "Basic"
	PremiumHealthPlan HealthInsurancePlan = "Premium"
	FamilyHealthPlan  HealthInsurancePlan = "Family"
)

// RetirementPlan Enum - Different retirement investment options
type RetirementPlan string

const (
	PensionFund      RetirementPlan = "Pension Fund"
	ProvidentFund    RetirementPlan = "Provident Fund"
	StockOptionsPlan RetirementPlan = "Stock Options"
)

// PerkType Enum - Various employee perks
type PerkType string

const (
	MealCard        PerkType = "Meal Card"
	TravelAllowance PerkType = "Travel Allowance"
	StockOptions    PerkType = "Stock Options"
	InternetStipend PerkType = "Internet Stipend"
)

// ExpenseType Enum - Categories for employee expense reimbursements
type ExpenseType string

const (
	TravelExpense        ExpenseType = "Travel"
	AccommodationExpense ExpenseType = "Accommodation"
	EducationExpense     ExpenseType = "Education"
	OfficeSupplies       ExpenseType = "Office Supplies"
)

// PerformanceRating Enum - Standard rating system for performance evaluation
type PerformanceRating string

const (
	Excellent PerformanceRating = "Excellent"
	Good      PerformanceRating = "Good"
	Average   PerformanceRating = "Average"
	Poor      PerformanceRating = "Poor"
)

// ReviewPeriod Enum - Specifies the review frequency
type ReviewPeriod string

const (
	Monthly   ReviewPeriod = "Monthly"
	Quarterly ReviewPeriod = "Quarterly"
	Annual    ReviewPeriod = "Annual"
)

// PromotionStatus Enum - Defines employee promotion decisions
type PromotionStatus string

const (
	NoChange          PromotionStatus = "No Change"
	PromotionApproved PromotionStatus = "Promotion Approved"
	NeedsImprovement  PromotionStatus = "Needs Improvement"
)
