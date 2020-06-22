module go.mindeco.de/ssb-validator

go 1.14

// We need our internal/extra25519 since agl pulled his repo recently.
// Issue: https://github.com/cryptoscope/ssb/issues/44
// Ours uses a fork of x/crypto where edwards25519 is not an internal package,
// This seemed like the easiest change to port agl's extra25519 to use x/crypto
// Background: https://github.com/agl/ed25519/issues/27#issuecomment-591073699
// The branch in use: https://github.com/cryptix/golang_x_crypto/tree/non-internal-edwards
replace golang.org/x/crypto => github.com/cryptix/golang_x_crypto v0.0.0-20200303113948-2939d6771b24

replace go.cryptoscope.co/ssb => /Users/cryptix/ssb/go-ssb

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/shurcooL/go v0.0.0-20200502201357-93f07166e636 // indirect
	github.com/shurcooL/go-goon v0.0.0-20170922171312-37c2f522c041
	go.cryptoscope.co/ssb v0.0.0-20200302095059-b4d663e8b635
)
