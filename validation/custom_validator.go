package validation

import (
	"fmt"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain"
)

// メール形式の正規表現
var emailRegex = regexp.MustCompile(`^[\w\.-]+@[\w\.-]+\.\w+$`)

// 電話番号（例: 090-1234-5678）
var phoneRegex = regexp.MustCompile(`^0\d{1,4}-\d{1,4}-\d{3,4}$`)

// 半角英数字
var alphaNumRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// ValidateEmailFormat メール形式バリデーション
func ValidateEmailFormat(value interface{}) error {
	email, _ := value.(string)
	if !emailRegex.MatchString(email) {
		return domain.WithInvalidArgument("invalid email format")
	}
	return nil
}

// ValidatePhoneFormat 電話番号形式（日本）のバリデーション
func ValidatePhoneFormat(value interface{}) error {
	phone, _ := value.(string)
	if !phoneRegex.MatchString(phone) {
		return domain.WithInvalidArgument("invalid phone number format")
	}
	return nil
}

// ValidateAlphaNumeric 半角英数字チェック
func ValidateAlphaNumeric(value interface{}) error {
	s, _ := value.(string)
	if !alphaNumRegex.MatchString(s) {
		return domain.WithInvalidArgument("must be alphanumeric characters only")
	}
	return nil
}

// ValidateFutureDate future date（未来日）チェック
func ValidateFutureDate(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return domain.WithInvalidArgument("invalid time format")
	}
	if !t.After(time.Now()) {
		return domain.WithInvalidArgument("must be a future date")
	}
	return nil
}

// ValidatePastDate past date（過去日）チェック
func ValidatePastDate(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return domain.WithInvalidArgument("invalid time format")
	}
	if !t.Before(time.Now()) {
		return domain.WithInvalidArgument("must be a past date")
	}
	return nil
}

// ValidateFixedLength 固定長の文字列（例: 6桁コードなど）

func ValidateFixedLength(length int) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if len(s) != length {
			return domain.WithInvalidArgument(
				fmt.Sprintf("must be exactly %d characters", length),
			)
		}
		return nil
	}
}

func validatePositiveUint64(value interface{}) error {
	v, ok := value.(uint64)
	if !ok {
		return domain.WithInvalidArgument("user_id must be a positive integer")
	}
	if v == 0 {
		return domain.WithInvalidArgument("user_id must be greater than 0")
	}
	return nil
}
