package builder

type VehilceProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehilceProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

func (f *ManufacturingDirector) Contructor() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

type CarBuilder struct {
	v VehilceProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehilceProduct {
	return c.v
}

type BikeBuilder struct {
	v VehilceProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehilceProduct {
	return c.v
}
