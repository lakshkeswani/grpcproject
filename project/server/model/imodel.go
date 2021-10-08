package model

type IModel interface {
	ImplIModel()
}

func (receiver Vehicle) ImplIModel() {

}

func (receiver Route) ImplIModel() {

}
func (receiver User) ImplIModel() {

}

func (receiver BusRouteStop) ImplIModel() {

}
func (receiver Station) ImplIModel() {

}

func (receiver Booking) ImplIModel() {

}
func (receiver Driver) ImplIModel() {

}

func (receiver Trip) ImplIModel() {

}
