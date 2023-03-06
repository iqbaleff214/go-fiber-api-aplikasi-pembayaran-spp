package service

import "errors"

var ErrUsernameMandatory = errors.New("username is mandatory")

var ErrPasswordMandatory = errors.New("password is mandatory")

var ErrPasswordIncorrect = errors.New("incorrect password")

var ErrUserNotFound = errors.New("user not found")

var ErrSignedToken = errors.New("unable to generate token")

var ErrNotFound = errors.New("data not found")

var ErrFailedToStore = errors.New("failed to store new data")

var ErrFailedToUpdate = errors.New("failed to update existing data")

var ErrFailedToDestroy = errors.New("failed to delete data")
