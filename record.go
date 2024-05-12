package main

import (
	"github.com/MarkKremer/microphone"
	"github.com/gopxl/beep/wav"
	"os"
	"os/signal"
)

func record() error {
	err := microphone.Init()
	if err != nil {
		return err
	}
	defer microphone.Terminate()

	stream, format, err := microphone.OpenDefaultStream(44100, 2)
	if err != nil {
		return err
	}
	defer stream.Close()

	f, err := os.Create("aa.wav")
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	go func() {
		<-sig
		stream.Stop()
		stream.Close()
	}()

	stream.Start()
	return wav.Encode(f, stream, format)
}
