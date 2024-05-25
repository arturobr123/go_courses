package main

import "fmt"

// PasswordProtector struct holds user details and the hashing algorithm to be used.
type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm // Strategy interface used for hashing algorithms.
}

// HashAlgorithm interface defines a method for hashing.
type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

// newPasswordProtector initializes a new PasswordProtector with the specified hashing algorithm.
func newPasswordProtector(user string, passwordName string, hashAlgorithm HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{user: user, passwordName: passwordName, hashAlgorithm: hashAlgorithm}
}

// setHashAlgorithm allows changing the hashing algorithm at runtime.
func (p *PasswordProtector) setHashAlgorithm(hashAlgorithm HashAlgorithm) {
	p.hashAlgorithm = hashAlgorithm
}

// hash performs the hashing operation using the current hashing strategy.
func (p *PasswordProtector) hash() {
	p.hashAlgorithm.Hash(p)
}

// SHA struct implements the HashAlgorithm interface using SHA hashing.
type SHA struct{}

func (s *SHA) Hash(p *PasswordProtector) {
	fmt.Println("Hashing using SHA for: ", p.passwordName)
}

// MD5 struct implements the HashAlgorithm interface using MD5 hashing.
type MD5 struct{}

func (s *MD5) Hash(p *PasswordProtector) {
	fmt.Println("Hashing using MD5 for: ", p.passwordName)
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	// Create a PasswordProtector with SHA algorithm.
	passwordProtector := newPasswordProtector("user", "password", sha)
	passwordProtector.hash() // Perform hashing.

	// Change the hashing algorithm to MD5.
	passwordProtector.setHashAlgorithm(md5)
	passwordProtector.hash() // Perform hashing with new algorithm.

	fmt.Println("Strategy Pattern")
}
