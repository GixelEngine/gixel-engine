package gixel

type CallbackFunc = func(obj1, obj2 *GxlObject)

type BaseGxlState struct {
	BaseGxlGroup
	cameras GxlCameraManager
}

func (s *BaseGxlState) Init(game *GxlGame) {
	s.cameras.Init(game)
	s.BaseGxlGroup.Init(game)
}

func (s *BaseGxlState) Cameras() *GxlCameraManager {
	return &s.cameras
}

func (s *BaseGxlState) OverlapsObjectGroup(obj GxlObject, grp GxlGroup, callbacks ...CallbackFunc) bool {
	overlapped := false

	grp.Range(func(idx int, value GxlBasic) bool {
		cobj, ok := value.(GxlObject)
		if !ok {
			return true
		}

		if !obj.Overlaps(cobj) {
			return true
		}

		overlapped = true
		for _, callback := range callbacks {
			callback(&obj, &cobj)
		}
		return true
	})

	return overlapped
}

func (s *BaseGxlState) OverlapsGroups(grp1, grp2 GxlGroup, callbacks ...CallbackFunc) bool {
	overlapped := false

	grp1.Range(func(idx int, value GxlBasic) bool {
		cobj, ok := value.(GxlObject)
		if !ok {
			return true
		}

		overlapped = overlapped || s.OverlapsObjectGroup(cobj, grp2, callbacks...)
		return true
	})

	return overlapped
}

func (s *BaseGxlState) Update(elapsed float64) error {
	err := s.BaseGxlGroup.Update(elapsed)
	if err != nil {
		return err
	}

	s.cameras.Update(elapsed)

	return nil
}

type GxlState interface {
	GxlGroup
	OverlapsObjectGroup(obj GxlObject, grp GxlGroup, callbacks ...CallbackFunc) bool
	OverlapsGroups(grp1, grp2 GxlGroup, callbacks ...CallbackFunc) bool
	Cameras() *GxlCameraManager
}
