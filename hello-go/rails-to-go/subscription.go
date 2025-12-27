package railstogo

import (
	"fmt"
	"strings"
	"time"
)

type SubscriptionStatus string

const (
	Active    SubscriptionStatus = "active"
	Inactive  SubscriptionStatus = "inactive"
	Cancelled SubscriptionStatus = "cancelled"
)

type Subscription struct {
	ID              string             // unique identifier
	CustomerID      string             // references a Customer
	Status          SubscriptionStatus // e.g., active, inactive, cancelled
	Currency        string             // e.g., "usd"
	AmountCents     int                // amount in cents
	NextBillingDate *time.Time         // date of next billing
	CancelledAt     *time.Time         // timestamp of cancellation, if any
	CancelReason    *string            // reason for cancellation, if any

	CreatedAt time.Time // timestamp of creation
	UpdatedAt time.Time // timestamp of last update
}

// "callback" replaced in go with explicit normalization function
func (s *Subscription) Normalize() {
	s.Currency = strings.ToUpper(s.Currency)

	if s.Status == "" {
		s.Status = Inactive
	}
}

// Validations
func (s *Subscription) Validate() error {
	if s.CustomerID == "" {
		return fmt.Errorf("CustomerID is required")
	}
	if s.AmountCents <= 0 {
		return fmt.Errorf("AmountCents must be greater than zero")
	}
	if s.Currency == "" {
		return fmt.Errorf("Currency is required")
	}
	if len(s.Currency) != 3 {
		return fmt.Errorf("Currency must be a 3-letter ISO code")
	}
	switch s.Status {
	case Active, Inactive, Cancelled:
		// valid status
		return nil
	default:
		return fmt.Errorf("Invalid subscription status: %s", s.Status)
	}
}

// Scope
func (s Subscription) Billable(now time.Time) bool {
	return s.Status == Active && s.NextBillingDate != nil && !s.NextBillingDate.After(now)
}
