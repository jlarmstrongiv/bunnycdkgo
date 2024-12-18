package bunnycdkgo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/jlarmstrongiv/bunnycdkgo/bunnycdkgo/jsii"
)

type Hello interface {
	SayHello() *string
}

// The jsii proxy struct for Hello
type jsiiProxy_Hello struct {
	_ byte // padding
}

func NewHello() Hello {
	_init_.Initialize()

	j := jsiiProxy_Hello{}

	_jsii_.Create(
		"bunny-cdk.Hello",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewHello_Override(h Hello) {
	_init_.Initialize()

	_jsii_.Create(
		"bunny-cdk.Hello",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_Hello) SayHello() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"sayHello",
		nil, // no parameters
		&returns,
	)

	return returns
}

