package response

type Login struct {
	Salt string
	Ephemeral_key string
}