package main

type QueryReq struct {
	Name string
	Age  int
}

type Handle func(req *QueryReq) (Handle, error)
