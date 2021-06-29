package common

type RunOrError func() error

func TryOrError(fun ...RunOrError) error {
	for _, r := range fun {
		if err := r(); err != nil {
			return err
		}
	}
	return nil
}
