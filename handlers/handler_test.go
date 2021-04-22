package handlers

import (
	"net/url"
	"testing"
)

type tests struct {
	link  string
	error string
	user1 string
	user2 string
}

func TestProcessUIds(t *testing.T) {
	testCases := testCasesInit()

	for _, test := range testCases {
		var link url.URL
		link.RawQuery = test.link
		params, err := url.ParseQuery(link.RawQuery)

		if err != nil {
			t.Errorf("err to parse Query %d", err)
			break
		}

		res, err := ProcessUIds(params)
		if err != nil {
			if err.Error() != test.error {
				t.Errorf("waiting %v, have: %d", test.error, err)
				break
			}
			break
		}

		if res.User1 != test.user1 {
			t.Errorf("For user1 waiting %v, have: %v", test.user1, res.User1)
			break
		}

		if res.User2 != test.user2 {
			t.Errorf("For user2 waiting %v, have: %v", test.user2, res.User2)
			break
		}
	}
}

func testCasesInit() (testCases []tests) {
	testCases = append(testCases, tests{
		link:  "user1=A123&user2=a123",
		user1: "A123",
		user2: "a123",
		error: "nil",
	}, tests{
		link:  "user1=A123&user2=a123&user3=sfe",
		user1: "nil",
		user2: "nil",
		error: "The number of params not equal 2.",
	}, tests{
		link:  "user3=A123&user1=a123",
		user1: "nil",
		user2: "nil",
		error: "Can not parse param user1.",
	}, tests{
		link:  "user1=A123&user3=a123",
		user1: "nil",
		user2: "nil",
		error: "Can not parse param user2.",
	})
	return
}
