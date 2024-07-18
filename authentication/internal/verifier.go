package internal

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

// mockgen -source=authentication/verifier.go -destination=authentication/internal/mock/verifier.go -package=mock

type TokenVerifier interface {
	// VerifyIDToken verifies the signature	and payload of the provided ID token.
	//
	// VerifyIDToken accepts a signed JWT token string, and verifies that it is current, issued for the
	// correct Firebase project, and signed by the Google Firebase services in the cloud. It returns
	// a Token containing the decoded claims in the input JWT. See
	// https://firebase.google.com/docs/auth/admin/verify-id-tokens#retrieve_id_tokens_on_clients for
	// more details on how to obtain an ID token in a client app.
	//
	// In non-emulator mode, this function does not make any RPC calls most of the time.
	// The only time it makes an RPC call is when Google public keys need to be refreshed.
	// These keys get cached up to 24 hours, and therefore the RPC overhead gets amortized
	// over many invocations of this function.
	//
	// This does not check whether or not the token has been revoked or disabled. Use `VerifyIDTokenAndCheckRevoked()`
	// when a revocation check is needed.
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)

	// VerifyIDTokenAndCheckRevoked verifies the provided ID token, and additionally checks that the
	// token has not been revoked or disabled.
	//
	// Unlike `VerifyIDToken()`, this function must make an RPC call to perform the revocation check.
	// Developers are advised to take this additional overhead into consideration when including this
	// function in an authorization flow that gets executed often.
	VerifyIDTokenAndCheckRevoked(ctx context.Context, idToken string) (*auth.Token, error)
}
