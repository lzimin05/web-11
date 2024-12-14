package usecase

func (u *Usecase) SelectNameQuery(name string) (string, error) {
	msg, err := u.p.SelectRandomHelloUser()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) InsertNameQuery(name string) error {

	err := u.p.InsertUser(name)
	if err != nil {
		return err
	}

	return nil
}
