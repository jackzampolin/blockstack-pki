package pkg

// HDNode is an interface to allow for switching between different nodes
type HDNode interface {
	GetAddress() string
	GetPKHex() string
	String() string
}
