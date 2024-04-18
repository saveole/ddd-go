package ch2

import "context"

type UserType = int
type SubscriptionType = int

const (
	unknownUserType UserType = iota
	lead
	customer
	churned
	lostLead
)

const (
	unknownSubscriptionType SubscriptionType = iota
	basic
	premium
	exclusive
)

type LeadRequest struct {
	email string
}

type Lead struct {
	id string
}

type LeadCreator interface {
	CreateLead(ctx context.Context, request LeadRequest) (Lead, error)
}

type Customer struct {
	leadID string
	userID string
}

func (c *Customer) UserID() string {
	return c.userID
}

func (c *Customer) SetUserID(userID string) {
	c.userID = userID
}

type LeadConvertor interface {
	Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error)
}

func (l *Lead) Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error) {
	// TODO implement me
	panic("implement me")
}

type UserAddRequest struct {
	UserType       UserType
	Email          string
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type UserModifyRequest struct {
	ID             string
	Email          string
	UserType       UserType
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type User struct {
	ID             string
	PaymentDetails PaymentDetails
}

type PaymentDetails struct {
	stripeTokenID string
}

type UserManager interface {
	AddUser(ctx context.Context, request UserAddRequest) (User, error)
	ModifyUser(ctx context.Context, request UserModifyRequest) (User, error)
}
