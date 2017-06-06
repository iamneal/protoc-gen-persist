// This file is generated by protoc-gen-persist
// Source File: examples/test_issue_32/service.proto
// DO NOT EDIT !
package test_issue_32

import (
	sql "database/sql"
	strings "strings"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
)

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// WARNING ! WARNING ! WARNING ! WARNING !WARNING ! WARNING !
// In order for your code to work you have to create a file
// in this package with the following content:
//
// type TestServiceImpl struct {
// 	SqlDB *sql.DB
// }
// WARNING ! WARNING ! WARNING ! WARNING !WARNING ! WARNING !
//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func NewTestServiceImpl(driver, connString string) (*TestServiceImpl, error) {
	db, err := sql.Open(driver, connString)
	if err != nil {
		return nil, err
	}
	return &TestServiceImpl{SqlDB: db}, nil
}

// sql unary GetTime
func (s *TestServiceImpl) GetTime(ctx context.Context, req *Request) (*Request, error) {
	var (
		Id  int32
		err error
	)

	beforeRes, err := TestHook(req)

	if err != nil {

		return nil, grpc.Errorf(codes.Unknown, err.Error())

	}
	if beforeRes != nil {

		return beforeRes, nil

	}

	err = s.SqlDB.QueryRow("SELECT time FROM timesource WHERE time = $1", req.Id).
		Scan(&Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, grpc.Errorf(codes.NotFound, "%+v doesn't exist", req)
		} else if strings.Contains(err.Error(), "duplicate key") {
			return nil, grpc.Errorf(codes.AlreadyExists, "%+v already exists", req)
		}
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}
	res := Request{

		Id: Id,
	}

	return &res, nil
}