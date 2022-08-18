package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Запросы и ответы
// Основной паттерн в Go kit это RPC
// Для каждого метода мы определяем структуры запроса и ответа,
// принимая все входящие и исходящие параметры соответственно
// For each method, we define request and response structs
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// Эндпоинты предоставляются самим пакетом и имеют вид
// type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
// их не нужно определять в коде
// Мы напишем простые адаптеры для преобразования каждого из методов нашей службы в конечную точку.
// Каждый адаптер принимает StringService и возвращает эндпоинт, соответствующий одному из методов.
// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}
