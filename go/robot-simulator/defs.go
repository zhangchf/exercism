package robot

// definitions used in step 1

var Step1Robot struct {
	X, Y int
	Dir
}

type Dir int

const (
	N Dir = iota
	E
	S
	W
)

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func Right() {
	Step1Robot.Dir = (Step1Robot.Dir+1)%4
}

func Left() {
	Step1Robot.Dir = (Step1Robot.Dir-1+4)%4
}

// additional definitions used in step 2

type Action Command
const (
	R Action = 'R'
	L = 'L'
	A = 'A'
	Over = 'O'
)
type Command byte // valid values are 'R', 'L', 'A'
type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

func StartRobot(cmd chan Command, act chan Action) {
	for command := range cmd {
		act <- Action(command)
	}
	close(act)
}

func Room(rect Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	for action := range act {
		robot.takeAction(rect, action)
	}
	rep <- robot
}

func (robot *Step2Robot) takeAction(roomExtent Rect, action Action) (msg string) {
	switch action {
	case R:
		robot.Dir = (robot.Dir+1)%4
	case L:
		robot.Dir = (robot.Dir-1+4)%4
	case A:
		switch robot.Dir {
		case N:
			if robot.Northing < roomExtent.Max.Northing {
				robot.Northing++
			} else {
				msg = "bump N wall"
			}
		case E:
			if robot.Easting < roomExtent.Max.Easting {
				robot.Easting++
			} else {
				msg = "bump E wall"
			}
		case S:
			if robot.Northing > roomExtent.Min.Northing {
				robot.Northing--
			} else {
				msg = "bump S wall"
			}
		case W:
			if robot.Easting > roomExtent.Min.Easting {
				robot.Easting--
			} else {
				msg = "bump W wall"
			}
		}
	}
	return
}

// additional definition used in step 3
type Action3 struct {
	name string
	action Action
}
type Step3Robot struct {
	Name string
	Step2Robot
}

func StartRobot3(name, script string, act chan Action3, log chan string) {
	for _, c := range script {
		act <- Action3{name, Action(c)}
	}
	act <- Action3{name, Over}
}

func Room3(extent Rect, robots []Step3Robot, act chan Action3, report chan []Step3Robot, log chan string) {
	if !checkRobots(robots, extent) {
		log <- "Invalid Robots"
	} else {
		overCmds := 0
		for action3 := range act {
			if action3.action == Over {
				overCmds++
				if overCmds == len(robots) {
					close(act)
				}
			}
			for i := range robots {
				if action3.name == robots[i].Name {
					robots[i].takeAction(extent, action3.action, log)
				}
			}
		}
	}
	report <- robots	
}

func checkRobots(robots []Step3Robot, roomExtent Rect) bool {
	len := len(robots)
	for i := 0; i < len; i++ {
		if robots[i].Name == "" || (robots[i].Easting < roomExtent.Min.Easting || robots[i].Northing < roomExtent.Min.Northing || robots[i].Easting > roomExtent.Max.Northing || robots[i].Northing > roomExtent.Max.Northing) {
			return false
		}
		for j := i+1; j < len; j++ {
			if robots[i].Name == robots[j].Name || (robots[i].Easting == robots[j].Easting && robots[i].Northing == robots[j].Northing) {
				return false
			}
		}
	}
	return true
} 

func (robot3 *Step3Robot) takeAction(roomExtent Rect, action Action, log chan string) {
	msg := robot3.Step2Robot.takeAction(roomExtent, action)
	if msg != "" {
		log <- robot3.Name + msg
	}
}
