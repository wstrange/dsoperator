package ds

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v2"
)

// DSConnection manages our ldap connection for status
type DSConnection struct {
	Username string
	Password string
	DN       string
}

// Connect to ldap
func (ds *DSConnection) Connect() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 1389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind(ds.Username, ds.Password)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Search demonstrates how to use the search interface
func (ds *DSConnection) Search() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "localhost", 1389))
	if err != nil {
		return err
	}

	defer l.Close()

	if err = l.Bind(ds.Username, ds.Password); err != nil {
		return err
	}

	searchRequest := ldap.NewSearchRequest(
		"cn=monitor", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 2000, 2000, false,
		"(objectClass=*)",    // The filter to apply
		[]string{"DN", "cn"}, // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return err
	}

	fmt.Printf("Got result %v", sr)

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("dn"))
	}

	return nil
}

// See https://github.com/go-ldap/ldap/blob/master/example_test.go
