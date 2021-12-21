package dockersdk

func WinVolOpt(name string) volume.VolumeCreateBody {
	return volume.VolumeCreateBody{
		Driver:     "local",
		DriverOpts: map[string]string{},
		Labels:     map[string]string{},
		Name:       name,
	}
}
