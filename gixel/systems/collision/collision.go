package collision

import (
	"math"

	"github.com/GixelEngine/gixel-engine/gixel"
	gm "github.com/GixelEngine/gixel-engine/gixel/math"
)

type imports interface {
	gixel.GxlState // only makes sense to apply on states
}

type Exports interface {
	LastX() float64
	LastY() float64
	Immovable() *bool
}

type Collision struct {
	subject            *imports
	objectsToUpdatePos map[int]gixel.GxlObject
	lastPosMap         map[int]*gm.GxlPoint
}

func (c *Collision) Init(subject imports) {
	c.subject = &subject
	c.objectsToUpdatePos = make(map[int]gixel.GxlObject)
	c.lastPosMap = make(map[int]*gm.GxlPoint)
}

func (c *Collision) Update(elapsed float64) {
	for _, obj := range c.objectsToUpdatePos {
		c.UpdateLastPos(obj)
	}
	c.objectsToUpdatePos = make(map[int]gixel.GxlObject)
}

type collisionAttributes interface {
	Immovable() *bool
}

func (c *Collision) UpdateLastPos(objs ...gixel.GxlObject) {
	for _, obj := range objs {
		x, y := obj.GetPosition()
		c.lastPosMap[obj.GetID()] = gm.NewPoint(x, y)
	}
}

func (c *Collision) GetLastPos(obj gixel.GxlObject) gm.GxlPoint {
	pos, ok := c.lastPosMap[obj.GetID()]

	if !ok {
		return *gm.NewPoint(*obj.X(), *obj.Y())
	}

	return *pos
}

// TODO: Figure out small delta collisions (1 pixel difference)
func (c *Collision) CollideObjects(obj1, obj2 gixel.GxlObject) {
	if obj1 == obj2 {
		return
	}

	c.objectsToUpdatePos[obj1.GetID()] = obj1
	c.objectsToUpdatePos[obj2.GetID()] = obj2

	attr1, ok := obj1.(collisionAttributes)

	immovable1 := false
	if ok {
		immovable1 = *attr1.Immovable()
	}

	attr2, ok := obj2.(collisionAttributes)

	immovable2 := false
	if ok {
		immovable2 = *attr2.Immovable()
	}

	if !obj1.Overlaps(obj2) {
		return
	}

	overlapX := c.computeOverlapX(obj1, obj2)
	overlapY := c.computeOverlapY(obj1, obj2)

	if immovable1 {
		*obj2.X() += overlapX
		*obj2.Y() += overlapY
	} else if immovable2 {
		*obj1.X() -= overlapX
		*obj1.Y() -= overlapY
	} else if !immovable1 && !immovable2 {
		*obj1.X() -= overlapX / 2
		*obj1.Y() -= overlapY / 2
		*obj2.X() += overlapX / 2
		*obj2.Y() += overlapY / 2
	}
}

func (c *Collision) CollideObjectGroup(obj gixel.GxlObject, grp gixel.GxlGroup) {
	grp.Range(func(idx int, value gixel.GxlBasic) bool {
		gobj, ok := value.(gixel.GxlObject)
		if !ok {
			return true
		}

		c.CollideObjects(obj, gobj)
		return true
	})
}

func (c *Collision) CollideGroups(grp1, grp2 gixel.GxlGroup) {
	objs := make([]gixel.GxlObject, 0)

	grp1.Range(func(idx int, value gixel.GxlBasic) bool {
		obj, ok := value.(gixel.GxlObject)
		if !ok {
			return true
		}

		objs = append([]gixel.GxlObject{obj}, objs...)

		c.CollideObjectGroup(obj, grp2)
		return true
	})

	for _, obj := range objs {
		c.CollideObjectGroup(obj, grp2)
	}
}

func (c *Collision) computeOverlapX(obj1, obj2 gixel.GxlObject) float64 {
	lastX1 := c.GetLastPos(obj1).X
	lastX2 := c.GetLastPos(obj2).X

	// log.Printf("X (%f %f)", lastX1, *obj1.X())

	// calculate how much overlap
	delta1 := *obj1.X() - lastX1
	delta2 := *obj2.X() - lastX2

	if delta1 == delta2 {
		return 0
	}

	delta1Abs := math.Abs(delta1)
	delta2Abs := math.Abs(delta2)
	maxOverlap := delta1Abs + delta2Abs + 4
	overlap := 0.0

	if delta1 > delta2 {
		overlap = *obj1.X() + float64(*obj1.W()) - *obj2.X()
		if overlap > maxOverlap {
			return 0
		}
	} else if delta1 < delta2 {
		overlap = *obj1.X() - float64(*obj2.W()) - *obj2.X()
		if -overlap > maxOverlap {
			return 0
		}
	}

	return overlap
}

func (c *Collision) computeOverlapY(obj1, obj2 gixel.GxlObject) float64 {
	lastY1 := c.GetLastPos(obj1).Y
	lastY2 := c.GetLastPos(obj2).Y

	// log.Printf("Y (%f %f)", lastY1, *obj1.Y())

	// calculate how much overlap
	delta1 := *obj1.Y() - lastY1
	delta2 := *obj2.Y() - lastY2

	if delta1 == delta2 {
		return 0
	}

	delta1Abs := math.Abs(delta1)
	delta2Abs := math.Abs(delta2)
	maxOverlap := delta1Abs + delta2Abs + 4
	overlap := 0.0

	if delta1 > delta2 {
		overlap = *obj1.Y() + float64(*obj1.H()) - *obj2.Y()
		if overlap > maxOverlap {
			return 0
		}
	} else if delta1 < delta2 {
		overlap = *obj1.Y() - float64(*obj2.H()) - *obj2.Y()
		if -overlap > maxOverlap {
			return 0
		}
	}

	return overlap
}
