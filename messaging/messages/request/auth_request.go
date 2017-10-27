package request

type AuthRequest struct {
    UserName string
    Email string
    EncryptedPass string
}