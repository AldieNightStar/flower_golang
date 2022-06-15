package flower

func Try(f func()) (outErr error) {
	defer func() {
		if r := recover(); r != nil {
			if err, errOk := r.(error); errOk {
				outErr = err
			}
		}
	}()
	f()
	return outErr
}
