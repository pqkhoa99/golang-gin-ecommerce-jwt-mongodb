package helpers

import "errors"

var (
	ErrCantFindProduct     = errors.New("can't find product")
	ErrCantDecodeProducts  = errors.New("can't find product")
	ErrUserIDIsNotValid    = errors.New("user is not valid")
	ErrCantUpdateUser      = errors.New("cannot add product to cart")
	ErrCantRemoveItem      = errors.New("cannot remove item from cart")
	ErrCantGetItem         = errors.New("cannot get item from cart ")
	ErrCantBuyCartItem     = errors.New("cannot update the purchase")
	ErrUserNotFound        = errors.New("user is not found")
	ErrUserAlreadyExists   = errors.New("User already exists")
	ErrPhoneIsAlreadyInUse = errors.New("Phone is already in use")
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
