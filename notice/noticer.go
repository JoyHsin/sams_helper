package notice

type NoticerSet struct {
	BarkSet     BarkSet  `yaml:"bark"`
	SoundSet    SoundSet `yaml:"sound"`
	FtqqSet     FTQQSet  `yaml:"ftqq"`
	EmailSet    EmailSet `yaml:"email"`
	NoticeType  int      `yaml:"noticeType"`
	NoticeTypes []int    `yaml:"noticeTypes"`
}

func Do(noticerSet NoticerSet) error {
	return DoMessage(noticerSet, "")
}

func DoMessage(noticerSet NoticerSet, message string) error {
	noticeTypes := noticerSet.NoticeTypes
	if len(noticeTypes) == 0 {
		noticeTypes = []int{noticerSet.NoticeType}
	}
	for _, noticeType := range noticeTypes {
		if err := doOne(noticerSet, noticeType, message); err != nil {
			return err
		}
	}
	return nil
}

func doOne(noticerSet NoticerSet, noticeType int, message string) error {
	switch noticeType {
	case 0:
		return nil
	case 1:
		return BarkPushMessage(noticerSet.BarkSet, message)
	case 2:
		return FTQQPushMessage(noticerSet.FtqqSet, message)
	case 3:
		return MacSoundMessage(noticerSet.SoundSet, message)
	case 4:
		return EmailPushMessage(noticerSet.EmailSet, message)
	default:
		return nil
	}
}
