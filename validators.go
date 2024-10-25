package lwhelper

import (
	"regexp"
	"strings"

	"golang.org/x/net/idna"
)

var (
	emailRegex = regexp.MustCompile(`^([a-zA-Z0-9.!#$%&*+/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*)?$`)
)

func ValidEmail(s string) bool {
	return emailRegex.MatchString(s)
}

func ValidDomain(s string) bool {
	_, err := idna.Lookup.ToASCII(s)
	return err == nil
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
