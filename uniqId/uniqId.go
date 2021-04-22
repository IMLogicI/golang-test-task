package uniqId

import (
	"golang.org/x/net/context"
	"sort"
	"strconv"
	"strings"
)

const (
	Separator = "s"
)

type Server struct {
}

func (s *Server) CreateUniqId(_ context.Context, uIds *UserIds) (*UniqId, error) {
	idsArray := []string{uIds.User1, uIds.User2}
	sort.Strings(idsArray)
	uniqIdArray := []string{strconv.Itoa(len(idsArray[0])), Separator, strconv.Itoa(len(idsArray[1]))}
	uniqIdArray = append(uniqIdArray, idsArray...)

	return &UniqId{Uid: strings.Join(uniqIdArray, "")}, nil
}
