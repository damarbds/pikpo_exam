package services

import (
	"pikpo_exam/pkg/db"
	"pikpo_exam/pkg/services/todo_service"

	"google.golang.org/grpc"
)

func RegisterService(server *grpc.Server, dbHandler db.IDatabaseAdapter) {
	todoService := todo_service.NewTodoService(dbHandler)

	todo_service.RegisterToDoServiceServer(server, todoService)
}
