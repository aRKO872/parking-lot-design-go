package miscallaneous

type Person struct {
	id      string
	name    string
	address Address
	phone   string
	email   string
}

func NewPerson(id, name string, address Address, phone, email string) *Person {
	return &Person{
		id:      id,
		name:    name,
		address: address,
		phone:   phone,
		email:   email,
	}
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Address() Address {
	return p.address
}

func (p *Person) SetAddress(address Address) {
	p.address = address
}

func (p *Person) Email() string {
	return p.email
}

func (p *Person) SetEmail(email string) {
	p.email = email
}

func (p *Person) Phone() string {
	return p.phone
}

func (p *Person) SetPhone(phone string) {
	p.phone = phone
}

func (p *Person) ID() string {
	return p.id
}

func (p *Person) SetID(id string) {
	p.id = id
}
