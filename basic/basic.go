package basic

import "errors"

type Db interface {
	Get(key string) (string, error)
	Put(key string, value string) error
}

// StoreMulti inserts multiple values into db
func StoreMulti(keys, values []string, db Db) error {
	if len(keys) != len(values) {
		return errors.New("length mismatch")
	}

	for i, key := range keys {
		err := db.Put(key, values[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// FetchMulti retrieves multiple values from db
func FetchMulti(keys []string, db Db) ([]string, error) {
	var res []string

	for _, key := range keys {
		val, err := db.Get(key)
		if err != nil {
			return nil, err
		}
		res = append(res, val)
	}
	return res, nil
}
