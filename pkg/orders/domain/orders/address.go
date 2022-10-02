package orders

import "errors"

type Address struct {
	name     string
	street   string
	city     string
	postCode string
	country  string
}

func NewAdderess(name string, street string, city string, postCode string, country string) (Address, error) {
	if len(name) == 0 {
		return Address{}, errors.New("name cannot be empty")
	}
	if len(street) == 0 {
		return Address{}, errors.New("street cannot be empty")
	}
	if len(city) == 0 {
		return Address{}, errors.New("city cannot be empty")
	}
	if len(postCode) == 0 {
		return Address{}, errors.New("postCode cannot be empty")
	}
	if len(country) == 0 {
		return Address{}, errors.New("country cannot be empty")
	}

	return Address{name, street, city, postCode, country}, nil
}

func (a Address) Name() string {
	return a.name
}

func (a Address) Street() string {
	return a.street
}

func (a Address) City() string {
	return a.city
}

func (a Address) PostCode() string {
	return a.postCode
}

func (a Address) Country() string {
	return a.country
}
