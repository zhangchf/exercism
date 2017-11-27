package react

const testVersion = 5

/************* cellImpl ************/
type cellImpl struct {
	data int
}

func (cell *cellImpl) Value() int {
	return cell.data
}

/************* inputCellImpl ************/
type inputCellImpl struct {
	cellImpl
	*reactorImpl
}

func (cell *inputCellImpl) SetValue(value int) {
	cell.setValue(value)
}

func (cell *inputCellImpl) setValue(value int) {
	if cell.data == value {
		return
	}
	cell.data = value
	for _, computeCell := range cell.updates[cell] {
		computeCell.updateValue()
	}
}

/************* computeCellImpl ************/
type computeCellImpl struct {
	cache int
	*reactorImpl
	callbacks []*func(int)
	//computeFunc interface{}
	computeFunc1 func(int) int
	computeFunc2 func(int, int) int
}

func (cell *computeCellImpl) Value() int {
	return cell.updateValue()
}

func (cell *computeCellImpl) updateValue() int {
	var v int
	dependencies := cell.dependencies[cell]
	if cell.computeFunc1 != nil {
		cell1Value := dependencies[0].Value()
		v = cell.computeFunc1(cell1Value)
	} else {
		cell1Value := dependencies[0].Value()
		cell2Value := dependencies[1].Value()
		v = cell.computeFunc2(cell1Value, cell2Value)
	}
	cell.setValue(v)
	return v
}

func (cell *computeCellImpl) setValue(value int) {
	if cell.cache == value {
		return
	}
	cell.cache = value
	for _, computeCell := range cell.updates[cell] {
		computeCell.updateValue()
	}
	for _, callback := range cell.callbacks {
		(*callback)(value)
	}
}

func (cell *computeCellImpl) AddCallback(callback func(int)) Canceler {
	cell.callbacks = append(cell.callbacks, &callback)
	canceler := cancelerImpl{computeCell: cell, callback: &callback}
	return &canceler
}

func (cell *computeCellImpl) cancelCallback(callback *func(int)) {
	for i, v := range cell.callbacks {
		if v == callback {
			cell.callbacks = append(cell.callbacks[:i], cell.callbacks[i+1:]...)
			break
		}
	}
}

/************* cancelerCellImpl ************/
type cancelerImpl struct {
	computeCell *computeCellImpl
	callback    *func(int)
}

func (canceler *cancelerImpl) Cancel() {
	canceler.computeCell.cancelCallback(canceler.callback)
}

/************* reactorImpl ************/
type reactorImpl struct {
	updates      map[Cell][]*computeCellImpl
	dependencies map[ComputeCell][]Cell
}

func (reactor *reactorImpl) CreateInput(initValue int) InputCell {
	inputCell := inputCellImpl{
		cellImpl:    cellImpl{data: initValue},
		reactorImpl: reactor,
	}
	return &inputCell
}

func (reactor *reactorImpl) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	computeCell := &computeCellImpl{reactorImpl: reactor, computeFunc1: compute}
	reactor.updates[cell] = append(reactor.updates[cell], computeCell)
	reactor.dependencies[computeCell] = append(reactor.dependencies[computeCell], cell)
	computeCell.updateValue()
	return computeCell
}

func (reactor *reactorImpl) CreateCompute2(cell1 Cell, cell2 Cell, compute func(int, int) int) ComputeCell {
	computeCell := &computeCellImpl{reactorImpl: reactor, computeFunc2: compute}
	reactor.updates[cell1] = append(reactor.updates[cell1], computeCell)
	reactor.updates[cell2] = append(reactor.updates[cell2], computeCell)
	reactor.dependencies[computeCell] = append(reactor.dependencies[computeCell], cell1, cell2)
	computeCell.updateValue()
	return computeCell
}

func New() Reactor {
	return &reactorImpl{
		updates:      make(map[Cell][]*computeCellImpl),
		dependencies: make(map[ComputeCell][]Cell),
	}
}
