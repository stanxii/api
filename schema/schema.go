package schema

import "io/ioutil"

// GetSchema returns the parsed GraphQL schema definition as a string
func GetSchema(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
