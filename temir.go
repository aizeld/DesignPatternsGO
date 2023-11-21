package main

import "fmt"

type IDeveloper interface {
	writeCode()
}

type IManager interface {
	manageProject()
}
type IDistributor interface {
	distrProject()
}

type TeamFactory interface {
	getDev() IDeveloper
	getMan() IManager
	getDist() IDistributor
}

func (o *OSDistributor) getDist() IDistributor {
	return o.getDist() ///new
}
func (o *OSDistributor) getDev() IDeveloper {
	return o.getDev() ///new
}
func (o *OSDistributor) getMan() IManager {
	return o.getMan() //new
}

//new

func newTeamFactory(teamName string) (TeamFactory, error) {
	if teamName == "OS" {
		return &OSTeam{}, nil //new
	}
	if teamName == "CPP" {
		return &CPlusPlus{}, nil //new
	}
	if teamName == "Dist" {
		return &OSDistributor{}, nil //new
	}
	return nil, fmt.Errorf("TEMIR HUESOS")
}

type CPlusPlus struct {
}

func (c *CPlusPlus) writeCode() {
	fmt.Println("CPP developer writes code")
}

func (c *CPlusPlus) getDev() IDeveloper {
	return c.getDev()
	///new
}

func (c *CPlusPlus) getDist() IDistributor {
	return c.getDist()
	///new
}

func (c *CPlusPlus) getMan() IManager {
	return c.getMan()
	///new
}

type OSManager struct {
}

func (o *OSManager) manageProject() {
	fmt.Println("OS manager manage the project")
}

type OSDistributor struct {
}

func (o *OSDistributor) distrProject() {
	fmt.Println("OS distributor distribute the project")
}

type OSTeam struct {
}

func (o *OSTeam) getDev() IDeveloper {
	return &CPlusPlus{}
}
func (o *OSTeam) getMan() IManager {
	return &OSManager{}
}
func (o *OSTeam) getDist() IDistributor {
	return &OSDistributor{}
}

func main() {
	// factory := &OSTeam{}
	// factory.getDev().writeCode()
	// factory.getDist().distrProject()
	// factory.getMan().manageProject()

	factory, _ := newTeamFactory("OS")

	factory.getDev().writeCode()
	factory.getDist().distrProject()

}
