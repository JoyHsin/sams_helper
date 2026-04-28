package notice

import (
	"fmt"
	"os/exec"
)

type SoundSet struct {
	Message string `yaml:"soundMessage"`
	Times   int    `yaml:"soundTimes"`
	Voice   string `yaml:"soundVoice"`
}

// MacSound only for mac
func MacSound(soundSet SoundSet) error {
	return MacSoundMessage(soundSet, "")
}

func MacSoundMessage(soundSet SoundSet, message string) error {
	if message == "" {
		message = soundSet.Message
	}
	for i := 0; i < soundSet.Times; i++ {
		err := exec.Command("say", message, fmt.Sprintf("--voice=%s", soundSet.Voice)).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
