package upgen

// character types
const (
	CharacterAlphanumeric = iota
	CharacterLetter
	CharacterLowercase
	CharacterUppercase
	CharacterNumeric
	CharacterSpecial
	CharacterHex
	CharacterBase32
	CharacterBase64
	WordBIP39
)

func GetCharacters(set int, avoid_ambiguous bool) string {
	switch set {
	case CharacterAlphanumeric:
		return GetCharacters(CharacterLetter, avoid_ambiguous) + GetCharacters(CharacterNumeric, avoid_ambiguous)
	case CharacterLetter:
		return GetCharacters(CharacterLowercase, avoid_ambiguous) + GetCharacters(CharacterUppercase, avoid_ambiguous)
	case CharacterLowercase:
		if avoid_ambiguous {
			return "abcdefghijkmnopqrstuvwxyz"
		} else {
			return "abcdefghijklmnopqrstuvwxyz"
		}
	case CharacterUppercase:
		if avoid_ambiguous {
			return "ABCDEFGHJKLMNPQRSTUVWXYZ"
		} else {
			return "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		}
	case CharacterNumeric:
		if avoid_ambiguous {
			return "23456789"
		} else {
			return "0123456789"
		}
	case CharacterSpecial:
		return "!@#$%^&*(),.-_=+"
	case CharacterHex:
		return "0123456789abcdef"
	case CharacterBase32:
		return "abcdefghijklmnopqrstuvwxyz234567"
	case CharacterBase64:
		return "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	default:
		return ""
	}
}
