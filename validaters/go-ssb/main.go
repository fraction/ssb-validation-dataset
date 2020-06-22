package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

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

	failed := 0
	checked := 0
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

		checked++
		gotRef, _, err := legacy.Verify(tc.Message, hmacKey)
		if tc.Valid {
			if err != nil {
				failed++
				check(err)
			}
		} else {
			if err == nil {
				failed++
				fmt.Printf("%05d: shouldnt pass (%s): %s\n", i, gotRef.ShortRef(), tc.Error)
				if !tc.ID.Equal(*gotRef) {
					fmt.Printf("%05d: key divergence! %s %s\n", i, gotRef.Ref(), tc.ID.Ref())
				}
			}
		}
	}

	fmt.Printf("failed on %d cases \n%f%% of checked (%d)\n%d ignored \n", failed, float64(failed*100/checked), checked, len(cases)-checked)
}
