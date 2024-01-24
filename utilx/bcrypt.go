package utilx

import "golang.org/x/crypto/bcrypt"

// TODO: wrap error for too long pw
func BEncrypt(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
}

func BCheck(pw string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pw))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}
