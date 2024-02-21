package miscallaneous

type Address struct {
	id      string
	zipCode string
	address string
	city    string
	state   string
	country string
}

func NewAddress(id, zipCode, address, city, state, country string) *Address {
	return &Address{
		id:      id,
		zipCode: zipCode,
		address: address,
		city:    city,
		state:   state,
		country: country,
	}
}

func (a *Address) ID() string {
	return a.id
}

func (a *Address) SetID(id string) {
	a.id = id
}

func (a *Address) ZipCode() string {
	return a.zipCode
}

func (a *Address) SetZipCode(zipCode string) {
	a.zipCode = zipCode
}

func (a *Address) Address() string {
	return a.address
}

func (a *Address) SetAddress(address string) {
	a.address = address
}

func (a *Address) City() string {
	return a.city
}

func (a *Address) SetCity(city string) {
	a.city = city
}

func (a *Address) State() string {
	return a.state
}

func (a *Address) SetState(state string) {
	a.state = state
}

func (a *Address) Country() string {
	return a.country
}

func (a *Address) SetCountry(country string) {
	a.country = country
}

