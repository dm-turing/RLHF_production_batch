func listVersions(file string) ([]int, error) {
	var versions []int
	files, err := ioutil.ReadDir(fmt.Sprintf("%s", file))
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}
		if fileInfo.Name() == "version" || !strings.HasPrefix(fileInfo.Name(), fmt.Sprintf("%s.", file)) {
			continue
		}
		version, err := strconv.Atoi(strings.TrimSuffix(fileInfo.Name(), ".txt"))
		if err == nil {
			versions = append(versions, version)
		}
	}

	return versions, nil
}