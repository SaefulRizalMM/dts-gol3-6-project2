package main

type Tasks struct {
	Id          string `form:"id" json:"id"`
	Description string `form:"description" json:"description"`
	Employee    string `form:"employee" json:"employee"`
	Deadline    string `form:"deadline" json:"deadline"`
	Status      int    `form:"status" json:"status"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Tasks
}
