package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ErrDummy = errors.New("dummy error")

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	err := outer()
	log.Error().Stack().Err(err).Msg("")

	fmt.Println("output from fmt.Printf")
	fmt.Printf("%+v", err)
}

func inner() error {

	// Record the stack.
	return errors.Wrap(ErrDummy, "inner error")
	//return ErrDummy
	//return errors.WithMessage(ErrDummy, "inner error")
}

func middle() error {
	err := inner()
	if err != nil {
		//return errors.Errorf("middle error: %v", err)
		//return errors.WithMessage(err, "middle error")
		//return errors.Wrap(err, "middle")
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		//	return errors.Errorf("outer error: %v", err)
		//return errors.Wrap(err, "outer")
		return errors.WithMessage(err, "outer error")
		//return err
	}
	return nil
}
