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
			switch recoverError := rec.(type) {
			case FrobError:
				resource.Defrob(recoverError.defrobTag)
				err = recoverError.inner
			case error:
				err = recoverError
			default:
				err = nil
			}
		}
		resource.Close()
	}()
	resource.Frob(input)

	return
}
