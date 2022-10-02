package orders

import "errors"

type Adderess struct {
	name     string
	street   string
	city     string
	postCode string
	country  string
}

func NewAdderess(name string, street string, city string, postCode string, country string) (Adderess, error) {
	if len(name) == 0 {
		return Adderess{}, errors.New("name cannot be empty")
	}
	if len(street) == 0 {
		return Adderess{}, errors.New("street cannot be empty")
	}
	if len(city) == 0 {
		return Adderess{}, errors.New("city cannot be empty")
	}
	if len(postCode) == 0 {
		return Adderess{}, errors.New("postCode cannot be empty")
	}
	if len(country) == 0 {
		return Adderess{}, errors.New("country cannot be empty")
	}

	return Adderess{name, street, city, postCode, country}, nil
}

func (a Adderess) Name() string {
	return a.name
}

func (a Adderess) Street() string {
	return a.street
}

func (a Adderess) City() string {
	return a.city
}

func (a Adderess) PostCode() string {
	return a.postCode
}

func (a Adderess) Country() string {
	return a.country
}
