package db

import "testing"

func Test_ConnectionStringToGormString(t *testing.T) {

	testCases := map[string]string{
		"postgres://admin:admin@192.168.99.100:5432/admin?sslmode=disable":   "host=192.168.99.100 user=admin dbname=admin password=admin port=5432 sslmode=disable",
		"postgres://admin:admin@localhost:5432/admin?sslmode=disable":        "host=localhost user=admin dbname=admin password=admin port=5432 sslmode=disable",
		"mysql://Adm00n:adm1n@localhost:1234/1234mydatabase?sslmode=disable": "host=localhost user=Adm00n dbname=1234mydatabase password=adm1n port=1234 sslmode=disable",
	}
	for k, v := range testCases {
		_, res := ConnectionStringToGormString(k)
		if res != v {
			t.Fatalf("want '%s', got '%s'", v, res)
		}
	}
}
