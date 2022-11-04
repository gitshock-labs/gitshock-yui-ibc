package types

const (
	// ModuleName defines the module name
	ModuleName = "incentives"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_incentives"

	// Version defines the current version the IBC module supports
	Version = "incentives-1"

	// PortID is the default port id that module binds to
	PortID = "incentives"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("incentives-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
