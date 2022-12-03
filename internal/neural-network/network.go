package neuralnetwork

import vc "github.com/AnimemeMops/hackathonProject/internal/vector-calculator"

type Network struct {
	model     *Model
	data      []*vc.IV
	age       int64
	batchSize int64
}

func NewNetwork(modelFile string, data []*vc.IV, age, batchSize int64) *Network {
	newNet := &Network{
		data:      data,
		age:       age,
		batchSize: batchSize,
	}
	if modelFile != "" {
		newNet.model = NewModel(modelFile)
		return newNet
	}
	newNet.feed()
	return newNet
}
func (n *Network) feed() (string, error) {
	return "", nil
}

func (n *Network) Evaluate(testSample []vc.IV) (rating float64) {
	return
}

func (n *Network) Segment(imgVector []vc.IV) string {
	return ""
}

type Model struct {
}

func NewModel(filePath string) *Model {
	return &Model{}
}
