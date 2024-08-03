package validators

import (
	"regexp"
	"sync"
)

type validationPatterns struct {
	phoneNumber *regexp.Regexp
	docNumber   *regexp.Regexp
}

var lock = &sync.Mutex{}
var validationsInstance *validationPatterns

func instance() *validationPatterns {
	if validationsInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if validationsInstance == nil {
			validationsInstance = &validationPatterns{
				phoneNumber: regexp.MustCompile(`^[1-9]\d{9}$`),
				docNumber:   regexp.MustCompile(`^[1-9]+\d{2,9}(-\d)?$`),
			}
		}
	}

	return validationsInstance
}

func IsDocumentNumber(docNumber string) bool {
	return instance().docNumber.MatchString(docNumber)
}

func IsPhoneNumber(phoneNumber string) bool {
	return instance().phoneNumber.MatchString(phoneNumber)
}
