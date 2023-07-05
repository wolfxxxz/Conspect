package main

import "fmt"

type Processor interface {
	getDescription() string
}

type Motherboard interface {
	getDescription() string
}

type VideoCard interface {
	getDescription() string
}

type ProcessorFactory interface {
	createIntelProcessor() Processor
	createAMDProcessor() Processor
}

type MotherboardFactory interface {
	createAsusMotherboard() Motherboard
	createGigabyteMotherboard() Motherboard
}

type VideoCardFactory interface {
	createNvidiaVideoCard() VideoCard
	createAMDVideoCard() VideoCard
}

// -Product-------------------------------
type IntelProcessor struct{}

func (p IntelProcessor) getDescription() string {
	return "I am an Intel Processor"
}

type AMDProcessor struct{}

func (p AMDProcessor) getDescription() string {
	return "I am an AMD Processor"
}

type AsusMotherboard struct{}

func (m AsusMotherboard) getDescription() string {
	return "I am an Asus Motherboard"
}

type GigabyteMotherboard struct{}

func (m GigabyteMotherboard) getDescription() string {
	return "I am a Gigabyte Motherboard"
}

type NvidiaVideoCard struct{}

func (v NvidiaVideoCard) getDescription() string {
	return "I am a Nvidia VideoCard"
}

type AMDVideoCard struct{}

func (v AMDVideoCard) getDescription() string {
	return "I am an AMD VideoCard"
}

// -Factory---------------------------------
type ProcessorFactoryChina struct{}

func (p ProcessorFactoryChina) createAMDProcessor() Processor {
	return AMDProcessor{}
}
func (p ProcessorFactoryChina) createIntelProcessor() Processor {
	return IntelProcessor{}
}

type MotherboardFactoryUSA struct{}

func (m MotherboardFactoryUSA) createAsusMotherboard() Motherboard {
	return AsusMotherboard{}
}
func (m MotherboardFactoryUSA) createGigabyteMotherboard() Motherboard {
	return GigabyteMotherboard{}
}

type VideoCardFactoryUkraine struct{}

func (v VideoCardFactoryUkraine) createNvidiaVideoCard() VideoCard {
	return NvidiaVideoCard{}
}
func (v VideoCardFactoryUkraine) createAMDVideoCard() VideoCard {
	return AMDVideoCard{}
}

type Computer struct {
	proc        Processor
	motherBoard Motherboard
	videoCard   VideoCard
}

func (c Computer) getDescription() string {
	return fmt.Sprintf("This computer consists of:\nProcessor: %s\nMotherboard: %s\nVideo Card: %s",
		c.proc.getDescription(), c.motherBoard.getDescription(), c.videoCard.getDescription())
}

func main() {
	var comp Computer
	procFactory := ProcessorFactoryChina{}
	processoR := procFactory.createAMDProcessor()
	comp.proc = processoR

	videoFactory := VideoCardFactoryUkraine{}
	videocard := videoFactory.createNvidiaVideoCard()
	comp.videoCard = videocard

	FactoryMotherboard := MotherboardFactoryUSA{}
	motherBoard := FactoryMotherboard.createGigabyteMotherboard()
	comp.motherBoard = motherBoard

	fmt.Println(comp.getDescription())
}
