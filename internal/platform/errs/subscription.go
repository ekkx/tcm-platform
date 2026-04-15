package errs

import "connectrpc.com/connect"

var (
	ErrSubscriptionNotFound = &Error{
		Code:        "subscription_not_found",
		Message:     "subscription not found",
		ConnectCode: connect.CodeNotFound,
	}
	ErrUsageLimitExceeded = &Error{
		Code:        "usage_limit_exceeded",
		Message:     "monthly usage limit exceeded",
		ConnectCode: connect.CodeResourceExhausted,
	}
	ErrCheckoutSessionFailed = &Error{
		Code:        "checkout_session_failed",
		Message:     "failed to create checkout session",
		ConnectCode: connect.CodeInternal,
	}
	ErrPortalSessionFailed = &Error{
		Code:        "portal_session_failed",
		Message:     "failed to create portal session",
		ConnectCode: connect.CodeInternal,
	}
	ErrUnlimitedUserCannotCheckout = &Error{
		Code:        "unlimited_user_cannot_checkout",
		Message:     "unlimited plan users cannot create checkout sessions",
		ConnectCode: connect.CodeFailedPrecondition,
	}
	ErrAlreadySubscribed = &Error{
		Code:        "already_subscribed",
		Message:     "already subscribed, use customer portal to change plan",
		ConnectCode: connect.CodeFailedPrecondition,
	}
	ErrNoActiveSubscription = &Error{
		Code:        "no_active_subscription",
		Message:     "active subscription required to make reservations",
		ConnectCode: connect.CodeFailedPrecondition,
	}
)
