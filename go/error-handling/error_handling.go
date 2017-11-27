package erratum

const testVersion = 2

func Use(opener ResourceOpener, inp string) (err error) {
	var res Resource
	for {
		res, err = opener()
		if _, ok := err.(TransientError); ok {
			continue
		}
		if err != nil {
			return err
		}
		break
	}

	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			if fe, ok := r.(FrobError); ok {
				res.Defrob(fe.defrobTag)
				err = fe
			}
			err = r.(error)
		}
	}()
	res.Frob(inp)

	return err
}
