package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"go.cryptoscope.co/ssb"
	"go.cryptoscope.co/ssb/message/legacy"
)

type Case struct {
	State   interface{}     `json:"state"`
	HmacKey interface{}     `json:"hmacKey"`
	Message json.RawMessage `json:"message"`
	Error   string          `json:"error"`
	Valid   bool            `json:"valid"`
	ID      ssb.MessageRef  `json:"id"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("../../data.json")
	check(err)
	defer f.Close()

	var cases []Case
	err = json.NewDecoder(f).Decode(&cases)
	check(err)

	for i, tc := range cases {
		var hmacKey *[32]byte

		if tc.HmacKey != nil {
			str, ok := tc.HmacKey.(string)
			if !ok {
				fmt.Printf("%05d: invalid hmac - skipping this: %v\n", i, tc.HmacKey)
				continue
			}

			keyData, err := base64.StdEncoding.DecodeString(str)
			if err == nil {
				var data [32]byte
				copy(data[:], keyData)
				hmacKey = &data
			} else {
				fmt.Printf("%05d: invalid hmac - skipping this: %s\n", i, err)
				continue
			}
		}

		gotRef, _, err := legacy.Verify(tc.Message, hmacKey)
		if tc.Valid {
			check(err)
		} else {
			if err == nil {
				fmt.Printf("%05d: shouldnt pass (%s): %s\n", i, gotRef.ShortRef(), tc.Error)
				if !tc.ID.Equal(*gotRef) {
					fmt.Printf("%05d: key divergence! %s %s\n", i, gotRef.Ref(), tc.ID.Ref())
				}
			}
		}
	}
}

// utils
type b64str []byte

func (s *b64str) UnmarshalJSON(data []byte) error {
	var strdata string
	err := json.Unmarshal(data, &strdata)
	if err != nil {
		return fmt.Errorf("b64str: json decode of string failed: %w", err)
	}
	decoded := make([]byte, len(strdata)) // will be shorter
	n, err := base64.StdEncoding.Decode(decoded, []byte(strdata))
	if err != nil {
		spew.Dump(data)
		spew.Dump(strdata)
		return fmt.Errorf("base64str: invalid base64 data: %w", err)
	}

	*s = decoded[:n]
	return nil
}
