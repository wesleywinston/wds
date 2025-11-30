package models

// User Roles
const (
	RoleAdmin  UserRole = "ADMIN"
	RoleBuyer  UserRole = "BUYER"
	RoleVendor UserRole = "VENDOR"
)

// User Statuses
const (
	StatusPending   AccountStatus = "PENDING_APPROVAL"
	StatusActive    AccountStatus = "ACTIVE"
	StatusInactive  AccountStatus = "INACTIVE"
	StatusSuspended AccountStatus = "SUSPENDED"
)

// --- Entity Compliance Statuses ---
const (
	ComplianceVerified AccountComplianceStatus = "VERIFIED"
	ComplianceExpired  AccountComplianceStatus = "EXPIRED"
	CompliancePending  AccountComplianceStatus = "PENDING"
)

// Order Status
const (
	OrderStatusPending    OrderStatus = "PENDING"
	OrderStatusAccepted   OrderStatus = "ACCEPTED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusShipped    OrderStatus = "SHIPPED"
	OrderStatusDelivered  OrderStatus = "DELIVERED"
	OrderStatusCancelled  OrderStatus = "CANCELLED"
	OrderStatusCompleted  OrderStatus = "COMPLETED"
)

const (
	PaymentStatusPending  PaymentStatus = "PENDING"
	PaymentStatusPaid     PaymentStatus = "PAID"
	PaymentStatusRefunded PaymentStatus = "REFUNDED"
)

const (
	PaymentMethodCash  PaymentMethod = "CASH"
	PaymentMethodCheck PaymentMethod = "CHECK"
	PaymentMethodOther PaymentMethod = "OTHER" // crypto ??
)
