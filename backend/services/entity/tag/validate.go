package tag

import ()

func CheckName(name string) error {
	if len(name) < 1 {
		return ErrorNameNotEmply
	}
	return nil
}
