package secretstring

type SecretString struct {
	secret string
	mask   string
}

func New(secret string) SecretString {
	return SecretString{
		secret: secret,
		mask:   "********",
	}
}

func NewWithMask(secret string, mask string) SecretString {
	return SecretString{
		secret: secret,
		mask:   mask,
	}
}

func (s SecretString) String() string {
	return s.mask
}

func (s SecretString) GetSecret() string {
	return s.secret
}
