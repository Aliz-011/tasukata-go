package models

type UserRole string
type StatusEnum string
type DeliveryStatus string
type PaymentMethod string

const (
	Admin    UserRole = "admin"
	Customer UserRole = "customer"
	Moderator UserRole = "moderator"

	Pending   StatusEnum = "pending"
	Ready     StatusEnum = "ready"
	OnTheWay  StatusEnum = "on the way"
	Canceled  StatusEnum = "cancelled"
	DeliveredStatusEnum  StatusEnum = "delivered"

	// DeliveryStatus values
	InProgress DeliveryStatus = "in_progress"
	Delivered  DeliveryStatus = "delivered"
	Failed     DeliveryStatus = "failed"

	// PaymentMethod values
	CashOnDelivery PaymentMethod = "cash_on_delivery"
	CreditCard     PaymentMethod = "credit_card"
	DebitCard      PaymentMethod = "debit_card"
)