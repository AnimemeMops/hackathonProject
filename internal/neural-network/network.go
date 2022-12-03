package neuralnetwork

type Network struct {
	model Model
}
type Model struct {
}

func NewNetwork() *Network {
	return &Network{}
}
func (n *Network) feed(data [][]float64, trainSample, epoch, butchSise int64) error {
	return nil
}

func (n *Network) Evaluate(testSample [][]float64) (rating float64) {
	return
}

func (n *Network) Analysis(imgVector []float64) string {
	return ""
}
