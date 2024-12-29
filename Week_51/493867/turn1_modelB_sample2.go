func historyCommand(filePath string) error {
	versionDirPath := filepath.Join(filepath.Dir(filePath), versionDir)
	versions, err := ioutil.ReadDir(versionDirPath)
	if err != nil {
		return err
	}

	for _, version := range versions {
		fmt.Println(version.Name())
	}
	return nil
}