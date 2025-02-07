package lwhelper

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/idna"
)

var (
	emailRegex  = regexp.MustCompile(`^([a-zA-Z0-9.!#$%&*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*)?$`)
	digitsRegex = regexp.MustCompile(`\D`)
)

func ValidEmail(s string) bool {
	return emailRegex.MatchString(s)
}

func ValidDomain(s string) bool {
	_, err := idna.Lookup.ToASCII(s)
	return err == nil
}

func ValidDomainOrWildcard(s string) bool {
	return ValidDomain(strings.TrimPrefix(s, "*."))
}

func GetDomainBaseName(domainName string) string {

	t := strings.Split(domainName, ".")
	if len(t) < 2 {
		return ""
	}

	tld := 2
	if strings.HasSuffix(domainName, ".gov.pl") ||
		strings.HasSuffix(domainName, ".com.pl") ||
		strings.HasSuffix(domainName, ".co.uk") ||
		strings.HasSuffix(domainName, ".com.br") {
		tld = 3
	}

	//

	return strings.Join(t[len(t)-tld:], ".")
}

//

func ValidPhone(s, countryPrefix string) string {

	// usuwam wszystko co nie jest cyfra
	s = digitsRegex.ReplaceAllString(s, "")

	// usuwam prefix
	s = strings.TrimPrefix(s, countryPrefix)

	// 000 000 000
	if len(s) == 9 {
		// nr jest poprawny
		return fmt.Sprintf("+%s %s", countryPrefix, s)
	}

	// nr jest NIEpoprawny
	return ""

}
