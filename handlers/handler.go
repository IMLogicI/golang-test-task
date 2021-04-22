package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"gRPC_test/uniqId"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func UserIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		uIds, err := ProcessUIds(params)
		if err!=nil{
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resultId, err := gRPCClient(uIds)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(resultId); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func ProcessUIds (params map[string][]string) (uniqId.UserIds, error){
	if len(params) > 2 || len(params) < 2 {
		return uniqId.UserIds{}, fmt.Errorf("The number of params not equal 2.")
	}

	var uIds uniqId.UserIds

	if _, ok := params["user1"]; !ok {
		return uniqId.UserIds{}, fmt.Errorf("Can not parse param user1.")
	}

	uIds.User1 = strings.Join(params["user1"], "")

	if _, ok := params["user2"]; !ok {
		return uniqId.UserIds{}, fmt.Errorf("Can not parse param user2.")
	}

	uIds.User2 = strings.Join(params["user2"], "")

	return  uIds, nil
}

func gRPCClient(uIds uniqId.UserIds) (*uniqId.UniqId, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect server, %v", err)
	}
	defer conn.Close()

	u := uniqId.NewUniqIdServiceClient(conn)

	response, err := u.CreateUniqId(context.Background(), &uIds)
	if err != nil {
		return &uniqId.UniqId{}, err
	}

	return response, nil
}
