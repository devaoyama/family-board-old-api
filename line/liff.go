package line

type Liff struct {
	Uid string
	Name string
	Picture string
}

func VerifiedIdToken(idToken string) *Liff {
	// todo:idTokenの有効性を確認する処理を実装する
	return &Liff{
		Uid: "123456789",
		Name: "タナカ",
		Picture: "123456789",
	}
}
