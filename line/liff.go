package line

type Liff struct {
	Uid string
	Name string
	Picture string
}

func VerifiedIdToken(idToken string) *Liff {
	return &Liff{
		Uid: "123456789",
		Name: "タナカ",
		Picture: "123456789",
	}
}
