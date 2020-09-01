// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitializeUserService(foo string, bar int) *UserService {
	mainMockUserRepo := NewMockUserRepo(foo, bar)
	userService := NewUserService(mainMockUserRepo)
	return userService
}
