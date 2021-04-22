package uniqId

import (
	"context"
	"testing"
)

type tests struct{
	testUids UserIds
	resultUniq UniqId
}

func TestCreateUniqId(t *testing.T) {
	testCases:=testCasesInit()
	s := Server{}

	for _, test:=range testCases{
		uniqId, _ := s.CreateUniqId(context.Background(),&test.testUids)
		if uniqId.Uid!=test.resultUniq.Uid{
			t.Errorf("waiting %v, have: %v", test.resultUniq.Uid, uniqId.Uid)
		}
	}

}

func testCasesInit() (testCases []tests){
	testCases = append(testCases,
		//simple case no specific
		tests{
		testUids:   UserIds{
			User1: "e3wedaadq2s",
			User2: "WEdfsdt5ed",
		},
		resultUniq: UniqId{
			Uid: "10s11WEdfsdt5ede3wedaadq2s",
		},
	},
	//u1 and u2 swap values
		tests{
			testUids: UserIds{
				User1: "WEdfsdt5ed",
				User2: "e3wedaadq2s",
			},
			resultUniq: UniqId{
				Uid: "10s11WEdfsdt5ede3wedaadq2s",
			},
		},
		//test alphabetic sort
		//criteria for u1+u2=uid=u2+u1
		tests{
			testUids: UserIds{
				User1: "a",
				User2: "A",
			},
			resultUniq: UniqId{
				Uid: "1s1Aa",
			},
		},
		//test similar simbols, different len, similar len1+len2
		//criteria for exclude aa+aaa=aaaaa=a+aaaa (joined id not uniq)
		//illustrate profit of len including in uniq id (not len sum)
		tests{
			testUids: UserIds{
				User1: "aa",
				User2: "aaa",
			},
			resultUniq: UniqId{
				Uid: "2s3aaaaa",
			},
		},
		tests{
			testUids: UserIds{
				User1: "a",
				User2: "aaaa",
			},
			resultUniq: UniqId{
				Uid: "1s4aaaaa",
			},
		},
		//test, when conc(len1,len2) not uniq
		//criteria for exclude 111 as 1:11 or 11:1
		//illustrate profit of separator
		tests{
			testUids: UserIds{
				User1: "a",
				User2: "basafgtndih",
			},
			resultUniq: UniqId{
				Uid: "1s11abasafgtndih",
			},
		},
		tests{
			testUids: UserIds{
				User1: "h",
				User2: "abasafgtndi",
			},
			resultUniq: UniqId{
				Uid: "11s1abasafgtndih",
			},
		})
	return
}