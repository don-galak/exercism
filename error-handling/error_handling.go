package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource
	for {
		resource, err = opener()
		if _, ok := err.(TransientError); ok {
			continue
		}

		if err != nil {
			return
		}
		break
	}

	defer func() {
		if rec := recover(); rec != nil {
			if recoverError, ok := rec.(FrobError); ok {
				resource.Defrob(recoverError.defrobTag)
			}
			err, _ = rec.(error)
		}
		resource.Close()
	}()
	resource.Frob(input)

	return
}
